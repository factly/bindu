import React from 'react';
import { Provider } from 'react-redux';
import store from './store/index.js';
// import logo from './logo.svg';
import './App.css';
import 'antd/dist/antd.css';
import Home from './pages/home/index.js';
import Templates from './pages/templates/index.js';
import Chart from './pages/charts/index.js';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import BasicLayout from './layout/basic/index.js';
const routes = [
  {
    path: '/templates',
    Component: Templates,
  },
  {
    path: '/chart/:id',
    Component: Chart,
  },
  {
    path: '/',
    Component: Home,
  },
];

function App() {
  return (
    <Provider store={store}>
      <Router basename={process.env.PUBLIC_URL}>
        <BasicLayout>
          <Switch>
            {routes.map((route) => (
              <Route key={route.path} exact path={route.path} component={route.Component} />
            ))}
          </Switch>
        </BasicLayout>
      </Router>
    </Provider>
  );
}

export default App;
