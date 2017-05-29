/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

function sendToServer(){
  if (document.getElementById("rs_url").value!="") {
    window.open(document.getElementById("rs_url").value, '_blank');}}

var jdata = null;
function sendToJsonServer(){
  if (jdata !== null) {
    var xhr = new XMLHttpRequest();
    var url = window.location.protocol+'//'+window.location.host+"/ndi/";
    if ("withCredentials" in xhr) {
      xhr.open("post", url, true);
    } else if (typeof XDomainRequest !== "undefined") {
      //for IE.
      xhr = new XDomainRequest();xhr.open("post", url);
    } else {
      alert("CORS not supported.");}
  
    xhr.onload = function() {
      var response = JSON.parse(xhr.responseText);
      if ("error" in response) {
        alert(response.error.message);
      } else {
        var result
        if ("result" in response) {
          result = response.result
        } else {
          result = response}
        if (typeof result === "object") {
          result = JSON.stringify(response.result);}
        var win = window.open("","_blank");
        win.document.write("<html><head><title>NDI RESULT</title></head><body>"+result+"</body></html>");
        win.document.close();
      }
    };

    xhr.onerror = function() {
      alert("POST ERROR");};
    
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.send(JSON.stringify(jdata));}}

function getUrl(urlfunc){
	var url = "@protocol//@server/@function?@code&@params&@data";
  url = url.replace("@protocol", window.location.protocol);
  url = url.replace("@server", window.location.host);
  url = url.replace("@function", "ndi/"+urlfunc);
  url = url.replace("@code", "code="+document.getElementById("code").value);
  return url;}
    
function getParams(urlfunc){
  var params = "database=@database|username=@username|password=@password|datatype=@datatype";
  params = params.replace("@database", document.getElementById("database").value);
  params = params.replace("@username", document.getElementById("username").value);
  params = params.replace("@password", document.getElementById("password").value);
  params = params.replace("@datatype", document.getElementById("datatype").value);
  if (document.getElementById("use_deleted")){
    if (document.getElementById("use_deleted").checked) {
  	  params = params+"|use_deleted";}}
  if (document.getElementById("insert_row")){
    if (urlfunc==="updateData" && document.getElementById("insert_row").value === "true") {
      params = params+"|insert_row";}}
  if (document.getElementById("insert_field")) {
    if (urlfunc==="updateData" && document.getElementById("insert_field").value === "true") {
      params = params+"|insert_field";}}
  if (document.getElementById("code")) {
    if (document.getElementById("code").value==="base64") {
      params = "params="+atob(params);
    } else {
      params = "params="+params;}}
  return params;}
  
function getParamsJson(urlfunc, jdata, server){
  if (!server){
    jdata.params[0].database = document.getElementById("database").value;
    jdata.params[0].username = document.getElementById("username").value;
    jdata.params[0].password = document.getElementById("password").value;}
  jdata.params[0].datatype = document.getElementById("datatype").value;
  if (document.getElementById("use_deleted")){
    if (document.getElementById("use_deleted").checked) {
      jdata.params[0].use_deleted = true;}}
  if (document.getElementById("insert_row")){
    if (urlfunc==="updateData" && document.getElementById("insert_row").value === "true") { 
      jdata.params[0].insert_row = true;}}
  if (document.getElementById("insert_field")) {
    if (urlfunc==="updateData" && document.getElementById("insert_field").value === "true") { 
      jdata.params[0].insert_field = true;}}
  return jdata;}

