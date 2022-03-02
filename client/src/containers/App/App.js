import React, { Component, Suspense, lazy } from 'react'

// Import CSS reset and Global Styles
import 'sanitize.css/sanitize.css';
import 'styles/theme.css';
import 'styles/global.css';
import 'styles/style.css';
import './App.css';

import { AppProvider } from './context'
import Spinner from 'components/Form/Spinner'

const MenuBar = lazy(() => import('containers/MenuBar'));

const Search = lazy(() => import('containers/Search'));
const Login = lazy(() => import('containers/Login'));
const Editor = lazy(() => import('containers/Editor'));
const Setting = lazy(() => import('containers/Setting'));
const Template = lazy(() => import('containers/Template'));

class App extends Component {

  componentDidMount() {
    /* istanbul ignore next */
    if(process.env.NODE_ENV !== 'production'){
      import('remotedev-react-state').then(({ default: connectToDevTools }) => {
        if(typeof window === 'object' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__){
          connectToDevTools(this, {name: "App"})
        }
      })  
    }
    window.addEventListener("scroll", this.onScroll.bind(this), {passive: true});
    //window.addEventListener('resize', this.onResize.bind(this), {passive: true})
    const [ current, params ] = this.getPath(window.location)
    if(current === "hash" && params.access_token){
      this.setHashToken(params)
    } else if(current === "search" && params.code){
      this.setCodeToken(params)
    } else {
      this.loadConfig()
      this.onResize()
    }
  }

  componentWillUnmount() {
    window.removeEventListener("scroll", this.onScroll.bind(this));
    window.removeEventListener("resize", this.onScroll.bind(this));
  }

  render() {
    const protector = () => {
      const { login, current } = this.state
      if(login.data){
        return (
          <div className={`${"main"} ${current.theme}`} >
            <MenuBar />
            {(current.module === "search")?<Search name={current.module} />:null}
            {(current.module === "edit")?<Editor name={current.module} />:null}
            {(current.module === "setting")?<Setting name={current.module} />:null}
            {(current.module === "template")?<Template name={current.module} />:null}
            {(current.modalForm)?current.modalForm:null}
          </div>
        )
      }
      return <Login />
    }
    
    return (
      <AppProvider value={{ data: this.state, setData: this.setData }}>
        <Suspense fallback={<Spinner />}>
          {protector()}
        </Suspense>
        {(this.state.current.request)?<Spinner />:null}
      </AppProvider>
    );
  }
}

export default App;
