
function showTemplate(ttype) {
  document.getElementById('flash').innerHTML = "";
  document.getElementById('flash').style.display = "none";
  document.getElementById('dv_xml_template').style.display = (ttype==="xml") ? "block" : "none";
  document.getElementById('dv_json_template').style.display = (ttype==="json") ? "block" : "none";
  document.getElementById('dv_js_template').style.display = (ttype==="js") ? "block" : "none";
  document.getElementById('dv_dbs_template').style.display = (ttype==="dbs") ? "block" : "none";
}

function showPanel() {
  document.getElementById('temp_panel').style.display =
    (document.getElementById('temp_panel').style.display === "none") ? "block" : "none";
}

function reportOutput(rpt, output) {
  switch (output) {
    case "win":
      rpt.save2DataUrl();
      break;
    case "prev":
      window.Preview.showReport(rpt.save2Pdf());
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
      break;}
}

function loadDbs(input,output) {
  document.getElementById('flash').innerHTML = "";
  document.getElementById('flash').style.display = "none";
  var ref_no = document.getElementById('dbs_no').value;
  var icon = '<i class="fa fa-exclamation-triangle fa-fw" aria-hidden="true"></i>&nbsp;'
  if (ref_no==="") {
    document.getElementById('flash').innerHTML = icon+"Missing Doc.No.";
    document.getElementById('flash').style.display = "block";
    return;}
  var server = "/npi/call/jsonrpc2/";
  var login = {database: document.getElementById('dbs_name').value, 
    username: document.getElementById('dbs_user').value, 
    password: document.getElementById('dbs_psw').value};
  var da = new window.NpiAdapter(server);
  da.callFunction(login, "getReport", {nervatype: document.getElementById('dbs_temp').value,
    refnumber:ref_no, output:input, orientation: document.getElementById('orientation').value}, 
    function(state,data){
    if (state==="ok") {
      if ("error_message" in data) {
        document.getElementById('flash').innerHTML = icon+data.error_message;
        document.getElementById('flash').style.display = "block";}
      else {
        var rpt = new window.Report(document.getElementById('orientation').value);
        rpt.loadDefinition(data.template);
        for(var i = 0; i < Object.keys(data.data).length; i++) {
          var pname = Object.keys(data.data)[i];
          rpt.setData(pname, data.data[pname]);}
        rpt.createReport();
        reportOutput(rpt, output);}} 
    else {
      document.getElementById('flash').innerHTML = icon+data;
      document.getElementById('flash').style.display = "block";}});
}

function createJS(rpt) {
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
  rpt.setData("logo", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABoAAAAaCAYAAACpSkzOAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAFxGAABcRgEUlENBAAAAB3RJTUUH4AcCFiIAfaA8WwAAAeBJREFUSMftlj9IVlEUwH/nfR/+G3QwaPlAg3D5pKkPDIoIghYVm4yCFKlPcGoQdGppa4mm4IJbQdFYRKDgoFR0o5ZwcL5LLoKCoPW+43JfvO/6ns9P0ckDj8c5l3t+55x73rkPzqUFqdRrTe/Tho0CRGeQ2FWA8ilmcgH4CUyeGFSp13DGpvUOoB+YSwDALwBp0VF6jwAKXAbGgIfAYOBi1Rl7oxDknXf56IaBKtDhnzagvcBHzRn7IxOUZFGp12aAaeDKMSv7D+gBdpyxzSBfmnZgGbh2wn5Yc8ZWEyXKOI+lY0C2RGQ9sD1OK/9BvlyjwPUWIZsi8lFVB1K2DWfsl/REKAdle94i5LuIbKnq/cC+mAR/ICMvvUcliMgroFNVb2csT4WGcoGeJbsisqCq4zmBLThj90JjmNFeQbt+i0ReqOpMDiR2xj7KmtYh6E/WZhH5WoqiCWCooTp/SDC3wrPJ7Dr/9Sdn8L5UKt0T5JOqVuNG401BSaedsSt5d08ULFwENoBNVb0Zx/FbRUeA7gLIU2esCWdjU/P41h4HZv3d8Q74DTw7QmP8Baacsa8Pg6S7bBvoSw9BP0yfAJ05gM/AXSDOO5cDGeVdC972ALgDXAJ2gQ/O2Jfnf0RnJvumbKT0gnMTFgAAAABJRU5ErkJggg==");
    
  return rpt;
}

function load_file(output) {
  if (document.getElementById('dv_dbs_template').style.display === "block") {
    return loadDbs("tmp",output);}
  var ctype = "xml";
  if (document.getElementById('dv_json_template').style.display === "block") {
    ctype = "json";}
  else if (document.getElementById('dv_js_template').style.display === "block") {
    ctype = "js";}
  var rpt = new window.Report(document.getElementById('orientation').value);
  switch (ctype) {
    case "js":
      createJS(rpt);
      break;
    case "xml":
      var xml_str = document.getElementById('dv_xml_template').value;
      rpt.loadDefinition(xml_str);
      break;
    case "json":
      var json_str = document.getElementById('dv_json_template').value;
      rpt.loadJsonDefinition(json_str);
      break;}
  rpt.createReport();
  reportOutput(rpt, output);
}

function nextPage() {
  window.Preview.onNextPage();
}

function prevPage() {
  window.Preview.onPrevPage();
}

//init
document.getElementById('dv_xml_template').value = atob(document.getElementById('xml-data').value);
document.getElementById('dv_json_template').value = atob(document.getElementById('json-data').value);
document.getElementById('js_code').innerHTML = createJS.toString();
window.hljs.highlightBlock(document.getElementById('js_code'));
document.getElementById('dbs_name').value = "demo";
document.getElementById('dbs_user').value = "demo";
document.getElementById('dbs_psw').value = "";
document.getElementById('dbs_no').value = "DMINV/00001";