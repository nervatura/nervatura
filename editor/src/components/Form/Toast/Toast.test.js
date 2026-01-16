import { fixture, expect } from '@open-wc/testing';

import './form-toast.js';
import { Template, Default, Error, Success } from  './Toast.stories.js';

describe('Toast', () => {

  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const toast = element.querySelector('#test_toast');
    expect(toast).to.exist;

    toast.close()
    toast.show({ style: Default.args.style, vlaue: Default.args.value, timeout: Default.args.timeout })
    toast.close()
    toast.show({ type: "missing" })

    const button = element.querySelector('#toast_show');
    button._onClick({
      stopPropagation: ()=>{}
    })
  })

  it('renders in the Error state', async () => {
    const element = await fixture(Template({
      ...Error.args
    }));
    const toast = element.querySelector('#test_toast');
    expect(toast).to.exist;

    toast.show({ style: Error.args.style, vlaue: Error.args.value, timeout: Error.args.timeout })
    toast.show({})
    toast.close()
  })

  it('renders in the Success state', async () => {
    const element = await fixture(Template({
      ...Success.args
    }));
    const toast = element.querySelector('#test_toast');
    expect(toast).to.exist;

    toast.show({ style: Success.args.style, vlaue: Success.args.value, timeout: Success.args.timeout })
  })

})