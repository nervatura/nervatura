/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/* global __dirname */

module.exports = function(callback){
  
var fs = require('fs');
var express = require('express');
var session = require('cookie-session')

var compression = require('compression');
var passport = require('passport');
var LocalStrategy = require('passport-local').Strategy;
var GoogleStrategy = require('passport-google-oauth20').Strategy;

var cors = require('cors');
var lusca = require('lusca');
var helmet = require('helmet');
var hpp = require('hpp');
var contentLength = require('express-content-length-validator');
var express_enforces_ssl = require('express-enforces-ssl');

var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var methodOverride = require('method-override');
var _ = require('lodash');
  
//routes
var ntura = require('./routes/ntura');
var demo = require('./routes/demo');
var npi = require('./routes/npi');
var ndi = require('./routes/ndi');
var nas = require('./routes/nas');
var report = require('./routes/report');
var wizard = require('./routes/wizard');
var custom = require('./routes/custom');

var app = express();
app.locals._ = _;
app.set('env', process.env.NODE_ENV || 'development');

var conf = require('./lib/node/conf.js');
app.set('conf', conf);

//host ip,port
app.set('host_type', process.env.HOST_TYPE || 'localhost');
app.set('port', conf.port || process.env.PORT || '8080');
if(process.env.OPENSHIFT_NODEJS_IP){
  app.set('host_type', 'openshift');
  app.set('ip', process.env.OPENSHIFT_NODEJS_IP);}
if(process.env.OPENSHIFT_NODEJS_PORT){
  app.set('host_type', 'openshift');
  app.set('port', process.env.OPENSHIFT_NODEJS_PORT);}

//data directory
if(process.env.OPENSHIFT_DATA_DIR){
  app.set('data_dir', process.env.OPENSHIFT_DATA_DIR);}
else if(process.env.NERVATURA_DATA_DIR){
  app.set('data_dir', process.env.NERVATURA_DATA_DIR);}
else{
  try {
    app.set('data_dir', conf.data_dir || 'data');
    console.log(path.join(__dirname, 'data'));
    fs.statSync(app.get('data_dir'));} 
  catch(e) {
    try {
      fs.statSync(path.join(__dirname, 'data'));} 
    catch(e) {
      fs.mkdirSync(path.join(__dirname, 'data'));}
    app.set('data_dir', path.join(__dirname, 'data'));}}
try {
  fs.statSync(path.join(app.get('data_dir'),'database'));} 
catch(e) {
  fs.mkdirSync(path.join(app.get('data_dir'),'database'));}
try {
  fs.statSync(path.join(app.get('data_dir'),'storage'));} 
catch(e) {
  fs.mkdirSync(path.join(app.get('data_dir'),'storage'));}
try {
  fs.statSync(path.join(app.get('data_dir'),'data'));} 
catch(e) {
  fs.mkdirSync(path.join(app.get('data_dir'),'data'));}
app.set('report_dir', path.join(__dirname, 'public','report'));

var version = require('./package.json').version;
app.set('version', version+'-NJS/EXP');
app.set('version_number', version);

var util = require('./lib/node/tools.js').DataOutput();
var lang = require('./lib/node/lang.js')[conf.lang];
app.locals.lang = lang;
app.set('storage', require('./lib/node/storage.js')(app, function(err,host_settings){
  if(!err){
    //check settings values
    for (var setting in conf.def_settings) {
      if(!host_settings[setting] || host_settings[setting]===""){
        if(setting === "session_secret"){
          host_settings[setting] = util.createKey();}
        else{
          host_settings[setting] = conf.def_settings[setting];}
        app.get('storage').updateSetting(
          {fieldname:setting,value:host_settings[setting],description:""});}
      if(setting.indexOf("host_restriction")>-1){
        if(host_settings[setting]===""){
          host_settings[setting] = [];}
        else{
          host_settings[setting] = host_settings[setting].split(",");}}
      else if(["session_cookie_max_age","max_content_length"].indexOf(setting)>-1){
        if(!isNaN(parseInt(host_settings[setting],10))){
           host_settings[setting] = parseInt(host_settings[setting],10);}
         else{
           host_settings[setting] = conf.def_settings[setting];}}}
    app.set('host_settings', host_settings);

    // view engine setup
    app.set('views', path.join(__dirname, 'views'));
    app.engine('.html', require('ejs').__express);
    app.engine('.xml', require('ejs').__express);
    app.set('view engine', 'ejs');
    if (app.get('env') === 'production') {
      app.enable('trust proxy');
      app.disable('x-powered-by');}
      
    app.use(compression());
    app.use(favicon(__dirname + '/public/favicon.ico'));
    app.use(express.static(path.join(__dirname, 'public')));
    app.use('/js', express.static(path.join(__dirname, 'lib/dist/js')));
    app.use('/css', express.static(path.join(__dirname, 'lib/dist/css')));
    
    app.use('/lib/w3', express.static(path.join(__dirname, 'node_modules/w3-css')));
    app.use('/lib/highlightjs', express.static(path.join(__dirname, 'node_modules/highlightjs')));
    app.use('/lib/jspdf', express.static(path.join(__dirname, 'node_modules/jspdf/dist')));
    app.use('/lib/pdfjs', express.static(path.join(__dirname, 'node_modules/pdfjs-dist/build')));
    app.use('/lib/icon', express.static(path.join(__dirname, 'node_modules/font-awesome')));
    app.use('/lib/flatpickr', express.static(path.join(__dirname, 'node_modules/flatpickr/dist')));

    app.use(logger((app.get('env')==='development')?'dev':'common'));
    app.use(bodyParser.json());
    app.use(bodyParser.urlencoded({extended: true}));
    app.use(methodOverride());
    app.use(cookieParser(host_settings.session_secret));
    app.use(session({name:'ntura', 
      secret: host_settings.session_secret,
      httpOnly: true, maxAge: host_settings.session_cookie_max_age,
      secure: (app.get('env') === 'production') ? true : false,  
      proxy: (app.get('env') === 'production') ? true : false}))
    
    app.use(cors());
    app.use('/npi', npi);
    app.use('/ndi', ndi);

    app.use(passport.initialize());
    app.use(passport.session());
    app.use(hpp());
    app.use(helmet());
    app.use(contentLength.validateMax({max: host_settings.max_content_length, status: 400, message: 'Too much content'}));
    if (app.get('env') === 'production') {
      app.use(express_enforces_ssl());}
    app.use(lusca.csrf({secret: host_settings.session_secret}));
    
    switch (conf.start_page) {
      case "static":
        app.use(express.static(path.join(__dirname, 'www')));
        break;
      case "custom":
        app.use('/', custom);
        break;
      default:
        app.use('/', ntura);
        break;}
    app.use('/ntura', ntura);
    app.use('/nas', nas);
    app.use('/report', report);
    app.use('/ndi/wizard', wizard);
    app.use('/ndi/demo', demo);
    
    // Configure the local strategy for use by Passport.
    passport.use(new LocalStrategy(
      conf.nas_login.local, app.get('storage').getUserFromName));
    if(conf.nas_login.google.clientID && conf.nas_login.google.clientSecret){
      passport.use(new GoogleStrategy( conf.nas_login.google,
        function(accessToken, refreshToken, profile, cb) {
          process.nextTick(function () {
            app.get('storage').getUserFromEmail(profile, accessToken, function (err, user, info) {
              return cb(err, user, info); });});}));}
    passport.serializeUser(function(user, cb) {cb(null, user.id);});
    passport.deserializeUser(app.get('storage').getUserFromId);
    
    // catch 404 and forward to error handler
    app.use(function(req, res, next) {
      res.locals.user = req.user;
      var err = new Error('Not Found');
      err.status = 404;
      next(err);});

    // error handlers

    // development error handler
    // will print stacktrace
    if (app.get('env') === 'development') {
      app.use(function(err, req, res, next) {
        res.status(err.status || 500);
        res.render('error', {
          message: err.message,
          error: err});});}

    // production error handler
    // no stacktraces leaked to user
    app.use(function(err, req, res, next) {
      res.status(err.status || 500);
      res.render('error', {
        message: err.message,
        error: {}});});
   
   return callback(err, app);}
  else{
    return callback(err, app);}}));}