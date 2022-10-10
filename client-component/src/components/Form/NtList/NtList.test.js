import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './nt-list.js';
import { Template, Default, BottomPagination, Filtered, TopPagination } from  './NtList.stories.js';

describe('NtList', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const testList = element.querySelector('#test_list');
    expect(testList).to.exist;

    // await expect(testList).shadowDom.to.be.accessible();
  })

  it('renders in the BottomPagination state', async () => {
    const element = await fixture(Template({
      ...BottomPagination.args
    }));
    const testList = element.querySelector('#test_list');
    expect(testList).to.exist;

  })

  it('renders in the TopPagination state', async () => {
    const onCurrentPage = sinon.spy()
    const element = await fixture(Template({
      ...TopPagination.args, onCurrentPage
    }));
    const testList = element.querySelector('#test_list');
    expect(testList).to.exist;

    const pagination = testList.shadowRoot.querySelector('#test_list_top_pagination')
    const btnLast = pagination.shadowRoot.querySelector('#pagination_btn_first')
    btnLast._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onCurrentPage);

    const pageSize = pagination.shadowRoot.querySelector('#pagination_page_size');
    pageSize._onInput({ target: { value: 10 } })
    expect(pageSize.value).to.equal(10);


  })

  it('renders in the Filtered state', async () => {
    const onEdit = sinon.spy()
    const onAddItem = sinon.spy()
    const onDelete = sinon.spy()
    const element = await fixture(Template({
      ...Filtered.args, onEdit, onAddItem, onDelete
    }));
    const testList = element.querySelector('#test_list');
    expect(testList).to.exist;

    const listRowIcon = testList.shadowRoot.querySelector('#row_edit_1')
    listRowIcon.click()
    sinon.assert.calledOnce(onEdit);

    const listRow = testList.shadowRoot.querySelector('#row_item_1')
    listRow.click()
    sinon.assert.calledTwice(onEdit);

    const listDelete = testList.shadowRoot.querySelector('#row_delete_1')
    listDelete.click()
    sinon.assert.calledOnce(onDelete);

    const btnAdd = testList.shadowRoot.querySelector('#btn_add')
    btnAdd._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onAddItem);

    const filter = testList.shadowRoot.querySelector('#filter');
    filter._onInput({ target: { value: "value" } })
    expect(filter.value).to.equal("value");

  })

})