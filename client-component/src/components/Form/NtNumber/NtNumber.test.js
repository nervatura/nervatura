import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './nt-number.js';
import { Template, Default, Integer } from  './NtNumber.stories.js';

describe('NtNumber', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onChange = sinon.spy()
    const onEnter = sinon.spy()
    const onBlur = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onChange, onEnter, onBlur
    }));
    const testNumber = element.querySelector('#test_number');
    expect(testNumber).to.exist;
    expect(testNumber.value).to.equal(0);

    testNumber._onInput({ target: { value: "123", valueAsNumber: 123 } })
    sinon.assert.calledOnce(onChange);
    expect(testNumber.value).to.equal(123);

    testNumber._onBlur()
    expect(testNumber.value).to.equal(123);

    testNumber._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.calledOnce(onEnter);

    testNumber._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 13 
    })
    sinon.assert.calledTwice(onEnter);

    testNumber._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 20 
    })
    sinon.assert.calledTwice(onEnter);

    await expect(testNumber).shadowDom.to.be.accessible();
  })

  it('renders in the Integer state', async () => {
    const element = await fixture(Template({...Integer.args}));
    const testNumber = element.querySelector('#test_number');
    expect(testNumber).to.exist;

    testNumber._onInput({ target: { value: "", valueAsNumber: NaN } })
    expect(testNumber.value).to.equal(0);

    testNumber._onInput({ target: { value: "-1", valueAsNumber: -1 } })
    expect(testNumber.value).to.equal(0);

    testNumber._onInput({ target: { value: "200", valueAsNumber: 200 } })
    expect(testNumber.value).to.equal(100);

    testNumber._onInput({ target: { value: "56", valueAsNumber: 56 } })
    expect(testNumber.value).to.equal(56);
  })

});
