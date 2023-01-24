import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-server.js';
import { Template, Default, DarkServer } from  './Server.stories.js';

describe('Server', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const server = element.querySelector('#server');
    expect(server).to.exist;

    const closeIcon = server.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnOK = server.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnCancel = server.shadowRoot.querySelector('#btn_cancel')
    btnCancel.click()
    sinon.assert.callCount(onModalEvent, 3);

    const inputRow = server.shadowRoot.querySelector('#row_1')
    const fieldStep = inputRow.shadowRoot.querySelector('#field_step').shadowRoot.querySelector('#field_step')
    fieldStep.click()
    expect(server.values.step).to.equal(1);

  })

  it('renders in the DarkServer state', async () => {
    const element = await fixture(Template({
      ...DarkServer.args
    }));
    const server = element.querySelector('#server');
    expect(server).to.exist;
  })

})