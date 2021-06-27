import React from 'react';
import { Provider } from 'react-redux';
import store from './store/index.js';
import './App.css';
import 'antd/dist/antd.css';

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import BasicLayout from './layout/basic/index.js';
import routes from './config/routesConfig';

function App() {
  return (
    <Provider store={store}>
      <Router basename={process.env.PUBLIC_URL}>
        <BasicLayout>
          <Switch>
            {Object.values(routes).map((route) => (
              <Route key={route.path} exact path={route.path} component={route.Component} />
            ))}
          </Switch>
        </BasicLayout>
      </Router>
    </Provider>
  );
}

export default App;
