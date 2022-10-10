import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './login-page.js';
import { Template, Default, DarkServer } from  './LoginPage.stories.js';

describe('LoginPage', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onPageEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onPageEvent
    }));
    const loginPage = element.querySelector('#login_page');
    expect(loginPage).to.exist;

    const loginButton = loginPage.shadowRoot.querySelector('#login')
    loginButton._onClick({
      stopPropagation: sinon.spy()
    })
    sinon.assert.calledOnce(onPageEvent);

    const username = loginPage.shadowRoot.querySelector('#username')
    username._onInput({ target: { value: "value" } })
    sinon.assert.calledTwice(onPageEvent);
    expect(username.value).to.equal("value");

    const password = loginPage.shadowRoot.querySelector('#password')
    password._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.calledThrice(onPageEvent);

    const database = loginPage.shadowRoot.querySelector('#database')
    database._onKeyEvent({ 
      stopPropagation: sinon.spy(), preventDefault: sinon.spy(),
      type: "keydown", keyCode: 13 
    })
    sinon.assert.callCount(onPageEvent, 4);

    const lang = loginPage.shadowRoot.querySelector('#lang');
    lang._onInput({ target: { value: "de" } })
    expect(lang.value).to.equal("de");
    loginButton.click()
    
    await expect(loginPage).shadowDom.to.be.accessible();
  })

  it('renders in the DarkServer state', async () => {
    const element = await fixture(Template({...DarkServer.args}));
    const loginPage = element.querySelector('#login_page');
    expect(loginPage).to.exist;

    const loginButton = loginPage.shadowRoot.querySelector('#login')
    loginButton.click()
  })

})