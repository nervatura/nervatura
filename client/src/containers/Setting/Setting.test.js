import React from 'react';
import { render, fireEvent, queryByAttribute } from '@testing-library/react'
import update from 'immutability-helper';

import Setting from './index';
import { store as app_store  } from 'config/app'
import { AppProvider } from 'containers/App/context'
import { Default as FormData } from 'components/Setting/Form/Form.stories'
import { Default as ViewData } from 'components/Setting/View/View.stories'
import { Default as SideBarData } from 'components/SideBar/Setting/Setting.stories'

import { appActions } from 'containers/App/actions'
import { settingActions } from 'containers/Setting/actions'

jest.mock("containers/App/actions");
jest.mock("containers/Setting/actions");

const getById = queryByAttribute.bind(null, 'id');

const store = update(app_store, {$merge: {
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
      audit_filter: SideBarData.args.auditFilter,
    }
  }
}})

describe('<Setting />', () => {
  beforeAll(() => {
   
  });

  beforeEach(() => {
    appActions.mockReturnValue({
      getText: ViewData.args.getText,
      showHelp: jest.fn(),
    })
    settingActions.mockReturnValue({
      setViewActions: jest.fn(),
      checkSetting: jest.fn(),
    })
  });
  
  afterAll(() => {
    jest.clearAllMocks();
  });

  it('renders in the Main state', () => {
    const setData = jest.fn((key, data, callback)=>{ if(callback){callback()} })

    const it_store = update(store, {$merge: {
      setting: {
        ...ViewData.args.data,
        type: "setting",
        group_key: "group_database",
      },
      current: {
        content: { type: 'setting' }
      }
    }})

    const { container, rerender } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Setting id="setting" paginationPage={3} />
      </AppProvider>
    );
    expect(getById(container, 'setting')).toBeDefined();

    //onEvent - setting
    const row_edit = getById(container, 'row_edit_0')
    fireEvent.click(row_edit)

    //companyForm
    const cmd_company = getById(container, 'cmd_company')
    fireEvent.click(cmd_company)
    expect(setData).toHaveBeenCalledTimes(2);

    //changeData
    const btn_next = getById(container, 'btn_next')
    fireEvent.click(btn_next)
    expect(setData).toHaveBeenCalledTimes(3);

    rerender(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Setting id="setting" />
      </AppProvider>
    )

  });

  it('renders Form state', () => {
    const setData = jest.fn((key, data, callback)=>{ if(callback){callback()} })

    const it_store = update(store, {$merge: {
      setting: {
        ...FormData.args.data,
        panel: {
          save: true,
          delete: true,
          new: true,
          help: "help"
        }
      }
    }})
    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Setting id="setting" />
      </AppProvider>
    );
    expect(getById(container, 'setting')).toBeDefined();

    //onEvent - app
    const cmd_help = getById(container, 'cmd_help')
    fireEvent.click(cmd_help)

  })

  it('renders in the Empty state', () => {
    const setData = jest.fn()
    render(
      <AppProvider value={{ data: store, setData: setData }}>
        <Setting id="setting" />
      </AppProvider>
    );

  });

})