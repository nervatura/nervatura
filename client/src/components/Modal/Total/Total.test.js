import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-total.js';
import { Template, Default, DarkTotal } from  './Total.stories.js';

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
    const total = element.querySelector('#total');
    expect(total).to.exist;

    const closeIcon = total.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnOK = total.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 2);

  })

  it('renders in the DarkTotal state', async () => {
    const element = await fixture(Template({
      ...DarkTotal.args
    }));
    const total = element.querySelector('#total');
    expect(total).to.exist;
  })

})