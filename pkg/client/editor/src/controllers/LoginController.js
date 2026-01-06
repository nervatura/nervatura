import { APP_MODULE, LOGIN_PAGE_EVENT, TOAST_TYPE } from '../config/enums.js'

export class LoginController {
  constructor(host) {
    this.host = host;
    this.onLogin = this.onLogin.bind(this)
    this.onPageEvent = this.onPageEvent.bind(this)
    host.addController(this);
  }

  async onLogin() {
    const { data, setData } = this.host.app.store
    const { requestData, resultError, currentModule, showToast, msg } = this.host.app
    var options = {
      method: "POST",
      data: {
        user_name: data[APP_MODULE.LOGIN].username, password: data[APP_MODULE.LOGIN].password,
        database: data[APP_MODULE.LOGIN].database
      }
    }
    var result = await requestData("/auth/login", options)
    if(!result.token){
      return resultError(result)
    }
    var resultData = {
      token: result.token,
      content: { 
        fkey: "setTemplate", 
        args: [{ type: "_sample" }]
      }
    }
    if(data[APP_MODULE.LOGIN].code){
      options = { method: "GET", token: result.token }
      result = await requestData("/config/"+data[APP_MODULE.LOGIN].code, options)
      if(!result.error){
        resultData.content = { 
          fkey: "setTemplate", 
          args: [{ type: "template", report: result }]
        }
      } else {
        showToast(TOAST_TYPE.ERROR, msg("", { id: "login_template_code_err" }))
      }
    }
    setData(APP_MODULE.LOGIN, {
      data: resultData
    })
    localStorage.setItem("database", data[APP_MODULE.LOGIN].database);
    localStorage.setItem("username", data[APP_MODULE.LOGIN].username);
    localStorage.setItem("server", data[APP_MODULE.LOGIN].server);
    localStorage.setItem("code", data[APP_MODULE.LOGIN].code);
    
    currentModule({ 
      data: { module: APP_MODULE.TEMPLATE }, 
      content: resultData.content 
    })
    
    window.history.replaceState(null, null, window.location.pathname)
  }

  onPageEvent({key, data}){
    const { setData } = this.host.app.store
    switch (key) {
      case LOGIN_PAGE_EVENT.CHANGE:
        setData(APP_MODULE.LOGIN, {
          [data.fieldname]: data.value 
        })
        break;

      case LOGIN_PAGE_EVENT.THEME:
      case LOGIN_PAGE_EVENT.LANG:
        setData("current", { 
          [key]: data 
        })
        localStorage.setItem([key], data);
        break;
      
      case LOGIN_PAGE_EVENT.LOGIN:
        this.onLogin()
        break;
    
      default:
        break;
    }
  }

}