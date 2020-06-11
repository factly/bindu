import sharedReducer from "./components/shared/index.js";
import barReducer from "./components/bar/index.js";

const chartReducer = (state = {}, action) => {
	
  switch (action.chart) {
    case 'shared':
      return sharedReducer(state, action);
    case 'bar':
      return barReducer(state, action);
    default:
      switch (action.type){
			case 'set-config':
				return {...state, spec: action.value}
			default:
				return state;
		}
  }
}

export default chartReducer;