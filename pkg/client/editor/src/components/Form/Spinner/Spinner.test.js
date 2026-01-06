import { html } from 'lit';
import { fixture, expect } from '@open-wc/testing';

import './form-spinner.js';

describe('Button', () => {
  it('renders in the Default state', async () => {
    const element = await fixture(html`<form-spinner></form-spinner>`);
    const spinner = element.shadowRoot.querySelector('.loading');
    expect(spinner).to.exist;
  })

});
