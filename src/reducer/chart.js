import sharedReducer from "./components/shared/index.js";
import barReducer from "./components/bar/index.js";
import lineReducer from "./components/line/index.js";
import pieReducer from "./components/pie/index.js";
import lineBarReducer from "./components/line_bar/index.js";

const chartReducer = (state = {}, action) => {
	
  switch (action.chart) {
    case 'shared':
      return sharedReducer(state, action);
    case 'bar':
      return barReducer(state, action);
    case 'line':
      return lineReducer(state, action);
    case 'pie':
      return pieReducer(state, action);
    case 'line_bar':
      return lineBarReducer(state, action);
    default:
      switch (action.type){
  			case 'set-config':
  				return {...state, spec: action.value}
        case 'set-options':
          return {...state, showOptions: !state.showOptions}
        case 'edit-chart-name':
          return {...state, isChartNameEditable: action.value}
        case 'set-chart-name':
          return {...state, chartName: action.value}
  			default:
  				return state;
		  }
  }
}

export default chartReducer;