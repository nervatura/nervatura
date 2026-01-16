var utils = require('./utils');

var path = require('path');
const protoLoader = require('@grpc/proto-loader');
const grpc = require('@grpc/grpc-js');

const protoPath = path.join("node_modules","nervatura","bin","proto","store.proto")

//const sslCreds = grpc.credentials.createSsl(sslCert);
const insecureCreds = grpc.credentials.createInsecure()

const pb = protoLoader.loadSync(protoPath, {
  keepCase: true, longs: Number, enums: String, defaults: false, oneofs: true
});
const nervatura = grpc.loadPackageDefinition(pb).nervatura

function rpcClient(ckey, meta, options, callback){
  const client = new nervatura.API(`localhost:${process.env.NT_GRPC_PORT}`, insecureCreds);
  client[ckey](options, meta, function(err, data){
    if(err && err.code === 14){
      client[ckey](options, meta, function(err, data){
        callback(err, data)
      })
    } else {
      callback(err, data)
    }
  });
}

exports.Database = function(options, callback) {
  var meta = new grpc.Metadata();
  meta.add('x-api-key', process.env.NT_API_KEY);
  rpcClient("Database", meta, options, function(err, data){
    if(err !== null){
      return callback(err.details, null)
    }
    var results = []
    if(data && data.data){
      results = utils.CheckJson(data.data)
    }
    callback(null, results)
  })
}

exports.CustomerUpdate = function(token, options, callback) {
  var meta = new grpc.Metadata();
  meta.add('x-api-key', process.env.NT_API_KEY);
  if(token){
    meta.add('Authorization', `Bearer ${token}`);
  }
  rpcClient("CustomerUpdate", meta, options, function(err, data){
    if(err !== null){
      return callback(err.details, null)
    }
    callback(null, data)
  })
}

exports.CustomerGet = function(token, options, callback) {
  var meta = new grpc.Metadata();
  meta.add('x-api-key', process.env.NT_API_KEY);
  if(token){
    meta.add('Authorization', `Bearer ${token}`);
  }
  rpcClient("CustomerGet", meta, options, function(err, data){
    if(err !== null){
      return callback(err.details, null)
    }
    callback(null, data)
  })
}

exports.CustomerQuery = function(token, options, callback) {
  var meta = new grpc.Metadata();
  meta.add('x-api-key', process.env.NT_API_KEY);
  if(token){
    meta.add('Authorization', `Bearer ${token}`);
  }
  rpcClient("CustomerQuery", meta, options, function(err, data){
    if(err !== null){
      return callback(err.details, null)
    }
    callback(null, data.data)
  })
}

exports.Delete = function(token, options, callback) {
  var meta = new grpc.Metadata();
  meta.add('x-api-key', process.env.NT_API_KEY);
  if(token){
    meta.add('Authorization', `Bearer ${token}`);
  }
  nt_model = nervatura.Model.type.value.filter(dtype => (dtype.name === options.model.toUpperCase()))
  if (nt_model.length > 0){
    options.model = nt_model[0].name
  }
  rpcClient("Delete", meta, options, function(err, data){
    if(err !== null){
      return callback(err.details, null)
    }
    callback(null, {})
  })
}

exports.Function = function(token, options, callback) {
  var meta = new grpc.Metadata();
  meta.add('x-api-key', process.env.NT_API_KEY);
  if(token){
    meta.add('Authorization', `Bearer ${token}`);
  }
  rpcClient("Function", meta, options, function(err, data){
    if(err !== null){
      return callback(err.details, null)
    }
    var results = {};
    if(data && data.data){
      results = utils.CheckJson(data.data)
    }
    callback(null, results)
  })
}

exports.View = function(token, options, callback) {
  var meta = new grpc.Metadata();
  meta.add('x-api-key', process.env.NT_API_KEY);
  if(token){
    meta.add('Authorization', `Bearer ${token}`);
  }
  rpcClient("View", meta, options, function(err, data){
    if(err !== null){
      return callback(err.details, null)
    }
    var results = [];
    if(data && data.data){
      results = utils.CheckJson(data.data)
    }
    callback(null, results)
  })
}