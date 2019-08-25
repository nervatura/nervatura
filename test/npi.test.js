const Nervastore  = require('../lib/nervastore')
const { getApi, getLogin, setData, loadDataSet, updateRecordSet, saveDataSet } = require('../lib/npi')();
const { ntconf } = require('./config')

describe('ndi', () => {
  let nstore = Nervastore(ntconf);
  it("getLogin", (done) => {
    getLogin(nstore, {database:"test", username:"admin"}, (validator) => {
      expect(validator.valid).toBeTruthy();
      done();
    })
  })
  it("setData", (done) => {
    setData(nstore, "table", 
      { login:{database:"test", username:"admin"}, 
        classAlias: "customer", filterStr: "", orderStr: ""}, (err, results) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("nervastore", (done) => {
    nstore.connect.getDatabaseSettings({}, (err, settings) => {
      expect(err).toBeNull();
      done();
    })
    /*
    nstore.valid.getGroupsId({
      groupname: ["transtype", "usergroup"]
    }, (err, groups) => {
      expect(err).toBeNull();
      done();
    })
    nstore.connect.getObjectAudit({
      nervatype: "customer"
    }, (err, info) => {
      expect(err).toBeNull();
      done();
    })
    nstore.connect.getDataAudit({}, (err, info) => {
      expect(err).toBeNull();
      done();
    })
    nstore.valid.getRefnumber(
      {nervatype: "item", ref_id: 6, rettype: ""}, (err, id, info) => {
      expect(err).toBeNull();
      done();
    })
    nstore.valid.getIdFromRefnumber(
      {nervatype: "ui_reportsources", refnumber: "ntr_invoice_en~head", extra_info: false}, (err, id, info) => {
      expect(err).toBeNull();
      done();
    })  
    setData(nstore, "delete", 
      { login:{database:"test", username:"admin"}, 
        record: {__tablename__: "customer", refnumber: "DMCUST/00001" } }, (err, results) => {
      expect(err).toBeNull();
      done();
    })
    */
  })
  
  it("setData", (done) => {
    setData(nstore, "function", 
      { login:{database:"test", username:"admin"}, 
        functionName: "getReport", paramList: {
          nervatype: "trans", refnumber:"DMINV/00001", 
          output: "pdf", orientation: "portrait" }}, (err, results) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("setData", (done) => {
    setData(nstore, "function", 
      { login:{database:"test", username:"admin"},
        functionName: "getReport", 
        paramList: {
          filters: { posdate: "2017-03-14" },
          output:"base64",
          reportkey: "xls_custpos_en" 
        }}, (err, results) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("loadDataSet", (done) => {
    loadDataSet(nstore, 
      { login:{database:"test", username:"admin"},
        dataSetInfo:[{infoName:"customer", infoType:"table", classAlias:"customer", 
        filterStr:"", orderStr:null}]}, (err, results) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("updateRecordSet", (done) => {
    updateRecordSet(nstore, "update",
      { login:{database:"test", username:"admin"}, method:"update",
        recordSet:[{ __tablename__:"customer", id:1, account:"12345678" }]}, (err, results) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("saveDataSet", (done) => {
    saveDataSet(nstore,
      { login:{database:"test", username:"admin"},
        dataSet:[ { updateType:"update", 
          recordSet:[{ __tablename__:"customer", id:1, account:"87654321" }]} ]}, (err, results) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("getApi", (done) => {
    getApi(nstore, { method:"loadTable", 
      params:{ login:{database:"test", username:"admin"}, 
        classAlias: "customer", filterStr: "", orderStr: ""} }, (result) => {
      expect(result.type).not.toBe("error");
      done();
    })
  })
})