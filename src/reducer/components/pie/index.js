import produce from "immer";

const pieReducer = (state = {}, action) => {
  switch (action.type) {
    case 'set-data-labels':
      return produce(state, draftState => {
          if(action.value){
            const layer = draftState.spec.layer[0];
            const {encoding, mark} = layer;
            draftState.spec.layer.push({
              "mark": { "type": "text", fontSize: 12, "radius": mark.outerRadius + 10},
              "encoding": {
                "theta": {...encoding.theta, "stack": true},
                "text": { "field": encoding.color.field, "type": encoding.color.type },
                "color": { ...encoding.color, "legend": null },
              }
            });
          } else {
            draftState.spec.layer.splice(1, 1);  
          }
      });
    case 'set-data-labels-size':
      return produce(state, draftState => {
        draftState.spec.layer[1].mark.fontSize = parseInt(action.value);
      });
    case 'set-outer-radius':
      return produce(state, draftState => {
        draftState.spec.layer[0].mark.outerRadius = parseInt(action.value);
      });
    case 'set-inner-radius':
      return produce(state, draftState => {
        draftState.spec.layer[0].mark.innerRadius = parseInt(action.value);
      });
    case 'set-corner-radius':
      return produce(state, draftState => {
        draftState.spec.layer[0].mark.cornerRadius = parseInt(action.value);
      });
    case 'set-padding-angle':
      return produce(state, draftState => {
        draftState.spec.layer[0].mark.padAngle = parseFloat(action.value);
      });

    case 'set-data-labels-position':
      return produce(state, draftState => {
        draftState.spec.layer[1].mark.radius = parseFloat(action.value);
      });
    default:
      return state;
  }
}

export default pieReducer;