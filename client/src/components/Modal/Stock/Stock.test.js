import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-stock.js';
import { Template, Default, DarkShipping } from  './Stock.stories.js';

describe('Total', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const stock = element.querySelector('#stock');
    expect(stock).to.exist;

    const closeIcon = stock.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnOK = stock.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 2);

  })

  it('renders in the DarkShipping state', async () => {
    const element = await fixture(Template({
      ...DarkShipping.args
    }));
    const stock = element.querySelector('#stock');
    expect(stock).to.exist;
  })

})