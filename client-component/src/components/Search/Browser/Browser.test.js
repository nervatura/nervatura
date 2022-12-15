import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './search-browser.js';
import { Template, Default, HideHeader, Columns, Filters, FormActions } from  './Browser.stories.js';

describe('Browser', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onBrowserEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onBrowserEvent
    }));
    const browser = element.querySelector('#browser');
    expect(browser).to.exist;

    const btnSearch = browser.shadowRoot.querySelector('#btn_search')
    btnSearch.click()
    sinon.assert.callCount(onBrowserEvent, 1);

    const btnBookmark = browser.shadowRoot.querySelector('#btn_bookmark')
    btnBookmark.click()
    sinon.assert.callCount(onBrowserEvent, 2);

    const btnExport = browser.shadowRoot.querySelector('#btn_export')
    btnExport.click()
    sinon.assert.callCount(onBrowserEvent, 3);

    const btnHelp = browser.shadowRoot.querySelector('#btn_help')
    btnHelp.click()
    sinon.assert.callCount(onBrowserEvent, 4);

    const btnViews = browser.shadowRoot.querySelector('#btn_views')
    btnViews.click()
    sinon.assert.callCount(onBrowserEvent, 4);

    const btnColumns = browser.shadowRoot.querySelector('#btn_columns')
    btnColumns.click()
    sinon.assert.callCount(onBrowserEvent, 5);

    const btnFilter = browser.shadowRoot.querySelector('#btn_filter')
    btnFilter.click()
    sinon.assert.callCount(onBrowserEvent, 6);

    const btnTotal = browser.shadowRoot.querySelector('#btn_total')
    btnTotal.click()
    sinon.assert.callCount(onBrowserEvent, 7);

    const table = browser.shadowRoot.querySelector('#browser_result')
    const rowIcon = table.shadowRoot.querySelector('#edit_customer-2')
    rowIcon.click()
    sinon.assert.callCount(onBrowserEvent, 8);
    const rowCell = table.shadowRoot.querySelector('#link_2')
    rowCell.click()
    sinon.assert.callCount(onBrowserEvent, 9);
  })

  it('renders in the HideHeader state', async () => {
    const onBrowserEvent = sinon.spy()
    const element = await fixture(Template({
      ...HideHeader.args, onBrowserEvent
    }));
    const browser = element.querySelector('#browser');
    expect(browser).to.exist;

    const btnHeader = browser.shadowRoot.querySelector('#btn_header')
    btnHeader.click()
    sinon.assert.callCount(onBrowserEvent, 1);

  })

  it('renders in the Columns state', async () => {
    const onBrowserEvent = sinon.spy()
    const element = await fixture(Template({
      ...Columns.args, onBrowserEvent
    }));
    const browser = element.querySelector('#browser');
    expect(browser).to.exist;

    const viewItem = browser.shadowRoot.querySelector('#view_CustomerAddressView')
    viewItem.click()
    sinon.assert.callCount(onBrowserEvent, 1);

    const colSurname = browser.shadowRoot.querySelector('#col_surname')
    colSurname.click()
    sinon.assert.callCount(onBrowserEvent, 2);

    const colStatus = browser.shadowRoot.querySelector('#col_status')
    colStatus.click()
    sinon.assert.callCount(onBrowserEvent, 3);

  })

  it('renders in the Filters state', async () => {
    const onBrowserEvent = sinon.spy()
    const element = await fixture(Template({
      ...Filters.args, onBrowserEvent
    }));
    const browser = element.querySelector('#browser');
    expect(browser).to.exist;

    const btnDelete = browser.shadowRoot.querySelector('#btn_delete_filter_0')
    btnDelete.click()
    sinon.assert.callCount(onBrowserEvent, 1);

    const filterName = browser.shadowRoot.querySelector('#filter_name_0')
    filterName._onInput({ target: { value: "sample_customer_date" } })
    sinon.assert.callCount(onBrowserEvent, 2);

    const filterType = browser.shadowRoot.querySelector('#filter_type_0')
    filterType._onInput({ target: { value: "==N" } })
    sinon.assert.callCount(onBrowserEvent, 3);

    const filterValueNumber = browser.shadowRoot.querySelector('#filter_value_0')
    filterValueNumber._onInput({ target: { value: "111", valueAsNumber: 111 } })
    sinon.assert.callCount(onBrowserEvent, 4);

    const filterValueDate = browser.shadowRoot.querySelector('#filter_value_1')
    filterValueDate._onInput({ target: { value: "2022-01-02" } })
    sinon.assert.callCount(onBrowserEvent, 5);

    const filterValueBool = browser.shadowRoot.querySelector('#filter_value_2')
    filterValueBool._onInput({ target: { value: "0" } })
    sinon.assert.callCount(onBrowserEvent, 6);

    const filterValueString = browser.shadowRoot.querySelector('#filter_value_3')
    filterValueString._onInput({ target: { value: "red" } })
    sinon.assert.callCount(onBrowserEvent, 7);
  })

  it('renders in the FormActions state', async () => {
    const onBrowserEvent = sinon.spy()
    const element = await fixture(Template({
      ...FormActions.args, onBrowserEvent
    }));
    const browser = element.querySelector('#browser');
    expect(browser).to.exist;

    const btnActionsNew = browser.shadowRoot.querySelector('#btn_actions_new')
    btnActionsNew.click()
    sinon.assert.callCount(onBrowserEvent, 1);

  })

})