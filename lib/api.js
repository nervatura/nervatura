/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2019, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var async = require("async");
var jwt = require('jsonwebtoken');
var argon2 = require('argon2');
var crypto = require('crypto');
var moment = require('moment');

var lang = require('./lang.js')[process.env.NT_LANG || "en"];
var ntura = require('./models.js');
var connect = require('./adapter.js').connect();
var models = require('./adapter.js').models();
var service = require('./service.js')();
var nervastore = require('./nervastore')

module.exports = function (nstore) {

  if(!nstore){
    nstore = nervastore({})
  }

  function api_view(params, _callback) {
    var conn;
    var trans;
    var results = {};
    async.waterfall(
      [
        function (callback) {
          if (!params || !Array.isArray(params)) {
            callback(lang.missing_params + ": queries");
          } else {
            conn = nstore.connect.getConnect();
            if (!conn) {
              callback(lang.not_connect);
            } else {
              callback(null);
            }
          }
        },

        function (callback) {
          var trans = connect.beginTransaction({
            connection: conn,
            engine: nstore.engine()
          });
          var results_lst = [];
          params.forEach(function (query) {
            results_lst.push(function (callback_) {
              var qname = query.key || "result";
              var sql = query.text || "";
              var params = query.values || [];
              trans.query(sql, params, function (err, data) {
                if (err) {
                  callback_(err);
                } else {
                  results[qname] = data ? data.rows : [];
                  callback_(null);
                }
              });
            });
          });
          if (results_lst.length > 0) {
            async.series(results_lst, function (err) {
              callback(err);
            });
          } else {
            callback(null);
          }
        }
      ],
      function (err) {
        if (trans) {
          trans.rollback();
        }
        if (conn) {
          conn.close();
        }
        _callback(err, results);
      }
    );
  }

  function api_post(params, _callback) {
    var conn;
    async.waterfall(
      [
        function (callback) {
          if (!params.nervatype) {
            callback(lang.missing_params + ": nervatype");
          } else if (!params.data) {
            callback(lang.missing_params + ": data");
          } else {
            if (!Array.isArray(params.data)) {
              params.data = [params.data];
            }
            callback(null);
          }
        },

        function (callback) {
          conn = nstore.connect.getConnect();
          if (!conn) {
            callback(lang.not_connect);
          } else {
            callback(null, { conn: conn });
          }
        },

        function (validator, callback) {
          if(params.nervatype === "trans"){
            var data_lst = [];
            params.data.forEach(function (record) {
              if(!(record.keys && record.keys.transtype && record.keys.customer_id)){
                var transtype_id = record.transtype || "null"
                var transtype_key = (record.keys) ? (record.keys.transtype) ? "'"+record.keys.transtype+"'" : "null" : "null"
                var customer_id = record.customer_id || "null"
                var custnumber = (record.keys) ? (record.keys.customer_id) ? "'"+record.keys.customer_id+"'" : "null" : "null"
                var trans_id = record.id || null
                if((transtype_key === "null") || ((transtype_key === "invoice") && (custnumber === "null"))){
                  data_lst.push(function (callback_) {
                    var _sql = [
                      { select: ["'trans' as rtype", "tt.groupvalue as transtype", "c.custnumber"],
                        from: "trans t",
                        inner_join: ["groups tt", "on", ["t.transtype", "=", "tt.id"]],
                        left_join: ["customer c", "on", ["t.customer_id", "=", "c.id"]],
                        where: ["t.id", "=", trans_id]},
                      { union_select: ["'groups' as rtype", "groupvalue as transtype", "null"], 
                        from: "groups",
                        where: [[[["groupname", "=", "'transtype'"],
                          ["and", "groupvalue", "=", transtype_key]]], ["or", ["id", "=", transtype_id]]]},
                      { union_select: ["'customer' as rtype", "null", "custnumber"],
                          from: "customer",
                          where: [[["id","=",customer_id]], ["or",["custnumber","=",custnumber]]]}
                    ]
                    conn.query(models.getSql(nstore.engine(), _sql), [], function ( err, data ) {
                      if(!err){
                        if(!record.keys){
                          record.keys = {}
                        }
                        var keys = {}
                        data.rows.forEach(row => {
                          keys[row.rtype] = [row.transtype, row.custnumber]
                        });
                        if(keys.groups){
                          if(!record.keys.transtype){
                            record.keys.transtype = keys.groups[0]
                          }
                        } else if(keys.trans){
                          if(!record.keys.transtype){
                            record.keys.transtype = keys.trans[0]
                          }
                        }
                        if(keys.customer){
                          if(!record.keys.customer_id){
                            record.keys.customer_id = keys.customer[1]
                          }
                        } else if(keys.trans && keys.trans[1]){
                          if(!record.keys.customer_id){
                            record.keys.customer_id = keys.trans[1]
                          }
                        }
                      }
                      callback_(err);
                    });
                  })
                }
              }
            })
            if (data_lst.length > 0) {
              async.series(data_lst, function (err) {
                callback(err, validator);
              });
            } else {
              callback(null, validator);
            }
          } else {
            callback(null, validator)
          }
        },

        function (validator, callback) {
          var data_lst = [];
          function getInfo(key, value, keys){
            var info = { fieldname: key, reftype: "id" };
            switch (key) {
              case "id":
                info["nervatype"] = params.nervatype;
                info["refnumber"] = value;
                break;

              case "ref_id":
              case "ref_id_1":
              case "ref_id_2":
                info["nervatype"] = value.split("/")[0];
                info["refnumber"] = value.replace(
                  value.split("/")[0] + "/",
                  ""
                );
                break;

              default:
                if (ntura.model[params.nervatype][key]) {
                  if (
                    ntura.model[params.nervatype][key].hasOwnProperty(
                      "references"
                    )
                  ) {
                    info["nervatype"] =
                      ntura.model[params.nervatype][key]["references"][0];
                    if (info["nervatype"] === "groups") {
                      switch (key) {
                        case "nervatype_1":
                        case "nervatype_2":
                          info["refnumber"] = "nervatype~" + value;
                          break;
                      
                        default:
                          info["refnumber"] = key + "~" + value;
                          break;
                      }
                    } else {
                      info["refnumber"] = value;
                      if((key === "customer_id") && (keys.transtype === "invoice")){
                        info["extra_info"] = true
                      }
                    }
                  } else if (
                    ntura.model[params.nervatype]["_key"][0] === key
                  ) {
                    if ((value === "numberdef") || (value[0] === "numberdef")) {
                      info["reftype"] = "numberdef";
                      if(Array.isArray(value) && value.length>1){
                        info["numberkey"] = value[1];
                      } else {
                        info["numberkey"] = key;
                      }
                      info["step"] = true;
                      info["insert_key"] = false;
                    } else {
                      info["nervatype"] = params.nervatype;
                      info["refnumber"] = value;
                      info["fieldname"] = "id";
                    }
                  }
                }
                if (
                  info["reftype"] === "id" &&
                  !info["nervatype"]
                ) {
                  info["nervatype"] = "invalid";
                  info["refnumber"] = value;
                }
                break;
            }
            return info
          }
          params.data.forEach(function (record) {
            if (record.keys) {
              for (const key in record.keys) {
                if (record.keys.hasOwnProperty(key)) {
                  data_lst.push(function (callback_) {
                    var info = getInfo(key, record.keys[key], record.keys);
                    if (info["reftype"] === "numberdef") {
                      nstore.connect.nextNumber(info, function (err, retnumber) {
                        if (!err) {
                          record[info["fieldname"]] = retnumber;
                        }
                        callback_(err, retnumber);
                      });
                    } else {
                      nstore.valid.getIdFromRefnumber(info, function (err, id, rfinfo) {
                        if (!err) {
                          record[info["fieldname"]] = id;
                          if(info["extra_info"]){
                            record.trans_custinvoice_compname = rfinfo.compname;
                            record.trans_custinvoice_comptax = rfinfo.comptax;
                            record.trans_custinvoice_compaddress = rfinfo.compaddress; 
                            record.trans_custinvoice_custname= rfinfo.custname;
                            record.trans_custinvoice_custtax = rfinfo.custtax;
                            record.trans_custinvoice_custaddress = rfinfo.custaddress;
                          }
                        }
                        callback_(err, id);
                      });
                    }
                  });
                }
              }
            }
          });
          if (data_lst.length > 0) {
            async.series(data_lst, function (err) {
              callback(err, validator);
            });
          } else {
            callback(null, validator);
          }
        },

        function (validator, callback) {
          var err
          params.data.forEach(function (record) {
            if (record.keys) {
              delete record.keys;
            }
            var model = ntura.model[params.nervatype]
            if(!record.id){
              for (const key in model) {
                switch (key) {
                  case "_access":
                  case "_key":
                  case "id":
                    break;
                  
                  case "crdate":
                    if(model.crdate.type === "datetime"){
                      record.crdate = new Date().toISOString()
                    } else if((model.crdate.type === "date") && !record.crdate){
                      record.crdate = new Date().toISOString().substr(0,10)
                    }
                    break;

                  case "cruser_id":
                    if(nstore.employee()){
                      record.cruser_id = nstore.employee().id
                    } else {
                      record.cruser_id = 1
                    }
                
                  default:
                    if(model[key].notnull && (typeof(model[key].default) === "undefined") && !record[key]){
                      err = lang.missing_required_field+" "+key
                    }
                    break;
                }
              }
              if((params.nervatype === "trans") && !record.trans_transcast){
                record.trans_transcast = "normal"
              }
            } else if(record.id && record.crdate && model.crdate){
              if(model.crdate.type === "datetime"){
                delete record.crdate
              }
            }
          })
          callback(err, validator)
        },

        function (validator, callback) {
          var results = [];
          var trans = connect.beginTransaction({
            connection: validator.conn,
            engine: nstore.engine()
          });
          var results_lst = [];
          params.data.forEach(function (record) {
            results_lst.push(function (callback_) {
              var update_params = {
                nervatype: params.nervatype,
                values: record,
                validate: true,
                insert_row: true,
                insert_field: true,
                transaction: trans,
                validator: validator
              };
              nstore.connect.updateData(update_params, function (
                err,
                record_id
              ) {
                if (err) {
                  callback_(err);
                } else {
                  results.push(record_id);
                  callback_(null, record_id);
                }
              });
            });
          });
          if(results_lst.length > 0){
            async.series(results_lst, function (err) {
              if (err) {
                if (trans.rollback) {
                  trans.rollback();
                }
                callback(err, results);
              } else if (results.length > 0) {
                if (trans.commit) {
                  trans.commit(function (cerr) {
                    callback(cerr || null, results);
                  });
                } else {
                  callback(err, results);
                }
              } else {
                callback(err, results);
              }
            });
          } else {
            callback(null, [])
          }
        }
      ],
      function (err, results) {
        if (conn) {
          conn.close();
        }
        _callback(err, results);
      }
    );
  }

  function api_get(params, _callback) {
    var conn;
    async.waterfall(
      [
        function (callback) {
          if (!params.nervatype) {
            callback(lang.missing_params + ": nervatype");
          } else if (!ntura.model.hasOwnProperty(params.nervatype)) {
            callback(lang.invalid_nervatype + params.nervatype);
          } else {
            params.sql = "select * from " + params.nervatype + " where 1=1 ";
            if (ntura.model[params.nervatype].hasOwnProperty("deleted")) {
              params.sql += " and deleted=0 ";
            }
            callback(null);
          }
        },

        function (callback) {
          if (params.ids) {
            params.sql += " and id in(" + params.ids + ")";
            callback(null);
          } else if (params.filter) {
            var filters = params.filter.split("|");
            params.where = [];
            var err = null;
            filters.forEach((filter, index) => {
              if (!err) {
                var fields = filter.split(";");
                var field = {};
                if (fields.length !== 3) {
                  err = lang.invalid_value + "- filter: " + filter;
                  return;
                }
                if (!ntura.model[params.nervatype].hasOwnProperty(fields[0])) {
                  err = lang.invalid_value + "- fieldname: " + fields[0];
                  return;
                }
                field.fieldname = fields[0];
                if (
                  !["==", "!=", "<", "<=", ">", ">=", "in"].includes(fields[1])
                ) {
                  err = lang.invalid_value + "- comparison: " + fields[1];
                  return;
                }
                field.comparison = fields[1];
                field.value = fields[2];
                field.fieldtype = ntura.model[params.nervatype][fields[0]].type;
                if (["string", "text", "date"].includes(field.fieldtype)) {
                  field.value = field.value.replace(/'/g, "");
                  var values = field.value.split(",");
                  field.value = "";
                  values.forEach(value => {
                    field.value += ",'" + value + "'";
                  });
                  field.value = field.value.substring(1);
                }
                params.where.push(field);
              }
            });
            callback(err);
          } else {
            callback(lang.missing_params + ": filter or path IDs");
          }
        },

        function (callback) {
          if (params.where) {
            params.where.forEach((filter, index) => {
              params.sql += " and ";
              params.sql += filter.fieldname;
              switch (filter.comparison) {
                case "==":
                  params.sql += "=" + filter.value;
                  break;
                case "!=":
                case "<":
                case "<=":
                case ">":
                case ">=":
                case "is":
                  params.sql += " " + filter.comparison + " " + filter.value;
                  break;
                case "in":
                  params.sql += " in(" + filter.value + ")";
                  break;
                default:
                  break;
              }
            });
          }
          callback(null);
        },

        function (callback) {
          conn = nstore.connect.getConnect();
          if (!conn) {
            callback(lang.not_connect);
          } else {
            conn.query(params.sql, [], function (err, data) {
              callback(err, data ? data.rows : []);
            });
          }
        },

        function (result, callback) {
          if (params.metadata === "true" && result.length > 0) {
            if ( ["address", "barcode", "contact", "currency", "customer", "employee", "event", "groups",
                  "item", "link", "log", "movement", "price", "place", "product", "project", "rate",
                  "tax", "tool", "trans" ].includes(params.nervatype)
            ) {
              var ids = [];
              result.forEach(row => {
                ids.push(row.id);
              });
              var _sql = {
                select: ["fv.*", "ft.groupvalue as fieldtype"],
                from: "fieldvalue fv",
                inner_join: [
                  ["deffield df", "on", ["fv.fieldname", "=", "df.fieldname"]],
                  ["groups nt", "on", ["df.nervatype", "=", "nt.id"]],
                  ["groups ft", "on", ["df.fieldtype", "=", "ft.id"]]
                ],
                where: [
                  ["fv.deleted", "=", "0"],
                  ["and", ["df.deleted", "=", "0"]],
                  ["and", ["nt.groupvalue", "=", "'" + params.nervatype + "'"]],
                  ["and", ["fv.ref_id", "in", [[], ids.join(",")]]]
                ],
                order_by: ["fv.fieldname", "fv.id"]
              };
              conn.query(models.getSql(nstore.engine(), _sql), [], function (
                err,
                data
              ) {
                callback(err, result, data ? data.rows : []);
              });
            } else {
              callback(null, result, []);
            }
          } else {
            callback(null, result, []);
          }
        },

        function (result, metadata, callback) {
          if (metadata.length > 0) {
            result.forEach(row => {
              row.metadata = [];
              var mdata = metadata.filter(function (meta) {
                return meta.ref_id === row.id;
              });
              mdata.forEach(data => {
                row.metadata.push({
                  id: data.id,
                  fieldname: data.fieldname,
                  fieldtype: data.fieldtype,
                  value: data.value,
                  notes: data.notes
                });
              });
            });
            callback(null, result);
          } else {
            callback(null, result);
          }
        }
      ],
      function (err, result) {
        if (conn) {
          conn.close();
        }
        _callback(err, result);
      }
    );
  }

  function api_delete(params, _callback) {
    async.waterfall(
      [
        function (callback) {
          if (!params.nervatype) {
            callback(lang.missing_params + ": nervatype");
          } else if (!params.id && !params.key) {
            callback(lang.missing_params + ": id or key");
          } else {
            callback(null);
          }
        },

        function (callback) {
          nstore.connect.deleteData(
            {
              nervatype: params.nervatype,
              ref_id: params.id,
              refnumber: params.key
            },
            function (err, id) {
              _callback(err, id);
            }
          );
        }
      ],
      function (err, id) {
        _callback(err, id);
      }
    );
  }

  function report(params, _callback) {
    params.conn = nstore.connect.getConnect();
    if (!params.conn) {
      _callback(lang.not_connect);
    } else {
      if (params.output === "data") {
        params.output = "tmp";
      }
      service.getReport(nstore, params, function (err, data) {
        params.conn.close();
        _callback(err, data);
      });
    }
  }

  function report_list(params, _callback) {
    var conn;
    async.waterfall(
      [
        function (callback) {
          conn = nstore.connect.getConnect();
          if (!conn) {
            callback(lang.not_connect);
          } else {
            var _sql = {
              select: ["id", "reportkey"],
              from: "ui_report"
            };
            conn.query(models.getSql(nstore.engine(), _sql), [], function (
              err,
              data
            ) {
              callback(err, data ? data.rows : []);
            });
          }
        },

        function (dbs_reports, callback) {
          params.dbs_reports = {};
          dbs_reports.forEach(report => {
            params.dbs_reports[report.reportkey] = report.id;
          });
          service.getReportFiles(nstore, params, function (err, files) {
            _callback(err, files);
          });
        }
      ],
      function (err, files) {
        if (conn) {
          conn.close();
        }
        _callback(err, files);
      }
    );
  }

  function report_install(params, _callback) {
    if (!params.reportkey) {
      _callback(lang.missing_params + ": reportkey");
    } else {
      service.installReport(nstore, { filename: params.reportkey }, function (
        err,
        report_id,
        reportkey
      ) {
        _callback(err, report_id);
      });
    }
  }

  function report_delete(params, _callback) {
    var conn;
    var trans;
    async.waterfall(
      [
        function (callback) {
          if (!params.reportkey) {
            callback(lang.missing_params + ": reportkey");
          } else {
            conn = nstore.connect.getConnect();
            if (!conn) {
              callback(lang.not_connect);
            } else {
              var _sql = {
                select: ["*"],
                from: "ui_report",
                where: ["reportkey", "=", "'"+params.reportkey+"'"]
              };
              conn.query(
                models.getSql(nstore.engine(), _sql),[],
                function (err, data) {
                  if (err) {
                    callback(err);
                  } else {
                    if (data.rowCount > 0) {
                      callback(null, data.rows[0].id);
                    } else {
                      callback(lang.missing_reportkey);
                    }
                  }
                }
              );
            }
          }
        },

        function (report_id, callback) {
          trans = connect.beginTransaction({
            connection: conn,
            engine: nstore.engine()
          });
          var _sql = {
            delete: "",
            from: "ui_reportfields",
            where: ["report_id", "=", report_id]
          };
          trans.query(models.getSql(nstore.engine(), _sql), [], function (
            err,
            data
          ) {
            if (err) {
              callback(err);
            } else {
              callback(null, report_id);
            }
          });
        },

        function (report_id, callback) {
          var _sql = {
            delete: "",
            from: "ui_reportsources",
            where: ["report_id", "=", report_id]
          };
          trans.query(models.getSql(nstore.engine(), _sql), [], function (
            err,
            data
          ) {
            if (err) {
              callback(err);
            } else {
              callback(null, report_id);
            }
          });
        },

        function (report_id, callback) {
          var _sql = {
            delete: "",
            from: "ui_message",
            where: ["secname", "like", "'"+params.reportkey + "%"+"'"]
          };
          trans.query(
            models.getSql(nstore.engine(), _sql),[],
            function (err, data) {
              if (err) {
                callback(err);
              } else {
                callback(null, report_id);
              }
            }
          );
        },

        function (report_id, callback) {
          var _sql = {
            delete: "",
            from: "ui_report",
            where: ["id", "=", report_id]
          };
          trans.query(models.getSql(nstore.engine(), _sql), [], function (
            err,
            data
          ) {
            if (err) {
              callback(err);
            } else {
              callback(null);
            }
          });
        }
      ],
      function (err) {
        if (err) {
          if (err.message) {
            err = err.message;
          }
        }
        if (!err && trans) {
          if (trans.commit) {
            trans.commit(function (cerr) {
              conn.close();
              _callback(cerr || null);
            });
          } else {
            conn.close();
            _callback(null);
          }
        } else {
          if (trans) {
            if (trans.rollback) {
              trans.rollback();
            }
          }
          if (conn) {
            conn.close();
          }
          _callback(err);
        }
      }
    );
  }

  function create_conn(database, _callback) {
    var alias = "NT_ALIAS_" + database.toUpperCase()
    if (process.env[alias]) {
      connect.createConnection({ connect: process.env[alias], pool: false }, function (err, conn) {
        if (err) {
          _callback(err);
        }
        else {
          conn.end();
          nstore.database(database)
          nstore.engine(process.env[alias].split("://")[0].replace("3",""))
          nstore.dbs_config(process.env[alias]);
          _callback(null);
        }
      });
    } else {
      _callback(lang.missing_database)
    }
  }

  function database_create(params, _callback) {
    var conn;
    var trans;
    var results = [];
    async.waterfall(
      [
        function (callback) {
          create_conn(params.database, function(err){
            callback(err)
          })
        },

        function (callback) {
          conn = nstore.connect.getConnect();
          if (!conn) {
            callback(lang.not_connect);
          } else {
            results.push({
              type: "create", start_process: true,
              database: params.database,
              stamp: moment().toISOString(),
              state: lang.log_start_process
            });
            callback(null);
          }
        },

        function (callback) {
          //drop all tables if exist
          trans = connect.beginTransaction({
            connection: conn,
            engine: nstore.engine()
          });
          results.push({
            type: "create",
            stamp: moment().toISOString(),
            state: lang.log_drop_table
          });
          var drop_lst = models.dropList(nstore.engine());
          var value_lst = [];
          drop_lst.forEach(function (sql) {
            value_lst.push(function (callback_) {
              trans.query(sql, [], function (err, data) {
                callback_(err, data);
              });
            });
          });
          async.series(value_lst, function (err, data) {
            if (!err && trans.commit) {
              trans.commit(function (cerr) {
                callback(cerr || null);
              });
            } else {
              if (trans.rollback) {
                trans.rollback();
              }
              callback(null);
            }
          });
        },

        function (callback) {
          //create all tables
          results.push({
            type: "create",
            stamp: moment().toISOString(),
            state: lang.log_create_table
          });
          trans = connect.beginTransaction({
            connection: conn,
            engine: nstore.engine()
          });
          var create_lst = models.modelList(nstore.engine());
          var value_lst = [];
          create_lst.forEach(function (sql) {
            value_lst.push(function (callback_) {
              trans.query(sql, [], function (err, data) {
                callback_(err, data);
              });
            });
          });
          async.series(value_lst, function (err, data) {
            callback(err);
          });
        },

        function (callback) {
          //create indexes
          results.push({
            type: "create",
            stamp: moment().toISOString(),
            state: lang.log_create_index
          });
          var index_lst = models.indexList(nstore.engine());
          var value_lst = [];
          index_lst.forEach(function (sql) {
            value_lst.push(function (callback_) {
              trans.query(sql, [], function (err, data) {
                callback_(err, data);
              });
            });
          });
          async.series(value_lst, function (err, data) {
            callback(err);
          });
        },

        function (callback) {
          if (!params.empty) {
            //insert data
            results.push({
              type: "create",
              stamp: moment().toISOString(),
              state: lang.log_init_data
            });
            var data_lst = models.dataList(nstore.engine());
            var value_lst = [];
            data_lst.forEach(function (sql) {
              value_lst.push(function (callback_) {
                trans.query(sql, [], function (err, data) {
                  callback_(err, data);
                });
              });
            });
            async.series(value_lst, function (err, data) {
              callback(err);
            });
          } else {
            callback(null);
          }
        },

        function (callback) {
          if (trans.commit) {
            trans.commit();
          }
          if (models.compact(nstore.engine()) !== null) {
            conn.query(models.compact(nstore.engine()), [], function (
              err
            ) {
              if (!err) {
                results.push({
                  type: "create",
                  stamp: moment().toISOString(),
                  state: "Rebuilding the database..."
                });
              }
              callback(err);
            });
          } else {
            callback(null);
          }
        }
      ],
      function (err) {
        if (err) {
          if (err.message) {
            err = err.message;
          }
        }
        if (!err) {
          results.push({
            type: "create",
            stamp: moment().toISOString(),
            state: lang.info_create_ok
          });
        } else {
          if (trans) {
            if (trans.rollback) {
              trans.rollback();
            }
          }
          results.push({
            type: "create",
            stamp: moment().toISOString(),
            error: err
          });
        }
        if (conn && !params.conn) {
          conn.close();
        }
        if(params.demo === "true"){
          database_demo({ results: results}, function(err, demo_results){
            if(!err){
              results.push({
                type: "create",end_process: true,
                stamp: moment().toISOString(),
                state: lang.log_end_process
              });
            }
            _callback(err, demo_results);
          })
        } else {
          results.push({
            type: "create", end_process: true,
            stamp: moment().toISOString(),
            state: lang.log_end_process
          });
          _callback(err, results);
        }
      }
    );
  }

  function database_demo(params, _callback) {
    var results = params.results || [];
    var data = JSON.parse(JSON.stringify(require('./demo.js')));
    async.waterfall(
      [
        function (callback) {
          results.push({
            type: "demo", start_process: true,
            stamp: moment().toISOString(),
            state: lang.log_start_process
          });
          //create 3 departments
          var post_params = { nervatype: "groups", data: data.groups };
          api_post(post_params, function (err, result) {
            if (!err) {
              results.push({
                type: "demo", start_group: true, end_group: true,
                datatype: post_params.nervatype, 
                result: result
              });
            }
            callback(err);
          });
        },

        function (callback) {
          //customer
          //-> def. 4 customer additional data (float,date,valuelist,customer types), 
          //-> create 3 customers, 
          //-> and more create and link to contacts, addresses and events
          async.series([
            function (callback_) {
              var post_params = { nervatype: "deffield", data: data.customer.deffield };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true, 
                    section: "customer", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "customer", data: data.customer.customer };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "address", data: data.customer.address };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "contact", data: data.customer.contact };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "event", data: data.customer.event };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //employee
          //-> def. 1 employee additional data (integer type), 
          //->create 1 employee, 
          //->and more create and link to contact, address and event
          async.series([
            function (callback_) {
              var post_params = { nervatype: "deffield", data: data.employee.deffield };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true, 
                    section: "employee", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "employee", data: data.employee.employee };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "address", data: data.employee.address };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "contact", data: data.employee.contact };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "event", data: data.employee.event };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //product
          //-> def. 3 product additional data (product,integer and valulist types),
          //->create 13 products,
          //->and more create and link to barcodes, events, prices, additional data
          async.series([
            function (callback_) {
              var post_params = { nervatype: "deffield", data: data.product.deffield };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true, 
                    section: "product", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "product", data: data.product.product };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "barcode", data: data.product.barcode };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "price", data: data.product.price };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "event", data: data.product.event };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //project
          //-> def. 2 project additional data, 
          //->create 1 project, 
          //->and more create and link to contact, address and event
          async.series([
            function (callback_) {
              var post_params = { nervatype: "deffield", data: data.project.deffield };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true, 
                    section: "project", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "project", data: data.project.project };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "address", data: data.project.address };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "contact", data: data.project.contact };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "event", data: data.project.event };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //tool
          //-> def. 2 tool additional data,
          //->create 3 tools,
          //->and more create and link to event and additional data
          async.series([
            function (callback_) {
              var post_params = { nervatype: "deffield", data: data.tool.deffield };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true, 
                    section: "tool", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "tool", data: data.tool.tool };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo",  
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "event", data: data.tool.event };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //create +1 warehouse
          var post_params = { nervatype: "place", data: data.place };
          api_post(post_params, function (err, result) {
            if (!err) {
              results.push({
                type: "demo", start_group: true, end_group: true,
                datatype: post_params.nervatype,
                result: result
              });
            }
            callback(err);
          });
        },

        function (callback) {
          //documents
          //offer, order, invoice, worksheet, rent
          async.series([
            function (callback_) {
              var post_params = { nervatype: "trans", data: data.trans_item.trans };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true,
                    section: "document(offer,order,invoice,rent,worksheet)",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "item", data: data.trans_item.item };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "link", data: data.trans_item.link };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //payments
          //bank and petty cash
          async.series([
            function (callback_) {
              var post_params = { nervatype: "trans", data: data.trans_payment.trans };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true,
                    section: "payment(bank,petty cash)",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "payment", data: data.trans_payment.payment };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "link", data: data.trans_payment.link };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //stock control
          //tool movement (for employee)
          //create delivery,stock transfer,correction
          //formula and production
          async.series([
            function (callback_) {
              var post_params = { nervatype: "trans", data: data.trans_movement.trans };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true,
                    section: "stock control(tool movement,delivery,stock transfer,correction,formula,production)",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "movement", data: data.trans_movement.movement };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", 
                    datatype: post_params.nervatype,
                    result: result
                  });
                  results.push({ datatype: params.datatype, result: result.data });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "link", data: data.trans_movement.link };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //sample menus and menufields
          async.series([
            function (callback_) {
              var post_params = { nervatype: "ui_menu", data: data.menu.ui_menu };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", start_group: true,
                    section: "sample menus",
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            },

            function (callback_) {
              var post_params = { nervatype: "ui_menufields", data: data.menu.ui_menufields };
              api_post(post_params, function (err, result) {
                if (!err) {
                  results.push({
                    type: "demo", end_group: true,
                    datatype: post_params.nervatype,
                    result: result
                  });
                }
                callback_(err);
              });
            }
          ], function (err) {
            callback(err);
          });
        },

        function (callback) {
          //load general reports and other templates
          report_list({}, function(err, reports){
            callback(err, reports);
          })
        },

        function (reports, callback) {
          var report_lst = []
          reports.forEach(function (report) {
            report_lst.push(function (callback_) {
              var rparams = { reportkey: report.reportkey} ;
              report_install(rparams, function(err, report_id){
                callback_(err, rparams.reportkey);
              })
            });
          });
          if (report_lst.length > 0) {
            async.series(report_lst, function (err, reports) {
              if (!err) {
                results.push({
                  type: "demo", start_group: true, end_group: true,
                  section: "report templates",
                  result: reports
                });
              }
              callback(err);
            });
          } else {
            callback(null);
          }
        }
      ],
      function (err) {
        if (err) {
          results.push({
            type: "demo",
            stamp: moment().toISOString(),
            error: err
          });
        }
        else {
          results.push({
            type: "demo", end_process: true,
            stamp: moment().toISOString(),
            state: lang.log_end_process
          });
        }
        _callback(err, results);
      }
    );
  }

  function api_function(params, _callback) {
    if (!params.key) {
      _callback(lang.missing_params + ": key");
    } else if (typeof service[params.key] === "undefined") {
      _callback(lang.unknown_method + " " + params.key);
    } else {
      params.values.conn = nstore.connect.getConnect();
      if (!params.values.conn) {
        _callback(lang.not_connect);
      } else {
        service[params.key](nstore, params.values, function (err, data) {
          _callback(err, data);
        });
      }
    }
  }

  function auth_token_login(params, _callback) {
    if (params.token) {
      var key = params.key || process.env.NT_TOKEN_KEY || ""
      var options = params.options || {}
      jwt.verify(params.token, key, options, function (err, token) {
        if (!err && token) {
          if (!token.database) {
            if (process.env.NT_DEFAULT_ALIAS) {
              params.database = process.env.NT_DEFAULT_ALIAS
            } else {
              err = lang.missing_database
            }
          } else {
            params.database = token.database
          }
          params.username = token.username || token.custnumber || token.email || token.phone_number;
          if (!params.username) {
            err = lang.missing_user;
          }
        }
        if (!err) {
          auth_user(params, function (err, result) {
            _callback(err, result)
          })
        } else {
          _callback(err, null);
        }
      })
    }
    else {
      _callback(lang.missing_token, null);
    }
  }

  function check_hashtable(_callback) {
    var conn
    var hashtable = process.env.NT_HASHTABLE || "ref17890714"
    async.waterfall([
      function (callback) {
        conn = nstore.connect.getConnect();
        if (!conn) {
          callback(lang.not_connect);
        } else {
          var sql = ""
          if((nstore.engine() === "sqlite3") || (nstore.engine() === "sqlite")){
            sql = "select * from sqlite_master WHERE name='"+hashtable+"'"
          } else {
            sql = "select * from information_schema.tables where table_name='"+hashtable+"'"
          }
          conn.query(sql, [], function (err, results) {
            if(!err){
              callback(null, (results.rowCount>0))
            } else{
              callback(err, false)
            }
          })
        }
      },

      function (valid, callback) {
        if(!valid){
          var text = models.getDataType(nstore.engine(), "text")
          var sql = "CREATE TABLE " + hashtable + " ( refname " + text + ", value " + text + ");"
          conn.query(sql, [], function (err, results) {
            callback(err, valid)
          })
        } else {
          callback(null, valid)
        }
      },

      function (valid, callback) {
        if(!valid){
          var sql = "CREATE UNIQUE INDEX " + hashtable + "_refname_idx ON " + hashtable + " (refname);"
          conn.query(sql, [], function (err, results) {
            callback(err, (err===null))
          })
        } else {
          callback(null, valid)
        }
      }
    ],
      function (err, valid) {
        if(err){
          if (err.message) { err = err.message; }
        }
        if (conn) {
          conn.close()
        }
        _callback(err, valid);
      });
  };

  function get_hashvalue(ustore, refname, _callback) {
    var conn
    async.waterfall([
      function (callback) {
        check_hashtable(function(err, valid){
          callback(err)
        })
      },

      function (callback) {
        conn = ustore.connect.getConnect();
        if (!conn) {
          callback(lang.not_connect);
        } else {
          var hashtable = process.env.NT_HASHTABLE || "ref17890714"
          var _sql = {
            select: "*", from: hashtable,
            where: ["refname", "=", "'" + refname + "'"]
          }
          conn.query(models.getSql(ustore.engine, _sql), [], function (err, results) {
            var value = null
            if (!err) {
              if (results.rowCount > 0) {
                value = results.rows[0].value
              }
            }
            callback(err, value)
          })
        }
      }
    ],
      function (err, value) {
        if(err){
          if (err.message) { err = err.message; }
        }
        if (conn) {
          conn.close()
        }
        _callback(err, value);
      });
  };

  function set_hashvalue(refname, value, _callback) {
    var conn
    async.waterfall([
      function (callback) {
        get_hashvalue(nstore, refname, function(err, current){
          callback(err, current)
        })
      },

      function (current, callback) {
        argon2.hash(value, { type: argon2.argon2id })
          .then(hash => {
            callback(null, current, hash)
          }).catch(err => {
            callback(err)
          })
      },

      function (current, hash, callback) {
        conn = nstore.connect.getConnect();
        if (!conn) {
          callback(lang.not_connect);
        } else {
          var hashtable = process.env.NT_HASHTABLE || "ref17890714"
          var sql = ""
          if(current === null){
            sql = "insert into "+hashtable+"(refname, value) values('"+refname+"','"+hash+"') "
          } else {
            sql = "update "+hashtable+" set value='"+hash+"' where refname='"+refname+"'"
          }
          conn.query(sql, [], function (err, results) {
            callback(err)
          })
        }
      }
    ],
      function (err, value) {
        if(err){
          if (err.message) { err = err.message; }
        }
        if (conn) {
          conn.close()
        }
        _callback(err);
      });
  };

  function auth_password(params, _callback) {
    var conn
    async.waterfall([
      function (callback) {
        if(!params.username && !params.custnumber){
          callback(lang.missing_required_field+" username or custnumber")
        } else if(!params.password || !params.confirm){
          callback(lang.missing_required_field+" password or confirm")
        } else if(params.password === ""){
          callback(lang.empty_password)
        } else if(!params.password || !params.confirm){
          callback(lang.missing_required_field+" password and confirm")
        } else if(params.password !== params.confirm){
          callback(lang.verify_password)
        } else {
          conn = nstore.connect.getConnect();
          if (!conn) {
            callback(lang.not_connect);
          } else {
            callback(null)
          }
        }
      },

      function (callback) {
        var refname = ""
        if(params.username && nstore.employee()){
          if(nstore.employee().username === params.username){
            refname = "employee"+nstore.employee().id
          }
        } else if(params.custnumber && nstore.customer()){
          if(nstore.customer().custnumber === params.custnumber){
            refname = "customer"+nstore.customer().id
          }
        }
        if(refname === ""){
          var _sql
          if(params.username){
            _sql = { 
              select: "*", from: "employee", 
              where: [["deleted", "=", "0"],["and","username","=","'"+params.username+"'"]]
            }
            refname = "employee"
          } else {
            _sql = { 
              select: "*", from: "customer", 
              where: [["deleted", "=", "0"],["and","custnumber","=","'"+params.username+"'"]]
            }
            refname = "customer"
          }
          conn.query(models.getSql(nstore.engine(), _sql), [], function (err, results) {
            if (!err) {
              if (results.rowCount === 0) {
                err = lang.missing_user
              } else {
                refname += results.rows[0].id
              }
            }
            callback(err, refname);
          });
        } else {
          callback(null, refname)
        }
      },

      function (refname, callback) {
        refname = crypto.createHash('md5').update(refname).digest("hex");
        set_hashvalue(refname, params.password, function(err){
          callback(err)
        })
      }
    ],
      function (err) {
        if(err){
          if (err.message) { err = err.message; }
        }
        if (conn) {
          conn.close()
        }
        _callback(err);
      });
  };

  function auth_user(params, _callback) {
    var conn
    async.waterfall([

      function (callback) {
        create_conn(params.database, function(err){
          callback(err)
        })
      },

      function (callback) {
        conn = nstore.connect.getConnect();
        if (!conn) {
          callback(lang.not_connect);
        } else {
          var _sql = {
            select: ["e.*", "ug.groupvalue"], from: "employee e",
            inner_join: ["groups ug", "on", ["e.usergroup", "=", "ug.id"]],
            where: [["e.inactive", "=", "0"], ["and", "e.deleted", "=", "0"],
            ["and", "e.username", "=", "'"+params.username+"'"]]
          }
          conn.query(models.getSql(nstore.engine(), _sql), [], function (err, results) {
            if (!err) {
              if (results.rowCount > 0) {
                nstore.employee(results.rows[0]);
              }
            }
            callback(err);
          });
        }
      },

      function (callback) {
        if (!nstore.employee()) {
          var _sql = {
            select: ["*"], from: "customer",
            where: [["inactive", "=", "0"], ["and", "deleted", "=", "0"],
            ["and", "custnumber", "=", "'"+params.username+"'"]]
          }
          conn.query(models.getSql(nstore.engine(), _sql), [], function (err, results) {
            if (!err) {
              if (results.rowCount > 0) {
                nstore.customer(results.rows[0]);
              }
            }
            callback(err);
          });
        }
        else {
          callback(null);
        }
      },

      function (callback) {
        if (!nstore.employee() && nstore.customer()) {
          var _sql = {
            select: ["e.*", "ug.groupvalue"], from: "employee e",
            inner_join: ["groups ug", "on", ["e.usergroup", "=", "ug.id"]],
            where: [["e.inactive", "=", "0"], ["and", "e.deleted", "=", "0"],
            ["and", "e.username", "=", "'guest'"]]
          }
          conn.query(models.getSql(nstore.engine(), _sql), ["guest"], function (err, results) {
            if (!err) {
              if (results.rowCount > 0) {
                nstore.employee(results.rows[0]);
              }
            }
            callback(err);
          });
        } else {
          callback(null)
        }
      },

    ],
      function (err) {
        if (conn) {
          conn.close()
        }
        if (err) {
          if (err.message) { err = err.message; }
        } else if (!nstore.employee()) {
          err = lang.unknown_user
        }
        _callback(err, (err) ? null : nstore);
      });
  };

  function auth_user_login(params, _callback) {
    async.waterfall([

      function (callback){
        if (!params.database) {
          if (process.env.NT_DEFAULT_ALIAS) {
            params.database = process.env.NT_DEFAULT_ALIAS
            callback(null)
          } else {
            callback(lang.missing_database)
          }
        } else {
          callback(null)
        }
      },
      
      function (callback) {
        auth_user(params, function(err, ustore){
          callback(err, ustore)
        })
      },

      function (ustore, callback) {
        var refname = ""
        if(ustore.customer()){
          refname = "customer"+ustore.customer().id
        } else{
          refname = "employee"+ustore.employee().id
        }
        refname = crypto.createHash('md5').update(refname).digest("hex");
        get_hashvalue(ustore, refname, function(err, current){
          var token = auth_token(ustore)
          callback(err, token, current, ustore)
        })
      },

      function (token, current, ustore, callback) {
        params.password = params.password || ""
        if(!current){
          if(params.password === ""){
            callback(null, token, ustore);
          } else {
            callback(lang.wrong_password);
          }
        } else {
          argon2.verify(current, params.password)
          .then((match) => {
            if (match) {
              callback(null, token, ustore);
            }
            else {
              callback(lang.wrong_password);
            }
          })
          .catch((err) => {
            callback(err);
          })
        }
      },

    ],
      function (err, token, ustore) {
        if (err) {
          if (err.message) { err = err.message; }
        }
        _callback(err, token, ustore);
      });

  }

  function auth_token(ustore, params) {
    params = params || {}
    if (ustore.employee()) {
      var token = { 
        iss: process.env.NT_TOKEN_ISS || "nervatura", 
        database: ustore.database() 
      }
      if (!ustore.customer()) {
        token.username = ustore.employee().username
      } else {
        token.custnumber = ustore.customer().custnumber
      }
      return {
        token: jwt.sign(token,
          params.key || process.env.NT_TOKEN_KEY || "",
          { expiresIn: params.exp || process.env.NT_TOKEN_EXP || "1h" }
        )
      };
    } else {
      return ""
    }
  }

  function decode_token(token) {
    return jwt.decode(token, { complete: true });
  }

  return {
    ApiPost: function (params, _callback) {
      return api_post(params, _callback);
    },
    ApiGet: function (params, _callback) {
      return api_get(params, _callback);
    },
    ApiDelete: function (params, _callback) {
      return api_delete(params, _callback);
    },
    Report: function (params, _callback) {
      return report(params, _callback);
    },
    ReportList: function (params, _callback) {
      return report_list(params, _callback);
    },
    ReportInstall: function (params, _callback) {
      return report_install(params, _callback);
    },
    ReportDelete: function (params, _callback) {
      return report_delete(params, _callback);
    },
    DatabaseCreate: function (params, _callback) {
      return database_create(params, _callback);
    },
    ApiView: function (params, _callback) {
      return api_view(params, _callback);
    },
    ApiFunction: function (params, _callback) {
      return api_function(params, _callback);
    },
    AuthTokenLogin: function (params, _callback) {
      return auth_token_login(params, _callback);
    },
    AuthUserLogin: function (params, _callback) {
      return auth_user_login(params, _callback);
    },
    AuthPassword: function (params, _callback) {
      return auth_password(params, _callback);
    },
    AuthToken: function (params) {
      return auth_token(nstore, params);
    },
    AuthDecode: function (params) {
      return decode_token(params);
    }
  };
};
