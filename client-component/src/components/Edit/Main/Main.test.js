import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './edit-main.js';
import { Template, Default, Report } from  './Main.stories.js';

describe('Edit-Main', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onEditEvent
    }));
    const main = element.querySelector('#main');
    expect(main).to.exist;

    let inputRow = main.shadowRoot.querySelector('#row_2')
    const customer = inputRow.shadowRoot.querySelector('#field_customer_id').shadowRoot.querySelector('#sel_show_customer_id')
    customer.click()
    sinon.assert.callCount(onEditEvent, 1);

    inputRow = main.shadowRoot.querySelector('#row_3')
    const curr = inputRow.shadowRoot.querySelector('#field_curr').shadowRoot.querySelector('#field_curr')
    curr._onInput({ target: { value: "USD" } })
    sinon.assert.callCount(onEditEvent, 2);

  })

  it('renders in the Report state', async () => {
    const element = await fixture(Template({
      ...Report.args
    }));
    const main = element.querySelector('#main');
    expect(main).to.exist;

  })

})