
function showTemplate(ttype) {
  if(document.getElementById('_flash')){
    document.getElementById('_flash').innerHTML = "";
    document.getElementById('flash_').style.display = "none"; }
  document.getElementById('dv_xml_template').style.display = 
    (ttype==="xml") ? "block" : "none";
  document.getElementById('dv_dbs_template').style.display = 
    (ttype==="dbs") ? "block" : "none";
}

function openWindow(wtype) {
  var url;
  switch (wtype) {
    case "template":
      url = "/report/template?"+document.getElementById('template').value+"=true";
      break;

    case "pdf":
      url = "/report/document?"+document.getElementById('orientation').value+"=true";
      url += "&"+document.getElementById('template').value+"=true";
      break;
    
    case "html":
      url = "/report/document?html=true&"+document.getElementById('orientation').value+"=true";
      url += "&"+document.getElementById('template').value+"=true";
      break;

    case "xml":
      url = "/report/document?data=true";
      url += "&"+document.getElementById('template').value+"=true";
      break;

    default:
      break;}
  window.open(url, '_blank');
}

function loadDbs(input, output) {
  if(document.getElementById('_flash')){
    document.getElementById('_flash').innerHTML = "";
    document.getElementById('flash_').style.display = "none"; }
  var ref_no = document.getElementById("dbs_no").value;
  if (ref_no === "") {
    document.getElementById('flash_').style.display = "block";
    document.getElementById('_flash').innerHTML = "Missing Doc.No.";
    return;}
  var server = "/npi/call/jsonrpc2/";
  var login = {database: document.getElementById('dbs_name').value, 
    username: document.getElementById('dbs_user').value, 
    password: document.getElementById('dbs_psw').value};
  var da = new window.NpiAdapter(server);
  da.callFunction(login, "getReport", {
    nervatype: document.getElementById('dbs_temp').value,
    refnumber:ref_no, output:input, orientation: document.getElementById('orientation').value}, 
    function(state,data){
    if (state==="ok") {
      if ("error_message" in data) {
        document.getElementById('flash_').style.display = "block";
        document.getElementById('_flash').innerHTML = data.error_message;
      }
      else {
        var arrayBuffer = new Uint8Array(data.template.data || data.template).buffer;
        window.Preview.showReport(arrayBuffer);
      }
    } 
    else {
      document.getElementById('flash_').style.display = "block";
      document.getElementById('_flash').innerHTML = data;
    }
  });
}

function nextPage() {
  window.Preview.onNextPage();
}

function prevPage() {
  window.Preview.onPrevPage();
}

function showPreview() {
  if (document.getElementById('dv_dbs_template').style.display === "block") {
    return loadDbs("pdf","pdf");
  }
  else{
    var url = "/report/document?"+document.getElementById('orientation').value+"=true";
    url += "&"+document.getElementById('template').value+"=true";
    window.Preview.showReport(url);
  }
}

//init
document.getElementById('dbs_name').value = "demo";
document.getElementById('dbs_user').value = "demo";
document.getElementById('dbs_psw').value = "";
document.getElementById('dbs_no').value = "DMINV/00001";