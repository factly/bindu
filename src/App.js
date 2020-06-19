import React from "react";
import { Provider } from "react-redux";
import store from "./store/index.js";
// import logo from './logo.svg';
import "./App.css";
import "antd/dist/antd.css";
import Home from "./pages/home/index.js";
import Templates from "./pages/templates/index.js";
import Chart from "./pages/charts/index.js";

import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

function App() {
  return (
  	<Provider store={store}>
		<Router>    
		    <Switch>
	          <Route path="/templates">
	            <Templates />
	          </Route>
	          <Route path="/chart/:id">
	            <Chart />
	          </Route>
	          <Route path="/">
	            <Home />
	          </Route>
	        </Switch>
	    </Router>
	</Provider>
  );
}

export default App;
