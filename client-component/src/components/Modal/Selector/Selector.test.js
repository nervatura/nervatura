import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-selector.js';
import { Template, Default, QuickView } from  './Selector.stories.js';

describe('MenuBar', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const selector = element.querySelector('#selector');
    expect(selector).to.exist;

    const closeIcon = selector.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnSearch = selector.shadowRoot.querySelector('#selector_btn_search')
    btnSearch.click()
    sinon.assert.callCount(onModalEvent, 2);

    const filter = selector.shadowRoot.querySelector('#selector_filter');
    filter._onInput({ target: { value: "value" } })
    expect(filter.value).to.equal("value");
    filter._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 13 
    })
    sinon.assert.callCount(onModalEvent, 3);

    const table = selector.shadowRoot.querySelector('#selector_result')
    const tableRow = table.shadowRoot.querySelector('#row_customer-2')
    tableRow.click()
    sinon.assert.callCount(onModalEvent, 4);
  })

  it('renders in the QuickView state', async () => {
    const element = await fixture(Template({
      ...QuickView.args
    }));
    const selector = element.querySelector('#selector');
    expect(selector).to.exist;
  })

})