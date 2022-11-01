import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './form-select.js';
import { Template, Default, NotNull } from  './Select.stories.js';

describe('Select', () => {
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
    const select = element.querySelector('#test_select');
    expect(select).to.exist;
    expect(select.value).to.equal('value1');

    select._onInput({ target: { value: "value2" } })
    sinon.assert.calledOnce(onChange);
    expect(select.value).to.equal('value2');

    select._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.calledOnce(onEnter);

    select._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 13 
    })
    sinon.assert.calledTwice(onEnter);

    select._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 20 
    })
    sinon.assert.calledTwice(onEnter);

    await expect(select).shadowDom.to.be.accessible();
  })

  it('renders in the NotNull state', async () => {
    const element = await fixture(Template({...NotNull.args}));
    const select = element.querySelector('#test_select');
    expect(select).to.exist;
  })

});
