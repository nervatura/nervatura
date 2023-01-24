import { fixture, expect } from '@open-wc/testing';

import './client-setting.js';
import { Template, Default, Form, Empty } from  './Setting.stories.js';

describe('Setting-Setting', () => {
  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const setting = element.querySelector('#setting');
    expect(setting).to.exist;

    const modalAudit = setting.modalAudit({})
    expect(modalAudit).to.exist;

    const modalMenu = setting.modalMenu({})
    expect(modalMenu).to.exist;

  })

  it('renders in the Form state', async () => {
    const element = await fixture(Template({
      ...Form.args
    }));
    const setting = element.querySelector('#setting');
    expect(setting).to.exist;

  })

  it('renders in the Empty state', async () => {
    const element = await fixture(Template({
      ...Empty.args
    }));
    const setting = element.querySelector('#setting');
    expect(setting).to.exist;

  })

})
