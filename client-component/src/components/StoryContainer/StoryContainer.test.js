import { html } from 'lit';
import { fixture, expect } from '@open-wc/testing';

import './story-container.js';

describe('NtIcon', () => {

  it('renders in the Default state', async () => {
    const element = await fixture(html`<story-container theme="dark" ></story-container>`);
    const container = element.shadowRoot.querySelector('div');
    expect(container).to.exist;

    await expect(element).shadowDom.to.be.accessible();
  })

  it('renders in the Invalid state', async () => {
    const element = await fixture(html`<story-container theme="invalid" ></story-container>`);
    const container = element.shadowRoot.querySelector('div');
    expect(container).to.exist;

    await expect(element).shadowDom.to.be.accessible();
  })

});
