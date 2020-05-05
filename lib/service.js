var async = require("async");
var moment = require('moment');
var fs = require('fs');
var os = require('os');
var path = require('path');
var xml2js = require('xml2js');
var xlsx = require('node-xlsx').default

var Report = require('nervatura-report/dist/report.node');
var connect = require('./adapter.js').connect();
var models = require('./adapter.js').models();

function call_menu_cmd(params, _callback) {
  //sample code
  var fnum1 = 0;
  if (typeof params.number_1 !== "undefined") {
    fnum1 = parseFloat(params.number_1);
  }
  var fnum2 = 0;
  if (typeof params.number_2 !== "undefined") {
    fnum2 = parseFloat(params.number_2);
  }
  _callback(null, "Successfully processed: " + (fnum1 + fnum2).toString());
}

function print_queue(nstore, params, _callback) {
  if (typeof params.items === "undefined") { params.items = []; }
  var conn = params.conn;
  async.waterfall([
    function (callback) {
      if (!params.printer) {
        callback(nstore.lang().missing_required_field + " printer");
      }
      else {
        var printer = {
          name: params.printer,
          options: { orientation: params.orientation || "P", size: params.size || "A4" }
        }
        if (printer.name === "mail") {
          printer.options.method = "mail";
        }
        var _sql = [
          {
            select: ["pt.fieldname as fieldname", "pt.value as value"], from: "tool too",
            inner_join: ["groups tg", "on", [["too.toolgroup", "=", "tg.id"],
            ["and", "tg.groupvalue", "=", "'printer'"]]],
            left_join: ["fieldvalue pt", "on", [["pt.ref_id", "=", "too.id"],
            ["and", "pt.fieldname", "in", [[], "'tool_printertype'", "'tool_printer_mail_smtp'",
              "'tool_printer_mail_sender'", "'tool_printer_mail_login'", "'tool_printer_mail_address'",
              "'tool_printer_mail_subject'", "'tool_printer_mail_message'"]],
            ["and", "pt.deleted", "=", "0"]]],
            where: [["too.deleted", "=", "0"], ["and", "too.serial", "=", "?"]]
          },
          {
            union_select: ["fieldname", "value"], from: "fieldvalue",
            where: [["ref_id", "is", "null"], ["and", "fieldname", "in", [[], "'printer_gsprint'",
              "'default_mail_smtp'", "'default_mail_sender'"]]]
          }]
        conn.query(models.getSql(nstore.engine(), _sql), [printer.name],
          function (err, data) {
            if (err) { callback(err); }
            else {
              data.rows.forEach(function (fvalue) {
                if (fvalue.fieldname === "tool_printertype") {
                  printer.options.method = fvalue.value;
                }
                else {
                  printer.options[fvalue.fieldname.replace("tool_", "")] = fvalue.value;
                }
              });
              if (!printer.options.method) {
                callback(nstore.lang().invalid_printer + ": " + printer.name);
              }
              else {
                callback(null, printer);
              }
            }
          });
      }
    },

    function (printer, callback) {
      if (printer.options.method) {
        switch (printer.options.method) {
          case "mail":
            try {
              printer.driver = require('nodemailer');
            }
            catch (err) {
              return callback(err);
            }
            if (params.smtp) {
              printer.options.printer_mail_smtp = params.smtp;
            }
            else if (!printer.options.printer_mail_smtp) {
              if (!printer.options.default_mail_smtp || printer.options.default_mail_smtp === ""
                || printer.options.default_mail_smtp === null) {
                return callback(nstore.lang().missing_required_field + " printer_mail_smtp");
              }
              else {
                printer.options.printer_mail_smtp = printer.options.default_mail_smtp;
              }
            }

            if (params.sender) {
              printer.options.printer_mail_sender = params.sender;
            }
            else if (!printer.options.printer_mail_sender) {
              if (!printer.options.default_mail_sender || printer.options.default_mail_sender === ""
                || printer.options.default_mail_sender === null) {
                return callback(nstore.lang().missing_required_field + " printer_mail_sender");
              }
              else {
                printer.options.printer_mail_sender = printer.options.default_mail_sender;
              }
            }

            if (params.address) {
              printer.options.printer_mail_address = params.address;
            }
            else if (!printer.options.printer_mail_address) {
              return callback(nstore.lang().missing_required_field + " printer_mail_address");
            }

            if (params.subject) {
              printer.options.printer_mail_subject = params.subject;
            }
            else if (!printer.options.printer_mail_subject) {
              printer.options.printer_mail_subject = "";
            }
            if (params.message) {
              printer.options.printer_mail_message = params.message;
            }
            else if (!printer.options.printer_mail_message) {
              printer.options.printer_mail_message = "";
            }
            callback(null, printer);
            break;

          case "local":
          case "network":
            try {
              printer.driver = require('printer');
            }
            catch (err) {
              return callback(err);
            }
            printer.printers = printer.driver.getPrinters();
            printer.printers.forEach(function (pe) {
              if (pe.name === printer.name) {
                printer.printer = pe;
              }
            });
            if (!printer.printer) {
              callback(nstore.lang().invalid_printer + ": " + printer.name);
            }
            else if (process.platform === "win32" && !printer.options.printer_gsprint) {
              callback(nstore.lang().missing_required_field + " printer_gsprint");
            }
            else {
              callback(null, printer);
            }
            break;
          default:
            callback(nstore.lang().invalid_fieldname + printer.options.method);
            break;
        }
      }
      else {
        callback(nstore.lang().missing_required_field + ": printertype");
      }
    },

    function (printer, callback) {
      var qid = []
      params.items.forEach(function (item) {
        qid.push(item.id);
      });
      var _sql = { select: ["*"], from: "ui_printqueue", where: ["id", "in", [[], qid.toString()]] }
      conn.query(models.getSql(nstore.engine(), _sql), [], function (error, data) {
        if (error) { callback(error); }
        else {
          printer.items = data.rows;
          callback(null, printer);
        }
      });
    },

    function (printer, callback) {
      var values_lst = [];
      printer.items.forEach(function (item) {
        values_lst.push(function (callback_) {
          params = {
            filters: { "@id": item.ref_id }, report_id: item.report_id, output: "pdf",
            orientation: printer.options.orientation, "size": printer.options.size
          }
          get_report(nstore, params, function (err, result) {
            callback_(err, result);
          });
        });
      });
      async.series(values_lst, function (err, results) {
        if (!err) {
          for (var i = 0; i < results.length; i++) {
            printer.items[i].data = results[i].template;
          }
        }
        callback(err, printer);
      });
    },

    function (printer, callback) {
      var doc_lst = [];
      switch (printer.options.method) {
        case "mail":
          var transporter = printer.driver.createTransport(printer.options.printer_mail_smtp);
          var mail = {
            from: printer.options.printer_mail_sender,
            to: printer.options.printer_mail_address,
            subject: printer.options.printer_mail_subject,
            text: printer.options.printer_mail_message,
            attachments: []
          }
          printer.items.forEach(function (item) {
            doc_lst.push(item.id);
            mail.attachments.push({
              filename: "ntura_report_" + item.id + ".pdf",
              content: item.data,
              contentType: 'application/pdf'
            })
          });
          transporter.sendMail(mail, function (err, info) {
            callback(err, doc_lst);
          });
          break;

        case "local":
        case "network":
          var orient = printer.options.orientation.toUpperCase();
          if (process.platform === "win32") {
            const exec = require('child_process').exec;
            var cmd = printer.options.printer_gsprint;
            cmd += " @tmp_file"
            cmd += " -printer " + printer.name;
            if (orient === "P" || orient === "PORTRAIT") {
              cmd += " -portrait";
            }
            else {
              cmd += " -landscape";
            }
            cmd += " -copies @qty";
            cmd += " -option -sPAPERSIZE=" + printer.options.size.toUpperCase();

            printer.items.forEach(function (item) {
              var tfile = path.join(os.tmpdir(), "ntura_report_" + item.id + ".pdf");
              var copies_lst = []; var idx = 0;
              for (var c = 0; c < item.qty; c++) {
                copies_lst.push(function (callback_) {
                  idx += 1;
                  exec(cmd.replace("@tmp_file", tfile).replace("@qty", item.qty), function (err, stdout, stderr) {
                    callback_(err, [item.id, idx]);
                  });
                });
              }
              doc_lst.push(function (_callback) {
                fs.writeFile(tfile, item.data, function (err) {
                  if (err) {
                    _callback(err);
                  }
                  else {
                    async.series(copies_lst, function (err, results) {
                      fs.unlink(tfile);
                      _callback(err, item.id);
                    });
                  }
                });
              });
            });
          }
          else {
            var print_options = {
              printer: printer.name,
              type: "PDF", options: {}
            }
            if (orient === "L" || orient === "LANDSCAPE") {
              print_options.options.landscape = "True";
            }
            print_options.options.media = printer.options.size.toUpperCase();
            printer.items.forEach(function (item) {
              print_options.data = item.data;
              var copies_lst = []; var idx = 0;
              for (var c = 0; c < item.qty; c++) {
                copies_lst.push(function (callback_) {
                  idx += 1;
                  print_options.docname = "ntura_report_" + item.id + "_" + idx;
                  print_options.success = function (id) {
                    callback_(null, [item.id, c + 1, id]);
                  }
                  print_options.error = function (err) {
                    callback_(err);
                  }
                  printer.driver.printDirect(print_options);
                });
              }
              doc_lst.push(function (_callback) {
                async.series(copies_lst, function (err, results) {
                  _callback(err, item.id);
                });
              });
            });
          }
          async.series(doc_lst, function (err, results) {
            callback(err, results);
          });
          break;
      }
    },

    function (results, callback) {
      var values_lst = [];
      results.forEach(function (id) {
        values_lst.push(function (callback_) {
          nstore.connect.deleteData({ nervatype: "ui_printqueue", ref_id: id },
            function (err, ref_id) {
              callback_(err, ref_id);
            });
        });
      });
      async.series(values_lst, function (err, results) {
        callback(err);
      });
    }
  ],
    function (err) {
      if (err) {
        if (err.message) { err = err.message; }
        _callback(err, null);
      }
      else { _callback(null, true); }
    })
}

