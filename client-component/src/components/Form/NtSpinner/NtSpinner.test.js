import { html } from 'lit';
import { fixture, expect } from '@open-wc/testing';

import './nt-spinner.js';

describe('NtButton', () => {
  it('renders in the Default state', async () => {
    const element = await fixture(html`<nt-spinner></nt-spinner>`);
    const spinner = element.shadowRoot.querySelector('.loading');
    expect(spinner).to.exist;
  })

});
