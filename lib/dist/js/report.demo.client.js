/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/*globals $ */
/*globals requirejs */
/*globals hljs */

requirejs.config({
  "baseUrl": "../javascripts",
  "paths": {
    "preview": "report.demo.preview",
    "xml_template": "report.demo.sample.xml",
    "json_template": "report.demo.sample.json",
    "npi": "NpiAdapter",
    "report": "Report",
    "pdf": "pdf.min",
    "jspdf": "jspdf.min",
    "hljs": "highlight.pack"
  },
  "shim": {
    "jspdf": {
      exports: "jsPDF"}
  }
});

requirejs(["preview","npi","report","xml_template","json_template","hljs"], 
  function(preview, NpiAdapter, Report, xml, json){
  
  function show_template(ttype) {
    $("#flash").html("");
    $("#flash").css("display","none");
    $("#dv_xml_template").css("display",(ttype==="xml") ? "block" : "none");
    $("#dv_json_template").css("display",(ttype==="json") ? "block" : "none");
    $("#dv_js_template").css("display",(ttype==="js") ? "block" : "none");
    $("#dv_dbs_template").css("display",(ttype==="dbs") ? "block" : "none");}
  
  function report_output(rpt, output) {
    switch (output) {
      case "win":
        rpt.save2DataUrl();
        break;
      case "prev":
        preview.showReport(rpt.save2Pdf());
        break;
      case "save":
        rpt.save2PdfFile("Report.pdf");
        break;
      case "xml_data":
        var xdata; 
        xdata = new Blob([rpt.save2Xml()], {type: 'text/xml'});
        window.URL.revokeObjectURL(xdata); xdata = window.URL.createObjectURL(xdata);
        window.open(xdata, '_blank');
        break;
      case "xml_temp":
        var xtdata; 
        xtdata = new Blob([rpt.getXmlTemplate()], {type: 'text/xml'});
        window.URL.revokeObjectURL(xtdata); xtdata = window.URL.createObjectURL(xtdata);
        window.open(xtdata, '_blank');
        break;
      case "json_temp":
        var jdata; 
        jdata = new Blob([JSON.stringify(rpt.template.elements, null, " ")], {type: 'application/json'});
        window.URL.revokeObjectURL(jdata); jdata = window.URL.createObjectURL(jdata);
        window.open(jdata, '_blank');
        break;}}
        
  function load_dbs(input,output) {
    $("#flash").html("");
    $("#flash").css("display","none");
    var ref_no = $("#dbs_no").val();
    if (ref_no==="") {
      $("#flash").html("Missing Doc.No.");
      $("#flash").css("display","block");
      return;}
    var server = "/npi/call/jsonrpc2/";
    var login = {database:$("#dbs_name").val(), username:$("#dbs_user").val(), password:$("#dbs_psw").val()};
    var da = new NpiAdapter(server);
    da.callFunction(login, "getReport", {nervatype:$("#dbs_temp").val(),
      refnumber:ref_no, output:input, orientation:$("#orientation").val()}, function(state,data){
      if (state==="ok") {
        if ("error_message" in data) {
          $("#flash").html(data.error_message);
          $("#flash").css("display","block");}
        else {
          var rpt = new Report($("#orientation").val());
          rpt.loadDefinition(data.template);
          for(var i = 0; i < Object.keys(data.data).length; i++) {
            var pname = Object.keys(data.data)[i];
            rpt.setData(pname, data.data[pname]);}
          rpt.createReport();
          report_output(rpt, output);}} 
      else {
        $("#flash").html(data);
        $("#flash").css("display","block");}});}
  
  function create_report(rpt) {
  
    //default values
    rpt.template.document.title = "Nervatura Report";
    rpt.template.margins["left-margin"] = 15;
    rpt.template.margins["top-margin"] = 15;
    rpt.template.margins["right-margin"] = 15;
    rpt.template.style["font-family"] = "times";
    
    //header
    var header = rpt.template.elements.header;
    var row_data = rpt.insertElement(header, "row", -1, {height: 10});
    rpt.insertElement(row_data, "image",-1,{src:"logo"});
    rpt.insertElement(row_data, "cell",-1,{
      name:"label", value:"labels.title", "font-style": "bolditalic", "font-size": 26, color: "#D8DBDA"});
    rpt.insertElement(row_data, "cell",-1,{
      name:"label", value:"Javascript Sample", "font-style": "bold", align: "right"});
    rpt.insertElement(header, "vgap", -1, {height: 2});
    rpt.insertElement(header, "hline", -1, {"border-color": 218});
    rpt.insertElement(header, "vgap", -1, {height: 2});
    
    //details
    var details = rpt.template.elements.details;
    rpt.insertElement(details, "vgap", -1, {height: 2});
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", width: "50%", "font-style": "bold", value: "labels.left_text", border: "LT", 
      "border-color": 218, "background-color": 245});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", "font-style": "bold", value: "labels.left_text", border: "LTR", 
      "border-color": 218, "background-color": 245});
    
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", width: "50%", value: "head.short_text", border: "L", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", value: "head.short_text", border: "LR", "border-color": 218});
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", width: "50%", value: "head.short_text", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", value: "head.short_text", border: "LBR", "border-color": 218});
    
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", width: "40", "font-style": "bold", value: "labels.left_text", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", align: "center", width: "30", "font-style": "bold", value: "labels.center_text", 
      border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", align: "right", width: "40", "font-style": "bold", value: "labels.right_text", 
      border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", "font-style": "bold", value: "labels.left_text", border: "LBR", "border-color": 218});
    
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", width: "40", value: "head.short_text", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "date", align: "center", width: "30", value: "head.date", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "amount", align: "right", width: "40", value: "head.number", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", value: "head.short_text", border: "LBR", "border-color": 218});
    
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", "font-style": "bold", value: "labels.left_text", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", width: "50", value: "head.short_text", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", "font-style": "bold", value: "labels.left_text", border: "LB", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", value: "head.short_text", border: "LBR", "border-color": 218});
    
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "long_text", "multiline": "true", value: "head.long_text", border: "LBR", "border-color": 218});
    
    rpt.insertElement(details, "vgap", -1, {height: 2});
    row_data = rpt.insertElement(details, "row", -1, {hgap: 2});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", value: "labels.left_text", "font-style": "bold", border: "1", "border-color": 245, 
      "background-color": 245});
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", value: "head.short_text", border: "1", "border-color": 218});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", value: "labels.left_text", "font-style": "bold", border: "1", "border-color": 245, "background-color": 245});
    rpt.insertElement(row_data, "cell",-1,{
      name: "short_text", value: "head.short_text", border: "1", "border-color": 218});
    
    rpt.insertElement(details, "vgap", -1, {height: 2});
    row_data = rpt.insertElement(details, "row", -1, {hgap: 2});
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", value: "labels.long_text", "font-style": "bold", border: "1", "border-color": 245, "background-color": 245});
    rpt.insertElement(row_data, "cell",-1,{
      name: "long_text", "multiline": "true", value: "head.long_text", border: "1", "border-color": 218});
    
    rpt.insertElement(details, "vgap", -1, {height: 2});
    rpt.insertElement(details, "hline", -1, {"border-color": 218});
    rpt.insertElement(details, "vgap", -1, {height: 2});
    
    row_data = rpt.insertElement(details, "row", -1, {"hgap": 3});
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "Barcode (Interleaved 2of5)", "font-style": "bold", "font-size": 10,
      "border": "1", "border-color": 245, "background-color": 245});
    rpt.insertElement(row_data, "barcode",-1,{"code-type": "ITF", "value": "1234567890", "visible-value":1});
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "Barcode (Code 39)", "font-style": "bold", "font-size": 10, 
      "border": "1", "border-color": 245, "background-color": 245});
    rpt.insertElement(row_data, "barcode",-1,{"code-type": "CODE_39", "value": "1234567890ABCDEF", "visible-value":1});
    
    rpt.insertElement(details, "vgap", -1, {height: 3});
    
    row_data = rpt.insertElement(details, "row");
    rpt.insertElement(row_data, "cell",-1,{
      name: "label", value: "Datagrid Sample", align: "center", "font-style": "bold", 
      border: "1", "border-color": 245, "background-color": 245});
    rpt.insertElement(details, "vgap", -1, {height: 2});
    
    var grid_data = rpt.insertElement(details, "datagrid", -1, {
      name: "items", databind: "items", border: "1", "border-color": 218, "header-background": 245, "footer-background": 245});
    rpt.insertElement(grid_data, "column",-1,{
      width: "8%", fieldname: "counter", align: "right", label: "labels.counter", footer: "labels.total"});
    rpt.insertElement(grid_data, "column",-1,{
      width: "20%", fieldname: "date", align: "center", label: "labels.center_text"});
    rpt.insertElement(grid_data, "column",-1,{
      width: "15%", fieldname: "number", align: "right", label: "labels.right_text", 
      footer: "items_footer.items_total", "footer-align": "right"});
    rpt.insertElement(grid_data, "column",-1,{
      fieldname: "text", label: "labels.left_text"});
    
    rpt.insertElement(details, "vgap", -1, {height: 5});
    rpt.insertElement(details, "html", -1, {fieldname: "html_text", 
      html: "<i>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</i> ={{html_text}} Nulla a pretium nunc, in cursus quam."});
    
    //footer
    var footer = rpt.template.elements.footer;
    rpt.insertElement(footer, "vgap", -1, {height: 2});
    rpt.insertElement(footer, "hline", -1, {"border-color": 218});
    row_data = rpt.insertElement(footer, "row", -1, {height: 10});
    rpt.insertElement(row_data, "cell",-1,{value: "Nervatura Report Template", "font-style": "bolditalic"});
    rpt.insertElement(row_data, "cell",-1,{value: "{{page}}", align: "right", "font-style": "bold"});
    
    //data
    rpt.setData("labels", {"title": "REPORT TEMPLATE", "left_text": "Short text", "center_text": "Centered text", 
                                        "right_text": "Right text", "long_text": "Long text", "counter": "No.", "total": "Total"});
    rpt.setData("head", {"short_text": "Lorem ipsum dolor", "number": "123 456", "date": "2015.01.01", 
                                      "long_text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim. Nulla a pretium nunc, in cursus quam."});
    rpt.setData("html_text", "<p><b>Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim.</b></p>");
    rpt.setData("items_footer", {"items_total": "3 703 680"});
    var items = [];
    for(var i=0; i<30; i++) {
      items.push({"text": "Lorem ipsum dolor", "number": "123 456", "date": "2015.01.01"});}
    rpt.setData("items", items);
    rpt.setData("logo", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAAAXNSR0IArs4c6QAAAAZiS0dEAP8A/wD/oL2nkwAAAAlwSFlzAAALEwAACxMBAJqcGAAAAAd0SU1FB9wIExQZM+QLBuYAAAQ5SURBVEjHtZVJbJRlGMd/77fM0ul0ykynLQPtjEOZWjptSisuJRorF6ocjEaDoDGSYKLiFmIMxoPGg0viRY0cNB6IUQ+YYCJEBLXQWqZgK6GU2taWLrS0dDrdO0vn+14PlYSlQAH9357L+89/eZ9HHDp0iP8TyqXDCV9Iuivvkf8lgXbp8JjPxUfHfseXs1yGs52sd+VgpueJqQYdPRE8sxniZgnElRbF8p1yXW6AgWgbDeeOoJgCu8PJtFbAc2W1tEeaxC0rADDTKVZmO+kxKwjawmwO5pJIGhzpP8je7m5KbycDgKnEDEnDYEO+G10xeT3Sxu6eIYJuP5OJffT5ihbNKD/klzckWLWuSibjUSyahZn5NBtXerkj0059dy+rPaUs1wvoiH3GXjkgIy6nbPV65PEMQzaNN0qLu+j6GZTfX5P+NPKjuqXIj8sV4vR0gv7ZJP2zCUpsEEdhW7GfztFOBqf6cQmV5ByUFpfxyr6veHpZpbhuBk/91qw+lGrDwIfXbqfGbkVXVQCiyRQvNZ5hY0EeIW+IkDfEm83dvL82wHRqAkOz3TiDXdqUeGLDLg73xfj5r4PoqopkwVa3RefFkkKeaWgDYHfXENV5LhRFRVMsZFvl0kLuP/aLCM+p4tfeLs4MtyEQgEQRgrs8Th70OtkeaefE2BQPeF1ICQ6Lk9WOOKMrlssltQjAavdwdqQbACkXrHXoOllWnWgiiVNApq4g/nV9+/pXaWrfw3Dh1Q1blCAr3kO5v5K0hPHUPHt6hinZf4JILE6R005CUTBNSJkmFxIp4qbG25veY3LyKJ9Pj8iGDIf0lFfIRT9aVjBHBsaK8Wb4qBuJMTyTpNTtoP2RdQD8PTXHjuOdfNwxiFWBHKsFj1WjPEvDJjIJOBXuCxQzdqpFLLoqJjwTsrbkYSyKBVPTsF6qUUpSpqRuOMaT9e0cq11LiSsTgLd+eheLq4YXKqtpra8Ti1o06ByQwYLVHB05yoHz+2mM1jGWGAPJQqOEoGsmzmvNnXxZHaL2yEKrvjv9LarmY2t4/WWPX6bAnm+TZ40kAlB0GDfO08dJptNxnl+xnYrcClrHZ9g7MMqOkI8PTvejaBo5Np253k/I879MaLBDXDPk+HBC5I9KkTcqhXdIitBIvtiau5Nszc3Xnd8A8GdsmnfKAnhtVnYW2iE1y8GhcbQZlUeDBSy5phdxobVNhPVnUdU0J893scXvITYbpbGvgVh6iIlUgiqPixbWcqb+sFjSur4Sued6hc0WkJIoX/zRQyI1icvhJG0YhA2DmJJDoxqgZoVXrhnsFjel4CIy5Ep+ONXKpjXVlKXconBcF8Epmyidd4hlFpU37vTxYccIuVX3ylsiUDNWse3ux+k4fvU181t1Nvu9pE3J9+eiNz6Zi8FfWSX7WpqveSoHg2EpgHC2g1hLk7hpgtvBP6lBrRsE+ni7AAAAAElFTkSuQmCC");
      
    return rpt;
  }
      
  function load_file(output) {
    if ($("#dv_dbs_template").css("display") === "block") {
      return load_dbs("tmp",output);}
    var ctype = "xml";
    if ($("#dv_json_template").css("display") === "block") {
      ctype = "json";}
    else if ($("#dv_js_template").css("display") === "block") {
      ctype = "js";}
    var rpt = new Report($("#orientation").val());
    switch (ctype) {
      case "js":
        create_report(rpt);
        break;
      case "xml":
        var xml_str = $("#dv_xml_template").val();
        rpt.loadDefinition(xml_str);
        break;
      case "json":
        var json_str = $("#dv_json_template").val();
        rpt.loadJsonDefinition(json_str);
        break;}
    rpt.createReport();
    report_output(rpt, output);}
  
  $("#prev").on("click", preview.onPrevPage);
  $("#next").on("click", preview.onNextPage);
  
  $("#set_xml_content").on("click", function() {show_template("xml");});
  $("#set_dbs_content").on("click", function() {show_template("dbs");});
  
  $("#client_prev_report").on("click", function() {load_file('prev');});
  $("#win_report").on("click", function() {load_file("win");});
  
  $("#save_pdf_report").on("click", function() {load_file("save");});
  $("#save_xml_data").on("click", function() {load_file("xml_data");});
  $("#panel_state").on("click", function() {
    $("#temp_panel").css("display", $("#temp_panel").css("display")==="none" ? "block" : "none");});
  $("#save_xml_temp").on("click", function() {load_file("xml_temp");});
  $("#save_json_temp").on("click", function() {load_file("json_temp");});
  
  $("#set_json_content").on("click", function() {show_template("json");});
  $("#set_js_content").on("click", function() {show_template("js");});

  $("#dv_xml_template").val(xml.Template);
  $("#dv_xml_template").css("height","300px");
  $("#dv_json_template").val(JSON.stringify(json, null, " "));
  $("#js_code").html(create_report.toString());
  hljs.highlightBlock($("#js_code")[0]);
  
  $("#dbs_name").val("demo");
  $("#dbs_user").val("demo");
  $("#dbs_psw").val("");
  $("#dbs_no").val("DMINV/00001");  
});
