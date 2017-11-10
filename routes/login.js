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

router.get('/',
  function(req, res){
    var data = {view:"login", username:"", flash:null};
    if (req.user) {
      data.username = req.user.username;}
    res.render("ntura/index.html", data);});

router.get('/insecure', function(req, res, next) {
  res.render("ntura/index.html", 
    {view: "login", insecure: true, username: "", 
     flash: "NAS "+req.app.locals.lang.insecure_err});});

router.post('/local', function(req, res, next) {
  if(!req.body.password || req.body.password===""){
    req.body.password = "empty";}
  passport.authenticate('local', function(err, user, info) {
    if (err) { return next(err); }
    if (!user) {
      res.render("ntura/index.html", {view:"login", username:info.username, flash:info.message});}
    else {
      req.logIn(user, function(err) {
        if (err) {return next(err);}
        if(user.dirty_password){
          return res.redirect('/nas/user/password');}
        else{
          return res.redirect('/nas/index');}});}
        })(req, res, next);});

router.get('/amazon', 
  passport.authenticate('amazon', { scope: ['profile', 'postal_code'] }));
      
router.all('/amazon/callback', function(req, res, next) {
  passport.authenticate('amazon', 
  function(err, user, info) {
    if (err) { return next(err); }
    if (!user && info) {
      res.render("ntura/index.html", {view:"login", username:info.username, flash:info.message});}
    else if (!user) {
      res.render("ntura/index.html", {view:"login"});}
    else {
      req.logIn(user, function(err) {
        if (err) {return next(err);}
        res.redirect('/nas/index');})
  }})(req, res, next);});

router.get('/azure',
  function(req, res, next) {
    passport.authenticate('azuread-openidconnect', 
      { response: res, scope: ['email'] } )(req, res, next);});

router.all('/azure/callback', function(req, res, next) {
  passport.authenticate('azuread-openidconnect', { response: res },
  function(err, user, info) {
    if (err) { return next(err); }
    if (!user && info) {
      res.render("ntura/index.html", {view:"login", username:info.username, flash:info.message});}
    else if (!user) {
      res.render("ntura/index.html", {view:"login"});}
    else {
      req.logIn(user, function(err) {
        if (err) {return next(err);}
        res.redirect('/nas/index');})
      }})(req, res, next);});

router.get('/google', 
  passport.authenticate('google', { scope: ['email'] }));
      
router.all('/google/callback', function(req, res, next) {
  passport.authenticate('google', 
  function(err, user, info) {
    if (err) { return next(err); }
    if (!user && info) {
      res.render("ntura/index.html", {view:"login", username:info.username, flash:info.message});}
    else if (!user) {
      res.render("ntura/index.html", {view:"login"});}
    else {
      req.logIn(user, function(err) {
        if (err) {return next(err);}
        res.redirect('/nas/index');})
  }})(req, res, next);});
  
router.get('/logout',
  function(req, res){
    req.logout();
    return res.redirect('/login');});

module.exports = router;