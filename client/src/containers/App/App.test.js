import { render, queryByAttribute } from '@testing-library/react'
import { create } from 'react-test-renderer';

import { Default as EditDefault } from 'components/SideBar/Edit/Edit.stories'
import InputBox from 'components/Modal/InputBox'

import App from './index';
import { guid, request, appActions } from './actions'

jest.mock("./actions");
const getById = queryByAttribute.bind(null, 'id');
const { location } = window;

let local = {
  lang: "en"
}

describe('<App />', () => {

  beforeEach(() => {
    delete window.location;
    window.location = { 
      assign: jest.fn(),
      pathname: "/"
    };

    global.Storage.prototype.setItem = jest.fn((key, value) => {
      local[key] = value
    })
    global.Storage.prototype.getItem = jest.fn((key) => local[key])

    appActions.mockReturnValue({
      getText: jest.fn(),
      resultError: jest.fn(),
    })
    request.mockReturnValue({})
    guid.mockReturnValue({})
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

  it('setHashToken', () => {
    const testRenderer = create(<App />);
    const app = testRenderer.getInstance()
    app.setHashToken({ access_token: "" })
    app.setHashToken({ path: "/", access_token: "" })
  })

  it('setCodeToken', () => {
    const testRenderer = create(<App />);
    let app = testRenderer.getInstance()
    request.mockReturnValue({
      access_token: "access_token"
    })
    app.loadConfig = ()=>({
      provider_token_callback: "/",
      provider_client_id: "client_id",
      provider_client_secret: "client_secret",
      provider_token_login: "/"
    })
    app.setCodeToken({ code: "code", path: "/" })
    app.setCodeToken({})
  })

  it('setCodeToken missing token', () => {
    const testRenderer = create(<App />);
    let app = testRenderer.getInstance()
    request.mockReturnValue({
    })
    app.loadConfig = ()=>({
      provider_token_callback: "/",
      provider_client_id: "client_id",
      provider_client_secret: "client_secret",
      provider_token_login: "/"
    })
    app.setCodeToken({ code: "code", path: "/" })
  })

  it('setCodeToken error', () => {
    const testRenderer = create(<App />);
    let app = testRenderer.getInstance()
    request.mockImplementation(() => {
      throw new Error();
    })
    app.loadConfig = ()=>({
      provider_token_login: "/",
      provider_token_callback: "/"
    })
    app.setCodeToken({})
    
    app.loadConfig = ()=>({
      provider_token_callback: "/"
    })
    app.setCodeToken({})
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
    app.loadConfig(true)
    app.loadConfig(false)
    request.mockImplementation(() => {
      throw new Error();
    })
    app.loadConfig(false)
  })

})