function createView(){
  if (document.getElementById("datatype").value=="") return;
  var url = getUrl("getData");
  var params = getParams("getData");
  var server_params = getParamsJson("getData", {params:[{},[]]}, true).params[0];
  jdata = {"id":1, "method":"getData", "jsonrpc":"2.0", params:[{},{}]};
  jdata = getParamsJson("getData", jdata);
	
  var filters = "";
  if (document.getElementById("datatype").value!="sql") {
    filters += "output="+document.getElementById("output").value;
    jdata.params[1].output = document.getElementById("output").value;
    ["view_nervatype","view_nervatype1","view_nervatype2"].forEach(function(ename) {
      if (document.getElementById(ename)){
        var element = document.getElementById(ename); 
        if (element.value !== "") {
          filters += "|"+element.name.replace("view_","")+"="+element.value;
          jdata.params[1][element.name.replace("view_","")] = element.value;}}});
    ["where","orderby","header","columns"].forEach(function(ename) {
      var element = document.getElementById(ename);
      if (element.value !== "") {
        filters += "|"+ename+"="+element.value;
        jdata.params[1][ename] = element.value;}});
    if (document.getElementById("cross_tab")) {
      if (document.getElementById("cross_tab").value === "true") {
        filters += "|cross_tab";
        jdata.params[1].cross_tab = true;}}
    if (document.getElementById("show_id")) {
      if (document.getElementById("show_id").value === "true") {
        filters += "|show_id";
        jdata.params[1].show_id = true;}}
    if (document.getElementById("no_deffield")) {
      if (document.getElementById("no_deffield").value === "true") {
        filters += "|no_deffield";
        jdata.params[1].no_deffield = true;}}}
	else {
		filters += "output="+document.getElementById("output").value;
		jdata.params[1].output = document.getElementById("output").value;
		filters += "|sql="+document.getElementById("sql").value;
		jdata.params[1].sql = document.getElementById("sql").value;}
	if (jdata.params[1].output === "xml" || jdata.params[1].output === "excel") {
	jdata.params[1].output = "json";}
		
	if (document.getElementById("code").value=="base64") {
	  filters = "filters="+atob(filters);} 
	else {
	  filters = "filters="+filters;}
	url = url.replace("@params", params);
	url = url.replace("@data", filters);
	document.getElementById("rs_url").value = url;
	document.getElementById("rs_server").value = 'getData('+
    JSON.stringify(server_params)+','+JSON.stringify(jdata.params[1])+')';
  document.getElementById("rs_json").value = "url: "+window.location.protocol+'//'+ 
	  window.location.host+"/ndi/\ndata: "+JSON.stringify(jdata);}
 
function get_input_lst(itype){
  var inputs = document.querySelectorAll("#rs_"+itype+" input[id=cb_"+itype+"]");
  var values = {}; var nom = document.getElementById("datatype").value;
  for (var i=0;i<inputs.length;i++) {
    if (inputs[i].checked){
      var _row = inputs[i].name.split("_");
      var row = _row[_row.length-1];
      if (typeof values[row] === "undefined"){
        values[row] = [];}
      if (itype === "update"){
        values[row].push({
        fieldname:inputs[i].name.replace("cb_"+itype+"_"+nom+"_","").replace("_"+row,""),
        value:document.querySelector("#"+inputs[i].name.replace("cb_"+itype+"_"+nom+"_","")).value
      });}
      else {
        var keys = document.querySelectorAll("#rs_"+itype+" input[id=delete_"+nom+"_"+row+"]");
        for (var k=0;k<keys.length;k++) {
          values[row].push({
            fieldname:keys[k].name.replace("delete_"+nom+"_","").replace("_"+row,""),
            value:keys[k].value});}
        keys = document.querySelectorAll("#rs_"+itype+" select[id=delete_"+nom+"_"+row+"]");
        for (var k=0;k<keys.length;k++) {
          values[row].push({
            fieldname:keys[k].name.replace("delete_"+nom+"_","").replace("_"+row,""),
            value:keys[k].value});}
      }};}
  return values;}
 
function createUpdate(itype){
	if (document.getElementById("datatype").value=="") return;
	
	var url = getUrl(itype+"Data");
	var params = getParams(itype+"Data");
  var server_params = getParamsJson(itype+"Data", {params:[{},[]]}, true).params[0];
	jdata = {"id":2, "method":itype+"Data", "jsonrpc":"2.0", params:[{},[]]};
  if (itype==="delete"){jdata.id = 3;}
  jdata = getParamsJson(itype+"Data", jdata);
  
	var data = ""; var inputs = get_input_lst(itype);
  for (var i in inputs) {
    var row = inputs[i];
    var data_row=""; var data_json_row = {};
    row.forEach(function(field) {
      data_row += "|"+field.fieldname+"="+field.value;
			data_json_row[field.fieldname] = field.value;});
    if (data_row!="") {
			data += '|'+data_row;
			jdata.params[1].push(data_json_row);}}
	
	if (document.getElementById("code").value=="base64") {
		data = "data="+atob(data.substring(2));} 
  else {
   data = "data="+data.substring(2);}
	
	url = url.replace("@params", params);
	url = url.replace("@data", data);
	document.getElementById("rs_url").value = url;
  document.getElementById("rs_json").value = "url: "+window.location.protocol+'//'+ 
	  window.location.host+"/ndi/\ndata: "+JSON.stringify(jdata);
  document.getElementById("rs_server").value = itype+'Data('+
    JSON.stringify(server_params)+','+JSON.stringify(jdata.params[1])+')';}
    