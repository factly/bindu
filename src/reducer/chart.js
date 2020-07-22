import sharedReducer from './components/shared/index.js';

const chartReducer = (state = {}, action) => {
  switch (action.chart) {
    case 'shared':
      return sharedReducer(state, action);
    default:
      switch (action.type) {
        case 'set-config':
          return { ...state, spec: action.value, mode: action.mode || 'vega-lite' };
        case 'set-options':
          return { ...state, showOptions: !state.showOptions };
        case 'edit-chart-name':
          return { ...state, isChartNameEditable: action.value };
        case 'set-chart-name':
          return { ...state, chartName: action.value };
        default:
          return state;
      }
  }
};

export default chartReducer;
