/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

module.exports = {
  port: 3000,
  lang: "en",
  secret: "abrakadabra",
  storage_type: "sqlite", //'level'(and PouchDB) or 'sqlite' or 'loki'
  session_cookie: 7*24*60*60*1000, //1 week
  pool_config: {min: 0, max: 10},
  host_restriction: {
    //localhost sample: ["::1","127.0.0.1","::ffff:127.0.0.1"]
    all:[],
    nas:[],
    npi:[],
    ndi:[]
  },
  engine: [{key:"sqlite", label:"SQLite"}, {key:"mysql", label:"MySQL"}, 
    {key:"postgres", label:"PostgreSQL"}, {key:"mssql", label:"MSSQL"}],
  users: [{username:"admin"}, {username:"demo"}],
  default_password: "12345",
  backup_dir: "public/backup",
  report_dir: "public/report",
  python2_path: "python", //Optional, for the server-side Nervatura Report (The path where to locate the "python" executable. Default: "python")
  python_script: "./lib/pylib", //The default path where to look for python 2 scripts.
  databases: [
    {alias:"demo", engine:"sqlite", 
     connect:{host:"", port:"", dbname:"demo", username:"", password:""},
     settings:{ndi_enabled:true, ndi_md5_password:false, ndi_encrypt_data:false, 
       ndi_encrypt_password:"", ndi_host_restriction:""}},
    {alias:"mydemo", engine:"mysql", 
      connect:{host:"localhost", port:3306, dbname:"demo", username:"root", password:"admin"},
     settings:{ndi_enabled:true, ndi_md5_password:false, ndi_encrypt_data:false, 
       ndi_encrypt_password:"", ndi_host_restriction:""}},
    {alias:"pgdemo", engine:"postgres", 
      connect:{host:"localhost", port:5432, dbname:"demo", username:"postgres", password:"admin"},
     settings:{ndi_enabled:true, ndi_md5_password:false, ndi_encrypt_data:false, 
       ndi_encrypt_password:"", ndi_host_restriction:""}},
    {alias:"msdemo", engine:"mssql", 
      connect:{host:"localhost", port:1433, dbname:"demo", username:"sa", password:"admin"},
     settings:{ndi_enabled:true, ndi_md5_password:false, ndi_encrypt_data:false, 
       ndi_encrypt_password:"", ndi_host_restriction:""}}],
  settings: [
    {fieldname: "login_enabled_lst", value: "",
      description: "Enabled admin login ip pattern. Comma-separated list."},
    {fieldname: "request_enabled_lst", value: "",
      description: "Enabled request ip pattern. Comma-separated list."}],
  datepicker: {
    dateFormat:"yy-mm-dd",
    defaultDate:0,
    changeMonth:true,
    changeYear:true}
}