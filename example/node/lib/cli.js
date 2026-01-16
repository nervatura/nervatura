var path = require('path');
const execFile = require('child_process').execFile;

var utils = require('./utils');

function cliApi(_args, callback){
  const child = execFile(path.join(utils.servicePath, utils.serviceFile), _args, {env: process.env}, 
    (error, stdout, stderr) => {
    if (error) {
      return callback(stderr, null)
    }
    const result = utils.CheckJson(stdout.toString().split("\n")[stdout.toString().split("\n").length-2], true)
    if((typeof result === "object") && result.code){
      if(result.code != 200 && result.code != 201 && result.code != 204){
        return callback(result.message, null)
      }
    }
    callback(null, result)
  });
}

exports.Database = function(options, callback) {
  cliApi(['-c', 'database', '-o', utils.EncodeOptions(options)], function(err, data){
    callback(err, data)
  })
}

exports.ResetPassword = function(options, callback) {
  cliApi(['-c', 'reset', '-o', utils.EncodeOptions(options)], function(err, data){
    callback(err, data)
  })
}

exports.Create = function(model, options, data, callback) {
  cliApi(['-c', 'create', '-m', model, '-o', utils.EncodeOptions(options), '-d', utils.EncodeOptions(data)], function(err, data){
    callback(err, data)
  })
}

exports.Update = function(model, options, data, callback) {
  cliApi(['-c', 'update', '-m', model, '-o', utils.EncodeOptions(options), '-d', utils.EncodeOptions(data)], function(err, data){
    callback(err, data)
  })
}

exports.Delete = function(model, options, callback) {
  cliApi(['-c', 'delete', '-m', model, '-o', utils.EncodeOptions(options)], function(err, data){
    callback(err, data)
  })
}

exports.Get = function(model, options, callback) {
  cliApi(['-c', 'get', '-m', model, '-o', utils.EncodeOptions(options)], function(err, data){
    callback(err, data)
  })
}

exports.Query = function(model, options, callback) {
  cliApi(['-c', 'query', '-m', model, '-o', utils.EncodeOptions(options)], function(err, data){
    callback(err, data)
  })
}

exports.View = function(options, callback) {
  cliApi(['-c', 'view', '-o', utils.EncodeOptions(options)], function(err, data){
    callback(err, data)
  })
}

exports.Function = function(options, callback) {
  cliApi(['-c', 'function', '-o', utils.EncodeOptions(options)], function(err, data){
    callback(err, data)
  })
}