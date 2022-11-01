import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './form-table.js';
import { Template, Default, BottomPagination, TopPagination, Filtered } from  './Table.stories.js';

describe('Table', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onRowSelected = sinon.spy()
    const onEditCell = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onRowSelected, onEditCell
    }));
    const testTable = element.querySelector('#test_table');
    expect(testTable).to.exist;

    const linkCell = testTable.shadowRoot.querySelector('#link_2')
    linkCell.click()
    sinon.assert.calledOnce(onEditCell);

    const tableRow = testTable.shadowRoot.querySelector('#row_1')
    tableRow.click()
    sinon.assert.calledOnce(onRowSelected);

    // await expect(testTable).shadowDom.to.be.accessible();
  })

  it('renders in the BottomPagination state', async () => {
    const element = await fixture(Template({
      ...BottomPagination.args
    }));
    const testTable = element.querySelector('#test_table');
    expect(testTable).to.exist;

  })

  it('renders in the TopPagination state', async () => {
    const element = await fixture(Template({
      ...TopPagination.args
    }));
    const testTable = element.querySelector('#test_table');
    expect(testTable).to.exist;

  })

  it('renders in the Filtered state', async () => {
    const onCurrentPage = sinon.spy()
    const onAddItem = sinon.spy()
    const element = await fixture(Template({
      ...Filtered.args, onCurrentPage, onAddItem
    }));
    const testTable = element.querySelector('#test_table');
    expect(testTable).to.exist;

    const pagination = testTable.shadowRoot.querySelector('#test_table_top_pagination')
    const btnLast = pagination.shadowRoot.querySelector('#pagination_btn_last')
    btnLast._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onCurrentPage);

    const pageSize = pagination.shadowRoot.querySelector('#pagination_page_size');
    pageSize._onInput({ target: { value: 10 } })
    expect(pageSize.value).to.equal(10);

    const filter = testTable.shadowRoot.querySelector('#filter');
    filter._onInput({ target: { value: "value" } })
    expect(filter.value).to.equal("value");

    const btnAdd = testTable.shadowRoot.querySelector('#btn_add')
    btnAdd._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onAddItem);

    const sortCol = testTable.shadowRoot.querySelector('.sort-none')
    sortCol.click()
    sortCol.click()

    testTable.sortAsc = true
    sortCol.click()
    
  })

})