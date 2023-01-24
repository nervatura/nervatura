import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './sidebar-search.js';
import { Template, Default, Office } from  './Search.stories.js';

describe('SideBar-Search', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onSideEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onSideEvent
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

    const btnView = sideBar.shadowRoot.querySelector('#btn_view_transitem')
    btnView.click()
    sinon.assert.callCount(onSideEvent, 1);

    const btnBrowser = sideBar.shadowRoot.querySelector('#btn_browser_transitem')
    btnBrowser.click()
    sinon.assert.callCount(onSideEvent, 2);

    const btnGroup = sideBar.shadowRoot.querySelector('#btn_group_customer')
    btnGroup.click()
    sinon.assert.callCount(onSideEvent, 3);

    const btnReport = sideBar.shadowRoot.querySelector('#btn_report')
    btnReport.click()
    sinon.assert.callCount(onSideEvent, 5);

    const btnOffice = sideBar.shadowRoot.querySelector('#btn_office')
    btnOffice.click()
    sinon.assert.callCount(onSideEvent, 6);

  })

  it('renders in the Office state', async () => {
    const onSideEvent = sinon.spy()
    const element = await fixture(Template({
      ...Office.args, onSideEvent
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

    const btnPrintqueue = sideBar.shadowRoot.querySelector('#btn_printqueue')
    btnPrintqueue.click()
    sinon.assert.callCount(onSideEvent, 1);

    const btnRate = sideBar.shadowRoot.querySelector('#btn_rate')
    btnRate.click()
    sinon.assert.callCount(onSideEvent, 2);

    const btnServercmd = sideBar.shadowRoot.querySelector('#btn_servercmd')
    btnServercmd.click()
    sinon.assert.callCount(onSideEvent, 3);

  })

})