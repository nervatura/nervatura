import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './sidebar-template.js';
import { Template, Default, Sample, Dirty } from  './Template.stories.js';

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

    const cmdBack = sideBar.shadowRoot.querySelector('#cmd_back')
    cmdBack.click()
    sinon.assert.callCount(onSideEvent, 1);

  })

  it('renders in the Sample state', async () => {
    const element = await fixture(Template({
      ...Sample.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

  it('renders in the Dirty state', async () => {
    const element = await fixture(Template({
      ...Dirty.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

})