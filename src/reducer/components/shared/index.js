import produce from "immer";
const sharedReducer = (state = {}, action) => {
  switch (action.type) {
    case 'set-width':
      return produce(state, draftState => {
          draftState.spec.width = parseFloat(action.value);
      });
    case 'set-height':
      return produce(state, draftState => {
          draftState.spec.height = parseFloat(action.value);
      });
    case 'set-title':
      return produce(state, draftState => {
          draftState.spec.title = action.value;
      });
    case 'set-color':
      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.scale.range[action.index] = action.value;
      });
    case 'set-background':
      return produce(state, draftState => {
          draftState.spec.background = action.value;
      });
    default:
      return state;
  }
}

export default sharedReducer;