function get_price_value(nstore, params, _callback) {
  if (typeof params.vendorprice === "undefined") { params.vendorprice = 0; }
  if (typeof params.product_id === "undefined") { params.product_id = null; }
  if (typeof params.posdate === "undefined") {
    params.posdate = moment().format("YYYY-MM-DD");
  }
  else {
    params.posdate = moment(params.posdate,"YYYY-MM-DD").format("YYYY-MM-DD");
  }
  if (typeof params.curr === "undefined") { params.curr = null; }
  if (typeof params.qty === "undefined") { params.qty = 0; }
  if (typeof params.customer_id === "undefined") { params.customer_id = null; }

  var fxprice = 0; var discount = 0;
  var conn = params.conn;
  async.waterfall([
    function (callback) {
      //best_listprice
      if (params.product_id === null) {
        callback(nstore.lang().missing_required_field + ": product_id");
      }
      else if (params.curr === null) {
        callback(nstore.lang().missing_required_field + ": curr");
      }
      else {
        var _sql = {
          select: ["min(p.pricevalue) as mp"],
          from: "price p",
          left_join: ["link l", "on", [["l.ref_id_1", "=", "p.id"],
          ["and", "l.nervatype_1", "=", [{
            select: ["id"], from: "groups",
            where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'price'"]]
          }]],
          ["and", "l.deleted", "=", "0"]]],
          where: [["p.deleted", "=", "0"], ["and", "p.discount", "is", "null"], ["and", "p.pricevalue", "<>", "0"],
          ["and", "l.ref_id_2", "is", "null"], ["and", "p.vendorprice", "=", "?"], ["and", "p.product_id", "=", "?"],
          ["and", "p.validfrom", "<=", "?"], ["and", [["p.validto", ">=", "?"], ["or", "p.validto", "is", "null"]]],
          ["and", "p.curr", "=", "?"], ["and", "p.qty", "<=", "?"]]
        }
        conn.query(models.getSql(nstore.engine(), _sql),
          [params.vendorprice, params.product_id, params.posdate, params.posdate,
          params.curr, params.qty], function (err, data) {
            if (data.rowCount > 0) {
              if (data.rows[0].mp !== null) {
                fxprice = data.rows[0].mp;
              }
              callback(err);
            }
          });
      }
    },

    function (callback) {
      //customer discount
      if (params.customer_id === null) { callback(null); }
      else {
        var _sql = { select: ["*"], from: "customer", where: ["id", "=", "?"] }
        conn.query(models.getSql(nstore.engine(), _sql), [params.customer_id], function (err, data) {
          if (data.rowCount > 0) {
            if (data.rows[0].discount !== null) {
              discount = data.rows[0].discount;
            }
          }
          callback(err);
        });
      }
    },

    function (callback) {
      //best_custprice
      if (params.customer_id === null) { callback(null); }
      else {
        var _sql = {
          select: ["min(p.pricevalue) as mp"],
          from: "price p",
          inner_join: ["link l", "on", [
            ["l.ref_id_1", "=", "p.id"],
            ["and", "l.nervatype_1", "=", [{
              select: ["id"], from: "groups",
              where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'price'"]]
            }]],
            ["and", "l.nervatype_2", "=", [{
              select: ["id"], from: "groups",
              where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'customer'"]]
            }]],
            ["and", "l.deleted", "=", "0"]]],
          where: [["p.deleted", "=", "0"], ["and", "p.discount", "is", "null"],
          ["and", "p.pricevalue", "<>", "0"], ["and", "p.vendorprice", "=", "?"],
          ["and", "p.product_id", "=", "?"], ["and", "p.validfrom", "<=", "?"],
          ["and", [["p.validto", ">=", "?"], ["or", "p.validto", "is", "null"]]],
          ["and", "p.curr", "=", "?"], ["and", "p.qty", "<=", "?"], ["and", "l.ref_id_2", "=", "?"]]
        }
        conn.query(models.getSql(nstore.engine(), _sql),
          [params.vendorprice, params.product_id, params.posdate,
          params.posdate, params.curr, params.qty, params.customer_id], function (err, data) {
            if (data.rowCount > 0) {
              if (data.rows[0].mp !== null && (data.rows[0].mp < fxprice || fxprice === 0)) {
                fxprice = data.rows[0].mp; discount = 0;
              }
            }
            callback(err);
          });
      }
    },

    function (callback) {
      //best_grouprice
      if (params.customer_id === null) { callback(null); }
      else {
        var _sql = {
          select: ["min(p.pricevalue) as mp"],
          from: "price p",
          inner_join: [
            ["link l", "on", [["l.ref_id_1", "=", "p.id"], ["and", "l.deleted", "=", "0"],
            ["and", "l.nervatype_1", "=", [{
              select: ["id"], from: "groups",
              where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'price'"]]
            }]],
            ["and", "l.nervatype_2", "=", [{
              select: ["id"], from: "groups",
              where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'groups'"]]
            }]]]],
            ["groups g", "on", [["g.id", "=", "l.ref_id_2"],
            ["and", "g.id", "in", [{
              select: ["l.ref_id_2"], from: "link l",
              where: [["l.deleted", "=", "0"],
              ["and", "l.nervatype_1", "=", [{
                select: ["id"], from: "groups",
                where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'customer'"]]
              }]],
              ["and", "l.nervatype_2", "=", [{
                select: ["id"], from: "groups",
                where: [["groupname", "=", "'nervatype'"], ["and", "groupvalue", "=", "'groups'"]]
              }]],
              ["and", "l.ref_id_1", "=", "?"]]
            }]]]]],
          where: [["p.deleted", "=", "0"], ["and", "p.discount", "is", "null"],
          ["and", "p.pricevalue", "<>", "0"], ["and", "p.vendorprice", "=", "?"],
          ["and", "p.product_id", "=", "?"], ["and", "p.validfrom", "<=", "?"],
          ["and", [["p.validto", ">=", "?"], ["or", "p.validto", "is", "null"]]],
          ["and", "p.curr", "=", "?"], ["and", "p.qty", "<=", "?"]]
        }
        conn.query(models.getSql(nstore.engine(), _sql),
          [params.customer_id, params.vendorprice, params.product_id,
          params.posdate, params.posdate, params.curr, params.qty], function (err, data) {
            if (data.rowCount > 0) {
              if (data.rows[0].mp !== null && (data.rows[0].mp < fxprice || fxprice === 0)) {
                fxprice = data.rows[0].mp; discount = 0;
              }
            }
            callback(err);
          });
      }
    }
  ],
    function (err) {
      if (err) { if (err.message) { err = err.message; } }
      _callback(err, { price: fxprice, discount: discount });
    })
}

