/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

var express = require('express');
var passport = require('passport');
var router = express.Router();

router.use(function (req, res, next) {
  next()
});

router.get('/', function(req, res, next) {
  res.redirect('index');
});

/*
router.get('/env', function(req, res, next) {
  var content = 'Version: ' + process.version + '\n<br/>\n' +
    'Env: {<br/>\n<pre>';
  for (var k in process.env) {
    content += '   ' + k + ': ' + process.env[k] + '\n';}
  content += '}\n</pre><br/>\n'
  res.send('<html>\n' +
    '  <head><title>Node.js Process Env</title></head>\n' +
    '  <body>\n<br/>\n' + content + '</body>\n</html>');
});
*/

router.get('/index', function(req, res, next) {
  res.render('default/index.html',{view:"index"});
});

router.get('/ndoc', function(req, res, next) {
  res.render('default/index.html',{view:"ndoc"});
});

router.get('/about', function(req, res, next) {
  res.render('default/index.html',{view:"about"});
});

router.get('/licenses', function(req, res, next) {
  res.render('default/index.html',{view:"licenses"});
});

router.get('/docs/nas', function(req, res, next) {
  res.render('docs/nas.html',{});
});

router.get('/docs/nom', function(req, res, next) {
  res.render('docs/nom.html',{});
});

router.get('/docs/ndi', function(req, res, next) {
  res.render('docs/ndi.html',{});
});

router.get('/docs/npi', function(req, res, next) {
  res.render('docs/npi.html',{});
});

router.get('/docs/ndr', function(req, res, next) {
  res.render('docs/ndr.html',{});
});

router.get('/docs/report', function(req, res, next) {
  res.render('docs/report.html',{});
});

function page_render(params){
  if (typeof params.dir === "undefined"){
    params.dir = "default"}
  if (typeof params.data === "undefined"){
    params.data = {view:params.page};}
  else{
    params.data.view = params.page;}
  params.res.render(params.dir+"/index.html",params.data);}

router.get('/login',
  function(req, res){
    var data = {username:"", flash:null};
    if (req.user) {
      data.username = req.user.username;}
    page_render({req:req, res:res, page:"login", data:data});});

router.post('/login', function(req, res, next) {
  if(!req.body.password || req.body.password===""){
    req.body.password = "empty";}
  passport.authenticate('local', function(err, user, info) {
    if (err) { return next(err); }
    if (!user) {
      page_render({req:req, res:res, page:"login", 
        data:{username:info.username, flash:info.message}});}
    else {
      req.logIn(user, function(err) {
        if (err) {return next(err);}
        if(user.dirty_password){
          return res.redirect('nas/user/password');}
        else{
          return res.redirect('nas/index');}});}
})(req, res, next);});
  
router.get('/logout',
  function(req, res){
    req.logout();
    return res.redirect('login');});
     
module.exports = router;
