import { APP_MODULE, SIDE_VISIBILITY, MENU_EVENT, MODAL_EVENT, EDITOR_EVENT, SIDE_EVENT } from '../config/enums.js'

export class MenuController {
  constructor(host) {
    this.host = host;
    this.deleteBookmark = this.deleteBookmark.bind(this)
    this.onMenuEvent = this.onMenuEvent.bind(this)
    this.onModalEvent = this.onModalEvent.bind(this)
    host.addController(this);
  }

  async deleteBookmark({ bookmark_id }){
    const { inputBox } = this.host.app.host
    const { data, setData } = this.host.app.store
    const { requestData, resultError, loadBookmark, msg } = this.host.app
    const { modalBookmark } = this.host
    const login = data[APP_MODULE.LOGIN]
    const  modalForm = inputBox({ 
      title: msg("", { id: "msg_warning" }), 
      message: msg("", { id: "msg_delete_text" }),
      infoText: msg("", { id: "msg_delete_info" }),
      onEvent: {
        onModalEvent: async ({ key }) => {
          setData("current", { modalForm: null })
          if(key === MODAL_EVENT.CANCEL){
            return setData("current", { modalForm: modalBookmark() })
          }
          const result = await requestData("/ui_userconfig", 
            { method: "DELETE", query: { id: bookmark_id } })
          if(result && result.error){
            return resultError(result)
          }
          const bookmark = await loadBookmark({user_id: login.data.employee.id})
          /* c8 ignore next 1 */
          return setData("current", { modalForm: modalBookmark(bookmark) })
        }
      }
    })
    setData("current", { modalForm })
  }

  onMenuEvent({key, data}){
    const { setData } = this.host.app.store
    const { current, setting } = this.host.app.store.data
    const { signOut, showHelp, currentModule } = this.host.app
    const { modalBookmark } = this.host
    switch (key) {
      case MENU_EVENT.SIDEBAR:
        setData("current", {
          side: (current.side === SIDE_VISIBILITY.SHOW) ? SIDE_VISIBILITY.HIDE : SIDE_VISIBILITY.SHOW 
        })
        break;

      case MENU_EVENT.MODULE:
        switch (data.value) {
          case APP_MODULE.LOGIN:
            signOut()
            break;
          
          case APP_MODULE.HELP:
            showHelp("")
            break;
    
          case APP_MODULE.BOOKMARK:
            setData("current", { modalForm: modalBookmark() })
            break;
        
          default:
            const values = { module: data.value, menu: "", side: SIDE_VISIBILITY.HIDE }
            let content = null
            if(data.value === APP_MODULE.SETTING && !setting.group_key){
              setData(APP_MODULE.SETTING, { group_key: "group_admin" })
              content = { fkey: "checkSetting", args: [{ type: 'setting' }, SIDE_EVENT.LOAD_SETTING] }
            }
            currentModule({ data: { ...values }, content })
            break;
        }
        break;

      case MENU_EVENT.SCROLL:
        window.scrollTo(0,0);
        break;

      default:
        break;
    }
  }

  onModalEvent({key, data}){
    const { currentModule } = this.host.app
    const { setData } = this.host.app.store
    const { search } = this.host.app.store.data
    switch (key) {
      case MODAL_EVENT.CANCEL:
        setData("current", { modalForm: null })
        break;

      case MODAL_EVENT.SELECTED:
        if((data.view === "bookmark") && (data.row.cfgroup === "browser")){
          const search_data = {
            ...search,
            filters: {
              ...search.filters,
              [data.row.view]: data.row.filters
            },
            columns: {
              ...search.columns,
              [data.row.view]: data.row.columns
            }
          }
          currentModule({ 
            data: { module: APP_MODULE.SEARCH, modalForm: null }, 
            content: { fkey: "showBrowser", args: [data.row.vkey, data.row.view, search_data] }
          })
        } else {
          currentModule({ 
            data: { module: APP_MODULE.EDIT, modalForm: null }, 
            content: { 
              fkey: "checkEditor", 
              args: [{ntype: data.row.ntype, ttype: data.row.transtype, id: data.row.id}, EDITOR_EVENT.LOAD_EDITOR] 
            }
          })
        }
        break;

      case MODAL_EVENT.DELETE:
        this.deleteBookmark(data)
        break;
    
      default:
        break;
    }
  }
}