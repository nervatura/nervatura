import sinon from 'sinon'

import { MenuController } from './MenuController.js'
import { store as storeConfig } from '../config/app.js'
import { APP_MODULE, SIDE_VISIBILITY, MENU_EVENT, MODAL_EVENT } from '../config/enums.js'

const host = { 
  addController: ()=>{},
  inputBox: (prm)=>(prm),
}
const store = {
  data: {
    ...storeConfig
  },
  setData: sinon.spy(),
  msg: (value)=>value,
  showToast: sinon.spy(),
}
const app = {
  store,
  requestData: () => ({ value: "OK" }),
  resultError: sinon.spy(),
  loadBookmark: () => ({}),
  signOut: sinon.spy(),
  showHelp: sinon.spy(),
}

describe('MenuController', () => {
  it('onMenuEvent', () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy()
      },
      signOut: sinon.spy(),
      showHelp: sinon.spy(),
    }
    const scrollTo = sinon.spy()
    Object.defineProperty(window, 'scrollTo', { value: scrollTo });

    let menu = new MenuController(host, testApp)
    menu.onMenuEvent({ key: MENU_EVENT.SIDEBAR, data: {} })
    sinon.assert.callCount(testApp.store.setData, 1);

    menu.onMenuEvent({ key: MENU_EVENT.MODULE, data: { value: APP_MODULE.LOGIN } })
    sinon.assert.callCount(testApp.signOut, 1);

    menu.onMenuEvent({ key: MENU_EVENT.MODULE, data: { value: APP_MODULE.HELP } })
    sinon.assert.callCount(testApp.showHelp, 1);

    menu.onMenuEvent({ key: MENU_EVENT.MODULE, data: { value: APP_MODULE.BOOKMARK } })
    sinon.assert.callCount(testApp.store.setData, 2);

    menu.onMenuEvent({ key: MENU_EVENT.MODULE, data: { value: APP_MODULE.EDIT } })
    sinon.assert.callCount(testApp.store.setData, 3);

    menu.onMenuEvent({ key: MENU_EVENT.MODULE, data: { value: APP_MODULE.SETTING } })
    sinon.assert.callCount(testApp.store.setData, 5);

    menu.onMenuEvent({ key: MENU_EVENT.SCROLL, data: {} })
    sinon.assert.callCount(scrollTo, 1);

    menu.onMenuEvent({ key: "missing", data: {} })
    
    testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy(),
        data: {
          ...app.store.data,
          current: {
            ...app.store.data.current,
            side: SIDE_VISIBILITY.SHOW
          }
        }
      },
    }

    menu = new MenuController(host, testApp)
    menu.onMenuEvent({ key: MENU_EVENT.SIDEBAR, data: {} })
    sinon.assert.callCount(testApp.store.setData, 1);

  })

  it('onModalEvent', () => {
    let testApp = {
      ...app,
      store: {
        ...app.store,
        setData: sinon.spy()
      },
      requestData: () => ({}),
      resultError: sinon.spy(),
      loadBookmark: () => ({}),
    }
    let menu = new MenuController(host, testApp)
    menu.onModalEvent({ key: MODAL_EVENT.CANCEL, data: {} })
    sinon.assert.callCount(testApp.store.setData, 1);

    menu.onModalEvent({ key: MODAL_EVENT.SELECTED, data: { view: "bookmark", row: {} } })
    sinon.assert.callCount(testApp.store.setData, 2);

    menu.onModalEvent({ 
      key: MODAL_EVENT.SELECTED, 
      data: { view: "bookmark", row: { cfgroup: "browser", filters: {}, columns: {} } } 
    })
    sinon.assert.callCount(testApp.store.setData, 3);

    testApp = {
      ...app,
      store: {
        ...app.store,
        setData: (key, data) => {
          // debugger;
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.CANCEL })
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK })
          }
        }
      },
      requestData: () => ({}),
      resultError: sinon.spy(),
      loadBookmark: () => ({}),
    }
    menu = new MenuController(host, testApp)
    menu.onModalEvent({ key: MODAL_EVENT.DELETE, data: { bookmark_id: 1, menubar: { _onBookmark: sinon.spy() } } })

    testApp = {
      ...app,
      store: {
        ...app.store,
        setData: (key, data) => {
          if(data && data.modalForm){
            data.modalForm.onEvent.onModalEvent({ key: MODAL_EVENT.OK })
          }
        }
      },
      requestData: () => ({ error: {} }),
      resultError: sinon.spy(),
      loadBookmark: () => ({}),
    }
    menu = new MenuController(host, testApp)
    menu.onModalEvent({ key: MODAL_EVENT.DELETE, data: { bookmark_id: 1, menubar: { _onBookmark: sinon.spy() } } })

    menu.onModalEvent({ key: "missing", data: {} })

  })

})