function get_report_files(nstore, params, _callback) {
  if (typeof params.report_dir === "undefined") {
    params.report_dir = nstore.report_dir();
  }
  if (typeof params.filters === "undefined") {
    params.filters = {};
  }
  if (typeof params.dbs_reports === "undefined") {
    params.dbs_reports = {};
  }
  if (typeof params.filters.repname === "undefined") {
    params.filters.repname = "";
  }
  if (typeof params.filters.group === "undefined") {
    params.filters.group = "";
  }
  var reports = [];

  async.waterfall([
    function (callback) {
      fs.readdir(params.report_dir, function (err, files) {
        callback(err, files);
      });
    },

    function (files, callback) {
      if (files.length > 0) {
        var parser = new xml2js.Parser();
        var templates = []; var parser_lst = [];
        files.forEach(function (fname) {
          parser_lst.push(function (callback_) {
            fs.readFile(path.join(params.report_dir, fname), function (err, data) {
              if (!err) {
                parser.parseString(data, function (err, result) {
                  if (err === null && (typeof result !== "undefined") && result !== null) {
                    if (typeof result.report !== "undefined") {
                      templates.push({ filename: fname, data: result.report });
                    }
                  }
                  callback_(null);
                });
              }
              else { callback_(null); }
            });
          });
        });
        async.series(parser_lst, function (err) {
          callback(null, files, templates);
        });
      }
      else { callback(null, files, []); }
    },

    function (files, templates, callback) {
      templates.forEach(function (template) {
        if (template.data["$"].reportkey !== "") {
          var report = {
            reportkey: template.data["$"].reportkey, repname: template.data["$"].repname,
            description: template.data["$"].description, reptype: template.data["$"].filetype,
            filename: template.filename
          };
          if (template.data["$"].nervatype === "trans") {
            report.label = template.data["$"].transtype;
          }
          else {
            report.label = template.data["$"].nervatype;
          }
          if (files.indexOf(template.filename.replace("xml", "png")) > -1) {
            report.preview = template.filename.replace(".xml", "");
          }
          if ((params.filters.alias !== "") && (typeof params.dbs_reports[report.reportkey] !== "undefined")) {
            report.installed = true;
          }
          else { report.installed = false; }
          var filter = true;
          if (params.filters.group !== "") {
            if (report.label.indexOf(params.filters.group) === -1) {
              filter = false;
            }
          }
          if (filter === true) {
            reports.push(report);
          }
        }
      });
      callback(null);
    }

  ],
    function (err) {
      if (err) { if (err.message) { err = err.message; } }
      _callback(err, reports);
    });
}

