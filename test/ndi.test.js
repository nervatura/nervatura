const Nervastore  = require('../lib/nervastore')
const { getApi, getData, updateData, deleteData } = require('../lib/ndi')();
const { ntconf } = require('./config')

describe('ndi', () => {
  let nstore = Nervastore(ntconf);

  it("getApi", (done) => {
    getApi(nstore, {
      id:1, method:"getData", jsonrpc:"2.0", 
      params:[
        {database:"test", username:"admin", password:"", datatype:"customer"},
        {output:"json"}] }, (result) => {
      expect(result.type).not.toBe("error");
      done();
    })
  })

  it("getData", (done) => {
    let _conn = nstore.connect.getConnect()
    getData(nstore, 
      {datatype:"customer", validator: {valid:true, message:"", conn:_conn}}, 
      {output:"json", check_audit: false }, (err, result) => {
      if (_conn !== null){
        _conn.close();}
      expect(err).toBeNull();
      done();
    })
  })

  it("updateData", (done) => {
    let _conn = nstore.connect.getConnect()
    updateData(nstore,
      {datatype:"currency", insert_row:true, validator: {valid:true, message:"", conn:_conn}},
      [{curr:"ABC", description:"TEST"}], (err, result) => {
      if (_conn !== null){
        _conn.close();}
      expect(err).toBeNull();
      done();
    })
  })

  it("deleteData", (done) => {
    let _conn = nstore.connect.getConnect()
    deleteData(nstore,
      {datatype:"currency", validator: {valid:true, message:"", conn:_conn}},
      [{curr:"ABC"}], (err, result) => {
      if (_conn !== null){
        _conn.close();}
      expect(err).toBeNull();
      done();
    })
  })

})