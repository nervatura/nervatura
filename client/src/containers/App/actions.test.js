import { Suspense } from 'react'
import { queryByAttribute } from '@testing-library/react'
import ReactDOM from 'react-dom';
import update from 'immutability-helper';

import { store as app_store  } from 'config/app'
import { request, guid, saveToDisk, getSql, appActions } from './actions'
import { toast } from 'react-toastify';
import InputBox from 'components/Modal/InputBox'

jest.mock("react-toastify");

const getById = queryByAttribute.bind(null, 'id');

it('getSql', () => {
  expect(getSql("sqlite", "select * from table")).toBeDefined();
  expect(getSql("sqlite3", "select * from table")).toBeDefined();
  expect(getSql("mysql", "select * from table")).toBeDefined();
  expect(getSql("postgres", "select * from table")).toBeDefined();
  expect(getSql("mssql", "select * from table")).toBeDefined();
  expect(getSql("", "select * from table")).toBeDefined();

  expect(getSql("postgres", {
    select: ["col1, col2"], from: "table t1",
    inner_join:["table2 t2","on",["t1.id","=","t2.id"]],
    left_join:["table3 t3","on",["t1.id","=","t3.id"]],
    where: [["t1.col1","=","0"],["and",["t3.col3","is",null]],["or",["t2.col2",">","1"]]]
  })).toBeDefined();

  expect(getSql("postgres", {
    update: "table", set: [[],["col1","=","?"],["col2","=","?"]], where: [["col1","=","0"]]
  })).toBeDefined();
  expect(getSql("postgres", {
    insert_into: ["table",[[],"col1","col2","col3"]], values:[[],"?","?","?"]
  })).toBeDefined();

  expect(getSql("sqlite3", {
    update: "table", set: [[],["col1","=","?"],["col2","=","?"]], where: [["col1","=","0"]]
  })).toBeDefined();
  expect(getSql("sqlite3", {
    insert_into: ["table",[[],"col1","col2","col3"]], values:[[],"?","?","?"], where: []
  })).toBeDefined();
});

describe('saveToDisk', () => {
  beforeEach(() => {
    jest.spyOn(document.body, 'appendChild');
  });
  afterEach(() => {
    jest.clearAllMocks();
  });
  it('call saveToDisk', () => {
    saveToDisk("fileUrl", "fileName")
    saveToDisk("fileUrl")
  });
})

it('guid', () => {
  expect(guid()).toBeDefined();
});

describe('request', () => {
  // Before each test, stub the fetch function
  beforeEach(() => {
    window.fetch = jest.fn();
  });

  afterEach(() => {
    jest.clearAllMocks();
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.hello).toBe('world');
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).toBe("hello");
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).toBe("hello,hello");
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.status).toBe(200);
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).toBeDefined();
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).toBeDefined();
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).toBeDefined();
    });

  });

  describe('stubbing 204 response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('', {
        status: 204,
        statusText: 'No Content'
      });

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should return null on 204 response', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).toBeDefined();
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.message).toBe('Unauthorized');
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

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should format the response correctly', async () => {
      const result = await request('/thisurliscorrect')
      expect(result.error).toBe('errordata');
    });

  });

  describe('stubbing 205 response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('', {
        status: 205,
        statusText: 'No Content'
      });

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should return null on 205 response', async () => {
      const result = await request('/thisurliscorrect')
      expect(result).toBeDefined();
    });

  });

  describe('stubbing error response', () => {
    // Before each test, pretend we got a successful response
    beforeEach(() => {
      const res = new Response('', {
        status: 404,
        statusText: 'Not Found',
      });

      window.fetch.mockReturnValue(Promise.resolve(res));
    });

    it('should catch errors', async () => {
      let result
      try {
        result = await request('/thisurliscorrect')
      } catch (error) {
        result = error.message
      }
      expect(result).toBe("Not Found");
    });

  });

});