function get_report(nstore, params, _callback) {
  var conn = params.conn;
  if (typeof params.orientation === "undefined") { params.orientation = "p"; }
  if (typeof params.size === "undefined") { params.size = "a4"; }
  if (typeof params.filters === "undefined") { params.filters = {}; }
  if (!conn) { conn = nstore.connect.getConnect(); }
  async.waterfall([
    function (callback) {
      if (!conn) {
        callback(nstore.lang().not_connect);
      }
      else {
        if (params.nervatype && params.refnumber && !params.filters["@id"]) {
          nstore.valid.getIdFromRefnumber({ nervatype: params.nervatype, refnumber: params.refnumber },
            function (err, ref_id, info) {
              if (err || ref_id === null) {
                callback(nstore.lang().not_exist);
              }
              else {
                params.filters["@id"] = ref_id;
                if (!params.reportkey && !params.report_id) {
                  //set default report
                  var _sql = {
                    select: ["r.*", "ft.groupvalue as reptype"],
                    from: "ui_report r",
                    inner_join: [["groups ft", "on", ["r.filetype", "=", "ft.id"]],
                    ["groups nt", "on", [
                      ["r.nervatype", "=", "nt.id"], ["and", "nt.groupvalue", "=", "'" + params.nervatype + "'"]]]]
                  }
                  if (params.nervatype === "trans") {
                    _sql.inner_join.push(["groups tt", "on", [
                      ["r.transtype", "=", "tt.id"], ["and", "tt.groupvalue", "=", "'" + info.transtype + "'"]]]);
                    _sql.inner_join.push(["groups dir", "on", [
                      ["r.direction", "=", "dir.id"], ["and", "dir.groupvalue", "=", "'" + info.direction + "'"]]]);
                  }
                  conn.query(models.getSql(nstore.engine(), _sql), [], function (err, data) {
                    if (err) { callback(err); }
                    else if (data.rows.length === 0) {
                      callback(nstore.lang().not_exist);
                    }
                    else { callback(null, { report: data.rows[0] }) }
                  });
                }
                else {
                  callback(null, {});
                }
              }
            });
        }
        else {
          callback(null, {});
        }
      }
    },

    function (results, callback) {
      if (!results.report) {
        if (!params.reportkey && !params.report_id) {
          callback(nstore.lang().missing_required_field + " reportkey or report_id");
        }
        else {
          var _sql = {
            select: ["r.*", "ft.groupvalue as reptype"],
            from: "ui_report r",
            inner_join: ["groups ft", "on", ["r.filetype", "=", "ft.id"]]
          }
          if (params.reportkey) {
            _sql.where = ["r.reportkey", "=", "'" + params.reportkey + "'"];
          }
          else {
            _sql.where = ["r.id", "=", params.report_id];
          }
          conn.query(models.getSql(nstore.engine(), _sql), [], function (err, data) {
            if (err) { callback(err); }
            else if (data.rows.length === 0) {
              callback(nstore.lang().not_exist);
            }
            else { callback(null, { report: data.rows[0] }) }
          });
        }
      }
      else {
        callback(null, results);
      }
    },

    function (results, callback) {
      var _sql = { select: ["*"], from: "ui_reportsources", where: ["report_id", "=", results.report.id] }
      conn.query(models.getSql(nstore.engine(), _sql), [], function (err, data) {
        if (err) { callback(err); }
        else {
          results.sources = data.rows;
          callback(null, results);
        }
      });
    },

    function (results, callback) {
      var _sql = {
        select: ["rf.fieldname as fieldname", "ft.groupvalue as fieldtype", "rf.dataset as dataset",
          "wt.groupvalue as wheretype", "rf.sqlstr as sqlstr"],
        from: "ui_reportfields rf",
        inner_join: [["groups ft", "on", ["rf.fieldtype", "=", "ft.id"]],
        ["groups wt", "on", ["rf.wheretype", "=", "wt.id"]]],
        where: ["rf.report_id", "=", results.report.id]
      }
      conn.query(models.getSql(nstore.engine(), _sql), [], function (err, data) {
        if (err) { callback(err); }
        else {
          results.fields = {}
          data.rows.forEach(function (field) {
            results.fields[field.fieldname] = {
              fieldtype: field.fieldtype, wheretype: field.wheretype, dataset: field.dataset, sql: field.sqlstr
            }
          });
          callback(null, results);
        }
      });
    },

    function (results, callback) {
      results.datarows = {};
      var _sql = {
        select: ["*"], from: "ui_message",
        where: ["secname", "in", [[], "'" + results.report.reportkey + "_report'"]]
      }
      results.sources.forEach(function (rs) {
        _sql.where[2].push("'" + results.report.reportkey + "_" + rs.dataset + "'");
      });
      conn.query(models.getSql(nstore.engine(), _sql), [], function (err, data) {
        if (err) { callback(err); }
        else {
          results.datarows.labels = {};
          data.rows.forEach(function (label) {
            if (label.secname === results.report.reportkey + "_report") {
              results.datarows.labels[label.fieldname] = label.msg;
            }
            else {
              results.sources.forEach(function (rs) {
                if (label.secname === results.report.reportkey + "_" + rs.dataset) {
                  rs.sqlstr = rs.sqlstr.replace(new RegExp("={{" + label.fieldname + "}}", "g"),label.msg);
                }
              });
            }
          });
          callback(null, results);
        }
      });
    },

    function (results, callback) {
      results.where_str = {};
      for (var fieldname in params.filters) {
        if (!results.fields[fieldname]) {
          if (fieldname === "@id") {
            results.sources.forEach(function (rs) {
              rs.sqlstr = rs.sqlstr.replace(/@id/g, params.filters["@id"]);
            });
          }
          else {
            return callback(nstore.lang().invalid_fieldname + " " + fieldname);
          }
        }
        else {
          var rel = " = ";
          if (results.fields[fieldname].fieldtype === "date") {
            params.filters[fieldname] = "'" + moment(params.filters[fieldname],"YYYY-MM-DD").format("YYYY-MM-DD") + "'";
          }
          if (results.fields[fieldname].fieldtype === "string") {
            if (params.filters[fieldname].substr(1, 1) !== "'") {
              params.filters[fieldname] = "'" + params.filters[fieldname] + "'";
            }
            rel = " like ";
          }
          results.sources.forEach(function (rs) {
            if (results.fields[fieldname].dataset === rs.dataset || results.fields[fieldname].dataset === null) {
              if (results.fields[fieldname].wheretype === "where") {
                var wkey = results.fields[fieldname].dataset;
                if (wkey === null) { wkey = "nods" }
                var fstr;
                if (results.fields[fieldname].sql === null || results.fields[fieldname].sql === "") {
                  fstr = fieldname + rel + params.filters[fieldname];
                }
                else {
                  fstr = results.fields[fieldname].sql.replace("@" + fieldname, params.filters[fieldname]);
                }
                if (!results.where_str[wkey]) {
                  results.where_str[wkey] = " and " + fstr;
                }
                else {
                  results.where_str[wkey] = results.where_str[wkey] + " and " + fstr;
                }
              }
              else {
                if (results.fields[fieldname].sql === null || results.fields[fieldname].sql === "") {
                  rs.sqlstr = rs.sqlstr.replace(new RegExp("@" + fieldname, "g"), params.filters[fieldname]);
                }
                else {
                  fstr = results.fields[fieldname].sql.replace(new RegExp("@" + fieldname, "g"), params.filters[fieldname]);
                  rs.sqlstr = rs.sqlstr.replace(new RegExp("@" + fieldname, "g"), fstr);
                }
              }
            }
          });
        }
      }
      callback(null, results);
    },

    function (results, callback) {
      var trows = 0; var values_lst = [];
      results.sources.forEach(function (rs) {
        if (results.where_str[rs.dataset]) {
          rs.sqlstr = rs.sqlstr.replace("@where_str", results.where_str[rs.dataset]);
        }
        if (results.where_str["nods"]) {
          rs.sqlstr = rs.sqlstr.replace("@where_str", results.where_str.nods);
        }
        rs.sqlstr = rs.sqlstr.replace("@where_str", "");
        values_lst.push(function (callback_) {
          conn.query(rs.sqlstr, [], function (err, data) {
            if (!err) {
              results.datarows[rs.dataset] = data.rows;
              trows += data.rows.length;
            }
            callback_(err, data);
          });
        });
      });
      async.series(values_lst, function (err, data) {
        results.datarows.title = results.report.repname;
        results.datarows.crtime = moment().format("YYYY-MM-DDT00:00:00");
        if (!err && trows === 0) {
          callback("NODATA");
        }
        else if (!err && results.datarows.ds) {
          if (results.datarows.ds.length === 0) {
            callback("NODATA");
          }
          else {
            callback(err, results);
          }
        }
        else {
          callback(err, results);
        }
      });
    },

    function (results, callback) {
      if (params.output === "tmp") {
        callback(null, {
          filetype: results.report.reptype,
          template: results.report.report, data: results.datarows
        });
      }
      else {
        if (results.report.reptype === "xls") {
          var book = []
          var template = JSON.parse(results.report.report)
          Object.keys(results.datarows).forEach(skey => {
            if (Array.isArray(results.datarows[skey])) {
              if (results.datarows[skey].length > 0) {
                var sheet = { name: skey, data: [[]] }
                var columns = []
                if (template[skey]) {
                  if (template[skey]["sheetName"]) {
                    sheet.name = template[skey]["sheetName"]
                  }
                  if (template[skey]["columns"]) {
                    columns = template[skey]["columns"]
                  }
                } else {
                  results.datarows[skey][0].forEach(colname => {
                    columns.push({ "name": colname })
                  });
                }
                if (results.datarows["labels"]) {
                  columns.forEach(col => {
                    if (results.datarows["labels"][col["name"]]) {
                      sheet.data[0].push(results.datarows["labels"][col["name"]])
                    } else {
                      sheet.data[0].push(col["name"])
                    }
                  });
                }
                results.datarows[skey].forEach(row => {
                  var srow = []
                  columns.forEach(col => {
                    srow.push(row[col["name"]])
                  });
                  sheet.data.push(srow)
                });
                book.push(sheet)
              }
            }
          });
          var output = xlsx.build(book);
          if (params.output === "base64") {
            output = Buffer.from(output).toString('base64')
          }
          callback(null, { filetype: "xlsx", template: output, data: null });
        } else if (results.report.reptype === "ntr") {
          var rpt = new Report(params.orientation, "mm", params.size);
          rpt.loadDefinition(results.report.report);
          for (var i = 0; i < Object.keys(results.datarows).length; i++) {
            var pname = Object.keys(results.datarows)[i];
            rpt.setData(pname, results.datarows[pname]);
          }
          rpt.createReport();
          if (params.output === "xml") {
            callback(null, {
              filetype: results.report.reptype,
              template: rpt.save2Xml(), data: null
            });
          } else if (params.output === "base64") {
            rpt.save2DataUrlString(function (xml) {
              callback(null, { filetype: results.report.reptype, template: xml, data: null });
            })
          } else {
            rpt.save2Pdf(function (pdf) {
              callback(null, { filetype: results.report.reptype, template: new Buffer(pdf), data: null });
            })
          }
        } else {
          callback(nstore.lang().invalid_fieldname + " " + results.report.reptype);
        }
      }
    }
  ],
    function (err, results) {
      if (err) { if (err.message) { err = err.message; } }
      if (conn && !params.conn) {
        conn.close();
      }
      _callback(err, results);
    });
}

