import React from 'react';

const AppStore = React.createContext({})

export const AppProvider = AppStore.Provider
export const AppConsumer = AppStore.Consumer
export default AppStore