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
		    	id: 0,
		    	name: "Area Chart",
		    	icon: "https://public.flourish.studio/uploads/26101372-8975-461f-80d6-1995eeae96f3.png"
		  	},
		  	{
		    	id: 1,
		    	name: "Stacked Area Chart",
		    	icon: "https://public.flourish.studio/uploads/9684ec0a-79f4-4e28-9721-d14f91504245.png"
		  	},
		  	{
		    	id: 2,
		    	name: "Stacked Area Chart(Proportional)",
		    	icon: "https://public.flourish.studio/uploads/9684ec0a-79f4-4e28-9721-d14f91504245.png"
		  	},
		  	{
		    	id: 3,
		    	name: "Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/c1c20bf2-cd95-4408-9857-cd5dd6fa4d08.png"
		  	},
		  	{
		    	id: 4,
		    	name: "Horizontal Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/58a046b3-d283-438f-8355-9e92964ea1df.png"
		  	},
		  	{
		    	id: 5,
		    	name: "Horizontal Stack Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/64f6878b-e0d1-4db1-8829-a34816820d65.png"
		  	},
		  	{
		    	id: 6,
		    	name: "Grouped Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/ebed7412-cc56-4ef0-a3ba-97865643f545.png"
		  	},
		  	{
		    	id: 7,
		    	name: "Grouped Line Chart",
		    	icon: "https://public.flourish.studio/uploads/edea3cce-ab40-4fa6-97d2-a1eb324c2f4c.png"
		  	},
		  	{
		    	id: 8,
		    	name: "Line Chart",
		    	icon: "https://public.flourish.studio/uploads/edea3cce-ab40-4fa6-97d2-a1eb324c2f4c.png"
		  	},
		  	{
		    	id: 9,
		    	name: "Line Chart (Projected)",
		    	icon: "https://public.flourish.studio/uploads/edea3cce-ab40-4fa6-97d2-a1eb324c2f4c.png"
		  	},
		  	{
		    	id: 10,
		    	name: "Pie Chart",
		    	icon: "https://public.flourish.studio/uploads/4192497e-38a7-4f63-8a85-42150ffe7dd4.png"
		  	},
		  	{
		    	id: 11,
		    	name: "Donut Chart",
		    	icon: "https://public.flourish.studio/uploads/4192497e-38a7-4f63-8a85-42150ffe7dd4.png"
		  	},
		  	{
		    	id: 12,
		    	name: "Line + Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/782dc711-8ef3-4d2e-8d7e-2e48e73c88e2.png"
		  	},
		  	{
		    	id: 13,
		    	name: "Diverging Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/dedc7386-d69c-4b68-bc28-10fa7552de26.png"
		  	},
		  	{
		    	id: 14,
		    	name: "Grouped Bar Chart(Proportional)",
		    	icon: "https://public.flourish.studio/uploads/ebed7412-cc56-4ef0-a3ba-97865643f545.png"
		  	},
		  	{
		    	id: 15,
		    	name: "Horizontal Grouped Bar Chart(Proportional)",
		    	icon: "https://public.flourish.studio/uploads/ebed7412-cc56-4ef0-a3ba-97865643f545.png"
		  	},
		  	{
		    	id: 16,
		    	name: "Grid of Pie Chart",
		    	icon: "https://public.flourish.studio/uploads/55a3ab6f-9b55-41cf-847d-e38cd17d2dd2.png"
		  	},
		  	{
		    	id: 17,
		    	name: "Grid of Bar Chart",
		    	icon: "https://public.flourish.studio/uploads/1c44c202-9de5-4fc6-a403-1c310cef0fa3.png"
		  	},
		  	{
		    	id: 18,
		    	name: "Grid of Line Chart",
		    	icon: "https://public.flourish.studio/uploads/4f6b9081-4ada-481f-9232-110615db716e.png"
		  	},
		  	{
		    	id: 19,
		    	name: "Grouped Bar",
		    	icon: "https://public.flourish.studio/uploads/b7f72bb3-9432-4a67-8082-f23370b6b412.png"
		  	}
		]
	}
}

const store = createStore(rootReducer, defaultStore);
window.store = store;

export default store;