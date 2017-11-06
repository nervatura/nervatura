/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var express = require('express');
var router = express.Router();

var nas = require('../lib/node/nas.js')();
var lang; var validator;

router.use(function (req, res, next) {
  validator = nas.validUser(req);
  lang = req.app.locals.lang;
  next()});

router.get('/', function(req, res, next) {
  res.redirect('nas/index');});

router.get('/index', function(req, res, next) {
  if (validator === "ok"){
    nas.pageRender({res:res, page:"index", 
      data:{subtitle:lang.title_home, flash:null, user:req.user}});}
  else {
    res.redirect(validator);}});  

router.get('/user/list', function(req, res, next) {
  if (validator === "ok"){
    nas.userList({res:res, req:req, next:next, data:{flash:null}});}
  else {
    res.redirect(validator);}});

router.post('/user/list', function(req, res, next) {
  if (validator === "ok"){
    if (req.body.update_cmd === "new"){
      req.app.settings.storage.newUser({username:req.body.username, password:req.body.password},
      function(err, message){
        if (err) {return next(err);}
        else {
          nas.userList({res:res, req:req, next:next, data:{flash:message}});}});}
    else if (req.body.update_cmd === "delete"){
      req.app.settings.storage.deleteUser(req.body.username,
      function(err, message){
        if (err) {return next(err);}
        else {
          nas.userList({res:res, req:req, next:next, data:{flash:message}});}});}
    else {
      nas.userList({res:res, req:req, next:next, data:{flash:null}});}}
  else {
    res.redirect(validator);}});

router.get('/user/password', function(req, res, next) {
  if (validator === "ok"){
    var flash = null;
    if(req.user.dirty_password){
      flash = lang.dirty_password;}
    nas.pageRender({res:res, page:"password", 
      data:{subtitle:lang.label_change_password, flash:flash}});}
  else {
    res.redirect(validator);}});

router.post('/user/password', function(req, res, next) {
  if (validator === "ok"){
    if(req.body.old === ""){
      req.body.old = "empty";}
    req.app.settings.storage.changePassword(req.user.username, req.body.old, req.body.new, req.body.verify, 
      function(err, message){
        if (err) {return next(err);}
        else {
          nas.pageRender({res:res, page:"password", 
            data:{subtitle:lang.label_change_password, flash:message}});}});}
  else {
    res.redirect(validator);}});

router.get('/database/list', function(req, res, next) {
  if (validator === "ok"){
    nas.databaseList({res:res, req:req, next:next, data:{flash:null}});}
  else {
    res.redirect(validator);}});

router.get('/database/edit', function(req, res, next) {
  if (validator === "ok"){
    var params = {res:res, req:req, page:"database", data:{
      flash:null, subtitle:lang.label_databases, view:"edit",
      engines:req.app.settings.conf.engine}};
    if (req.query.alias === ""){
      params.data.data = {alias:"", engine:"sqlite", 
        connect:{host:"", port:"", dbname:"", username:"", password:""},
        settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}};
      nas.pageRender(params);}
    else {
      req.app.settings.storage.getDbsFromAlias(req.query.alias,
        function(err, data){
          if (err) {return next(err);}
          else {
            params.data.data = data;
            nas.pageRender(params);}});}}
  else {
    res.redirect(validator);}});

router.post('/database/list', function(req, res, next) {
  if (validator === "ok"){
    if (req.body.update_cmd === "update"){
      req.app.settings.storage.updateDatabase(req.body, function(err, message){
        if (err) {return next(err);}
        else {
          nas.databaseList({res:res, req:req, next:next, data:{flash:message}});}});}
    else if (req.body.update_cmd === "delete"){
      req.app.settings.storage.deleteDatabase(req.body.alias, function(err, message){
        if (err) {return next(err);}
        else {
          nas.databaseList({res:res, req:req, next:next, data:{flash:message}});}});}
    else {
      nas.databaseList({res:res, req:req, next:next, data:{flash:null}});}}
  else {
    res.redirect(validator);}});
    
router.get('/database/create', function(req, res, next) {
  if (validator === "ok"){
    nas.databaseList({res:res, req:req, next:next,
      data:{view:"create", subtitle:lang.label_creation, form:{}, flash:null}});}
  else {
    res.redirect(validator);}});

router.post('/database/create', function(req, res, next) {
  if (validator === "ok"){
    req.setTimeout(req.app.settings.conf.long_timeout);
    var nstore = require('../lib/node/nervastore.js')(req, res);
    nas.createDatabase(nstore, {database:req.body.alias}, function(err, logstr){
      var form = req.body; form.message = logstr;
      nas.databaseList({res:res, req:req, next:next,
      data:{view:"create", subtitle:lang.label_creation, form:form, flash:null}});});}
  else {
    res.redirect(validator);}});

