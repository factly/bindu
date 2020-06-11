import { createStore } from "redux";
import rootReducer from "../reducer/index.js";

const defaultStore = {
	chart: {
		spec: {}
	},
	templates: {
		options: [
			{
		    	name: "Grouped Bar Chart",
		    	path: "GroupedBarChart",
		    	icon: "https://public.flourish.studio/uploads/ebed7412-cc56-4ef0-a3ba-97865643f545.png"
		  	}
		],
		selectedOption: -1
	}
}

const store = createStore(rootReducer, defaultStore);
window.store = store;

export default store;