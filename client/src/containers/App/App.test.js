import { render, queryByAttribute } from '@testing-library/react'
import { create } from 'react-test-renderer';

import { toast } from 'react-toastify';
import { Default as EditDefault } from 'components/SideBar/Edit/Edit.stories'
import InputBox from 'components/Modal/InputBox'

import App from './index';
import { guid, request, appActions } from './actions'

jest.mock("./actions");
jest.mock("react-toastify");

const getById = queryByAttribute.bind(null, 'id');
const { location } = window;

let local = {
  lang: "en"
}

describe('<App />', () => {

  beforeEach(() => {
    delete window.location;
    window.location = { 
      replace: jest.fn(),
      pathname: "/"
    };
    global.Storage.prototype.setItem = jest.fn((key, value) => {
      local[key] = value
    })
    global.Storage.prototype.getItem = jest.fn((key) => local[key])

    appActions.mockReturnValue({
      getText: jest.fn(),
      tokenValidation: jest.fn(),
      setCodeToken: jest.fn(),
      resultError: jest.fn(),
    })
    request.mockReturnValue({})
    guid.mockReturnValue({})
    toast.mockReturnValue({
      error: jest.fn(),
      warning: jest.fn(),
      success: jest.fn(),
      info: jest.fn(),
    })
  });

  afterEach(() => {
    jest.clearAllMocks();
    window.location = location;
    global.Storage.prototype.setItem.mockReset()
    global.Storage.prototype.getItem.mockReset()
  });

  it('renders without crashing', () => {
    const { container } = render(<App />);
    expect(getById(container, 'login')).toBeDefined();

    const getPath = App.prototype.getPath
    App.prototype.getPath = () => {
      return["hash", { access_token: "ABC012" }]
    }
    render(<App />)

    App.prototype.getPath = () => {
      return["search", { code: "ABC012" }]
    }
    render(<App />)
    App.prototype.getPath = getPath

  });

  it('renders in the Search state', () => {
    const data = {
      login: {
        username: 'admin',
        data: {
          audit_filter: EditDefault.args.auditFilter,
        }
      },
      current: {
        module: "search",
        request: true
      }
    }
    const { container } = render(<App data={data} />);
    expect(getById(container, 'btn_search')).toBeDefined();

  });

  it('renders in the Edit state', () => {
    const data = {
      login: {
        username: 'admin',
        data: {
          audit_filter: EditDefault.args.auditFilter,
          edit_new: EditDefault.args.newFilter
        }
      },
      current: {
        module: "edit"
      }
    }
    render(<App data={data} />);

  });

  it('renders in the Setting state', () => {
    const data = {
      login: {
        username: 'admin',
        data: {
          audit_filter: EditDefault.args.auditFilter,
          edit_new: EditDefault.args.newFilter
        }
      },
      current: {
        module: "setting"
      }
    }
    render(<App data={data} />);

  });

  it('renders in the Template state', () => {
    const data = {
      login: {
        username: 'admin',
        data: {
          audit_filter: EditDefault.args.auditFilter,
          edit_new: EditDefault.args.newFilter
        }
      },
      current: {
        module: "template"
      }
    }
    render(<App data={data} />);

  });

  it('renders in the Modal state', () => {
    const data = {
      login: {
        username: 'admin',
        data: {
          audit_filter: EditDefault.args.auditFilter,
          edit_new: EditDefault.args.newFilter
        }
      },
      current: {
        modalForm: <InputBox />
      }
    }
    render(<App data={data} />);

  });

  it('setData', () => {
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.setData("test", "value", ()=>{})
    app.setData("test", "value")
    app.setData("ui", {})
    app.setData("ui", {}, ()=>{})
    app.setData()
  })

  it('onScroll', () => {
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.onScroll()
    app.onScroll()
  })

  it('getPath', () => {
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.getPath({hash: "abc"})
    app.getPath({search: "abc"})
    app.getPath({pathname: "abc/abc"})
  })

  it('onResize', () => {
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.onResize()
    app.onResize()
  })

  it('loadConfig', () => {
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.loadConfig()
    request.mockReturnValue({
      locales: {
        de: {}
      }
    })
    app.loadConfig()

    request.mockImplementation(() => {
      throw new Error();
    })
    app.loadConfig(false)
  })

  it('loadConfig setCodeToken', () => {
    window.location = { 
      replace: jest.fn(),
      pathname: "/",
      search: "?code=g0ZGZmNjVmOWIjNTk2NTk4ZTYyZGI3",
      hash: ""
    };
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.loadConfig()
  })

  it('loadConfig tokenValidation', () => {
    window.location = { 
      replace: jest.fn(),
      pathname: "/",
      search: "",
      hash: "#access_token=g0ZGZmNjVmOWIjNTk2NTk4ZTYyZGI3"
    };
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.loadConfig()
  })

})
