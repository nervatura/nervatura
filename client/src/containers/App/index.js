import update from 'immutability-helper';

import { default as AppComponent } from "./App";
import { store } from 'config/app'
import { request, appActions } from './actions'

class App extends AppComponent {

  constructor(props) {
    super(props);

    this.state = update(store, {$merge: props.data||{} })
    this.setData = this.setData.bind(this)
  }

  setData(key, data, callback) {
    if(key && this.state[key] && typeof data === "object" && data !== null){
      this.setState({ [key]: update(this.state[key], {$merge: data}) }, 
        ()=>{ if(callback) {callback()} })
    } else if(key){
      this.setState({ [key]:  data }, ()=>{ if(callback) {callback()} })
    }
  }

  getPath(location) {
    const getParams = (prmString) => {
      let params = {}
      prmString.split('&').forEach(prm => {
        params[prm.split("=")[0]] = prm.split("=")[1]
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

  onResize() {
    if((this.state.current.clientHeight !== window.innerHeight) || 
      (this.state.current.clientWidth !== window.innerWidth)){
      this.setData("current", { clientHeight: window.innerHeight, clientWidth: window.innerWidth })
    }
  }

  onScroll() {
    const scrollTop = ((document.body.scrollTop > 100) || (document.documentElement.scrollTop > 100))
    if(this.state.current.scrollTop !== scrollTop){
      this.setData("current", { scrollTop: scrollTop })
    }
  }

  async loadConfig(){
    const app = appActions(this.state, this.setData)
    let config = update({}, {$merge: this.state.session})
    try {
      const result = await request(this.state.login.server+"/config", {
        method: "GET",
        headers: { "Content-Type": "application/json" }
      })
      if(result.locales && (typeof result.locales == "object")){
        config = update(config, {locales: {$merge: result.locales }})
      }
      this.setData("session", config )
      if(localStorage.getItem("lang") && config.locales[localStorage.getItem("lang")] 
        && (localStorage.getItem("lang") !== this.state.current.lang)){
          this.setData("current", {lang: localStorage.getItem("lang")} )
        }
      const [ current, params ] = this.getPath(window.location)
      if(current === "hash" && params.access_token){
        app.tokenValidation(params)
      }
      if(current === "search" && params.code){
        app.setCodeToken(params)
      }
    } catch (error) {
      app.resultError(error)
    }
  }
}

export default App;