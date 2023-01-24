import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './client-login.js';
import { Template, Default, DarkServer } from  './Login.stories.js';

describe('Login', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onPageEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onPageEvent
    }));
    const Login = element.querySelector('#login_page');
    expect(Login).to.exist;

    const loginButton = Login.shadowRoot.querySelector('#login')
    loginButton._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onPageEvent);

    const username = Login.shadowRoot.querySelector('#username')
    username._onInput({ target: { value: "value" } })
    sinon.assert.calledTwice(onPageEvent);
    expect(username.value).to.equal("value");

    const password = Login.shadowRoot.querySelector('#password')
    password._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.calledThrice(onPageEvent);

    const database = Login.shadowRoot.querySelector('#database')
    database._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.callCount(onPageEvent, 4);

    const theme = Login.shadowRoot.querySelector('#theme')
    theme.click()
    sinon.assert.callCount(onPageEvent, 5);

    const lang = Login.shadowRoot.querySelector('#lang');
    lang._onInput({ target: { value: "de" } })
    expect(lang.value).to.equal("de");
    loginButton.click()
    
    await expect(Login).shadowDom.to.be.accessible();
  })

  it('renders in the DarkServer state', async () => {
    const element = await fixture(Template({...DarkServer.args}));
    const Login = element.querySelector('#login_page');
    expect(Login).to.exist;

    const loginButton = Login.shadowRoot.querySelector('#login')
    loginButton.click()
  })

})