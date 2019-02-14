
const Nervastore  = require('../lib/nervastore')
const { createDatabase, createDemo, getDbsReports, reportDelete, reportInstall, getApi } = require('../lib/nas')();
const { ntconf } = require('./config')

describe('nas', () => {
  let nstore = Nervastore(ntconf);

  it("createDatabase", (done) => {
    createDatabase(nstore, {database:"test", logtype:"json"}, (err, log) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("createDemo", (done) => {
    createDemo(Nervastore(ntconf), {database:"test", username:"admin", logtype:"json"}, (result) => {
      expect(result[result.length-1].error).toBeUndefined();
      done();
    })
  })
  it("getDbsReports", (done) => {
    getDbsReports(nstore, {}, (err, results) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("reportDelete", (done) => {
    reportDelete(nstore, {reportkey:"ntr_customer_en"}, (err) => {
      expect(err).toBeNull();
      done();
    })
  })
  it("reportInstall", (done) => {
    reportInstall(nstore, {reportkey:"ntr_customer_en"}, (err, report_id, reportkey) => {
      expect(err).toBeNull();
      done();
    })
  })
  
  it("getApi: database/create", (done) => {
    getApi(nstore, {method:"database/create" , alias:"test", output:"json"}, (result) => {
      expect(result.type).not.toBe("error");
      done();
    })
  })
  it("getApi: database/demo", (done) => {
    getApi(Nervastore(ntconf), 
      {method:"database/demo" , database:"test", username:"admin", output:"json"}, (result, params) => {
      expect(result.data[result.data.length-1].error).toBeUndefined();
      done();
    })
  })
  it("getApi: report/list", (done) => {
    getApi(nstore, {method:"report/list" , alias:"test", output:"json"}, (result, params) => {
      expect(result.type).not.toBe("error");
      done();
    })
  })
  it("getApi: report/delete", (done) => {
    getApi(nstore, 
      {method:"report/delete" , alias:"test", reportkey:"ntr_customer_en", output:"json"}, (result, params) => {
      expect(result.type).not.toBe("error");
      done();
    })
  })
  it("getApi: report/install", (done) => {
    getApi(nstore, 
      {method:"report/install" , alias:"test", reportkey:"ntr_customer_en", output:"json"}, (result, params) => {
      expect(result.type).not.toBe("error");
      done();
    })
  })

})

