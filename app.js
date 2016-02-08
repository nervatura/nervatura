/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2015, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/* global __dirname */

var express = require('express');
var session = require('express-session');

var compression = require('compression');
var passport = require('passport');
var Strategy = require('passport-local').Strategy;
var cors = require('cors');
var lusca = require('lusca');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
  
//routes
var index = require('./routes/index');
var demo = require('./routes/demo');
var npi = require('./routes/npi');
var ndi = require('./routes/ndi');
var nas = require('./routes/nas');
var report = require('./routes/report');
var wizard = require('./routes/wizard');

//ntura
var conf = require('./lib/node/conf.js');
var util = require('./lib/node/tools.js').DataOutput();
var version = require('./package.json').version;
var lang = require('./lib/node/lang.js')[conf.lang];
var storage = require('./lib/node/storage.js')(conf, lang);

var app = express();

app.set('port', conf.port);
app.set('storage', storage);
app.set('version', version+"-NJS/EXP");
app.set('version_number', version);
app.set('conf', conf);
app.set('backup_dir', path.join(__dirname, conf.backup_dir));
app.set('report_dir', path.join(__dirname, conf.report_dir));
app.locals.lang = lang;
// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.engine('.html', require('ejs').__express);
app.engine('.xml', require('ejs').__express);
app.set('view engine', 'ejs');

app.use(compression());
app.use(favicon(__dirname + '/public/favicon.ico'));
app.use(express.static(path.join(__dirname, 'public')));
app.use('/javascripts', express.static(path.join(__dirname, 'lib/dist/js')));
app.use('/stylesheets', express.static(path.join(__dirname, 'lib/dist/css')));
app.use('/javascripts', express.static(path.join(__dirname, 'node_modules/jquery/dist')));
app.use('/javascripts', express.static(path.join(__dirname, 'node_modules/jquery-mobile-babel-safe/js')));
app.use('/stylesheets', express.static(path.join(__dirname, 'node_modules/jquery-mobile-babel-safe/css')));

app.use(logger('dev'));
app.use(cookieParser());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));
//app.use(fileUpload());
if (conf.storage_type==='level' && util.checkOptional('level-session-store')){
  //optional: Linux, Mac OS, FreeBSD
  //data-store: leveldown and PouchDB (/storage/data/)
  var LevelSession = require('level-session-store')(session);
  app.use(session({name:"ntura", secret: conf.secret, resave: true, saveUninitialized: true, 
    store: new LevelSession('./storage/session'), 
    cookie: {httpOnly:true, secure:true, maxAge:conf.session_cookie}}));}
else if (conf.storage_type==='sqlite' && util.checkOptional('connect-sqlite3')){
  //optional: Windows, Linux, Mac OS, FreeBSD
  //data-store: lokiJS (/storage/data.json)
  var SQLiteStore = require('connect-sqlite3')(session);
  app.use(session({name:"ntura", secret: conf.secret, resave: true, saveUninitialized: true,
    store: new SQLiteStore({dir:'./storage',db:'session'}), 
    cookie: {httpOnly:true, secure:false, maxAge:conf.session_cookie}}));}
else if (conf.storage_type==='gcloud' && util.checkOptional('cloud-datastore-session')){
  //optional: Google Cloud Platform
  //data-store: Google Cloud Datastore
}
else {
  //default: MemoryStore for debugging and developing
  //data-store: lokiJS (/storage/data.json)
  app.use(session({name:"ntura", secret: conf.secret, resave: true, saveUninitialized: true}));}

app.use(cors());
app.use('/npi', npi);
app.use('/ndi', ndi);

app.use(lusca({
  csrf: true,
  csp: { /* ... */},
  xframe: 'SAMEORIGIN',
  p3p: 'ABCDEF',
  hsts: {maxAge: 31536000, includeSubDomains: true, preload: true},
  xssProtection: true}));
app.use(passport.initialize());
app.use(passport.session());
  
app.use('/', index);
app.use('/nas', nas);
app.use('/report', report);
app.use('/ndi/wizard', wizard);
app.use('/ndi/demo', demo);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
});

// error handlers

// development error handler
// will print stacktrace
if (app.get('env') === 'development') {
    app.use(function(err, req, res, next) {
        res.status(err.status || 500);
        res.render('error', {
            message: err.message,
            error: err
        });
    });
}

// production error handler
// no stacktraces leaked to user
app.use(function(err, req, res, next) {
    res.status(err.status || 500);
    res.render('error', {
        message: err.message,
        error: {}
    });
});

// Configure the local strategy for use by Passport.
passport.use(new Strategy(storage.getUserFromName));
passport.serializeUser(function(user, cb) {cb(null, user.id);});
passport.deserializeUser(storage.getUserFromId);
    
module.exports = app;
