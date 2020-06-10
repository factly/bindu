import { createStore } from "redux";
import rootReducer from "../reducer/index.js";

const store = createStore(rootReducer, {
	spec: {}
});

window.store = store;

export default store;