/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

module.exports = {
  //Initial setup
  databases: [
    {alias:"demo", engine:"sqlite", 
     connect:{host:"", port:"", dbname:"demo", username:"", password:""},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"mydemo", engine:"mysql", 
      connect:{host:"localhost", port:3306, dbname:"nervatura", username:"root", password:"admin"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"pgdemo", engine:"postgres", 
      connect:{host:"localhost", port:5432, dbname:"nervatura", username:"postgres", password:"admin"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"msdemo", engine:"mssql", 
      connect:{host:"localhost", port:1433, dbname:"nervatura", username:"sa", password:"sadmin"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}}],
  databases_docker: [
    {alias:"demo", engine:"sqlite", 
     connect:{host:"", port:"", dbname:"demo", username:"", password:""},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"mydemo", engine:"mysql", 
      connect:{host:"$MYSQL_PORT_3306_TCP_ADDR", port:"$MYSQL_PORT_3306_TCP_PORT", 
      dbname:"$MYSQL_ENV_MYSQL_DATABASE", username:"$MYSQL_ENV_MYSQL_USER", 
      password:"$MYSQL_ENV_MYSQL_PASSWORD"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"pgdemo", engine:"postgres", 
      connect:{host:"$POSTGRES_PORT_5432_TCP_ADDR", port:"$POSTGRES_PORT_5432_TCP_PORT", 
      dbname:"$POSTGRES_ENV_POSTGRES_DB", username:"$POSTGRES_ENV_POSTGRES_USER", 
      password:"$POSTGRES_ENV_POSTGRES_PASSWORD"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"msdemo", engine:"mssql", 
      connect:{host:"192.168.56.1", port:1433, dbname:"nervatura", username:"sa", password:"sadmin"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}}],
  databases_gae: [
    {alias:"demo", engine:"sqlite", 
     connect:{host:"", port:"", dbname:"demo", username:"", password:""},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"mydemo", engine:"mysql", 
      connect:{host:"/cloudsql/$INSTANCE_CONNECTION_NAME", port:"", 
      dbname:"$SQL_DATABASE", username:"$SQL_USER", password:"$SQL_PASSWORD"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"pgdemo", engine:"postgres", 
      connect:{host:"/cloudsql/$INSTANCE_CONNECTION_NAME", port:"", 
      dbname:"$SQL_DATABASE", username:"$SQL_USER", password:"$SQL_PASSWORD"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}}],
  databases_openshift: [
    {alias:"demo", engine:"sqlite", 
     connect:{host:"", port:"", dbname:"demo", username:"", password:""},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"mydemo", engine:"mysql", 
      connect:{host:"$OPENSHIFT_MYSQL_DB_HOST", port:"$OPENSHIFT_MYSQL_DB_PORT", 
      dbname:"$OPENSHIFT_APP_NAME", username:"$OPENSHIFT_MYSQL_DB_USERNAME", 
      password:"$OPENSHIFT_MYSQL_DB_PASSWORD"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"pgdemo", engine:"postgres", 
      connect:{host:"$OPENSHIFT_POSTGRESQL_DB_HOST", port:"$OPENSHIFT_POSTGRESQL_DB_PORT", 
      dbname:"$OPENSHIFT_APP_NAME", username:"$OPENSHIFT_POSTGRESQL_DB_USERNAME", 
      password:"$OPENSHIFT_POSTGRESQL_DB_PASSWORD"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}}],
  settings: [],
  users: [{username:"admin"}],
  //Application const and UI/GUI settings
  start_page: "ntura", //static, custom, ntura
  data_dir: "data",
  engine: [{key:"sqlite", label:"SQLite"}, {key:"mysql", label:"MySQL"}, 
    {key:"postgres", label:"PostgreSQL"}, {key:"mssql", label:"MSSQL"}],
  pool_config: {min: 0, max: 10},
  port: 8080,
  python_script: "./lib/pylib",
  lang: "en",
  data_store: "lokijs", //lokijs,pouchdb,datastore
  nas_login: {
    local: {
      session: false
    },
    google:{
      clientID: process.env.GOOGLE_CLIENT_ID,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET,
      callbackURL: "/nas/google/callback",
      session: false
    }
  },
  email_providers: {
    smtp: {
      host: process.env.SMTP_HOST || "YOUR_SMTP_HOST",
      port: process.env.SMTP_PORT || "YOUR_SMTP_PORT",
      secure: process.env.SMTP_SECURE || true, // use TLS
      auth: {
        user: process.env.SMTP_USER || "YOUR_SMTP_USER",
        pass: process.env.SMTP_PASSWORD || "YOUR_SMTP_PASSWORD"
      }
    },
    mailjet: {
      clientID: process.env.MJ_APIKEY_PUBLIC || "YOUR_MJ_APIKEY_PUBLIC",
      clientSecret: process.env.MJ_APIKEY_PRIVATE || "YOUR_MJ_APIKEY_PRIVATE"
    }
  },
  def_settings: {
    nas_auth: "local", //local, google
    python2_path: "python",
    session_cookie_max_age: 7*24*60*60*1000,
    max_content_length: 20000,
    session_secret:"", all_host_restriction: "",
    nas_host_restriction: "", npi_host_restriction: "", ndi_host_restriction: ""},
  get_value: function(value,filter) {
    if(this[value+"_"+filter]){
      return this[value+"_"+filter];}
    else{
      return this[value];}}
}