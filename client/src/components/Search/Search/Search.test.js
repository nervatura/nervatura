import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './client-search.js';
import { Template, Default, Browser } from  './Search.stories.js';

describe('SideBar-Search', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const search = element.querySelector('#search');
    expect(search).to.exist;

    const selector = search.shadowRoot.querySelector('#selector_customer')
    const table = selector.shadowRoot.querySelector('#selector_result')
    const tableRow = table.shadowRoot.querySelector('#row_customer-2')
    tableRow.click()
    sinon.assert.callCount(onModalEvent, 1);

    const serverFrm = search.modalServer({ cmd: {}, fields: [], values: {} })
    expect(serverFrm).to.exist;
    const totalFrm = search.modalTotal({})
    expect(totalFrm).to.exist;
  })

  it('renders in the Browser state', async () => {
    const element = await fixture(Template({
      ...Browser.args
    }));
    const search = element.querySelector('#search');
    expect(search).to.exist;
  })

})