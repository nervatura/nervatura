import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './sidebar-setting.js';
import { Template, Default, DatabaseGroup, UserGroup, FormItemAll, FormItemNew, FormItemRead, AdminGroup } from  './Setting.stories.js';

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

    const numberdef = sideBar.shadowRoot.querySelector('#cmd_numberdef')
    numberdef.click()
    sinon.assert.callCount(onSideEvent, 1);

    const groupDatabase = sideBar.shadowRoot.querySelector('#group_database_group')
    groupDatabase.click()
    sinon.assert.callCount(onSideEvent, 2);

  })

  it('renders in the AdminGroup state', async () => {
    const element = await fixture(Template({
      ...AdminGroup.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

  it('renders in the DatabaseGroup state', async () => {
    const element = await fixture(Template({
      ...DatabaseGroup.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

  it('renders in the UserGroup state', async () => {
    const element = await fixture(Template({
      ...UserGroup.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

  it('renders in the FormItemAll state', async () => {
    const element = await fixture(Template({
      ...FormItemAll.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

  it('renders in the FormItemNew state', async () => {
    const element = await fixture(Template({
      ...FormItemNew.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

  it('renders in the FormItemRead state', async () => {
    const element = await fixture(Template({
      ...FormItemRead.args
    }));
    const sideBar = element.querySelector('#side_bar');
    expect(sideBar).to.exist;

  })

})