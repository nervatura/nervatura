import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-shipping.js';
import { Template, Default, DarkShipping } from  './Shipping.stories.js';

describe('Shipping', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const shipping = element.querySelector('#shipping');
    expect(shipping).to.exist;

    const closeIcon = shipping.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnCancel = shipping.shadowRoot.querySelector('#btn_cancel')
    btnCancel.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnOK = shipping.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 3);

    const batchNo = shipping.shadowRoot.querySelector('#batch_no');
    batchNo._onInput({ target: { value: "abc123" } })
    expect(batchNo.value).to.equal("abc123");

    const qty = shipping.shadowRoot.querySelector('#qty');
    qty._onInput({ target: { valueAsNumber: 123 } })
    expect(qty.value).to.equal(123);
  })

  it('renders in the DarkShipping state', async () => {
    const element = await fixture(Template({
      ...DarkShipping.args
    }));
    const shipping = element.querySelector('#shipping');
    expect(shipping).to.exist;

  })

})