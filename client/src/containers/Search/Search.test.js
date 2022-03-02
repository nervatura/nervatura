import React from 'react';
import { render, fireEvent, queryByAttribute, waitFor, screen } from '@testing-library/react'
import update from 'immutability-helper';

import Search from './index';
import { store as app_store  } from 'config/app'
import { AppProvider } from 'containers/App/context'
import { Default as SelectorData } from 'components/Modal/Selector/Selector.stories'
import { Default as BrowserData, Filters, HavingFilter, Columns, FormActions } from 'components/Browser/Browser.stories'
import { Default as SideBarData } from 'components/SideBar/Search/Search.stories'

import { appActions, getSql, saveToDisk } from 'containers/App/actions'
import { searchActions } from 'containers/Search/actions'

jest.mock("containers/App/actions");
jest.mock("containers/Search/actions");

const getById = queryByAttribute.bind(null, 'id');

const store = update(app_store, {$merge: {
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
      audit_filter: SideBarData.args.auditFilter,
      edit_new: SideBarData.args.newFilter
    }
  }
}})

describe('<Search />', () => {

  beforeEach(() => {
    appActions.mockReturnValue({
      getText: BrowserData.args.getText,
      showHelp: jest.fn(),
      requestData: jest.fn(async () => ({})),
      resultError: jest.fn(),
      saveBookmark: jest.fn(),
      quickSearch: jest.fn(async () => ({})),
      getDataFilter: jest.fn(() => ([""])),
      getUserFilter: jest.fn(() => ({ where: [""], params:[] })),
    })
    searchActions.mockReturnValue({
      showServerCmd: jest.fn(),
      showBrowser: jest.fn(),
      getFilterWhere: jest.fn(() => ([])),
      defaultFilterValue: jest.fn(() => (""))
    })
    getSql.mockReturnValue({
      sql: "",
      prmCount: 1
    })
    saveToDisk.mockReturnValue()
  });
  
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('renders without crashing', async () => {
    const setData = jest.fn()
    const it_store = update(store, {
      search: {$merge: {
        seltype: "selector",
        result: SelectorData.args.result, 
        vkey: null, 
        qview: "customer", 
        qfilter: "", 
        page: 0
      }}
    })

    const { container, rerender } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );
    expect(getById(container, 'selector_customer')).toBeDefined();

    //quickSearch
    const btn_search = getById(container, 'btn_search')
    fireEvent.click(btn_search)
    await waitFor(() =>{ 
      expect(setData).toHaveBeenCalledTimes(1);
    })
    
    //editRow
    const row_item = screen.getAllByRole('row')[1]
    fireEvent.click(row_item)

    //SideBar quickView
    const btn_transitem = getById(container, 'btn_view_transitem')
    fireEvent.click(btn_transitem)
    expect(setData).toHaveBeenCalledTimes(4);

    rerender(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    )
  });

  it('useEffect', async () => {
    const setData = jest.fn((key, data, callback)=>{ 
      if(callback){callback()} 
    })
    const it_store = update(store, {
      current: {$merge: {
        content: []
      }}
    })

    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );
    expect(getById(container, 'selector_customer')).toBeDefined();

  })

  it('selector quickSearch err. and editRow servercmd', () => {
    appActions.mockReturnValue({
      getText: BrowserData.args.getText,
      showHelp: jest.fn(),
      requestData: jest.fn(async () => ({})),
      resultError: jest.fn(),
      saveBookmark: jest.fn(),
      quickSearch: jest.fn(async () => ({ error: {} })),
      getDataFilter: jest.fn(),
      getUserFilter: jest.fn(),
    })
    searchActions.mockReturnValue({
      showServerCmd: jest.fn(),
      showBrowser: jest.fn(),
      getFilterWhere: jest.fn(),
    })
    const setData = jest.fn()
    const it_store = update(store, {
      search: {$merge: {
        seltype: "selector",
        result: [
          { description: 'Server function example', id: 'servercmd//1', label: 'Server function example' },
          { description: 'Internet URL example', id: 'servercmd//2', label: 'Internet URL example' }
        ], 
        vkey: null, 
        qview: "servercmd", 
        qfilter: "", 
        page: 0,
        group_key: "office"
      }}
    })

    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );
    
    //quickSearch
    const btn_search = getById(container, 'btn_search')
    fireEvent.click(btn_search)
    expect(setData).toHaveBeenCalledTimes(0);

    //editRow
    const row_item = screen.getAllByRole('row')[1]
    fireEvent.click(row_item)

    //SideBar checkEditor
    const btn_printqueue = getById(container, 'btn_printqueue')
    fireEvent.click(btn_printqueue)

  })

  it('browser', async () => {
    let local = {
      paginationPage: 2
    }
    global.URL.createObjectURL = jest.fn();
    global.Storage.prototype.getItem = jest.fn((key) => local[key])
    const setData = jest.fn((key, data, callback)=>{ 
      if((key === "current") && data.modalForm ){
        const { container } = render(data.modalForm);
        // onOK
        const btn_ok = getById(container, 'btn_ok')
        fireEvent.click(btn_ok)
      }
      if(callback){callback()} 
    })
    const it_store = update(store, {
      search: {$merge: {
        seltype: "browser",
        ...BrowserData.args.data
      }}
    })

    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );
    expect(getById(container, 'browser_customer')).toBeDefined();

    //browserView
    const btn_search = getById(container, 'btn_search')
    fireEvent.click(btn_search)
    await waitFor(() =>{ 
      expect(setData).toHaveBeenCalledTimes(1);
    })

    //bookmarkSave
    const btn_bookmark = getById(container, 'btn_bookmark')
    fireEvent.click(btn_bookmark)

    //exportResult
    const btn_export = getById(container, 'btn_export')
    fireEvent.click(btn_export)

    //showTotal
    const btn_total = getById(container, 'btn_total')
    fireEvent.click(btn_total)
    expect(setData).toHaveBeenCalledTimes(3);

    //changeData
    const btn_columns = getById(container, 'btn_columns')
    fireEvent.click(btn_columns)
    expect(setData).toHaveBeenCalledTimes(4);

    //onEdit
    const row_item = getById(container, 'edit_customer//2')
    fireEvent.click(row_item)
    expect(setData).toHaveBeenCalledTimes(5);

    //onPage
    const page_2 = screen.getByText("2")
    fireEvent.click(page_2)
    expect(setData).toHaveBeenCalledTimes(6);
    
    global.Storage.prototype.getItem.mockReset()
  })

  it('browserView err.', async () => {
    appActions.mockReturnValue({
      getText: BrowserData.args.getText,
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
      quickSearch: jest.fn(async () => ({})),
      getDataFilter: jest.fn(() => ([])),
      getUserFilter: jest.fn(() => ({ where: [], params:[] })),
    })
    searchActions.mockReturnValue({
      showServerCmd: jest.fn(),
      showBrowser: jest.fn(),
      getFilterWhere: jest.fn(() => ([])),
    })
    const setData = jest.fn()
    const it_store = update(store, {
      search: {$merge: {
        seltype: "browser",
        ...HavingFilter.args.data
      }}
    })

    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );
    expect(getById(container, 'browser_customer')).toBeDefined();

    //browserView
    const btn_search = getById(container, 'btn_search')
    fireEvent.click(btn_search)
    await waitFor(() =>{ 
      expect(setData).toHaveBeenCalledTimes(0);
    })

  })

  it('showTotal deffield, editFilter, deleteFilter, addFilter, onEdit', async () => {
    const setData = jest.fn()
    const it_store = update(store, {
      search: {$merge: {
        seltype: "browser",
        ...Filters.args.data,
        page: 0
      }}
    })
    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );

    //showTotal
    const btn_total = getById(container, 'btn_total')
    fireEvent.click(btn_total)
    expect(setData).toHaveBeenCalledTimes(1);

    //editFilter
    const filter_value_string = getById(container, 'filter_value_3')
    fireEvent.change(filter_value_string, {target: {value: "red"}})
    expect(setData).toHaveBeenCalledTimes(2);

    const filter_name = getById(container, 'filter_name_0')
    fireEvent.change(filter_name, {target: {value: "sample_customer_date"}})
    expect(setData).toHaveBeenCalledTimes(3);

    fireEvent.change(filter_name, {target: {value: "sample_customer_valuelist"}})
    expect(setData).toHaveBeenCalledTimes(4);

    fireEvent.change(filter_name, {target: {value: "custname"}})
    expect(setData).toHaveBeenCalledTimes(5);

    const filter_type = getById(container, 'filter_type_0')
    fireEvent.change(filter_type, {target: {value: "==N"}})
    expect(setData).toHaveBeenCalledTimes(6);

    //deleteFilter
    const btn_delete = getById(container, 'btn_delete_filter_0')
    fireEvent.click(btn_delete)
    expect(setData).toHaveBeenCalledTimes(7);

    //addFilter
    const btn_filter = getById(container, 'btn_filter')
    fireEvent.click(btn_filter)
    expect(setData).toHaveBeenCalledTimes(8);

    //onEdit
    const link_cell = getById(container, 'link_31')
    fireEvent.click(link_cell)
    expect(setData).toHaveBeenCalledTimes(9);

  })

  it('showBrowser, setColumns', async () => {
    const setData = jest.fn()
    const it_store = update(store, {
      search: {$merge: {
        seltype: "browser",
        ...Columns.args.data
      }}
    })

    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );

    //showBrowser
    const view_item = getById(container, 'view_CustomerAddressView')
    fireEvent.click(view_item)

    //setColumns
    const col_surname = getById(container, 'col_surname')
    fireEvent.click(col_surname)
    expect(setData).toHaveBeenCalledTimes(1);

    const col_status = getById(container, 'col_status')
    fireEvent.click(col_status)
    expect(setData).toHaveBeenCalledTimes(2);

  })

  it('setFormActions', async () => {
    const setData = jest.fn()
    const it_store = update(store, {
      search: {$merge: {
        seltype: "browser",
        ...FormActions.args.data
      }}
    })

    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Search />
      </AppProvider>
    );

    //setFormActions
    const btn_actions_new = getById(container, 'btn_actions_new')
    fireEvent.click(btn_actions_new)
    expect(setData).toHaveBeenCalledTimes(1);

  })

})
