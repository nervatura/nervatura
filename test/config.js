const settings = {
  DATA_DIR: "test"
}

const databases = {
  "test": {
    "engine": "sqlite", 
    "connect": {
      "host": "", 
      "port": "", 
      "dbname": "test", 
      "username": "", 
      "password": "" },
    "settings":{
      "db_ssl": false,
      "ndi_enabled": true, 
      "encrypt_password": "", 
      "dbs_host_restriction": "" }
  }
}

const { basicStore } = require('../lib/storage')
const Nervastore  = require('../lib/nervastore')

const conf = require('../lib/conf')(settings);
const lang = require('../lib/lang')[conf.lang];

const storage = basicStore({ data_store: conf.data_store, databases: databases,
  conf: conf, lang: lang, data_dir: conf.data_dir, host_type: conf.host_type});

exports.ntconf = { 
  conf: conf, data_dir: conf.data_dir, report_dir: conf.report_dir,
  host_ip: "", host_settings: conf.def_settings, storage: storage,
  lang: lang };