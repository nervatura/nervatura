/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var out = require('./tools.js').DataOutput();

module.exports = function(conf, lang, init) {
  switch (conf.storage_type) {
    case "loki":
      return lokiStore(conf, lang, init);
      break;
    case "level":
      return levelStore(conf, lang, init);
      break;
    default:
      return lokiStore(conf, lang, init);
      break;}}

function lokiStore(conf, lang, init) {
  var LokiJS = require('lokijs');
  var self = this;
  var store = new LokiJS('./storage/data.json');
  if (typeof init === "undefined"){
    init = function(err, storage){}}
  store.loadDatabase({}, function (err) {
    if (store.collections.length === 0){
      var index;
      if (!store.getCollection("dbs")){
        var dbs = store.addCollection("dbs",{unique:"key"});
        for (index = 0; index < conf.databases.length; index++) {
          dbs.insert({key:conf.databases[index].alias, 
            alias:conf.databases[index].alias, engine:conf.databases[index].engine,
            connect:out.cryptedValue(conf.databases[index].alias, JSON.stringify(conf.databases[index].connect), "hex"), 
            settings:conf.databases[index].settings, time_stamp:new Date().toISOString()});}}
      if (!store.getCollection("user")){
        var user = store.addCollection("user",{unique:"key"});
        for (index = 0; index < conf.users.length; index++) {
          user.insert({key:conf.users[index].username, username:conf.users[index].username,
            password:get_md5_password(conf.default_password)});}}
      if (!store.getCollection("setting")){
        var setting = store.addCollection("setting",{unique:"key"});
        for (index = 0; index < conf.settings.length; index++) {
          setting.insert({key:conf.settings[index].fieldname, fieldname:conf.settings[index].fieldname,
            value:conf.settings[index].value, description:conf.settings[index].description});}}
      store.saveDatabase(function(err){
        init(null, self);});}
    else {init(null, self);}});

function get_groups(gtype){
  return JSON.parse(JSON.stringify(store.getCollection(gtype).data));}  
  
function get_md5_password(password){
  if (typeof password === "undefined" || password === "" || password === null) {
    password = "";}
  else {password = out.createHash(password,"hex");}
  return password;}

function get_users(callback){
  var users = [];
  store.getCollection("user").data.forEach(function(user) {
    users.push(user.username);});
  return callback(null,users);}
  
function get_user(params, callback){
  var _user;
  try {
    if (params.hasOwnProperty("id")){
      _user = store.getCollection("user").get(params.id);}
    else {
      _user = store.getCollection("user").by("key",params.username);}} 
  catch (error) {
    return callback(null, null);}
  
  if (!_user) {
    return callback(null, false, {message:lang.unknown_user, username:params.username});}
  else {
    var user = {id:_user.$loki, username:_user.username};
    if (params.hasOwnProperty("password")){
      if (_user.password != get_md5_password(params.password)) {
        return callback(null, false, {message:lang.wrong_password, username:params.username});}
      else {return callback(null, user);}}
    else {
      return callback(null, user);}}}

function update_data(collection, key, values, callback){
  var data = collection.by("key",key);
  if (!data) {
    var _data = {"key":key, doc:{}};
    data = collection.insert(_data);}
  for (var fieldname in values) {
    switch (fieldname) {
      case "password":
        if (!values.hasOwnProperty(fieldname) || values[fieldname] === "" 
          || values[fieldname] === null){
          values[fieldname] = conf.default_password;}
          data[fieldname] = get_md5_password(values[fieldname]);
        break;
      default:
        if (!values.hasOwnProperty(fieldname) || values[fieldname] === null){
          values[fieldname] = "";}
        data[fieldname] = values[fieldname];
        break;}}
  store.saveDatabase(function(err){
    if(err) {
      return callback(null, lang.update_error);}
    else {
      return callback(null, lang.update_ok);}});}

function get_dbs_doc(values){
  var vdata = {connect:{}, settings:{}};
  for (var fieldname in values) {
    switch (fieldname) {
      case "ndi_enabled":
      case "ndi_md5_password":
      case "ndi_encrypt_data":
        if (values[fieldname]==="on"){
          vdata.settings[fieldname] = true;}
        else {
          vdata.settings[fieldname] = false;}
        break;
      case "ndi_encrypt_password":
      case "ndi_host_restriction":
        vdata.settings[fieldname] = values[fieldname];
        break;
      case "port":
        vdata.connect[fieldname] = parseInt(values[fieldname],10);
        if (isNaN(vdata.connect[fieldname])){
          vdata.connect[fieldname] = 0;}
        break;
      case "host":
      case "dbname":
      case "username":
      case "password":
        vdata.connect[fieldname] = values[fieldname];
        break;
      default:
        vdata[fieldname] = values[fieldname];
        break;}}
  vdata.connect = out.cryptedValue(values.alias, JSON.stringify(vdata.connect), "hex");
  vdata.time_stamp = new Date().toISOString();
  return vdata;}

function delete_data(collection, key, callback){
  var data = collection.by("key",key);
  if (!data) {
    if (collection.name === "user"){
      return callback(null,lang.unknown_user);}
    else {
      return callback(null, null);}}
  else {
    if (collection.name === "user" && collection.data.length === 1){
      return callback(null, lang.user_delete_err);}
    else {
      collection.remove(data);
      store.saveDatabase(function(err){
        if(err) {
          return callback(null, lang.update_error);}
        else {
          return callback(null, lang.data_deleted);}});}}}
              
function change_password(username, old_passw, new_passw, ver_passw, callback){
  var data = store.getCollection("user").by("key",username);
  if (!data) {
    return callback(null, lang.unknown_user);}
  else {
    if (data.password != get_md5_password(old_passw)) {
      return callback(null, lang.wrong_password);}
    else if (new_passw === null || new_passw === "") {
      return callback(null, lang.empty_password);}
    else if (new_passw !== ver_passw) {
      return callback(null, lang.verify_password);}
    else {
      data.password = get_md5_password(new_passw);
      store.saveDatabase(function(err){
        if(err) {
          return callback(null, lang.update_error);}
        else {
          return callback(null, lang.changed_password);}});}}}
    
return {
  getDataStore: function(){
    return store;},
  getUsers: function(callback){
    return get_users(callback);},
  getSettings: function(callback){
    return callback(null, get_groups("setting"));},
  getDatabases: function(callback){
    return callback(null, get_groups("dbs"));},  
  
  newUser: function(values, callback){
    if (!values.hasOwnProperty("username") || values.username === "" 
      || values.username === null){
      return callback(null, lang.invalid_fieldname);}
    else {
      return update_data(store.getCollection("user"), values.username, values, callback);}},
  changePassword: function(username, old_passw, new_passw, ver_passw, callback){
    return change_password(username, old_passw, new_passw, ver_passw, callback);},
  updateDatabase: function(values, callback){
    if (!values.hasOwnProperty("alias") || values.alias === "" || values.alias === null){
      return callback(null, lang.invalid_fieldname);}
    else {
      return update_data(store.getCollection("dbs"), values.alias, get_dbs_doc(values), callback);}},
  updateSetting: function(values, callback){
    if (!values.hasOwnProperty("fieldname") || values.fieldname === "" 
      || values.fieldname === null){
      return callback(null, lang.invalid_fieldname);}
    else {
      return update_data(store.getCollection("setting"), values.fieldname, values, callback);}},
          
  deleteUser: function(username, callback){
    return delete_data(store.getCollection("user"), username, callback);},
  deleteSetting: function(fieldname, callback){
    return delete_data(store.getCollection("setting"), fieldname, callback);},
  deleteDatabase: function(alias, callback){
    return delete_data(store.getCollection("dbs"), alias, callback);},
      
  getUserFromName: function(username, password, callback){
    get_user({username:username, password:password}, callback);},
  getUserFromId: function(id, callback){
    get_user({id:id}, callback);},
  getSettingFromFieldname: function(fieldname, callback){
    return callback(null, store.getCollection("setting").by("key",fieldname));},
  getDbsFromAlias: function(alias, callback){
    var _data = store.getCollection("dbs").by("key",alias);
    if (!_data || !alias) {
      return callback("missing",null);}
    else {
      var data = JSON.parse(JSON.stringify(_data));
      data.connect = JSON.parse(out.decipherValue(alias, data.connect, "hex"));
      return callback(null, data);}}}}