describe('appActions', () => {

  beforeEach(() => {
    window.fetch = jest.fn()
    toast.mockReturnValue({
      error: jest.fn(),
      warning: jest.fn(),
      success: jest.fn(),
      info: jest.fn(),
    })
    jest.spyOn(document.body, 'appendChild');
  });
  
  afterEach(() => {
    jest.clearAllMocks();
  });
  
  it('getText', () => {
    let langText = appActions(app_store, jest.fn()).getText("en", "en")
    expect(langText).toBe("English")
  });
  
  it('showToast', () => {
    appActions(app_store, jest.fn()).showToast({autoClose: true, type: "error", message: "message"})
    appActions(app_store, jest.fn()).showToast({autoClose: false, type: "warning", message: "message"})
    appActions(app_store, jest.fn()).showToast({autoClose: false, type: "success", message: "message"})
    appActions(app_store, jest.fn()).showToast({autoClose: true, type: "info", message: "message"})
    appActions(app_store, jest.fn()).showToast({autoClose: true, type: "default", message: "message"})
  });

  it('resultError', () => {    
    appActions(app_store, jest.fn()).resultError({})
    appActions(app_store, jest.fn()).resultError({error:"error"})
    appActions(app_store, jest.fn()).resultError({error:{message: "message"}})
  });

  it('signOut', () => {
    const setData = jest.fn()
    appActions(app_store, setData).signOut()
    expect(setData).toHaveBeenCalledTimes(1);
  });

  it('requestData 200', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    const it_store = update(app_store, {
      session: {$merge: {
        configServer: true
      }},
      login: {$merge: {
        data: {
          token: "token"
        }
      }}
    })
    const options = {
      data: {
        value: "value"
      },
      query: { 
        id: 1 
      }
    }
    
    const resultData = await appActions(it_store, setData).requestData("/test", options, false)
    expect(resultData.hello).toBe("world");
    expect(setData).toHaveBeenCalledTimes(2);

  });

  it('requestData 401', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 401,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    const it_store = update(app_store, {
      session: {$merge: {
        configServer: false
      }}
    })
    const options = {
      token: "token",
      headers: {}
    }
    
    const resultData = await appActions(it_store, setData).requestData("/test", options, true)
    expect(resultData.error.message).toBe("Unauthorized");
    expect(setData).toHaveBeenCalledTimes(1);

  });

  it('requestData 400', async () => {
    const res = new Response('{"code":400, "message":"errordata"}', {
      status: 400,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    const it_store = update(app_store, {
      session: {$merge: {
        configServer: false
      }}
    })
    const options = {}
    
    const resultData = await appActions(it_store, setData).requestData("/test", options, true)
    expect(resultData.error.message).toBe("errordata");
    expect(setData).toHaveBeenCalledTimes(0);

  });

  it('requestData error', async () => {
    const res = new Response('', {
      status: 404,
      statusText: 'Not Found',
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    const it_store = update(app_store, {
      session: {$merge: {
        configServer: false
      }}
    })
    const options = {}
    
    let resultData = await appActions(it_store, setData).requestData("/test", options, true)
    expect(resultData.error.message).toBe("Not Found");
    expect(setData).toHaveBeenCalledTimes(0);

    resultData = await appActions(it_store, setData).requestData("/test", options, false)
    expect(resultData.error.message).toBe("Not Found");
    expect(setData).toHaveBeenCalledTimes(2);

  });

  it('getAuditFilter', () => {
    const setData = jest.fn()
    const it_store = update(app_store, {
      login: {$merge: {
        data: {
          audit: [
            {
              inputfilter: 108,
              inputfilterName: 'update',
              nervatype: 10,
              nervatypeName: 'customer',
              subtype: null,
              subtypeName: null,
              supervisor: 1
            },
            {
              inputfilter: 107,
              inputfilterName: 'readonly',
              nervatype: 31,
              nervatypeName: 'trans',
              subtype: 62,
              subtypeName: 'inventory',
              supervisor: 0
            },
            {
              inputfilter: 106,
              inputfilterName: 'disabled',
              nervatype: 28,
              nervatypeName: 'report',
              subtype: 6,
              subtypeName: null,
              supervisor: 0
            },
            {
              inputfilter: 106,
              inputfilterName: 'disabled',
              nervatype: 18,
              nervatypeName: 'menu',
              subtype: 1,
              subtypeName: 'nextNumber',
              supervisor: 0
            }
          ]
        }
      }}
    })
    
    let audit = appActions(it_store, setData).getAuditFilter("trans", "inventory")
    expect(audit[0]).toBe("readonly");
    audit = appActions(it_store, setData).getAuditFilter("menu", "nextNumber")
    expect(audit[0]).toBe("disabled");
    audit = appActions(it_store, setData).getAuditFilter("report", 6)
    expect(audit[0]).toBe("disabled");
    audit = appActions(it_store, setData).getAuditFilter("customer")
    expect(audit[0]).toBe("update");
    audit = appActions(it_store, setData).getAuditFilter("product")
    expect(audit[0]).toBe("all");

  });

  it('createHistory trans', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    const it_store = update(app_store, {
      edit: {$merge: {
        current: {
          type: "trans",
          transtype: "invoice",
          item: {
            id: 5,
            transnumber: "DMINV/00001"
          }
        },
        template: {
          options: {
            title: "INVOICE",
            title_field: "transnumber"
          }
        }
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }}
    })
    
    await appActions(it_store, setData).createHistory("save")
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('createHistory customer', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    const it_store = update(app_store, {
      edit: {$merge: {
        current: {
          type: "customer",
          item: {
            id: 1,
            custnumber: "CUST/00001"
          }
        },
        template: {
          options: {
            title: "CUSTOMER",
            title_field: "custnumber"
          }
        }
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }},
      bookmark: {$merge: {
        history: {
          employee_id: 1,
          section: "history",
          cfgroup: "2022-01-03T22:27:00+02:00",
          cfname: 1,
          cfvalue: "[{\"datetime\":\"2022-01-03T22:21:32+02:00\",\"type\":\"save\",\"ntype\":\"trans\",\"transtype\":\"invoice\",\"id\":5,\"title\":\"INVOICE | DMINV/00001\"}]",
        }
      }}
    })
    
    await appActions(it_store, setData).createHistory("save")
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('createHistory error', async () => {
    const res = new Response('{"code":400, "message":"errordata"}', {
      status: 400,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    const it_store = update(app_store, {
      edit: {$merge: {
        current: {
          type: "customer",
          item: {
            id: 1,
          }
        },
        template: {
          options: {
            title: "CUSTOMER",
          }
        }
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }},
      bookmark: {$merge: {
        history: {
          employee_id: 1,
          section: "history",
          cfgroup: "2022-01-03T22:27:00+02:00",
          cfname: 1,
          cfvalue: "[{},{},{},{},{},{},{},{},{},{}]",
        }
      }}
    })
    
    await appActions(it_store, setData).createHistory("save")
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('loadBookmark 1.', async () => {
    const res = new Response('[{"section":"history"},{"section":"bookmark"}]', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn((key, data, callback)=>{ if(callback){callback()} })
    
    await appActions(app_store, setData).loadBookmark({ token:"token", user_id: 1, callback: ()=>{} })
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('loadBookmark 2.', async () => {
    const res = new Response('[{"section":"bookmark"}]', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn((key, data, callback)=>{ if(callback){callback()} })
    
    await appActions(app_store, setData).loadBookmark({ token:"token", user_id: 1 })
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('loadBookmark error', async () => {
    const res = new Response('{"code":400, "message":"errordata"}', {
      status: 400,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn()
    
    await appActions(app_store, setData).loadBookmark({ token:"token", user_id: 1, callback: ()=>{} })
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('showHelp', () => {
    appActions(app_store, jest.fn()).showHelp("help")
  });

  it('saveBookmark browser', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn((key, data, callback)=>{ 
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(
          <Suspense fallback={<InputBox {...data.modalForm.props} />}>
            {data.modalForm}
          </Suspense>, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        
        // onCancel
        const btn_cancel = getById(container, 'btn_cancel')
        btn_cancel.dispatchEvent(new MouseEvent('click', {bubbles: true}));
      }
      if(callback){callback()} 
    })
    const it_store = update(app_store, {
      search: {$merge: {
        vkey: "customer",
        filters: {
          CustomerView: [
          ],
        },
        columns: {
          CustomerView: {
            custnumber: true,
            custname: true,
            address: true,
          },
        },
        view: "CustomerView",
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }},
    })
    
    appActions(it_store, setData).saveBookmark(['browser', 'Customer Data'])
    expect(setData).toHaveBeenCalledTimes(4);

    appActions(it_store, setData).saveBookmark(['browser', ''])
    expect(setData).toHaveBeenCalledTimes(7);

  });

  it('saveBookmark editor trans 1.', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn((key, data, callback)=>{ 
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(
          <Suspense fallback={<InputBox {...data.modalForm.props} />}>
            {data.modalForm}
          </Suspense>, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        
      }
      if(callback){callback()} 
    })
    const it_store = update(app_store, {
      edit: {$merge: {
        current: {
          type: "trans",
          transtype: "invoice",
          item: {
            id: 5,
            transnumber: "DMINV/00001",
            transdate: "2020-12-10",
          }
        },
        dataset: {
          trans: [
            {
              custname: "First Customer Co."
            }
          ]
        }
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }},
    })
    
    appActions(it_store, setData).saveBookmark(['editor', 'trans', 'transnumber'])
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('saveBookmark editor trans 2.', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn((key, data, callback)=>{ 
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(
          <Suspense fallback={<InputBox {...data.modalForm.props} />}>
            {data.modalForm}
          </Suspense>, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        
      }
      if(callback){callback()} 
    })
    const it_store = update(app_store, {
      edit: {$merge: {
        current: {
          type: "trans",
          transtype: "receipt",
          item: {
            id: 5,
            transnumber: "DMINV/00001",
            transdate: "2020-12-10",
          }
        },
        dataset: {
          trans: [
            {
              custname: null
            }
          ]
        }
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }},
    })
    
    appActions(it_store, setData).saveBookmark(['editor', 'trans', 'transnumber'])
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('saveBookmark editor customer', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn((key, data, callback)=>{ 
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(
          <Suspense fallback={<InputBox {...data.modalForm.props} />}>
            {data.modalForm}
          </Suspense>, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        
      }
      if(callback){callback()} 
    })
    const it_store = update(app_store, {
      edit: {$merge: {
        current: {
          type: "customer",
          transtype: "",
          item: {
            id: 2,
            custnumber: "DMCUST/00001",
            custname: "First Customer Co.",
          }
        },
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }},
    })
    
    appActions(it_store, setData).saveBookmark(['editor', 'customer', 'custname', 'custnumber'])
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('saveBookmark error', async () => {
    const res = new Response('{"code":400, "message":"errordata"}', {
      status: 400,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setData = jest.fn((key, data, callback)=>{ 
      if((key === "current") && data.modalForm ){
        const container = document.createElement('div');
        ReactDOM.render(
          <Suspense fallback={<InputBox {...data.modalForm.props} />}>
            {data.modalForm}
          </Suspense>, container);

        // onOK
        const btn_ok = getById(container, 'btn_ok')
        btn_ok.dispatchEvent(new MouseEvent('click', {bubbles: true}));
        
      }
      if(callback){callback()} 
    })
    const it_store = update(app_store, {
      edit: {$merge: {
        current: {
          type: "trans",
          transtype: "invoice",
          item: {
            id: 5,
            transnumber: "DMINV/00001",
            transdate: "2020-12-10",
          }
        },
        dataset: {
          trans: [
            {
              custname: "First Customer Co."
            }
          ]
        }
      }},
      login: {$merge: {
        data: {
          employee: {
            id: 1
          }
        }
      }},
    })
    
    appActions(it_store, setData).saveBookmark(['editor', 'trans', 'transnumber'])
    expect(setData).toHaveBeenCalledTimes(3);

  });

  it('getDataFilter', () => {
    const it_store = update(app_store, {
      login: {$merge: {
        data: {
          audit: []
        }
      }}
    })
    let result = appActions(it_store, jest.fn()).getDataFilter("transitem", [])
    expect(result.length).toBe(0)
    result = appActions(it_store, jest.fn()).getDataFilter("transpayment", [])
    expect(result.length).toBe(0)
    result = appActions(it_store, jest.fn()).getDataFilter("transmovement", [], "")
    expect(result.length).toBe(0)
    result = appActions(it_store, jest.fn()).getDataFilter("transmovement", [], "InventoryView")
    expect(result.length).toBe(0)
  })

  it('getDataFilter disabled', () => {
    const it_store = update(app_store, {
      login: {$merge: {
        data: {
          audit: [
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'offer', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'order', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'worksheet', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'rent', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'invoice', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'bank', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'cash', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'delivery', supervisor: 0 },
            { inputfilterName: 'disabled', nervatypeName: 'trans', subtypeName: 'inventory', supervisor: 0 },
          ]
        }
      }}
    })
    let result = appActions(it_store, jest.fn()).getDataFilter("transitem", [])
    expect(result.length).toBe(10)
    result = appActions(it_store, jest.fn()).getDataFilter("transpayment", [])
    expect(result.length).toBe(4)
    result = appActions(it_store, jest.fn()).getDataFilter("transmovement", [], "")
    expect(result.length).toBe(4)
    result = appActions(it_store, jest.fn()).getDataFilter("transmovement", [], "InventoryView")
    expect(result.length).toBe(0)
  })

  it('getUserFilter', () => {
    let it_store = update(app_store, {
      login: {$merge: {
        data: {
          employee: {
            id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
            usergroupName: 'admin'
          },
          audit: [],
        }
      }},
    })
    let filter = appActions(it_store, jest.fn()).getUserFilter("missing")
    expect(filter.params.length).toBe(0)
    filter = appActions(it_store, jest.fn()).getUserFilter("customer")
    expect(filter.params.length).toBe(0)

    it_store = update(it_store, {login: {data: {$merge: {
      transfilterName: "usergroup"
    }}}})
    filter = appActions(it_store, jest.fn()).getUserFilter("customer")
    expect(filter.params.length).toBe(0)
    filter = appActions(it_store, jest.fn()).getUserFilter("transitem")
    expect(filter.params.length).toBe(1)

    it_store = update(it_store, {login: {data: {$merge: {
      transfilterName: "own"
    }}}})
    filter = appActions(it_store, jest.fn()).getUserFilter("customer")
    expect(filter.params.length).toBe(0)
    filter = appActions(it_store, jest.fn()).getUserFilter("transitem")
    expect(filter.params.length).toBe(1)
  })

  it('quickSearch', async () => {
    const it_store = update(app_store, {
      login: {$merge: {
        data: {
          employee: {
            id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
            usergroupName: 'admin'
          },
          audit: [],
          transfilterName: "usergroup"
        }
      }},
    })
    let data = await appActions(it_store, jest.fn()).quickSearch("customer", "")
    expect(data).toBeDefined()

    data = await appActions(it_store, jest.fn()).quickSearch("transitem", "item")
    expect(data).toBeDefined()
  })

  it('onSelector', async () => {
    const res = new Response('{"hello":"world"}', {
      status: 200,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));

    const setSelector = jest.fn()
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    const it_store = update(app_store, {
      login: {$merge: {
        data: {
          employee: {
            id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
            usergroupName: 'admin'
          },
          audit: [],
          transfilterName: "usergroup"
        }
      }},
    })
    const formProps = await appActions(it_store, setData).onSelector("customer", "", setSelector)
    formProps.onSearch("")
    formProps.onClose()
    formProps.onSelect({},"")
    expect(setData).toHaveBeenCalledTimes(3);
    expect(setSelector).toHaveBeenCalledTimes(1);
  })

  it('onSelector error', async () => {
    const res = new Response('{"code":400, "message":"errordata"}', {
      status: 400,
      headers: {
        'Content-type': 'application/json',
      },
    });
    window.fetch.mockReturnValue(Promise.resolve(res));
    const setSelector = jest.fn()
    const setData = jest.fn((key, data, callback)=>{
      if(callback){callback()}
    })
    const it_store = update(app_store, {
      login: {$merge: {
        data: {
          employee: {
            id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
            usergroupName: 'admin'
          },
          audit: [],
          transfilterName: "usergroup"
        }
      }},
    })
    const formProps = await appActions(it_store, setData).onSelector("customer", "", setSelector)
    formProps.onSearch("")
    expect(setData).toHaveBeenCalledTimes(1);
  })

})