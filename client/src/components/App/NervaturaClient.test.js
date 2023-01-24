import { html } from 'lit';
import { fixture, expect } from '@open-wc/testing';

import { StateController } from '../../controllers/StateController.js'
import { store as storeConfig } from '../../config/app.js'
import { APP_MODULE } from '../../config/enums.js'

import './nervatura-client.js';

describe('Nervatura Client', () => {
  it('renders without crashing', async () => {
    const client = await fixture(html`<nervatura-client id="client" ></nervatura-client>`);
    expect(client.id).to.equal("client")
    client._onScroll()
    client.inputBox({ onEvent: {} })
    
  })

  it('form-spinner', async () => {
    const state = new StateController({ addController: ()=>{} }, {...storeConfig,
      current: {
        ...storeConfig.current,
        request: true
      }
    })
    const client = await fixture(html`<nervatura-client id="client" .state=${state} ></nervatura-client>`);
    expect(client.id).to.equal("client")
  })

  it('protector SEARCH', async () => {
    const state = new StateController({ addController: ()=>{} }, {...storeConfig,
      current: {
        ...storeConfig.current,
        modalForm: html`<div></div>`,
        module: APP_MODULE.SEARCH
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {}
      }
    })
    const client = await fixture(html`<nervatura-client id="client" .state=${state} ></nervatura-client>`);
    expect(client.id).to.equal("client")
  })

  it('protector EDIT', async () => {
    const state = new StateController({ addController: ()=>{} }, {...storeConfig,
      current: {
        ...storeConfig.current,
        module: APP_MODULE.EDIT
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {}
      }
    })
    const client = await fixture(html`<nervatura-client id="client" .state=${state} ></nervatura-client>`);
    expect(client.id).to.equal("client")
  })

  it('protector SETTING', async () => {
    const state = new StateController({ addController: ()=>{} }, {...storeConfig,
      current: {
        ...storeConfig.current,
        module: APP_MODULE.SETTING
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {}
      }
    })
    const client = await fixture(html`<nervatura-client id="client" .state=${state} ></nervatura-client>`);
    expect(client.id).to.equal("client")
  })

  it('protector TEMPLATE', async () => {
    const state = new StateController({ addController: ()=>{} }, {...storeConfig,
      current: {
        ...storeConfig.current,
        module: APP_MODULE.TEMPLATE
      },
      [APP_MODULE.LOGIN]: {
        ...storeConfig[APP_MODULE.LOGIN],
        data: {}
      }
    })
    const client = await fixture(html`<nervatura-client id="client" .state=${state} ></nervatura-client>`);
    expect(client.id).to.equal("client")
  })

})