function send_email(nstore, params, _callback) {
  var conn = params.conn;
  if (!conn) { conn = nstore.connect.getConnect(); }
  async.waterfall([
    function (callback) {
      if (!conn) {
        callback(nstore.lang().not_connect);
      }
      else {
        if (params.email.attachments) {
          var values_lst = [];
          var output_format = "pdf"
          switch (params.provider) {
            case "mailjet":
            case "smtp":
              output_format = "base64";
              break;
            default:
              break;
          }
          params.email.attachments.forEach(function (item) {
            values_lst.push(function (callback_) {
              var repdata = {
                output: output_format, reportkey: item.reportkey,
                report_id: item.report_id, filters: { "@id": item.ref_id },
                nervatype: item.nervatype, refnumber: item.refnumber
              }
              get_report(nstore, repdata, function (err, result) {
                if (!err) {
                  result.filename = item.filename;
                }
                callback_(err, result);
              });
            });
          });
          async.series(values_lst, function (err, results) {
            callback(err, results);
          });
        }
        else {
          callback(null, []);
        }
      }
    },

    function (reports, callback) {
      switch (params.provider) {
        case "smtp":
          try {
            var smtp = {
              host: process.env.NT_SMTP_HOST || null,
              port: process.env.NT_SMTP_PORT || null,
              secure: process.env.NT_SMTP_SECURE || true, // use TLS
              auth: {
                user: process.env.NT_SMTP_USER || null,
                pass: process.env.NT_SMTP_PASSWORD || null
              }
            }
            var transporter = require('nodemailer').createTransport(smtp);
            var to_ = ""
            params.email.recipients.forEach(function (recipient) {
              to_ += ";" + recipient.email;
            });
            var mail = {
              from: params.email.from,
              to: to_.substr(1),
              subject: params.email.subject || "",
              text: params.email.text || "",
              html: params.email.html || "",
              attachments: []
            }
            for (var i = 0; i < reports.length; i++) {
              mail.attachments.push({
                contentType: "application/pdf",
                filename: reports[i].filename || "docs_" + i.toString() + ".pdf",
                content: reports[i].template, encoding: 'base64'
              });
            }
            transporter.sendMail(mail, function (err, info) {
              callback(err, info);
            });
          } catch (err) {
            callback(err.message);
          }
          break;

        case "mailjet":
          try {
            var publickey = params.publickey || process.env.NT_MJ_APIKEY_PUBLIC;
            var privatekey = params.privatekey || process.env.NT_MJ_APIKEY_PRIVATE;
            if (!publickey || !privatekey) {
              return callback(nstore.lang().missing_apikey);
            }

            var Mailjet = require('node-mailjet').connect(publickey, privatekey);
            var emailData = {
              FromEmail: params.email.from,
              FromName: params.email.name || "",
              Subject: params.email.subject || "",
              "Text-part": params.email.text || "",
              "Html-part": params.email.html || "",
              Recipients: params.email.recipients,
              Attachments: []
            }
            for (i = 0; i < reports.length; i++) {
              emailData.Attachments.push({
                "Content-Type": "application/pdf",
                Filename: reports[i].filename || "docs_" + i.toString() + ".pdf",
                Content: reports[i].template
              });
            }
            Mailjet.post('send').request(emailData, function (err, result) {
              if (result) { result = JSON.parse(result.response.text) }
              callback(err, result)
            });
          } catch (err) {
            callback(err.message);
          }
          break;

        default:
          callback(nstore.lang().invalid_provider);
          break;
      }
    }
  ],
    function (err, results) {
      if (err) { if (err.message) { err = err.message; } }
      if (conn && !params.conn) {
        conn.close();
      }
      _callback(err, results);
    });
}

