import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-audit.js';
import { Template, Default, Existing, Disabled } from  './Audit.stories.js';

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
    const audit = element.querySelector('#audit');
    expect(audit).to.exist;

    const closeIcon = audit.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnCancel = audit.shadowRoot.querySelector('#btn_cancel')
    btnCancel.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnOK = audit.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 3);

    const nervatype = audit.shadowRoot.querySelector('#nervatype');
    nervatype._onInput({ target: { value: "18" } })
    expect(nervatype.value).to.equal("18");

    nervatype._onInput({ target: { value: "28" } })
    expect(nervatype.value).to.equal("28");

    nervatype._onInput({ target: { value: "10" } })
    expect(nervatype.value).to.equal("10");

    const inputfilter = audit.shadowRoot.querySelector('#inputfilter');
    inputfilter._onInput({ target: { value: "106" } })
    expect(inputfilter.value).to.equal("106");

    const supervisor = audit.shadowRoot.querySelector('#supervisor')
    supervisor.click()
    expect(audit.supervisor).to.equal(0);

  })

  it('renders in the Existing state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Existing.args, onModalEvent
    }));
    const audit = element.querySelector('#audit');
    expect(audit).to.exist;

    const subtype = audit.shadowRoot.querySelector('#subtype');
    subtype._onInput({ target: { value: "58" } })
    expect(subtype.value).to.equal("58");

  })

  it('renders in the Disabled state', async () => {
    const element = await fixture(Template({
      ...Disabled.args
    }));
    const audit = element.querySelector('#audit');
    expect(audit).to.exist;

  })

})