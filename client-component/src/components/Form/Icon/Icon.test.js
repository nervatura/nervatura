import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './form-icon.js';
import { Template, Default, ColorPointer, InvalidIcon } from  './Icon.stories.js';

describe('Icon', () => {
  beforeEach(async () => {
    
  });
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const icon = element.querySelector('#test_icon');
    expect(icon).to.exist;

    await expect(element).shadowDom.to.be.accessible();
  })

  it('renders in the ColorPointer state', async () => {
    const element = await fixture(Template({
      ...ColorPointer.args
    }));
    const icon = element.querySelector('#test_icon');
    expect(icon).to.exist;
  })

  it('renders in the InvalidIcon state', async () => {
    const element = await fixture(Template({
      ...InvalidIcon.args
    }));
    const icon = element.querySelector('#test_icon');
    expect(icon).to.exist;
  })

});
