import { APP_MODULE, TOAST_TYPE } from '../config/enums.js'
import { getSql, request, saveToDisk } from './Utils.js'

export class AppController {
  constructor(host) {
    this.host = host;
    this.modules = {}
    
    this.getSql = getSql
    this.request = request
    this.saveToDisk = saveToDisk

    this.currentModule = this.currentModule.bind(this)
    this.getSetting = this.getSetting.bind(this)
    this.msg = this.msg.bind(this)
    this.requestData = this.requestData.bind(this)
    this.resultError = this.resultError.bind(this)
    this.showHelp = this.showHelp.bind(this)
    this.showToast = this.showToast.bind(this)
    this.signOut = this.signOut.bind(this)
    host.addController(this);
  }

  hostConnected(){
    const { state, setData } = this.host
    this.store = {
      data: state.data,
      setData
    };
    
    this._loadConfig(window.location)
  }

  async _loadConfig(location){
    const { data, setData } = this.store
    const getPath = () => {
      const getParams = (prmString) => {
        const params = {}
        prmString.split('&').forEach(prm => {
          const index = String(prm).indexOf("=")
          const fname = String(prm).substring(0,(index >0)?index:String(prm).length)
          const value = ((index > -1) && (index < String(prm).length)) ? String(prm).substring(index+1) : ""
          params[fname] = value
        });
        return params
      }
      if(location.hash){
        return ["hash", getParams(location.hash.substring(1))]
      }
      if(location.search){
        return ["search", getParams(location.search.substring(1))]
      }
      const path = location.pathname.substring(1).split("/")
      return [path[0], path.slice(1)]
    }

    const [ current, params ] = getPath()
    if (params.session) {
      this.request(`/client/api/template/${params.session}/${params.code}`, {
        method: "POST",
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      }).then(result => {
        setData(APP_MODULE.LOGIN, {
          data: {
            token: result.token,
          }
        })
        window.history.replaceState(null, null, window.location.pathname)
        return this.currentModule({ 
          data: { module: APP_MODULE.TEMPLATE }, 
          content: { 
            fkey: "setTemplate", 
            args: [{ type: "template", report: result.report }]
          } 
        })
      }).catch(error => {
        setData("error", error )
        this.showToast(TOAST_TYPE.ERROR, error.message)
      })
    }
    
    import('../components/Login/client-login.js');
  }

  async currentModule({data, content}) {
    const { setData } = this.store

    const modules = {
      template: async ()=>{
        const { TemplateController } = await import('./TemplateController.js');
        await import('../components/Template/Template/client-template.js');

        this.modules.template = new TemplateController(this.host)
      }
    }

    if(!this.modules[data.module]){
      await modules[data.module]()
    }
    setData("current", {...data})
    if(content && this.modules[data.module]){
      this.modules[data.module][content.fkey](...content.args)
    }
  }

  getSetting(key) {
    const { ui } = this.store.data
    if(key === "ui"){
      const values = {...ui}
        Object.keys(values).forEach(ikey => {
          if(localStorage.getItem(ikey)){
            values[ikey] = localStorage.getItem(ikey)
          }
        });
        return values
    }
    return localStorage.getItem(key) || ui[key] || "";
  }

  msg(defaultValue, props) {
    let value = defaultValue
    const {locales} = this.store.data.session
    const {lang} = this.store.data.current
    if(locales[lang] && locales[lang][props.id]){
      value = locales[lang][props.id]
    } else if((lang !== "en") && locales.en[props.id]) {
      value = locales.en[props.id]
    }
    return value
  }

  async requestData(path, params, silent) {
    const { data, setData } = this.store
    let options = params
    try {
      if (!silent){
        setData("current", {
          request: true 
        })
      }
      let url = (data.session.serverURL === "SERVER")?
        data.session.apiPath+path : data.login.server+path
      const token = (data.login.data) ? data.login.data.token : options.token || ""
      if (!options.headers)
        options = {
          ...options, 
          headers: {}
        }
      options = {
        ...options, 
        headers: {
          ...options.headers,
          "Content-Type": "application/json"
        }
      }
      if(token !== ""){
        options = {
          ...options, 
          headers: {
            ...options.headers,
            "Authorization": `Bearer ${token}`
          }
        }
      }
      if (options.data){
        options = {
          ...options, 
          body: JSON.stringify(options.data) 
        }
      }
      if(options.query){
        const query = new URLSearchParams();
        Object.keys(options.query).forEach(key => {
          query.append(key, options.query[key])
        });
        url += `?${query.toString()}`
      }
      
      const result = await this.request(url, options)
      if (!silent) {
        setData("current", { 
          request: false 
        })
      }
      if(result && result.code && result.message){
        if(result.code === 401){
          this.signOut()
        }
        return { error: { message: result.message }, data: null }
      }
      return result
    } catch (err) {
      if(!silent){
        setData("current", { 
          request: false 
        })
      }
      return { error: { message: err.message }, data: null }
    }
  }

  resultError(result) {
    const { setData } = this.store
    if(result.error){
      setData("error", result.error )
    }
    if(result.error && result.error.message){
      this.showToast(TOAST_TYPE.ERROR, result.error.message)
    } else {
      this.showToast(TOAST_TYPE.ERROR, 
        this.msg("Internal Server Error", { id: "error_internal" }) )
    }
    return false
  }

  showHelp(key) {
    const { data } = this.store
    const element = document.createElement("a")
    element.setAttribute("href", data.session.helpPage+key)
    element.setAttribute("target", "_blank")
    document.body.appendChild(element)
    element.click()
  }

  showToast(type, value, toastTimeout) {
    const { current } = this.store.data
    const timeout = (typeof(toastTimeout) !== "undefined") ? toastTimeout : this.getSetting("toastTimeout")
    if(current.toast){
      current.toast.show({
        type, value, timeout
      })
    }
  }

  signOut() {
    const { data, setData } = this.store
    /* c8 ignore next 4 */
    if(data[APP_MODULE.LOGIN].callback){
      window.location.replace(data[APP_MODULE.LOGIN].callback)
      return
    }
    setData(APP_MODULE.LOGIN, { 
      data: null, token: null 
    })
  }

}