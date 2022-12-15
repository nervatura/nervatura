
export class StateController {
  constructor(host, config) {
    this.host = host;
    this._data = config || {};
    host.addController(this);
  }

  get data(){
    return this._data
  }

  set data(props) {
    const { key, value, update } = props
    if(this._data[key] && (typeof(value) === "object")){
      this._data[key] = {...this._data[key], ...value}
      if(update !== false){
        this.host.requestUpdate();
      }
    }
  }

}
