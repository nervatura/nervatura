/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
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

router.get('/index', function(req, res, next) {
  res.render('default/index.html',{});
});

router.get('/ndoc', function(req, res, next) {
  res.render('default/ndoc.html',{});
});

router.get('/about', function(req, res, next) {
  res.render('default/about.html',{});
});

router.get('/licenses', function(req, res, next) {
  res.render('default/licenses.html',{});
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
    params.data = {};}
  params.res.render(params.dir+"/"+params.page+".html",params.data);}

router.get('/login',
  function(req, res){
    var data = {username:"", flash:null};
    if (req.user) {
      data.username = req.user.username;}
    page_render({req:req, res:res, page:"login", data:data});});

router.post('/login', function(req, res, next) {
  passport.authenticate('local', function(err, user, info) {
    if (err) { return next(err); }
    if (!user) {
      page_render({req:req, res:res, page:"login", 
        data:{username:info.username, flash:info.message}});}
    else {
      req.logIn(user, function(err) {
        if (err) { return next(err); }
        return res.redirect('nas/index');});}
  })(req, res, next);});
  
router.get('/logout',
  function(req, res){
    req.logout();
    return res.redirect('login');});
     
module.exports = router;
