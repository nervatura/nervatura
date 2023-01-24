import sinon from 'sinon'
import { expect } from '@open-wc/testing';

import { getSql, saveToDisk, request } from './Utils.js'

it('getSql', () => {
  expect(getSql("sqlite", "select * from table")).to.exist;
  expect(getSql("sqlite3", "select * from table")).to.exist;
  expect(getSql("mysql", "select * from table")).to.exist;
  expect(getSql("postgres", "select * from table")).to.exist;
  expect(getSql("mssql", "select * from table")).to.exist;
  expect(getSql("", "select * from table")).to.exist;

  expect(getSql("postgres", {
    select: ["col1, col2"], from: "table t1",
    inner_join:["table2 t2","on",["t1.id","=","t2.id"]],
    left_join:["table3 t3","on",["t1.id","=","t3.id"]],
    where: [["t1.col1","=","0"],["and",["t3.col3","is",null]],["or",["t2.col2",">","1"]]]
  })).to.exist;

  expect(getSql("postgres", {
    update: "table", set: [[],["col1","=","?"],["col2","=","?"]], where: [["col1","=","0"]]
  })).to.exist;
  expect(getSql("postgres", {
    insert_into: ["table",[[],"col1","col2","col3"]], values:[[],"?","?","?"]
  })).to.exist;

  expect(getSql("sqlite3", {
    update: "table", set: [[],["col1","=","?"],["col2","=","?"]], where: [["col1","=","0"]]
  })).to.exist;
  expect(getSql("sqlite3", {
    insert_into: ["table",[[],"col1","col2","col3"]], values:[[],"?","?","?"], where: []
  })).to.exist;
});

describe('saveToDisk', () => {
  beforeEach(() => {
    sinon.spy(document.body, 'appendChild');
  });
  afterEach(() => {
    sinon.restore();
  });
  it('call saveToDisk', () => {
    saveToDisk("fileUrl", "fileName")
    saveToDisk("fileUrl")
  });
})

describe('request', () => {
  // Before each test, stub the fetch function
  beforeEach(() => {
    window.fetch = sinon.spy();
  });

  afterEach(() => {
    sinon.restore();
  });

  describe('stubbing successful json response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('{"hello":"world"}', {
        status: 200,
        headers: {
          'Content-type': 'application/json',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.hello).to.equal('world');
    });

  });

  describe('stubbing successful text response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('hello', {
        status: 200,
        headers: {
          'Content-type': 'text/plain',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).to.equal("hello");
    });

  });

  describe('stubbing successful csv response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('hello,hello', {
        status: 200,
        headers: {
          'Content-type': 'text/csv',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).to.equal("hello,hello");
    });

  });

  describe('stubbing successful default response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('default', {
        status: 200,
        headers: {
          'Content-type': 'unknown',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.status).to.equal(200);
    });

  });

  describe('stubbing successful pdf response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response(new Blob(["pdf"], { type: 'application/pdf'}), {
        status: 200,
        headers: {
          'Content-type': 'application/pdf',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).to.exist;
    });

  });

  describe('stubbing successful xml response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response(new Blob(["<xml></xml>"], { type: 'text/xml'}), {
        status: 200,
        headers: {
          'Content-type': 'application/xml',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).to.exist;
    });

  });

  describe('stubbing successful excel response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response(new Blob(["<xml></xml>"], { type: 'text/xml'}), {
        status: 200,
        headers: {
          'Content-type': 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).to.exist;
    });

  });

  describe('stubbing 204 response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response(null, {
        status: 204,
        statusText: 'No Content'
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should return null on 204 response', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).to.null
    });

  });

  describe('stubbing 401 response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('{"hello":"world"}', {
        status: 401,
        headers: {
          'Content-type': 'application/json',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.message).to.equal('Unauthorized');
    });

  });

  describe('stubbing 400 response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('{"error":"errordata"}', {
        status: 400,
        headers: {
          'Content-type': 'application/json',
        },
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.error).to.equal('errordata');
    });

  });

  describe('stubbing 205 response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response(null, {
        status: 205,
        statusText: 'No Content'
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should return null on 205 response', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).to.null;
    });

  });

  describe('stubbing error response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('', {
        status: 404,
        statusText: 'Not Found',
      });

      window.fetch = sinon.spy(()=>(Promise.resolve(res)))
    });

    it('should catch errors', async () => {
      let result
      try {
        result = await request('/thisurliscorrect')
      } catch (error) {
        result = error.message
      }
      expect(result).to.equal("Not Found");
    });

  });

});