router.get('/database/export', function(req, res, next) {
  if (validator === "ok"){
    nas.databaseList({res:res, req:req, next:next,
      data:{view:"export", subtitle:lang.label_export, form:{}, flash:null}});}
  else {
    res.redirect(validator);}});

router.post('/database/export', function(req, res, next) {
  if (validator === "ok"){
     req.setTimeout(req.app.settings.conf.long_timeout);
    var nstore = require('../lib/node/nervastore.js')(req, res);
    nas.exportDatabase(nstore, 
      {database:req.body.alias, filename:req.body.filename, format:req.body.format,
       version: req.app.settings.version_number, export_dir:req.app.get('data_dir')+'/data'}, 
      function(err, logstr, result){
        if(err || req.body.filename !== "download"){
          var form = req.body; form.message = logstr;
          nas.databaseList({res:res, req:req, next:next,
            data:{view:"export", subtitle:lang.label_export, 
              form:form, flash:null}});}
        else {
          res.send(result);}});}
  else {
    res.redirect(validator);}});

router.get('/database/import', function(req, res, next) {
  if (validator === "ok"){
    nas.importList({res:res, req:req, next:next,
      data:{view:"import", import_dir:req.app.get('data_dir')+'/data',
        subtitle:lang.label_import, form:{}, filenames:[], flash:null}});}
  else {
    res.redirect(validator);}});

router.post('/database/import', function(req, res, next) {
  if (validator === "ok"){
    req.setTimeout(req.app.settings.conf.long_timeout);
    var nstore = require('../lib/node/nervastore.js')(req, res);
    nas.importDatabase(nstore, 
      {database:req.body.alias, filename:req.body.filename, 
       import_dir:req.app.get('data_dir')+'/data'}, 
      function(err, logstr){
        var form = req.body; form.message = logstr;
        nas.importList({res:res, req:req, next:next,
          data:{view:"import", import_dir:req.app.get('data_dir')+'/data',
            subtitle:lang.label_import, form:form, filenames:[], flash:null}});});}
  else {
    res.redirect(validator);}});

router.get('/report', function(req, res, next) {
  if (validator === "ok"){
    nas.reportRender({res:res, storage:req.app.settings.storage, 
      data:{subtitle:lang.label_reports, form:{}, flash:null}});}
  else {
    res.redirect(validator);}});

router.post('/report', function(req, res, next) {
  if (validator === "ok"){
    var form = req.body;
    if (req.body.update_cmd !== ""){
      if (req.body.update_database !== ""){
        form.alias = req.body.update_database;}
      if (req.body.update_group !== ""){
        form.group = req.body.update_group;}}
    var nstore = require('../lib/node/nervastore.js')(req, res);
    if (form.alias !== ""){
      nstore.local.setEngine({database:form.alias}, function(err, result){
        form.engine = result;
        if (err){
          nas.reportList(nstore, form, err);}
        else {
          if (req.body.update_cmd === "install"){
            nas.reportInstall(nstore, form);}
          else if (req.body.update_cmd === "delete"){
            nas.reportDelete(nstore, {reportkey:form.update_reportkey},
              function(err){nas.reportList(nstore, form, err);});}
          else if (req.body.update_cmd === "edit"){
            nas.reportEdit(nstore, form);}
          else if (req.body.update_cmd === "update"){
            nas.reportUpdate(nstore, form, 
              function(err){nas.reportList(nstore, form, err);});}
          else {
            nas.reportList(nstore, form, null);}}});}
    else {
      nas.reportList(nstore, form, null);}}
  else {
    res.redirect(validator);}});
    
router.get('/setting/list', function(req, res, next) {
  if (validator === "ok"){
    nas.settingList({res:res, req:req, next:next, data:{flash:null}});}
  else {
    res.redirect(validator);}});

router.post('/setting/update', function(req, res, next) {
  if (validator === "ok"){
    var setting_err = nas.validSetting(req.app.settings.conf, req.body, lang)
    if(setting_err === null){
      req.app.settings.storage.updateSetting(req.body, function(err, message){
        if (err) {return next(err);}
        else {
          req.app.get('host_settings')[req.body.fieldname] = req.body.value
          nas.settingList({res:res, req:req, next:next, data:{flash:message}});}});}
    else{
      nas.settingList({res:res, req:req, next:next, data:{flash:setting_err}});}}
  else {
    res.redirect(validator);}});
    
router.post('/setting/delete', function(req, res, next) {
  if (validator === "ok"){
    req.app.settings.storage.deleteSetting(req.body.fieldname,
    function(err, message){
      if (err) {return next(err);}
      else {
        nas.settingList({res:res, req:req, next:next, data:{flash:message}});}});}
  else {
    res.redirect(validator);}});
 
module.exports = router;