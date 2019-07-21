/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2019, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var out = require('./tools.js').DataOutput();

module.exports.basicStore = function(options) {
  var store = {}

  function get_md5_password(password){
    if (typeof password === "undefined" || password === "" || password === null) {
      password = "";}
    else {password = out.createHash(password,"hex");}
    return password;}

  var result = {
    getDataStore: function(){
      return store;},
      
    getUsers: function(callback){
      return callback(null, [store.user[0].username]);},
    getSettings: function(callback){
      return callback(null, store.setting);},
    getDatabases: function(callback){
      return callback(null, store.dbs);},  
    
    newUser: function(values, callback){
      if(callback)
        return callback(null, options.lang.disabled_update);},
    changePassword: function(username, old_passw, new_passw, ver_passw, callback){
      if(callback)
        return callback(null, options.lang.disabled_update);},
    updateDatabase: function(values, callback){
      if(callback)
        return callback(null, options.lang.disabled_update);},
    updateSetting: function(values, callback){
      if (typeof values.fieldname === "undefined" || values.fieldname === "" 
        || values.fieldname === null){
          if(callback)
            return callback(null, options.lang.invalid_fieldname);}
      else {
        store.settings[values.fieldname] = values.value;
        for (let index = 0; index < store.setting.length; index++) {
          if(store.setting[index].fieldname === values.fieldname){
            store.setting[index].value = values.value;
            if(callback)
              return callback(null, options.lang.update_ok);}}
        store.setting.push(
          { id: new Date().getTime(), fieldname: values.fieldname,
            value: values.value, description: values.description}) 
        if(callback)
          return callback(null, options.lang.update_ok);}
    },
            
    deleteUser: function(username, callback){
      if(callback)
        return callback(null, options.lang.disabled_update);},
    deleteSetting: function(fieldname, callback){
      if(callback)
        return callback(null, options.lang.disabled_update);},
    deleteDatabase: function(alias, callback){
      if(callback)
        return callback(null, options.lang.disabled_update);},
      
    getUserFromName: function(username, password, callback){
      if(store.user[0].username === username){
        if (password !== null){
          if (store.user[0].password != get_md5_password(password)) {
            return callback(null, false, {message: options.lang.wrong_password, username: username});}
          else {
            return callback(null, {id:store.user[0].id, username:store.user[0].username});}}}
      else {
        return callback(null, false, {message:options.lang.unknown_user, username: username});}},
    getUserFromId: function(id, callback){
      if(store.user[0].id === id){
        return callback(null, {id:store.user[0].id, username:store.user[0].username});}
      else {
        return callback(null, false, {message:options.lang.unknown_user, id: id});}},
    getUserFromEmail: function(email, profile, token, callback){
      if(store.user[0].username === email){
        return callback(null, {id:store.user[0].id, username:store.user[0].username});}
      else {
        return callback(null, false, {message:options.lang.unknown_user, email: email});}},
    getSettingFromFieldname: function(fieldname, callback){
      if(store.settings[fieldname]){
        return callback(null, store.settings[fieldname]); }
      else {
        return callback(options.lang.missing_fieldname, null); }},
    getDbsFromAlias: function(alias, callback){
      var _db
      store.dbs.forEach(dbs => {
        if(dbs.alias === alias){
          _db = dbs;}});
      if(_db){
        return callback(null, _db);}
      else {
        return callback(options.lang.missing_database+" "+alias, null);
      }
    }
  }

  store = {
    dbs: [],
    user: [
      { id: new Date().getTime(),
        username: "admin",
        password: get_md5_password("empty") }],
    setting: [], settings: {} }
  for (const alias in options.databases) {
    store.dbs.push(
      { id: new Date().getTime(),
        alias: alias, 
        engine: options.databases[alias].engine,
        connect: { host: options.databases[alias].connect.host, 
          port: options.databases[alias].connect.port, 
          dbname: options.databases[alias].connect.dbname, 
          username: options.databases[alias].connect.username, 
          password: options.databases[alias].connect.password }, 
        settings: { db_ssl: options.databases[alias].settings.db_ssl,
          ndi_enabled: options.databases[alias].settings.ndi_enabled, 
          encrypt_password: options.databases[alias].settings.encrypt_password, 
          dbs_host_restriction: options.databases[alias].settings.dbs_host_restriction }, 
        time_stamp: new Date().toISOString() })}
  var settings = {}
  for (var index = 0; index < options.conf.settings.length; index++) {
    store.setting.push(
      { id: new Date().getTime(),
        fieldname: options.conf.settings[index].fieldname,
        value: options.conf.settings[index].value, 
        description: options.conf.settings[index].description});
    store.settings[options.conf.settings[index].fieldname] = options.conf.settings[index].value; 
  }
  if(options.callback){ options.callback(null, result, settings); }
  return result;
}    
