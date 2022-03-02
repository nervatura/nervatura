import React from 'react';
import { render, fireEvent, queryByAttribute } from '@testing-library/react'
import update from 'immutability-helper';

import MenuBar from './index';
import { store as app_store  } from 'config/app'
import { AppProvider } from 'containers/App/context'
import { BookmarkData } from 'components/Modal/Bookmark/Bookmark.stories'

import { appActions } from 'containers/App/actions'

jest.mock("containers/App/actions");

const getById = queryByAttribute.bind(null, 'id');

const store = update(app_store, {$merge: {
  bookmark: BookmarkData.args.bookmark,
  login: {
    data: {
      employee: {
        id: 1, empnumber: 'admin', username: 'admin', usergroup: 114,
        usergroupName: 'admin'
      },
    }
  }
}})

const scrollTo = jest.fn()

describe('<MenuBar />', () => {

  beforeEach(() => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      signOut: jest.fn(),
      showHelp: jest.fn(),
      requestData: jest.fn(async () => ({})),
      resultError: jest.fn(),
      loadBookmark: jest.fn(({user_id, callback})=>{ 
        if(callback){callback()} 
      }),
    })
    Object.defineProperty(global.window, 'scrollTo', { value: scrollTo });
  });
  
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('renders without crashing', () => {
    const setData = jest.fn((key, data, callback)=>{ 
      if(callback){callback()} 
    })
    let it_store = update(store, {
      current: {$merge: {
        scrollTop: true
      }}
    })

    const { container, rerender } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <MenuBar />
      </AppProvider>
    );
    expect(getById(container, 'mnu_sidebar')).toBeDefined();

    const mnu_scroll = getById(container, 'mnu_scroll')
    fireEvent.click(mnu_scroll)
    expect(scrollTo).toHaveBeenCalledTimes(1);

    const mnu_sidebar = getById(container, 'mnu_sidebar')
    fireEvent.click(mnu_sidebar)
    expect(setData).toHaveBeenCalledTimes(1);

    const mnu_logout_large = getById(container, 'mnu_logout_large')
    fireEvent.click(mnu_logout_large)
    expect(setData).toHaveBeenCalledTimes(1);

    const mnu_help_large = getById(container, 'mnu_help_large')
    fireEvent.click(mnu_help_large)
    expect(setData).toHaveBeenCalledTimes(1);

    const mnu_setting_large = getById(container, 'mnu_setting_large')
    fireEvent.click(mnu_setting_large)
    expect(setData).toHaveBeenCalledTimes(4);

    const mnu_edit_large = getById(container, 'mnu_edit_large')
    fireEvent.click(mnu_edit_large)
    expect(setData).toHaveBeenCalledTimes(5);


    it_store = update(it_store, {
      current: {$merge: {
        side: "show"
      }}
    })
    rerender(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <MenuBar />
      </AppProvider>
    )
    fireEvent.click(mnu_sidebar)
    expect(setData).toHaveBeenCalledTimes(6);
  });

  it('showBookmarks - bookmark', () => {
    const setData = jest.fn((key, data, callback)=>{ 
      if((key === "current") && data.modalForm ){
        const { container } = render(data.modalForm);
        // onSelect bookmark element
        let row_item = getById(container, 'row_item_1')
        fireEvent.click(row_item)
        // onClose
        const close = getById(container, 'closeIcon')
        fireEvent.click(close)
      }
      if(callback){
        callback()
      }
    })

    const { container } = render(
      <AppProvider value={{ data: store, setData: setData }}>
        <MenuBar bookmarkView="bookmark" />
      </AppProvider>
    );

    const mnu_bookmark_large = getById(container, 'mnu_bookmark_large')
    fireEvent.click(mnu_bookmark_large)
    expect(setData).toHaveBeenCalledTimes(4);
  })

  it('showBookmarks - history', () => {
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const { container } = render(data.modalForm);
        // onSelect bookmark element
        let row_item = getById(container, 'row_item_0')
        fireEvent.click(row_item)
      }
      if(callback){
        callback()
      }
    })

    const { container } = render(
      <AppProvider value={{ data: store, setData: setData }}>
        <MenuBar bookmarkView="history" />
      </AppProvider>
    );

    const mnu_bookmark_large = getById(container, 'mnu_bookmark_large')
    fireEvent.click(mnu_bookmark_large)
    expect(setData).toHaveBeenCalledTimes(3);
  })

  it('showBookmarks - onDelete cancel+error', () => {
    appActions.mockReturnValue({
      getText: jest.fn(),
      requestData: jest.fn(async () => ({ error: {} })),
      resultError: jest.fn(),
    })
    let cancel = false
    let ok = false
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const { container } = render(data.modalForm);
        // showBookmarks - delete icon
        let row_item = getById(container, 'row_delete_1')
        if(row_item){
          fireEvent.click(row_item)
        }
        // InputBox
        let btn_cancel = getById(container, 'btn_cancel')
        if(btn_cancel && !cancel){
          cancel = true
          fireEvent.click(btn_cancel)
        }
        let btn_ok = getById(container, 'btn_ok')
        if(btn_ok && !ok){
          ok = true
          fireEvent.click(btn_ok)
        }
      }
      if(callback){
        callback()
      }
    })

    const { container } = render(
      <AppProvider value={{ data: store, setData: setData }}>
        <MenuBar bookmarkView="bookmark" />
      </AppProvider>
    );

    const mnu_bookmark_large = getById(container, 'mnu_bookmark_large')
    fireEvent.click(mnu_bookmark_large)
    expect(setData).toHaveBeenCalledTimes(4);
  })

  it('showBookmarks - onDelete ok', () => {
    let ok = false
    const setData = jest.fn((key, data, callback)=>{
      if((key === "current") && data.modalForm ){
        const { container } = render(data.modalForm);
        // showBookmarks - delete icon
        let row_item = getById(container, 'row_delete_1')
        if(row_item){
          fireEvent.click(row_item)
        }
        // InputBox
        let btn_ok = getById(container, 'btn_ok')
        if(btn_ok && !ok){
          ok = true
          fireEvent.click(btn_ok)
        }
      }
      if(callback){
        callback()
      }
    })

    const { container } = render(
      <AppProvider value={{ data: store, setData: setData }}>
        <MenuBar bookmarkView="bookmark" />
      </AppProvider>
    );

    const mnu_bookmark_large = getById(container, 'mnu_bookmark_large')
    fireEvent.click(mnu_bookmark_large)
    expect(setData).toHaveBeenCalledTimes(2);
  })

});
