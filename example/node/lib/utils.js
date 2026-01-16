var path = require('path');
var spawn = require('child_process').spawn;

var servicePath = path.join("node_modules","nervatura","bin")
var serviceFile = (process.platform === "win32") ? "nervatura.exe" : "nervatura"
var serviceLib = (process.platform === "win32") ? "nervatura.dll" : "nervatura.so"

exports.servicePath = servicePath
exports.serviceFile = serviceFile
exports.serviceLib = serviceLib

exports.CheckJson = function(data, dataResult) {
  try { 
    const result = JSON.parse(data) 
    return result; 
  } catch (error) {
    if(dataResult){
      return data
    }
    return {code:400, message: String(data)};
  } 
}

exports.EncodeOptions = function(data) {
  return JSON.stringify(data)
}

exports.ToBytes = (string) => {
	const buffer = Buffer.from(string, 'utf8');
	const result = Array(buffer.length);
	for (var i = 0; i < buffer.length; i++) {
		result[i] = buffer[i];
	}
	return result;
};

exports.StartService = function(consolLog, callback){
  const controller = new AbortController();
  const { signal } = controller;
  const child = spawn(path.join(servicePath, serviceFile), { env: { ...process.env }, signal });
  var result = false;
  child.stdout.on('data', (chunk) => {
    if(consolLog){
      console.log(chunk.toString());
    }
    if(!result){
      result = true
      callback(null, controller)
    } 
  });
  child.stderr.on('data', function(err) {
    if(!result){
      result = true
      callback(err, controller)
    }
  });
  child.on('error', function (err) {
    if(!result){
      callback(err, controller)
    }
  });
}

exports.GetApi = function(token, api_type, fName, options, apiCallback) {
  var http = require('./rest');
  var cli = require('./cli');
  var rpc = require('./rpc');

  switch (api_type) {
    case "cli":
      cli[fName](token, options, apiCallback)
      break;
    
    case "rpc":
      rpc[fName](token, options, apiCallback)
      break;

    default:
      http[fName](token, options, function(data){
        const err = (typeof(data.code) !== "undefined")
          ? data.message : null
          apiCallback(err, (err) ? null : data) 
      })
      break;
  }
}