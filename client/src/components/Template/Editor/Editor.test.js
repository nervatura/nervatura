import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './template-editor.js';
import { Template, Default, Details, Row, Cell, Data, StringData, ListData, ListDataItem, TableData, TableDataItem, Meta, VGap } from  './Editor.stories.js';

describe('SideBar-Search', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    // navButton
    const btnNext = editor.shadowRoot.querySelector('#btn_next')
    btnNext.click()
    sinon.assert.callCount(onTemplateEvent, 2);

    // mapButton
    const btn_tmpReport = editor.shadowRoot.querySelector('#btn_tmp_report')
    btn_tmpReport.click()
    sinon.assert.callCount(onTemplateEvent, 3);

    // tabButton
    const btnData = editor.shadowRoot.querySelector('#btn_data')
    btnData.click()
    sinon.assert.callCount(onTemplateEvent, 4);

    const inputRow = editor.shadowRoot.querySelector('#row_0')
    const title = inputRow.shadowRoot.querySelector('#field_title').shadowRoot.querySelector('#field_title')
    title._onInput({ target: { value: "test value" } })
    sinon.assert.callCount(onTemplateEvent, 5);

  })

  it('renders in the Details state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...Details.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    const addItem = editor.shadowRoot.querySelector('#btn_add_item')
    addItem.click()
    sinon.assert.callCount(onTemplateEvent, 2);

  })

  it('renders in the Row state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...Row.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    const addItem = editor.shadowRoot.querySelector('#btn_add_item')
    addItem.click()
    sinon.assert.callCount(onTemplateEvent, 2);

    const selItem = editor.shadowRoot.querySelector('#sel_add_item');
    selItem._onInput({ target: { value: "image" } })
    expect(selItem.value).to.equal("image");

  })

  it('renders in the Cell state', async () => {
    const element = await fixture(Template({
      ...Cell.args
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

  })

  it('renders in the Data state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...Data.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    const dataList = editor.shadowRoot.querySelector('#data_list_items')
    const addItem = dataList.shadowRoot.querySelector('#btn_add')
    addItem.click()
    sinon.assert.callCount(onTemplateEvent, 1);

    const rowEdit = dataList.shadowRoot.querySelector('#row_edit_0')
    rowEdit.click()
    sinon.assert.callCount(onTemplateEvent, 2);

    const rowDelete = dataList.shadowRoot.querySelector('#row_delete_0')
    rowDelete.click()
    sinon.assert.callCount(onTemplateEvent, 3);

  })

  it('renders in the StringData state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...StringData.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

  })

  it('renders in the ListData state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...ListData.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    const dataList = editor.shadowRoot.querySelector('#data_list_items')
    const addItem = dataList.shadowRoot.querySelector('#btn_add')
    addItem.click()
    sinon.assert.callCount(onTemplateEvent, 1);

    const rowEdit = dataList.shadowRoot.querySelector('#row_edit_0')
    rowEdit.click()
    sinon.assert.callCount(onTemplateEvent, 2);

    const rowDelete = dataList.shadowRoot.querySelector('#row_delete_0')
    rowDelete.click()
    sinon.assert.callCount(onTemplateEvent, 3);

  })

  it('renders in the ListDataItem state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...ListDataItem.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

  })

  it('renders in the TableData state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...TableData.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    const dataTable = editor.shadowRoot.querySelector('#data_table_items')
    const addItem = dataTable.shadowRoot.querySelector('#btn_add')
    addItem.click()
    sinon.assert.callCount(onTemplateEvent, 1);

    const tableRow = dataTable.shadowRoot.querySelector('#row_0')
    tableRow.click()
    sinon.assert.callCount(onTemplateEvent, 2);

    const tableRowDelete = dataTable.shadowRoot.querySelector('#delete_0')
    tableRowDelete.click()
    sinon.assert.callCount(onTemplateEvent, 3);

  })

  it('renders in the TableDataItem state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...TableDataItem.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

  })

  it('renders in the Meta state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...Meta.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

  })

  it('renders in the VGap state', async () => {
    const onTemplateEvent = sinon.spy()
    const element = await fixture(Template({
      ...VGap.args, onTemplateEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

  })

})