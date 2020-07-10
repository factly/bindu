import produce from "immer";

const lineReducer = (state = {}, action) => {
  switch (action.type) {
    case 'set-line-dots':
      return produce(state, draftState => {

          if(action.value){
            draftState.spec.layer[1].mark.point = {
              shape: "circle",
              size: 40,
              filled: true
            };
          } else {
            delete draftState.spec.layer[1].mark["point"]; 
          }
      });
    case 'set-line-dot-shape':
      return produce(state, draftState => {

          draftState.spec.layer[1].mark.point.shape = action.value;
      });
    case 'set-line-dot-size':
      return produce(state, draftState => {

          draftState.spec.layer[1].mark.point.size = parseInt(action.value);
      });
    case 'set-line-dots-hollow':
      return produce(state, draftState => {
          draftState.spec.layer[1].mark.point.filled = !action.value;
      });
    case 'set-line-witdth':
      return produce(state, draftState => {
          draftState.spec.layer[1].mark.strokeWidth = parseInt(action.value);
      });
    case 'set-line-opacity':
      return produce(state, draftState => {
          draftState.spec.layer[1].mark.opacity = parseFloat(action.value);
      });
    case 'set-line-curve':
      return produce(state, draftState => {
          draftState.spec.layer[1].mark.interpolate = action.value;
      });
    case 'set-line-dashed':
      return produce(state, draftState => {
          draftState.spec.layer[1].mark.strokeDash = parseInt(action.value);
      });
    default:
      return state;
  }
}

export default lineReducer;