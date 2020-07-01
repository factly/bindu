import produce from "immer";

const barReducer = (state = {}, action) => {
  switch (action.type) {
    case 'set-opacity':
      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.opacity.value = parseFloat(action.value);
      });

    case 'set-corner-radius':
      return produce(state, draftState => {
          draftState.spec.layer[0].mark.cornerRadius = parseInt(action.value);
      });
    default:
      return state;
  }
}

export default barReducer;