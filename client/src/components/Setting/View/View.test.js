import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './setting-view.js';
import { Template, Default, Table, ReadOnlyList, ReadOnlyTable } from  './View.stories.js';

describe('Setting-Form', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onSettingEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onSettingEvent
    }));
    const view = element.querySelector('#setting_view');
    expect(view).to.exist;

    const viewList = view.shadowRoot.querySelector('#view_list')
    const btnAdd = viewList.shadowRoot.querySelector('#btn_add')
    btnAdd.click()
    sinon.assert.callCount(onSettingEvent, 1);

    const rowEdit = viewList.shadowRoot.querySelector('#row_edit_0')
    rowEdit.click()
    sinon.assert.callCount(onSettingEvent, 2);

    const rowItem = viewList.shadowRoot.querySelector('#row_item_0')
    rowItem.click()
    sinon.assert.callCount(onSettingEvent, 3);

    const rowDelete= viewList.shadowRoot.querySelector('#row_delete_0')
    rowDelete.click()
    sinon.assert.callCount(onSettingEvent, 4);

    const pagination = viewList.shadowRoot.querySelector('#view_list_top_pagination')
    const btnNext = pagination.shadowRoot.querySelector('#pagination_btn_next')
    btnNext.click()
    sinon.assert.callCount(onSettingEvent, 5);

  })

  it('renders in the Table state', async () => {
    const onSettingEvent = sinon.spy()
    const element = await fixture(Template({
      ...Table.args, onSettingEvent
    }));
    const view = element.querySelector('#setting_view');
    expect(view).to.exist;

    const viewTable = view.shadowRoot.querySelector('#view_table')
    const btnAdd = viewTable.shadowRoot.querySelector('#btn_add')
    btnAdd.click()
    sinon.assert.callCount(onSettingEvent, 1);

    const editIcon = viewTable.shadowRoot.querySelector('#edit_1')
    editIcon.click()
    sinon.assert.callCount(onSettingEvent, 2);

    const editRow = viewTable.shadowRoot.querySelector('#row_1')
    editRow.click()
    sinon.assert.callCount(onSettingEvent, 3);

  })

  it('renders in the ReadOnlyList state', async () => {
    const element = await fixture(Template({
      ...ReadOnlyList.args
    }));
    const view = element.querySelector('#setting_view');
    expect(view).to.exist;

  })

  it('renders in the ReadOnlyTable state', async () => {
    const onSettingEvent = sinon.spy()
    const element = await fixture(Template({
      ...ReadOnlyTable.args, onSettingEvent
    }));
    const view = element.querySelector('#setting_view');
    expect(view).to.exist;

    const pagination = view.shadowRoot.querySelector('#view_table').shadowRoot.querySelector('#view_table_top_pagination')
    const btnNext = pagination.shadowRoot.querySelector('#pagination_btn_next')
    btnNext.click()
    sinon.assert.callCount(onSettingEvent, 1);

  })

})