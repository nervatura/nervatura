import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './edit-meta.js';
import { Template, Default, NewField, Customer, ReadOnly } from  './Meta.stories.js';

describe('Edit-Meta', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onEditEvent
    }));
    const meta = element.querySelector('#meta');
    expect(meta).to.exist;

    let inputRow = meta.shadowRoot.querySelector('#row_0')
    const transitem = inputRow.shadowRoot.querySelector('#field_trans_transitem_link').shadowRoot.querySelector('#sel_show_fieldvalue_value')
    transitem.click()
    sinon.assert.callCount(onEditEvent, 1);

    inputRow = meta.shadowRoot.querySelector('#row_1')
    const valuelist = inputRow.shadowRoot.querySelector('#field_4e451b7f-72d1-b19c-7cbe-2c80495b5a8e').shadowRoot.querySelector('#field_4e451b7f-72d1-b19c-7cbe-2c80495b5a8e')
    valuelist._onInput({ target: { value: "red" } })
    sinon.assert.callCount(onEditEvent, 2);

  })

  it('renders in the NewField state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...NewField.args, onEditEvent
    }));
    const meta = element.querySelector('#meta');
    expect(meta).to.exist;

    const btnNew = meta.shadowRoot.querySelector('#btn_new')
    btnNew.click()
    sinon.assert.callCount(onEditEvent, 1);

    const pagination = meta.shadowRoot.querySelector('#meta_top_pagination').shadowRoot.querySelector('#pagination_btn_first')
    pagination.click()

    const pageSize = meta.shadowRoot.querySelector('#meta_top_pagination').shadowRoot.querySelector('#pagination_page_size');
    pageSize._onInput({ target: { value: 10 } })
    expect(pageSize.value).to.equal(10);

  })

  it('renders in the Customer state', async () => {
    const element = await fixture(Template({
      ...Customer.args
    }));
    const meta = element.querySelector('#meta');
    expect(meta).to.exist;

  })

  it('renders in the ReadOnly state', async () => {
    let element = await fixture(Template({
      ...ReadOnly.args
    }));
    let meta = element.querySelector('#meta');
    expect(meta).to.exist;

    element = await fixture(Template({
      ...ReadOnly.args, pageSize: 2
    }));
    meta = element.querySelector('#meta');
    expect(meta).to.exist;

  })

})