function levelStore(conf, lang, init) {
  var PouchDB = require('pouchdb');
  var self = this;
  var store = new PouchDB('./storage/data');
  if (typeof init === "undefined"){
    init = function(err, storage){}}
  store.info(function(err, info) {
    if (!err) {
      if (info.doc_count === 0){
        var items = []; var index;
        for (index = 0; index < conf.databases.length; index++) {
          items.push({_id:"dbs_"+conf.databases[index].alias, alias:conf.databases[index].alias, 
            engine:conf.databases[index].engine,
            connect:out.cryptedValue(conf.databases[index].alias, JSON.stringify(conf.databases[index].connect), "hex"), 
            settings:conf.databases[index].settings, 
            time_stamp:new Date().toISOString()});}
        for (index = 0; index < conf.users.length; index++) {
          items.push({_id:"user_"+conf.users[index].username, username:conf.users[index].username,
            password:get_md5_password(conf.default_password)});}
        for (index = 0; index < conf.settings.length; index++) {
          items.push({_id:"setting_"+conf.settings[index].fieldname, fieldname:conf.settings[index].fieldname,
            value:conf.settings[index].value, description:conf.settings[index].description});}
        store.bulkDocs(items, function(err, response){init(err, null);});}
        else {init(null, self);}}
    else{init(err, null);}});
  
function get_md5_password(password){
  if (typeof password === "undefined" || password === "" || password === null) {
    password = "";}
  else {password = out.createHash(password,"hex");}
  return password;}

function get_groups(gtype, callback){
  store.allDocs({include_docs: true, attachments: true, startkey: gtype, endkey: gtype+'\uffff'}, 
    function(err, data) {
      var rows = [];
      data.rows.forEach(function(row) {
        rows.push(row.doc);});
      callback(err,rows);});}

function get_users(callback){
    get_groups("user",function(err, data){
      if (err) {return callback(err,null);}
      var users = [];
      data.forEach(function(user) {
        users.push(user.username);});
      return callback(err,users);});}

function get_from_key(key, callback){
  store.get(key, function (err, data) {return callback(err, data);});}
  
function get_user(username, password, callback){
  store.get(username, function (err, data) {
    if (err) {
      if (err.status === 404) {
        return callback(null, false, {message:lang.unknown_user, username:username.split("_")[1]});}
      return callback(err);}
    var user = {id:data._id, username:data.username};
    if (password !== null){
      if (data.password != get_md5_password(password)) {
        return callback(null, false, {message:lang.wrong_password, username:username.split("_")[1]});}
      else {return callback(null, user);}}
    else {
      return callback(null, user);}});}

function update_data(key, values, callback){
  store.get(key, function (err, data) {
    if (err) {
      if (err.status === 404) {
        data = {}; values._id = key;}}
    for (var fieldname in values) {
      switch (fieldname) {
        case "password":
          if (!values.hasOwnProperty(fieldname) || values[fieldname] === "" 
            || values[fieldname] === null){
            values[fieldname] = conf.default_password;}
            data[fieldname] = get_md5_password(values[fieldname]);
          break;
        default:
          if (!values.hasOwnProperty(fieldname) || values[fieldname] === null){
            values[fieldname] = "";}
          data[fieldname] = values[fieldname];
          break;}}
    store.put(data, function(err, response) {
      if (err) {return callback(err, null);}
      else if (response.ok) {
        return callback(null, lang.update_ok);}
      else {
        return callback(null, lang.update_error);}});});}

function get_dbs_doc(values){
  var vdata = {connect:{}, settings:{}};
  for (var fieldname in values) {
    switch (fieldname) {
      case "ndi_enabled":
      case "ndi_md5_password":
      case "ndi_encrypt_data":
        if (values[fieldname]==="on"){
          vdata.settings[fieldname] = true;}
        else {
          vdata.settings[fieldname] = false;}
        break;
      case "ndi_encrypt_password":
      case "ndi_host_restriction":
        vdata.settings[fieldname] = values[fieldname];
        break;
      case "port":
        vdata.connect[fieldname] = parseInt(values[fieldname],10);
        if (isNaN(vdata.connect[fieldname])){
          vdata.connect[fieldname] = 0;}
        break;
      case "host":
      case "dbname":
      case "username":
      case "password":
        vdata.connect[fieldname] = values[fieldname];
        break;
      default:
        vdata[fieldname] = values[fieldname];
        break;}}
  vdata.connect = out.cryptedValue(values.alias, JSON.stringify(vdata.connect), "hex");
  vdata.time_stamp = new Date().toISOString();
  return vdata;}
      
function delete_user(username, callback){
  get_groups("user",function(err, data){
    if (err) {return callback(err,null);}
    if (data.length === 1){
      return callback(null, lang.user_delete_err);}
    var user = null;
    data.forEach(function(doc) {
      if (doc.username === username){
        user = doc;}});
    if (user !== null){
      store.remove(user._id, user._rev, function(err, response) {
        if (err) {return callback(err, err);}
        else if (response.ok) {
          return callback(null, lang.data_deleted);}
        else {
          return callback(null, lang.update_error);}});}
    else {
      return callback(null,lang.unknown_user);}});}

function delete_data(key, callback){
  get_from_key(key,function(err, data){
    if (err) {
      if (err.status === 404) {
        return callback(null, null);}
      return callback(err, err);}
    store.remove(data._id, data._rev, function(err, response) {
      if (err) {return callback(err, err);}
      else if (response.ok) {
        return callback(null, lang.data_deleted);}
      else {
        return callback(null, lang.update_error);}});});}
              
function change_password(username, old_passw, new_passw, ver_passw, callback){
  store.get(username, function (err, data) {
    if (err) {
      if (err.status === 404) {
        return callback(null, lang.unknown_user);}
      return callback(err, err);}
    if (data.password != get_md5_password(old_passw)) {
      return callback(null, lang.wrong_password);}
    else if (new_passw === null || new_passw === "") {
      return callback(null, lang.empty_password);}
    else if (new_passw !== ver_passw) {
      return callback(null, lang.verify_password);}
    else {
      data.password = get_md5_password(new_passw);
      store.put(data, function(err, response) {
        if (err) {return callback(err, err);}
        else if (response.ok) {
          return callback(null, lang.changed_password);}
        else {
          return callback(null, lang.update_error);}});}});}
    
return {
  getDataStore: function(){
    return store;},
  getUsers: function(callback){
    return get_users(callback);},
  getSettings: function(callback){
    return get_groups("setting",callback);},
  getDatabases: function(callback){
    return get_groups("dbs",callback);},  
  
  newUser: function(values, callback){
    if (typeof values.username === "undefined" || values.username === "" 
      || values.username === null){
      return callback(null, lang.invalid_fieldname);}
    else {
      return update_data("user_"+values.username, values, callback);}},
  changePassword: function(username, old_passw, new_passw, ver_passw, callback){
    return change_password("user_"+username, old_passw, new_passw, ver_passw, callback);},
  updateDatabase: function(values, callback){
    if (typeof values.alias === "undefined" || values.alias === "" || values.alias === null){
      return callback(null, lang.invalid_fieldname);}
    else {
      return update_data("dbs_"+values.alias, get_dbs_doc(values), callback);}},
  updateSetting: function(values, callback){
    if (typeof values.fieldname === "undefined" || values.fieldname === "" 
      || values.fieldname === null){
      return callback(null, lang.invalid_fieldname);}
    else {
      return update_data("setting_"+values.fieldname, values, callback);}},
          
  deleteUser: function(username, callback){
    return delete_user(username, callback);},
  deleteSetting: function(fieldname, callback){
    return delete_data("setting_"+fieldname, callback);},
  deleteDatabase: function(alias, callback){
    return delete_data("dbs_"+alias, callback);},
      
  getUserFromName: function(username, password, callback){
    get_user("user_"+username, password, callback);},
  getUserFromId: function(id, callback){
    get_user(id, null, callback);},
  getSettingFromFieldname: function(fieldname, callback){
    store.get("setting_"+fieldname, function (err, data) {
      if (err) {return callback(err,null);}
      else {return callback(err, data);}});},
  getDbsFromAlias: function(alias, callback){
    store.get("dbs_"+alias, function (err, data) {
      if (err) {return callback(err,null);}
      else {
        data.connect = JSON.parse(out.decipherValue(alias, data.connect, "hex"));
        return callback(err, data);}});}}}           
