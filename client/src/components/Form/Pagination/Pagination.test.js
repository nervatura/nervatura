import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './form-pagination.js';
import { Template, Default, Items } from  './Pagination.stories.js';

describe('Pagination', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const testPagination = element.querySelector('#test_pagination');
    expect(testPagination).to.exist;

    await expect(testPagination).shadowDom.to.be.accessible();
  })

  it('renders in the Items state', async () => {
    const onEvent = sinon.spy()
    const element = await fixture(Template({
      ...Items.args, onEvent
    }));
    const testPagination = element.querySelector('#test_pagination');
    expect(testPagination).to.exist;

    testPagination._onEvent("gotoPage", 1, true)
    sinon.assert.calledOnce(onEvent);

    const input = testPagination.shadowRoot.querySelector('#pagination_input_goto')
    input._onInput({ target: { value: 50 } })
    sinon.assert.calledTwice(onEvent);

    const btnFirst = testPagination.shadowRoot.querySelector('#pagination_btn_first');
    expect(btnFirst).to.exist;
    btnFirst.click()

    const btnNext = testPagination.shadowRoot.querySelector('#pagination_btn_next');
    expect(btnNext).to.exist;
    btnNext.click()

    const btnPrevious = testPagination.shadowRoot.querySelector('#pagination_btn_previous');
    expect(btnPrevious).to.exist;
    btnPrevious.click()

    const btnLast = testPagination.shadowRoot.querySelector('#pagination_btn_last');
    expect(btnLast).to.exist;
    btnLast.click()

    const pageSize = testPagination.shadowRoot.querySelector('#pagination_page_size');
    pageSize._onInput({ target: { value: 5 } })
    expect(pageSize.value).to.equal(5);
  })

})