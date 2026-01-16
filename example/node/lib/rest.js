const http = require('http');
var utils = require('./utils');

function httpApi(params, callback){
  const options = {
    protocol: "http:", hostname: "localhost", port: process.env.NT_HTTP_PORT,
    path: "/api/v6"+params.path, method: params.method,
    headers: {
      "Content-Type": "application/json",
    }
  };
  const post_data = (params.data) ? utils.EncodeOptions(params.data) : null
  if(params.token){
    options.headers["Authorization"] = "Bearer "+params.token
  }
  if(process.env.NT_API_KEY){
    options.headers["X-API-Key"] = process.env.NT_API_KEY
  }
  if(post_data){
    options.headers["Content-Length"] = post_data.length
  }
  const hreq = http.request(options, (hres) => {
    let data = "";

    hres.on("data", (chunk) => {
        data += chunk;
    });

    hres.on("end", () => {
      callback(utils.CheckJson(data))
    });

  }).on("error", (err) => {
    callback({code:400, message: err.message})
  });
  
  if(post_data){
    hreq.write(post_data);
  }
  hreq.end();
}

exports.Get = function(token, path, query, callback) {
  var path = "/"+path;
  var query_str = "";
  for(var key in query){
    query_str += "&"+key+"="+query[key];
  }
  if(query_str){
    path += "?"+encodeURIComponent(query_str);
  }
  httpApi({ path, method: "GET", token }, function(data){
    callback(data)
  })
}

exports.Post = function(token, path, post_data, callback) {
  var path = "/"+path;
  httpApi({ path, method: "POST", token, data: post_data }, function(data){
    callback(data)
  })
}

exports.Put = function(token, path, post_data, callback) {
  var path = "/"+path;
  httpApi({ path, method: "PUT", token, data: post_data }, function(data){
    callback(data)
  })
}

exports.Delete = function(token, path, callback) {
  var path = "/"+path;
  httpApi({ path, method: "DELETE", token }, function(data){
    callback(data)
  })
}