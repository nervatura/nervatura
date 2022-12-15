import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './edit-item.js';
import { Template, Default, NewItem } from  './Item.stories.js';

describe('Edit-Item', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onEditEvent
    }));
    const item = element.querySelector('#item');
    expect(item).to.exist;

    let inputRow = item.shadowRoot.querySelector('#row_0')
    const product = inputRow.shadowRoot.querySelector('#field_product_id').shadowRoot.querySelector('#sel_show_product_id')
    product.click()
    sinon.assert.callCount(onEditEvent, 1);

    inputRow = item.shadowRoot.querySelector('#row_3')
    const curr = inputRow.shadowRoot.querySelector('#field_qty').shadowRoot.querySelector('#field_qty')
    curr._onInput({ target: { value: "2" } })
    sinon.assert.callCount(onEditEvent, 2);

  })

  it('renders in the NewItem state', async () => {
    const element = await fixture(Template({
      ...NewItem.args
    }));
    const item = element.querySelector('#item');
    expect(item).to.exist;

  })

})