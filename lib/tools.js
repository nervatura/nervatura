/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2019, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/* global __dirname */
/* global process */
/* global Buffer */

var crypto = require('crypto');
var service = require('./service.js')();
    
//DataOutput

function check_optional(mname) {
  try {
    require.resolve(mname); 
    return true;} 
  catch(e){}
    return false;}

function create_cipher(key, value, decode){
  try {
    var cipher = crypto.createCipher("des-ecb", key);
    var crypted = cipher.update(value,'utf8',decode); crypted += cipher.final(decode);
    return crypted;} 
  catch (error) {
    return value;}}

function create_decipher(key, value, decode){
  try {
    var decipher = crypto.createDecipher("des-ecb", key);
    var dec = decipher.update(value,decode,'utf8'); dec += decipher.final('utf8');
    return dec;} 
  catch (error) {
    return value;}}

function create_hash(value, decode){
  value = crypto.createHash('md5').update(value).digest(decode);
  return value;}

function create_random_key(size){
  var value = crypto.randomBytes(size || 32).toString("base64");;
  return value;}

function replace_all(text, sold, snew) {
  text = text.replace(sold,snew);
  if (text.indexOf(sold)>-1) {return replace_all(text, sold, snew);}
  else return text;}

function zero_pad(x,y){
  y = Math.max(y-1,0);
  var n = (x / Math.pow(10,y)).toFixed(y);
  return n.replace('.','');}

function round(n,dec) {
  n = parseFloat(n);
  if(!isNaN(n)){
    if(!dec) dec= 0;
    var factor= Math.pow(10,dec);
    return Math.floor(n*factor+((n*factor*10)%10>=5?1:0))/factor;}
  else{return n;}}
    
function get_iso_date(cdate,nosep) {
  if (typeof cdate === "undefined") {
    cdate = new Date();}
  if (nosep){
    return cdate.getFullYear()+zero_pad(cdate.getMonth()+1,2)+zero_pad(cdate.getDate(),2);}
  else {
    return cdate.getFullYear()+"-"+zero_pad(cdate.getMonth()+1,2)+"-"+zero_pad(cdate.getDate(),2);}}

function get_iso_datetime(cdate,full,nosep) {
  if (typeof cdate === "undefined") {
    cdate = new Date();}
  if (typeof full === "undefined") {
    full = true;}
  if (full) {
    if (nosep){
      return get_iso_date(cdate,nosep)+zero_pad(cdate.getHours(),2)+
        zero_pad(cdate.getMinutes(),2)+zero_pad(cdate.getSeconds(),2);}
    else {
      return get_iso_date(cdate)+"T"+zero_pad(cdate.getHours(),2)+":"+
        zero_pad(cdate.getMinutes(),2)+":00";}}
  else {
    return get_iso_date(cdate)+"T00:00:00";}}

function get_valid_date(value){
  var year,mo,day;
  value = replace_all(value.toString(),"'","");
  if (value!=="" && value!==null){
    if (value.length>=4) {
      year = parseInt(value.substring(0,4),10);
      if (isNaN(year)) {year = new Date().getFullYear();}
      if (year<1900) {year = 1900;}
      if (year>2200) {year = 2200;}}
    else {
      year = new Date().getFullYear();}
    if (value.length>=7) {
      mo = parseInt(value.substring(5,7),10);
      if (isNaN(mo)) {mo = 1;}
      if (mo<1) {mo = 1;}
      if (mo>12) {mo = 12;}}
    else {mo = 1;}
    if (value.length>=10) {
      day = parseInt(value.substring(8,10),10);
      if (isNaN(day)) {day = 1;}}
    else {day = 1;}
    return get_iso_date(new Date(year,mo-1,day));}
  else {
    return value;}}

function get_valid_datetime(value){
  var year,mo,day,ho,min;
  value = replace_all(value.toString(),"'","");
  if (value!=="" && value!==null){
    if (value.length>=4) {
      year = parseInt(value.substring(0,4),10);
      if (isNaN(year)) {year = new Date().getFullYear();}
      if (year<1900) {year = 1900;}
      if (year>2200) {year = 2200;}}
    else {
      year = new Date().getFullYear();}
    if (value.length>=7) {
      mo = parseInt(value.substring(5,7),10);
      if (isNaN(mo)) {mo = 1;}
      if (mo<1) {mo = 1;}
      if (mo>12) {mo = 12;}}
    else {mo = 1;}
    if (value.length>=10) {
      day = parseInt(value.substring(8,10),10);
      if (isNaN(day)) {day = 1;}}
    else {day = 1;}
    if (value.length>=13) {
      ho = parseInt(value.substring(11,13),10);
      if (isNaN(ho)) {ho = 0;}}
    else {ho = 0;}
    if (value.length>=16) {
      min = parseInt(value.substring(14,16),10);
      if (isNaN(min)) {min = 0;}}
    else {min = 0;}
    return get_iso_datetime(new Date(year,mo-1,day,ho,min,0));}
  else {
    return value;}}

function get_valid_path(){
  return __dirname;
}

exports.NervaTools = function() {
  return {
    callMenuCmd: function(nstore, params, _callback) {
      service.callMenuCmd(params, _callback);},
    printQueue: function(nstore, params, _callback) {
      service.printQueue(nstore, params, _callback);},
    getPriceValue: function(nstore, params, _callback) {
      service.getPriceValue(nstore, params, _callback);},
    nextNumber: function(nstore, params, _callback) {
      nstore.connect.nextNumber(params, _callback);},
    getReport: function(nstore, params, _callback) {
      service.getReport(nstore, params, _callback);},
    sendEmail: function(nstore, params, _callback) {
      service.sendEmail(nstore, params, _callback);},
    getReportFiles: function(nstore, params, _callback) {
      service.getReportFiles(nstore, params, _callback);},
    installReport: function(nstore, params, _callback) {
      service.installReport(nstore, params, _callback);}
};}

exports.DataOutput = function() {
  return {
    checkOptional: function(mname) {
      return check_optional(mname);},
    cryptedValue: function(key, value, decode) {
      return create_cipher(key, value, decode);},
    decipherValue: function(key, value, decode) {
      return create_decipher(key, value, decode);},
    createHash: function(value, decode) {
      return create_hash(value, decode);},
    createKey: function(size) {
      return create_random_key(size);},
    
    getISODate: function(cdate,nosep) {
      return get_iso_date(cdate,nosep);},
    getISODateTime: function(cdate,full,nosep) {
      return get_iso_datetime(cdate,full,nosep);},
    getValidDate: function(value) {
      return get_valid_date(value);},
    getValidDateTime: function(value) {
      return get_valid_datetime(value);},
    getValidPath: function() {
      return get_valid_path();},  
    replaceAll: function(text, sold, snew) {
      return replace_all(text, sold, snew);},
    zeroPad: function(x,y) {
      return zero_pad(x,y);},
    Round: function(n, dec) {
      return round(n, dec);} 
};}

