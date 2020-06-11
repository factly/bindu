import { mergeDeep } from "../../../utils/index.js";
const sharedReducer = (state = {}, action) => {
	const spec = mergeDeep({}, state.spec);
  switch (action.type) {
    case 'set-width':
      spec.width = parseInt(action.value);
      return {...state, spec: spec }
    case 'set-height':
      spec.height = parseInt(action.value);
      return {...state, spec: spec}
    default:
      return state;
  }
}

export default sharedReducer;