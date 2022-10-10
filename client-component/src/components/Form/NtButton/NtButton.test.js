import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './nt-button.js';
import { Template, Default, PrimaryDark, Secondary, SecondaryDark, Border, 
  BorderDark, SmallDisabled, ButtonStyle, Invalid } from  './NtButton.stories.js';

describe('NtButton', () => {
  beforeEach(async () => {
    
  });
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onClick = sinon.spy();
    const onEnter = sinon.spy();
    const element = await fixture(Template({
      ...Default.args, onClick, onEnter
    }));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;
    const value = button.shadowRoot.querySelector('#value')
    expect(value.assignedNodes({flatten: true})[0].textContent).to.equal('Primary');

    button._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onClick);

    button._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.calledOnce(onEnter);

    button._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 13 
    })
    sinon.assert.calledTwice(onEnter);

    button._onKeyEvent({ 
      stopPropagation: sinon.spy(), 
      type: "keypress", keyCode: 20 
    })
    sinon.assert.calledTwice(onEnter);

    await expect(button).shadowDom.to.be.accessible();
  })

  it('renders in the PrimaryDark state', async () => {
    const element = await fixture(Template({...PrimaryDark.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;

    await expect(button).shadowDom.to.be.accessible();
  })

  it('renders in the Secondary state', async () => {
    const element = await fixture(Template({...Secondary.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;

    await expect(button).shadowDom.to.be.accessible();
  })

  it('renders in the SecondaryDark state', async () => {
    const element = await fixture(Template({...SecondaryDark.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;

    await expect(button).shadowDom.to.be.accessible();
  })

  it('renders in the Border state', async () => {
    const element = await fixture(Template({...Border.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;

    await expect(button).shadowDom.to.be.accessible();
  })

  it('renders in the BorderDark state', async () => {
    const element = await fixture(Template({...BorderDark.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;

    await expect(button).shadowDom.to.be.accessible();
  })

  it('renders in the SmallDisabled state', async () => {
    const element = await fixture(Template({...SmallDisabled.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;
  })

  it('renders in the ButtonStyle state', async () => {
    const element = await fixture(Template({...ButtonStyle.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;
  })

  it('renders in the Invalid state', async () => {
    const element = await fixture(Template({...Invalid.args}));
    const button = element.querySelector('#test_button');
    expect(button).to.exist;
  })

});
