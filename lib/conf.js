/*
This file is part of the Nervatura Framework
http://nervatura.com
Copyright © 2011-2019, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var path = require('path');
var util = require('./tools.js').DataOutput()

module.exports = function(_settings) {
  var self = this;
  self.init_settings = _settings;

  function check_value(key, defvalue){
    key = String(key);
    if(key.indexOf("$")>-1){
      key = key.substr(key.indexOf("$")+1);}
    var value = (typeof(defvalue) === "undefined") ? "" : defvalue;
    if(typeof(process.env[key]) !== "undefined"){
      value = process.env[key];}
    else if(typeof(self.init_settings[key]) !== "undefined"){
      value = self.init_settings[key];}
    switch (key) {
      case "REPORT_DIR":
        if((value === defvalue) || (value === "")) {
          value = path.join(__dirname, path.join("..","report")) }
        break;

      case "SESSION_SECRET":
        if((value === "") || (value === "auto")){
          value = util.createKey();}
        break;
      
      case "LONG_TIMEOUT":
      case "COOKIE_MAX_AGE":
      case "CONTENT_LENGTH":
        if(!isNaN(parseInt(value,10))){
          value = parseInt(value,10);}
        break;

      default:
        break;
    }
    switch (value) {
      case "true":
        return true;

      case "false":
        return false;

      default:
        return value;
    }
  }
  return {
    settings: [],
    users: [{username:"admin"}],
    engine: [{key:"sqlite", label:"SQLite"}, {key:"mysql", label:"MySQL"}, 
      {key:"postgres", label:"PostgreSQL"}, {key:"mssql", label:"MSSQL"}],
    pool_config: {min: 0, max: 10},

    host_type: check_value("HOST_TYPE", "localhost"),
    
start_page: check_value("START_PAGE", "ntura"), //static, custom, ntura
    lang: check_value("APP_LANG", "en"),
    check_admin: check_value("CHECK_ADMIN", false),
    admin_key: check_value("ADMIN_KEY", ""),

    data_dir: check_value("DATA_DIR", "data"),
    report_dir: check_value("REPORT_DIR", path.join("..","report")),
    python2_path: check_value("PYTHON2_PATH", "python"),
    python_script: path.join(__dirname, check_value("PYTHON_SCRIPT", path.join("..","pylib"))),
    data_store: check_value("DATA_STORE", "memory"), //lokijs,pouchdb,datastore,memory
    
    long_timeout: check_value("LONG_TIMEOUT", (4 * 60 * 1000)), //create/export/import database & demo data
    session_secret: check_value("SESSION_SECRET", ""),
    session_cookie_max_age: check_value("COOKIE_MAX_AGE", 7*24*60*60*1000),
    max_content_length: check_value("CONTENT_LENGTH", 20000),
    
    nas_login: {
      local: {
        session: false
      },
      amazon:{
        clientID: check_value("AMAZON_CLIENT_ID", null),
        clientSecret: check_value("AMAZON_CLIENT_SECRET", null),
        callbackURL: "/login/amazon/callback",
        session: false
      },
      azure:{
        clientID: check_value("AZURE_CLIENT_ID", null),
        clientSecret: check_value("AZURE_CLIENT_SECRET", null),
        redirectUrl: check_value("AZURE_REDIRECT_URL", null),
        identityMetadata: check_value("AZURE_IDENTITY_URL", null),
        responseType: "code id_token",
        responseMode: "form_post",
        allowHttpForRedirectUrl: true,
        session: false
      },
      google:{
        clientID: check_value("GOOGLE_CLIENT_ID", null),
        clientSecret: check_value("GOOGLE_CLIENT_SECRET", null),
        callbackURL: "/login/google/callback",
        session: false
      }
    },
    email_providers: {
      smtp: {
        host: check_value("SMTP_HOST", null),
        port: check_value("SMTP_PORT", null),
        secure: check_value("SMTP_SECURE", true), // use TLS
        auth: {
          user: check_value("SMTP_USER", null),
          pass: check_value("SMTP_PASSWORD", null)
        }
      },
      mailjet: {
        clientID: check_value("MJ_APIKEY_PUBLIC", null),
        clientSecret: check_value("MJ_APIKEY_PRIVATE", null)
      }
    },
    def_settings: {
      nas_auth: check_value("NAS_AUTH", "local"), //local, amazon, azure, google

      all_host_restriction: check_value("ALL_HOST_RESTRICTION", ""),
      nas_host_restriction: check_value("NAS_HOST_RESTRICTION", ""),
      npi_host_restriction: check_value("NPI_HOST_RESTRICTION", ""),
      ndi_host_restriction: check_value("NDI_HOST_RESTRICTION", "")
    },
    checkValue: function(key, defvalue){
      return check_value(key, defvalue);
    }
  }
}