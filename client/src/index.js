// Import all the third party stuff
import React from 'react';
import ReactDOM from 'react-dom';

import * as serviceWorkerRegistration from './service-workerRegistration';
import { ClearCacheProvider } from 'react-clear-cache';

// Import root app
import App from './containers/App';

ReactDOM.render(
    <React.StrictMode>
      <ClearCacheProvider duration={5000} auto={true} basePath={window.location.pathname}>
        <App />
      </ClearCacheProvider>
    </React.StrictMode>, 
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorkerRegistration.register();
