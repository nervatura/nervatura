import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-menu.js';
import { Template, Default, Disabled } from  './Menu.stories.js';

describe('Audit', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const menu = element.querySelector('#menu');
    expect(menu).to.exist;

    const closeIcon = menu.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnCancel = menu.shadowRoot.querySelector('#btn_cancel')
    btnCancel.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnOK = menu.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 3);

    const fieldname = menu.shadowRoot.querySelector('#fieldname');
    fieldname._onInput({ target: { value: "value" } })
    expect(fieldname.value).to.equal("value");

    const description = menu.shadowRoot.querySelector('#description');
    description._onInput({ target: { value: "value" } })
    expect(description.value).to.equal("value");

    const fieldtype = menu.shadowRoot.querySelector('#fieldtype');
    fieldtype._onInput({ target: { value: "37" } })
    expect(fieldtype.value).to.equal("37");

    const orderby = menu.shadowRoot.querySelector('#orderby')
    orderby._onInput({ target: { valueAsNumber: 123 } })
    expect(orderby.value).to.equal(123);

  })

  it('renders in the Disabled state', async () => {
    const element = await fixture(Template({
      ...Disabled.args
    }));
    const menu = element.querySelector('#menu');
    expect(menu).to.exist;

  })

})