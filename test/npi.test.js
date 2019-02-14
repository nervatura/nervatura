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