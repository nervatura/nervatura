import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './form-datetime.js';
import { Template, Default, DateTime, Time } from  './DateTime.stories.js';

describe('Date', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onChange = sinon.spy()
    const onEnter = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onChange, onEnter
    }));
    const testDate = element.querySelector('#test_date');
    expect(testDate).to.exist;
    expect(testDate.value).to.equal("");

    testDate._onInput({ target: { value: "2024-12-24" } })
    sinon.assert.calledOnce(onChange);
    expect(testDate.value).to.equal("2024-12-24");

    testDate._onBlur()
    expect(testDate.value).to.equal("2024-12-24");

    testDate._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.calledOnce(onEnter);

    testDate._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 13 
    })
    sinon.assert.calledTwice(onEnter);

    testDate._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 20 
    })
    sinon.assert.calledTwice(onEnter);

    testDate._onFocus()

    testDate._defaultValue()

    await expect(testDate).shadowDom.to.be.accessible();
  })

  it('renders in the DateTime state', async () => {
    const element = await fixture(Template({...DateTime.args}));
    const testDate = element.querySelector('#test_date');
    expect(testDate).to.exist;

    testDate._onInput({ target: { value: "2024-12-24" } })
    expect(testDate.value).to.equal("2024-12-24");

    testDate._onInput({ target: { value: "" } })
    expect(testDate.value).to.not.equal("");
  })

  it('renders in the Time state', async () => {
    const element = await fixture(Template({...Time.args}));
    const testDate = element.querySelector('#test_date');
    expect(testDate).to.exist;

    testDate._defaultValue()
  })

  it('renders in the invalid type state', async () => {
    const element = await fixture(Template({
      ...Default.args,
      type: "number"
    }));
    const testDate = element.querySelector('#test_date');
    expect(testDate).to.exist;
  })

});
