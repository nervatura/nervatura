import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './form-input.js';
import { Template, Default, Color, File, Password } from  './Input.stories.js';

describe('Input', () => {
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
    const testInput = element.querySelector('#test_input');
    expect(testInput).to.exist;
    expect(testInput.value).to.equal("");

    testInput._onInput({ target: { value: "value" } })
    sinon.assert.calledOnce(onChange);
    expect(testInput.value).to.equal("value");

    testInput._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.calledOnce(onEnter);

    testInput._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 13 
    })
    sinon.assert.calledTwice(onEnter);

    testInput._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 20 
    })
    sinon.assert.calledTwice(onEnter);

    await expect(testInput).shadowDom.to.be.accessible();
  })

  it('renders in the Color state', async () => {
    const element = await fixture(Template({...Color.args}));
    const testInput = element.querySelector('#test_input');
    expect(testInput).to.exist;
  })

  it('renders in the File state', async () => {
    const onChange = sinon.spy()
    const element = await fixture(Template({
      ...File.args, onChange
    }));
    const testInput = element.querySelector('#test_input');
    expect(testInput).to.exist;

    testInput._onInput({ target: { files: [] } })
    sinon.assert.calledOnce(onChange);
  })

  it('renders in the Password state', async () => {
    const element = await fixture(Template({...Password.args}));
    const testInput = element.querySelector('#test_input');
    expect(testInput).to.exist;
  })

  it('renders in the invalid type state', async () => {
    const element = await fixture(Template({
      ...Default.args,
      type: "number"
    }));
    const testInput = element.querySelector('#test_input');
    expect(testInput).to.exist;
  })

});
