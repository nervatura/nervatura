const fs = require('fs')
var path = require('path');
var async = require("async");

if(!fs.existsSync(path.join(__dirname, ".env"))){
  fs.copyFileSync(path.join(__dirname, ".env.example"),path.join(__dirname, ".env"))
}
require('dotenv').config()

var utils = require('./lib/utils');

var api_map = {
  http: require('./lib/rest'),
  cli: require('./lib/cli'),
  rpc: require('./lib/rpc')
}

fn_map = {
  cli: {
    Database: [{ alias: "demo", demo: true }],
    Function: [{name: "test", values: {}}],
    ResetPassword: [{alias: "demo", code: "USR0000000000N1"}],
    Create: ["customer", {alias: "demo"}, {code: "CLI0000000000N1", customer_name: "Node.js Test"}],
    Update: ["customer", {alias: "demo", code: "CLI0000000000N1"}, {customer_meta: {account: "1234567890"}}],
    Get: ["customer", {alias: "demo", code: "CLI0000000000N1"}],
    Query: ["customer", {alias: "demo", customer_type: "CUSTOMER_COMPANY"}],
    View: [{alias: "demo", name: "VIEW_CUSTOMER_EVENTS", filters:[{field:"like_subject", value:"visit"}, {field:"place", value:"City1"}], limit:10}],
    Delete: ["customer", {alias: "demo", code: "CLI0000000000N1"}],
  },
  http: {
    Post: ["", "customer", {code: "REST0000000000N1", customer_name: "Node Test"}],
    Put: ["", "customer/REST0000000000N1", {customer_name: "Test Customer"}],
    Get: ["", "customer", {"customer_type": "CUSTOMER_COMPANY"}],
    Delete: ["", "customer/REST0000000000N1"],
  },
  rpc: {
    //Database: [{alias: "demo", demo: true }],
    CustomerUpdate: ["", {code: "RPC0000000000N1", customer_name: "Node Test"}],
    CustomerGet: ["", {code: "RPC0000000000N1"}],
    CustomerQuery: ["", {customer_type: "CUSTOMER_COMPANY"}],
    Delete: ["", {code: "RPC0000000000N1", model: "customer"}],
    Function: ["", {function: "product_price", args: {product_code:"PRD0000000000N1", currency_code:"EUR", price_type:"PRICE_CUSTOMER"}}],
    View: ["", {name: "VIEW_CUSTOMER_EVENTS", filters:[{field:"like_subject", value:"visit"}, {field:"place", value:"City1"}], limit:10}],
  }
}

async.waterfall([
  function(callback) {
    const start_time = Date.now()
    const api_type = "cli"
    var fn_lst = [];
    Object.keys(fn_map[api_type]).forEach(func_name => {
      fn_lst.push(function(callback_){
        api_map[api_type][func_name](...fn_map[api_type][func_name], function(err, result){
          if(err){
            console.log(api_type+" "+func_name+" error: "+err)
          } else {
            console.log(api_type+" "+func_name+" OK")
          }
          callback_()
        })
      });
    });            
    async.series(fn_lst,function() {
      console.log("--------------------")
      console.log(`${api_type} time ${Math.floor(Date.now() - start_time)}`);
      console.log("--------------------")
      callback(null)
    });
  },

  function(callback) {
    utils.StartService(false, (err, controller)=>{
      callback(err, controller)
    })
  },

  function(controller, callback) {
    const start_time = Date.now()
    const api_type = "rpc"
    var fn_lst = [];
    Object.keys(fn_map[api_type]).forEach(func_name => {
      fn_lst.push(function(callback_){
        api_map[api_type][func_name](...fn_map[api_type][func_name], function(err, result){
          if(err){
            console.log(api_type+" "+func_name+" error: "+err)
          } else {
            console.log(api_type+" "+func_name+" OK")
          }
          callback_()
        })
      });
    });            
    async.series(fn_lst,function() {
      console.log("--------------------")
      console.log(`${api_type} time ${Math.floor(Date.now() - start_time)}`);
      console.log("--------------------")
      callback(null, controller)
    });
  },

  function(controller, callback) {
    const start_time = Date.now()
    const api_type = "http"
    var fn_lst = [];
    Object.keys(fn_map[api_type]).forEach(func_name => {
      fn_lst.push(function(callback_){
        api_map[api_type][func_name](...fn_map[api_type][func_name], function(data){
          const err = (typeof(data.code) !== "undefined")
            ? data.message : null
          if(err){
            console.log(api_type+" "+func_name+" error: "+err)
          } else {
            console.log(api_type+" "+func_name+" OK")
          }
          callback_()
        })
      });
    });            
    async.series(fn_lst,function() {
      console.log("--------------------")
      console.log(`${api_type} time ${Math.floor(Date.now() - start_time)}`);
      console.log("--------------------")
      callback(null, controller)
    });
  }

],  
function(err, controller) {
  if(err){
    console.log(err)
  }
  if(controller){
    controller.abort()
  }
})