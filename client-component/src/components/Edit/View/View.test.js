import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './edit-view.js';
import { Template, Default, List, ReadOnly, Empty, DeleteOnly, ReadOnlyList } from  './View.stories.js';

describe('Edit-Meta', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onEditEvent
    }));
    const view = element.querySelector('#view');
    expect(view).to.exist;

    const table = view.shadowRoot.querySelector('#view_table')
    const btnAdd = table.shadowRoot.querySelector('#btn_add')
    btnAdd.click()
    sinon.assert.callCount(onEditEvent, 1);

    const editIcon = table.shadowRoot.querySelector('#edit_18')
    editIcon.click()
    sinon.assert.callCount(onEditEvent, 2);

    const deleteIcon = table.shadowRoot.querySelector('#delete_18')
    deleteIcon.click()
    sinon.assert.callCount(onEditEvent, 3);

  })

  it('renders in the List state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...List.args, onEditEvent
    }));
    const view = element.querySelector('#view');
    expect(view).to.exist;

    const viewList = view.shadowRoot.querySelector('#view_list')
    const rowEdit = viewList.shadowRoot.querySelector('#row_edit_0')
    rowEdit.click()
    sinon.assert.callCount(onEditEvent, 1);

    const rowDelete = viewList.shadowRoot.querySelector('#row_delete_0')
    rowDelete.click()
    sinon.assert.callCount(onEditEvent, 2);

    const btnAdd = viewList.shadowRoot.querySelector('#btn_add')
    btnAdd.click()
    sinon.assert.callCount(onEditEvent, 3);

  })

  it('renders in the ReadOnly state', async () => {
    const element = await fixture(Template({
      ...ReadOnly.args
    }));
    const view = element.querySelector('#view');
    expect(view).to.exist;
  })

  it('renders in the Empty state', async () => {
    const element = await fixture(Template({
      ...Empty.args
    }));
    const view = element.querySelector('#view');
    expect(view).to.exist;
  })

  it('renders in the DeleteOnly state', async () => {
    const element = await fixture(Template({
      ...DeleteOnly.args
    }));
    const view = element.querySelector('#view');
    expect(view).to.exist;
  })

  it('renders in the ReadOnlyList state', async () => {
    const element = await fixture(Template({
      ...ReadOnlyList.args
    }));
    const view = element.querySelector('#view');
    expect(view).to.exist;
  })

})