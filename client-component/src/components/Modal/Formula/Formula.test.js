import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-formula.js';
import { Template, Default, Disabled } from  './Formula.stories.js';

describe('Formula', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const formula = element.querySelector('#formula');
    expect(formula).to.exist;

    const closeIcon = formula.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnCancel = formula.shadowRoot.querySelector('#btn_cancel')
    btnCancel.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnOK = formula.shadowRoot.querySelector('#btn_ok')
    btnOK.click()
    sinon.assert.callCount(onModalEvent, 3);

    const selFormula = formula.shadowRoot.querySelector('#sel_formula');
    selFormula._onInput({ target: { value: "" } })
    expect(selFormula.value).to.equal("");
  })

  it('renders in the Disabled state', async () => {
    const element = await fixture(Template({
      ...Disabled.args
    }));
    const formula = element.querySelector('#formula');
    expect(formula).to.exist;
  })

})