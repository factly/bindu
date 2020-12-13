import React from 'react';
import { Provider } from 'react-redux';
import store from './store/index.js';
// import logo from './logo.svg';
import './App.css';
import 'antd/dist/antd.css';
import Home from './pages/home/index.js';
import Templates from './pages/templates/index.js';
import Tags from './pages/tags/index.js';
import TagsCreate from './pages/tags/CreateTag';
import TagsEdit from './pages/tags/EditTag';
import Categories from './pages/categories/index.js';
import CategoriesCreate from './pages/categories/CreateCategory';
import CategoriesEdit from './pages/categories/EditCategory';
import SavedCharts from './pages/charts/saved.js';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import BasicLayout from './layout/basic/index.js';
import EditCharts from './pages/charts/EditCharts.js';
import CreateCharts from './pages/charts/CreateChart.js';

const routes = [
  {
    path: '/templates',
    Component: Templates,
  },
  {
    path: '/chart/:id',
    Component: CreateCharts,
  },
  {
    path: '/charts/saved',
    Component: SavedCharts,
  },
  {
    path: '/charts/:id/edit',
    Component: EditCharts,
  },
  {
    path: '/',
    Component: Home,
  },
  {
    path: '/tags',
    Component: Tags,
  },
  {
    path: '/tags/:id/edit',
    Component: TagsEdit,
  },
  {
    path: '/tags/create',
    Component: TagsCreate,
  },
  {
    path: '/categories',
    Component: Categories,
  },
  {
    path: '/categories/create',
    Component: CategoriesCreate,
  },
  {
    path: '/categories/:id/edit',
    Component: CategoriesEdit,
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
