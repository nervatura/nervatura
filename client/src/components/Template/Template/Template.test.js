import { fixture, expect } from '@open-wc/testing';

import './client-template.js';
import { Template, Default } from  './Template.stories.js';

describe('Template-Template', () => {
  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const template = element.querySelector('#template');
    expect(template).to.exist;

    const modalTemplate = template.modalTemplate({})
    expect(modalTemplate).to.exist;

  })

})
