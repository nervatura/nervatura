import React from 'react';
import { render, fireEvent, queryByAttribute } from '@testing-library/react'
import update from 'immutability-helper';

import Editor from './index';
import { store as app_store  } from 'config/app'
import { AppProvider } from 'containers/App/context'
import { Default as EditorData } from 'components/Editor/Editor/Editor.stories'
import { NewField as EditorMeta } from 'components/Editor/Meta/Meta.stories'
import { Default as SideBarData } from 'components/SideBar/Edit/Edit.stories'

import { appActions } from 'containers/App/actions'
import { editorActions } from 'containers/Editor/actions'

jest.mock("containers/App/actions");
jest.mock("containers/Editor/actions");

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
  },
  edit: {
    ...EditorData.args
  }
}})

describe('<Editor />', () => {
  beforeAll(() => {
    Object.defineProperty(global.document, 'execCommand', { value: jest.fn() });
  });

  beforeEach(() => {
    appActions.mockReturnValue({
      getText: EditorData.args.getText,
      onSelector: jest.fn()
    })
    editorActions.mockReturnValue({
      editItem: jest.fn(),
      checkEditor: jest.fn(),
      prevTransNumber: jest.fn()
    })
  });
  
  afterAll(() => {
    jest.clearAllMocks();
  });

  it('renders in the Main state', () => {
    const setData = jest.fn((key, data, callback)=>{ if(callback){callback()} })
    let it_store = update(store, {$merge: {
      current: {
        content: { 
          ntype: "trans", ttype: "invoice", id: 5, item: {} 
        }
      }
    }})

    const { container, rerender } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Editor id="edit" />
      </AppProvider>
    );
    expect(getById(container, 'edit')).toBeDefined();

    //onSelector
    const btn_selector = getById(container, 'sel_show_customer_id')
    fireEvent.click(btn_selector)

    //prevTransNumber
    //const cmd_arrow_left = getById(container, 'cmd_arrow_left')
    //fireEvent.click(cmd_arrow_left)

    //editItem
    const field_notes = getById(container, 'field_notes')
    fireEvent.change(field_notes, {target: {value: "test data"}})

    //changeCurrentData
    const btn_fieldvalue = getById(container, 'btn_fieldvalue')
    fireEvent.click(btn_fieldvalue)
    expect(setData).toHaveBeenCalledTimes(2);

    //changeData
    const state_new = getById(container, 'state_new')
    fireEvent.click(state_new)
    expect(setData).toHaveBeenCalledTimes(3);

    rerender(
      <AppProvider value={{ data: store, setData: setData }}>
        <Editor />
      </AppProvider>
    )

    it_store = update(store, {edit: {current: {$merge: {
      item: null
    }}}})
    render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Editor id="edit" />
      </AppProvider>
    )
  });

  it('renders in the Metadata state', () => {
    const setData = jest.fn()
    const it_store = update(store, {$merge: {
      edit: {
        ...EditorMeta.args,
        side_view: "edit",
        panel: {
          help: true, arrow: true
        }
      }
    }})

    const { container } = render(
      <AppProvider value={{ data: it_store, setData: setData }}>
        <Editor id="edit" />
      </AppProvider>
    );
    expect(getById(container, 'edit')).toBeDefined();

    const cmd_help = getById(container, 'cmd_help')
    fireEvent.click(cmd_help)

    const cmd_arrow_left = getById(container, 'cmd_arrow_left')
    fireEvent.click(cmd_arrow_left)

  });

  it('renders in the Empty state', () => {
    const setData = jest.fn()
    render(
      <AppProvider value={{ data: store, setData: setData }}>
        <Editor id="edit" />
      </AppProvider>
    );

  });

})