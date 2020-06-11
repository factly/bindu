import { mergeDeep } from "../../../utils/index.js";
const barReducer = (state = {}, action) => {
	const spec = mergeDeep({}, state.spec);
  switch (action.type) {
    case 'set-opacity':
      spec.layer[0].encoding.opacity.value = parseFloat(action.value);
      return {...state, spec: spec }
    case 'set-corner-radius':
      spec.layer[0].mark.cornerRadius = parseInt(action.value);
      return {...state, spec: spec }
    case 'set-xaxis-title':
      spec.layer[0].encoding.x.axis.title = action.value;
      return {...state, spec: spec }
    case 'set-xaxis-position':
      spec.layer[0].encoding.x.axis.orient = action.value;
      return {...state, spec: spec }
    case 'set-xaxis-label-format':
      spec.layer[0].encoding.x.axis.format = action.value;
      return {...state, spec: spec }
    case 'set-xaxis-label-color':
      spec.layer[0].encoding.x.axis.labelColor = action.value;
      return {...state, spec: spec }
    case 'set-yaxis-title':
      spec.layer[0].encoding.y.axis.title = action.value;
      return {...state, spec: spec }
    case 'set-yaxis-position':
      spec.layer[0].encoding.y.axis.orient = action.value;
      return {...state, spec: spec }
    case 'set-yaxis-label-format':
      spec.layer[0].encoding.y.axis.format = action.value;
      return {...state, spec: spec }
    case 'set-yaxis-label-color':
      spec.layer[0].encoding.y.axis.labelColor = action.value;
      return {...state, spec: spec }
    default:
      return state;
  }
}

export default barReducer;