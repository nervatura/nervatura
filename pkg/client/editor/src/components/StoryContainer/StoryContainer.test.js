import { html } from 'lit';
import { fixture, expect } from '@open-wc/testing';

import './story-container.js';
import { APP_THEME } from '../../config/enums.js'

describe('Icon', () => {

  it('renders in the Default state', async () => {
    const element = await fixture(html`<story-container theme="${APP_THEME.DARK}" ></story-container>`);
    const container = element.shadowRoot.querySelector('div');
    expect(container).to.exist;

    await expect(element).shadowDom.to.be.accessible();
  })

});
