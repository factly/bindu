import { createStore } from "redux";
import rootReducer from "../reducer/index.js";

const defaultStore = {
	chart: {
		spec: {},
		showOptions: true,
		chartName: "Untitled"
	},
	templates: {
		options: [
			{
		    	name: "Area Chart",
		    	icon: "https://public.flourish.studio/uploads/26101372-8975-461f-80d6-1995eeae96f3.png"
		  	},
		  	{
		    	name: "Stacked Area Chart",
		    	icon: "https://public.flourish.studio/uploads/9684ec0a-79f4-4e28-9721-d14f91504245.png"
		  	},
		  	{
		    	name: "Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/c1c20bf2-cd95-4408-9857-cd5dd6fa4d08.png"
		  	},
		  	{
		    	name: "Horizontal Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/58a046b3-d283-438f-8355-9e92964ea1df.png"
		  	},
		  	{
		    	name: "Horizontal Stack Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/64f6878b-e0d1-4db1-8829-a34816820d65.png"
		  	},
		  	{
		    	name: "Grouped Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/ebed7412-cc56-4ef0-a3ba-97865643f545.png"
		  	},
		  	{
		    	name: "Grouped Line Chart",
		    	icon: "https://public.flourish.studio/uploads/edea3cce-ab40-4fa6-97d2-a1eb324c2f4c.png"
		  	},
		  	{
		    	name: "Line Chart",
		    	icon: "https://public.flourish.studio/uploads/edea3cce-ab40-4fa6-97d2-a1eb324c2f4c.png"
		  	},
		  	{
		    	name: "Pie Chart",
		    	icon: "https://public.flourish.studio/uploads/4192497e-38a7-4f63-8a85-42150ffe7dd4.png"
		  	},
		  	{
		    	name: "Line + Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/782dc711-8ef3-4d2e-8d7e-2e48e73c88e2.png"
		  	},
		  	{
		    	name: "Diverging Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/dedc7386-d69c-4b68-bc28-10fa7552de26.png"
		  	}
		]
	}
}

const store = createStore(rootReducer, defaultStore);
window.store = store;

export default store;