import React from "react";
import { Provider } from "react-redux";
import store from "./store/index.js";
// import logo from './logo.svg';
import "./App.css";
import "antd/dist/antd.css";
import Home from "./pages/home/index.js";

function App() {
  return (
  	<Provider store={store}>
	    <div className="App">
	      <Home />
	    </div>
	</Provider>
  );
}

export default App;
