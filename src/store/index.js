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
		    	icon: "https://public.flourish.studio/uploads/ebed7412-cc56-4ef0-a3ba-97865643f545.png"
		  	},
		  	{
		    	name: "Line Chart",
		    	icon: "https://public.flourish.studio/uploads/edea3cce-ab40-4fa6-97d2-a1eb324c2f4c.png"
		  	},
		  	{
		    	name: "Pie Chart",
		    	icon: "https://public.flourish.studio/uploads/4192497e-38a7-4f63-8a85-42150ffe7dd4.png"
		  	}
		]
	}
}

const store = createStore(rootReducer, defaultStore);
window.store = store;

export default store;