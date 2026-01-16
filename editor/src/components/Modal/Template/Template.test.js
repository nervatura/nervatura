import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-template.js';
import { Template, Default, DarkForm } from  './Template.stories.js';

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
    const template = element.querySelector('#template');
    expect(template).to.exist;

    const closeIcon = template.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnCancel = template.shadowRoot.querySelector('#btn_cancel')
    btnCancel.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnOK = template.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 3);

    const fieldname = template.shadowRoot.querySelector('#name');
    fieldname._onInput({ target: { value: "name" } })
    expect(fieldname.value).to.equal("name");

    template._onTextInput({ target: { value: "col3,col4,col5" } })
    expect(template.columns).to.equal("col3,col4,col5");

    const fieldtype = template.shadowRoot.querySelector('#type');
    fieldtype._onInput({ target: { value: "list" } })
    expect(fieldtype.value).to.equal("list");

  })

  it('renders in the DarkForm state', async () => {
    const element = await fixture(Template({
      ...DarkForm.args
    }));
    const template = element.querySelector('#template');
    expect(template).to.exist;

  })

})