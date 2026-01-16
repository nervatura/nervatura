import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-inputbox.js';
import { Template, Default, InputValue, DefaultCancel } from  './InputBox.stories.js';

describe('InputBox', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const inputbox = element.querySelector('#inputbox');
    expect(inputbox).to.exist;
  })

  it('renders in the InputValue state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...InputValue.args, onModalEvent
    }));
    const inputbox = element.querySelector('#inputbox');
    expect(inputbox).to.exist;

    const btnOK = inputbox.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 1);
  })

  it('renders in the DefaultCancel state', async () => {
    const element = await fixture(Template({
      ...DefaultCancel.args
    }));
    const inputbox = element.querySelector('#inputbox');
    expect(inputbox).to.exist;
  })
})