function install_report(nstore, params, _callback) {
  var conn = params.conn; var trans;
  async.waterfall([
    function (callback) {
      if (!conn) { conn = nstore.connect.getConnect(); }
      if (!conn) {
        callback(nstore.lang().not_connect);
      }
      else {
        trans = connect.beginTransaction({ connection: conn, engine: nstore.engine() });
        var parser = new xml2js.Parser();
        if (params.filename) {
          fs.readFile(path.join(nstore.report_dir(), params.filename + ".xml"), function (err, data) {
            if (err) {
              callback(nstore.lang().invalid_template);
            }
            else {
              parser.parseString(data, function (err, result) {
                callback(err, result);
              });
            }
          });
        }
        else {
          callback(nstore.lang().invalid_template);
        }
      }
    },

    function (template, callback) {
      if (typeof template.report !== "undefined") {
        //check reportkey
        var _sql = {
          select: ["*"], from: "ui_report",
          where: ["reportkey", "=", "?"]
        }
        trans.query(models.getSql(nstore.engine(), _sql), [template.report["$"].reportkey],
          function (error, data) {
            if (error) { callback(error); }
            else {
              if (data.rowCount > 0) {
                callback(nstore.lang().exists_template)
              }
              else {
                callback(null, template);
              }
            }
          });
      }
      else {
        callback(nstore.lang().invalid_template);
      }
    },

    function (template, callback) {
      var _sql = {
        select: ["*"], from: "groups",
        where: ["groupname", "in", [[], "'nervatype'", "'transtype'",
          "'direction'", "'filetype'", "'fieldtype'", "'wheretype'"]]
      }
      trans.query(models.getSql(nstore.engine(), _sql), [],
        function (error, data) {
          if (error) { callback(error); }
          else {
            var groups = {};
            data.rows.forEach(function (group) {
              if (typeof groups[group.groupname] === "undefined") {
                groups[group.groupname] = {};
              }
              groups[group.groupname][group.groupvalue] = group.id;
            });
            callback(null, template, groups);
          }
        });
    },

    function (template, groups, callback) {
      //report template
      var report = template.report["$"];
      var values = { reportkey: report.reportkey, repname: report.repname, description: report.description }
      values.nervatype = groups.nervatype[report.nervatype];
      values.filetype = groups.filetype[report.filetype];
      if (typeof report.transtype !== "undefined") {
        values.transtype = groups.transtype[report.transtype];
      }
      if (typeof report.direction !== "undefined") {
        values.direction = groups.direction[report.direction];
      }
      if (typeof report.label !== "undefined") {
        values.label = report.label;
      }
      if (typeof template.report.template !== "undefined") {
        values.report = template.report.template[0];
      }
      nstore.connect.updateData(
        { nervatype: "ui_report", values: values, validate: false, insert_row: true, transaction: trans },
        function (err, id) {
          if (err) { callback(err); }
          else { callback(null, template.report, groups, id); }
        });
    },

    function (template, groups, report_id, callback) {
      //reportsources
      if (typeof template.dataset !== "undefined") {
        var ds_values = {};
        template.dataset.forEach(function (dataset) {
          if ((dataset["$"].engine === "") || (dataset["$"].engine === nstore.engine())) {
            if ((typeof ds_values[dataset["$"].name] === "undefined") || (dataset["$"].engine === nstore.engine())) {
              ds_values[dataset["$"].name] = {
                "report_id": report_id,
                "dataset": dataset["$"].name, "sqlstr": dataset["_"]
              }
            }
          }
        });
        if (Object.keys(ds_values).length > 0) {
          var ds_names = Object.keys(ds_values); var ds_lst = [];
          ds_names.forEach(function (dsname) {
            ds_lst.push(function (callback_) {
              nstore.connect.updateData(
                { nervatype: "ui_reportsources", values: ds_values[dsname], validate: false, insert_row: true, transaction: trans },
                function (err, id) {
                  callback_(err, id);
                });
            });
          });
          async.series(ds_lst, function (err, id) {
            callback(err, template, groups, report_id);
          });
        }
        else { callback(null, template, groups, report_id); }
      }
      else { callback(null, template, groups, report_id); }
    },

    function (template, groups, report_id, callback) {
      //reportfields
      if (typeof template.field !== "undefined") {
        var fields = [];
        template.field.forEach(function (field) {
          var values = { report_id: report_id }
          if (field["$"].fieldname) {
            values.fieldname = field["$"].fieldname;
          }
          if (field["$"].description) {
            values.description = field["$"].description;
          }
          if (field["$"].orderby) {
            values.orderby = field["$"].orderby;
          }
          if (field["$"].dataset) {
            values.dataset = field["$"].dataset;
          }
          if (field["$"].defvalue) {
            values.defvalue = field["$"].defvalue;
          }
          if (field["$"].valuelist) {
            values.valuelist = field["$"].valuelist;
          }
          if (field["$"].fieldtype) {
            values.fieldtype = groups.fieldtype[field["$"].fieldtype];
          }
          if (field["$"].wheretype) {
            values.wheretype = groups.wheretype[field["$"].wheretype];
          }
          if (field["_"] && field["_"] !== "") {
            values.sqlstr = field["_"];
          }
          fields.push(values);
        });
        if (fields.length > 0) {
          var fields_lst = [];
          fields.forEach(function (values) {
            fields_lst.push(function (callback_) {
              nstore.connect.updateData(
                { nervatype: "ui_reportfields", values: values, validate: false, insert_row: true, transaction: trans },
                function (err, id) {
                  callback_(err, id);
                });
            });
          });
          async.series(fields_lst, function (err, id) {
            callback(err, template, groups, report_id);
          });
        }
        else { callback(null, template, groups, report_id); }
      }
      else { callback(null, template, groups, report_id); }
    },

    function (template, groups, report_id, callback) {
      //message
      if (typeof template.message !== "undefined") {
        var messages = [];
        template.message.forEach(function (message) {
          messages.push({
            secname: template["$"].reportkey + "_" + message["$"].secname,
            fieldname: message["$"].fieldname, msg: message["_"]
          });
        });
        if (messages.length > 0) {
          var messages_lst = [];
          messages.forEach(function (values) {
            messages_lst.push(function (callback_) {
              nstore.connect.updateData(
                { nervatype: "ui_message", values: values, validate: false, insert_row: true, transaction: trans },
                function (err, id) {
                  callback_(err, id);
                });
            });
          });
          async.series(messages_lst, function (err, id) {
            callback(err, report_id, template["$"].reportkey);
          });
        }
        else { callback(null, report_id, template["$"].reportkey); }
      }
      else { callback(null, report_id, template["$"].reportkey); }
    }

  ],
    function (err, report_id, reportkey) {
      if (err) { if (err.message) { err = err.message; } }
      if (err && trans) {
        if (trans.rollback) {
          trans.rollback(function (error) {
            if (!error && !params.conn) {
              conn.close();
            }
          });
        }
        _callback(err);
      }
      else if (!err && trans) {
        if (trans.commit) {
          trans.commit(function (cerr) {
            if (!params.conn) {
              conn.close();
            }
            _callback(cerr || null, report_id, reportkey);
          });
        }
        else {
          if (!params.conn) {
            conn.close();
          }
          _callback(null, report_id, reportkey);
        }
      }
      else {
        _callback(err, report_id, reportkey);
      }
    });
}

module.exports = function() {
  return {
    callMenuCmd: function(nstore, params, _callback) {
      call_menu_cmd(params, _callback);},
    printQueue: function(nstore, params, _callback) {
      print_queue(nstore, params, _callback);},
    getPriceValue: function(nstore, params, _callback) {
      get_price_value(nstore, params, _callback);},
    nextNumber: function(nstore, params, _callback) {
      nstore.connect.nextNumber(params, _callback);},
    getReport: function(nstore, params, _callback) {
      get_report(nstore, params, _callback);},
    sendEmail: function(nstore, params, _callback) {
      send_email(nstore, params, _callback);},
    getReportFiles: function(nstore, params, _callback) {
      get_report_files(nstore, params, _callback);},
    installReport: function(nstore, params, _callback) {
      install_report(nstore, params, _callback);}
};}