import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './sidebar-edit.js';
import { Template, Default, NewItem, NewPayment, NewMovement, NewResource, Document, 
  DocumentDeleted, DocumentCancellation, DocumentClosed, DocumentReadonly, DocumentNoOptions } from  './Edit.stories.js';

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

    const stateEdit = sideBar.shadowRoot.querySelector('#state_edit')
    stateEdit.click()
    sinon.assert.callCount(onSideEvent, 1);

    const group = sideBar.shadowRoot.querySelector('#new_transitem_group')
    group.click()
    sinon.assert.callCount(onSideEvent, 2);

  })

  it('renders in the NewItem state', async () => {
    const element = await fixture(Template({
      ...NewItem.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the NewPayment state', async () => {
    const element = await fixture(Template({
      ...NewPayment.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the NewMovement state', async () => {
    const element = await fixture(Template({
      ...NewMovement.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the NewResource state', async () => {
    const element = await fixture(Template({
      ...NewResource.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the Document state', async () => {
    const element = await fixture(Template({
      ...Document.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the DocumentDeleted state', async () => {
    const element = await fixture(Template({
      ...DocumentDeleted.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the DocumentCancellation state', async () => {
    const element = await fixture(Template({
      ...DocumentCancellation.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the DocumentClosed state', async () => {
    const element = await fixture(Template({
      ...DocumentClosed.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the DocumentReadonly state', async () => {
    const element = await fixture(Template({
      ...DocumentReadonly.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

  it('renders in the DocumentNoOptions state', async () => {
    const element = await fixture(Template({
      ...DocumentNoOptions.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;
  })

})