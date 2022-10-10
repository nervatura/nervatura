import { fixture, expect } from '@open-wc/testing';

import './nt-toast.js';
import { Template, Default, Error, Success, Invalid } from  './NtToast.stories.js';

describe('NtToast', () => {

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const toast = element.querySelector('#test_toast');
    expect(toast).to.exist;

    const button = element.querySelector('#toast_show');
    button._onClick({
      stopPropagation: ()=>{}
    })

    toast.toggleToastClass()
    toast.show()
    toast.close()
  })

  it('renders in the Error state', async () => {
    const element = await fixture(Template({
      ...Error.args
    }));
    const toast = element.querySelector('#test_toast');
    expect(toast).to.exist;
  })

  it('renders in the Success state', async () => {
    const element = await fixture(Template({
      ...Success.args
    }));
    const toast = element.querySelector('#test_toast');
    expect(toast).to.exist;
  })

  it('renders in the Invalid state', async () => {
    const element = await fixture(Template({
      ...Invalid.args
    }));
    const toast = element.querySelector('#test_toast');
    expect(toast).to.exist;
  })

})