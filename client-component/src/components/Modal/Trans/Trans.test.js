import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-trans.js';
import { Template, Default, Order, Worksheet } from  './Trans.stories.js';

describe('Bookmark', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const trans = element.querySelector('#trans');
    expect(trans).to.exist;

    const closeIcon = trans.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnCancel = trans.shadowRoot.querySelector('#btn_cancel')
    btnCancel.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnOK = trans.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 3);

    const transtype = trans.shadowRoot.querySelector('#transtype');
    transtype._onInput({ target: { value: "worksheet" } })
    expect(transtype.value).to.equal("worksheet");

    const direction = trans.shadowRoot.querySelector('#direction');
    direction._onInput({ target: { value: "in" } })
    expect(direction.value).to.equal("in");
  })

  it('renders in the Order state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Order.args, onModalEvent
    }));
    const trans = element.querySelector('#trans');
    expect(trans).to.exist;

    const transtype = trans.shadowRoot.querySelector('#transtype');
    transtype._onInput({ target: { value: "receipt" } })
    expect(transtype.value).to.equal("receipt");

    const refno = trans.shadowRoot.querySelector('#refno')
    refno.click()

    const netto = trans.shadowRoot.querySelector('#netto')
    netto.click()
    
    const from = trans.shadowRoot.querySelector('#from')
    from.click()
    
    const btnOK = trans.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 1);
  })

  it('renders in the Worksheet state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Worksheet.args, onModalEvent
    }));
    const trans = element.querySelector('#trans');
    expect(trans).to.exist;

    const transtype = trans.shadowRoot.querySelector('#transtype');
    transtype._onInput({ target: { value: "receipt" } })
    expect(transtype.value).to.equal("receipt");

    const btnOK = trans.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 1);
  })

})