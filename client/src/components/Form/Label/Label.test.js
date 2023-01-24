import { fixture, expect } from '@open-wc/testing';

import './form-label.js';
import { Template, Default, LeftIcon, RightIcon, Centered, LabelStyle } from  './Label.stories.js';

describe('Label', () => {

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const label = element.querySelector('#test_label');
    expect(label).to.exist;

    await expect(label).shadowDom.to.be.accessible();
  })

  it('renders in the LeftIcon state', async () => {
    const element = await fixture(Template({
      ...LeftIcon.args
    }));
    const label = element.querySelector('#test_label');
    expect(label).to.exist;
  })

  it('renders in the RightIcon state', async () => {
    const element = await fixture(Template({
      ...RightIcon.args
    }));
    const label = element.querySelector('#test_label');
    expect(label).to.exist;
  })

  it('renders in the Centered state', async () => {
    const element = await fixture(Template({
      ...Centered.args
    }));
    const label = element.querySelector('#test_label');
    expect(label).to.exist;
  })

  it('renders in the LabelStyle state', async () => {
    const element = await fixture(Template({
      ...LabelStyle.args
    }));
    const label = element.querySelector('#test_label');
    expect(label).to.exist;
  })

});
