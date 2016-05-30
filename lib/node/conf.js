/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
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
      connect:{host:"localhost", port:3306, dbname:"demo", username:"root", password:"admin"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"pgdemo", engine:"postgres", 
      connect:{host:"localhost", port:5432, dbname:"demo", username:"postgres", password:"admin"},
     settings:{ndi_enabled:true, encrypt_password:"", dbs_host_restriction:""}},
    {alias:"msdemo", engine:"mssql", 
      connect:{host:"localhost", port:1433, dbname:"demo", username:"sa", password:"sadmin"},
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
  data_dir: "data",
  datepicker: {
    dateFormat:"yy-mm-dd",
    defaultDate:0,
    changeMonth:true,
    changeYear:true},
  engine: [{key:"sqlite", label:"SQLite"}, {key:"mysql", label:"MySQL"}, 
    {key:"postgres", label:"PostgreSQL"}, {key:"mssql", label:"MSSQL"}],
  pool_config: {min: 0, max: 10},
  port: 3000,
  python_script: "./lib/pylib",
  lang: "en",
  data_store: "lokijs", //lokijs,pouchdb
  session_store: "sqlite", //leveldown,sqlite,memory
  def_settings: {
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