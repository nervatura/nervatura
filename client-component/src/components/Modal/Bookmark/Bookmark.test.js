import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-bookmark.js';
import { Template, Default, BookmarkData, DarkTheme } from  './Bookmark.stories.js';

describe('Bookmark', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const bookmark = element.querySelector('#bookmark');
    expect(bookmark).to.exist;
  })

  it('renders in the BookmarkData state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...BookmarkData.args, onModalEvent
    }));
    const bookmark = element.querySelector('#bookmark');
    expect(bookmark).to.exist;

    const list = bookmark.shadowRoot.querySelector('#bookmark_list')
    const row = list.shadowRoot.querySelector('#row_item_1')
    row.click()
    sinon.assert.callCount(onModalEvent, 1);

    const rowDelete = list.shadowRoot.querySelector('#row_delete_1')
    rowDelete.click()
    sinon.assert.callCount(onModalEvent, 2);

    const closeIcon = bookmark.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 3);

    const btnHistory = bookmark.shadowRoot.querySelector('#btn_history')
    btnHistory.click()
    expect(bookmark.tabView).to.equal("history");

    const btnBookmark = bookmark.shadowRoot.querySelector('#btn_bookmark')
    btnBookmark.click()
    expect(bookmark.tabView).to.equal("bookmark");
  })

  it('renders in the DarkTheme state', async () => {
    const element = await fixture(Template({
      ...DarkTheme.args
    }));
    const bookmark = element.querySelector('#bookmark');
    expect(bookmark).to.exist;
  })

})