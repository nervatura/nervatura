/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/*globals $ */
/*globals requirejs */

requirejs.config({
  "baseUrl": "../javascripts",
  "paths": {
    "preview": "report.demo.preview",
    "npi": "NpiAdapter",
    "pdf": "pdf.min"
  }
});

requirejs(["preview","npi"], function(preview, NpiAdapter){
  
  function show_template(ttype) {
    $("#flash").html("");
    $("#flash").css("display","none");
    $("#dv_xml_template").css("display",(ttype==="xml") ? "block" : "none");
    $("#dv_dbs_template").css("display",(ttype==="dbs") ? "block" : "none");}
  
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
          var arrayBuffer = new Uint8Array(data.template.data || data.template).buffer;
          preview.showReport(arrayBuffer);}} 
      else {
        $("#flash").html(data);
        $("#flash").css("display","block");}});}
      
  $("#prev").on("click", preview.onPrevPage);
  $("#next").on("click", preview.onNextPage);
  
  $("#set_xml_content").on("click", function() {show_template("xml");});
  $("#set_dbs_content").on("click", function() {show_template("dbs");});
  
  $("#server_prev_report").on("click", function() {
    if ($("#dv_dbs_template").css("display") === "block") {
	    return load_dbs("pdf","pdf");}
    else{
      var url = "/report/document?"+$("#orientation").val()+"=true";
      url += "&"+$("#template").val()+"=true";
      preview.showReport(url);}});
  $("#html_report").on("click", function() {
    var url = "/report/document?html=true&"+$("#orientation").val()+"=true";
    url += "&"+$("#template").val()+"=true";
    window.open(url, '_blank');});
  $("#xml_data").on("click", function() {
    var url = "/report/document?data=true";
    url += "&"+$("#template").val()+"=true";
    window.open(url, '_blank');});
  $("#open_template").on("click", function() {
    var url = "/report/template?"+$("#template").val()+"=true";
    window.open(url, '_blank');});
  $("#pdf_report").on("click", function() {
    var url = "/report/document?"+$("#orientation").val()+"=true";
    url += "&"+$("#template").val()+"=true";
    window.open(url, '_blank');});
    
  $("#dbs_name").val("demo");
  $("#dbs_user").val("demo");
  $("#dbs_psw").val("");
  $("#dbs_no").val("DMINV/00001");

});
