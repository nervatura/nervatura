import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './client-menubar.js';
import { Template, Default, ScrollTop } from  './MenuBar.stories.js';

describe('MenuBar', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onMenuEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onMenuEvent
    }));
    const menuBar = element.querySelector('#menu_bar');
    expect(menuBar).to.exist;

    const search = menuBar.shadowRoot.querySelector('#mnu_search_large')
    search.click()
    sinon.assert.callCount(onMenuEvent, 1);

    const logout = menuBar.shadowRoot.querySelector('#mnu_logout_large')
    logout.click()
    sinon.assert.callCount(onMenuEvent, 2);

    const bookmark = menuBar.shadowRoot.querySelector('#mnu_bookmark_large')
    bookmark.click()
    sinon.assert.callCount(onMenuEvent, 3);

    const modalBookmark = menuBar.modalBookmark()
    expect(modalBookmark).to.exist;
  })

  it('renders in the ScrollTop state', async () => {
    const onMenuEvent = sinon.spy()
    const element = await fixture(Template({
      ...ScrollTop.args, onMenuEvent
    }));
    const menuBar = element.querySelector('#menu_bar');
    expect(menuBar).to.exist;

    const scroll = menuBar.shadowRoot.querySelector('#mnu_scroll')
    scroll.click()
    sinon.assert.callCount(onMenuEvent, 1);

    const sidebar = menuBar.shadowRoot.querySelector('#mnu_sidebar')
    sidebar.click()
    sinon.assert.callCount(onMenuEvent, 2);
  })

})