import produce from "immer";
const sharedReducer = (state = {}, action) => {
  switch (action.type) {
  
    /**
     * GENERAL PROPERTIES
     */

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
      const namePath;
      return produce(state, draftState => {
          if (draftState.spec.layer[0].encoding.color.hasOwnProperty("field")){
            draftState.spec.layer[0].encoding.color.scale.range[action.index] = action.value;
          } else {
            draftState.spec.layer[0].encoding.color.value = action.value;
          }
      });
    case 'set-background':
      return produce(state, draftState => {
          draftState.spec.background = action.value;
      });
  
    /**
     * X AXIS PROPERTIES
     */
    case 'set-xaxis-title':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.x.axis.title = action.value;
      });

    case 'set-xaxis-position':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.x.axis.orient = action.value;
      });

    case 'set-xaxis-label-format':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.x.axis.format = action.value;
      });

    case 'set-xaxis-label-color':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.x.axis.labelColor = action.value;
      });
    
    /**
     * Y AXIS PROPERTIES
     */
    case 'set-yaxis-title':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.y.axis.title = action.value;
      });

    case 'set-yaxis-position':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.y.axis.orient = action.value;
      });

    case 'set-yaxis-label-format':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.y.axis.format = action.value;
      });

    case 'set-yaxis-label-color':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.y.axis.labelColor = action.value;
      });
    
    /**
     * LEGEND LABEL AXIS PROPERTIES
     */
    case 'set-legend-label-position':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.labelAlign = action.value;
      });
    case 'set-legend-label-baseline':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.labelBaseline = action.value;
      });
    case 'set-legend-label-color':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.labelColor = action.value;
      });

    /**
     * LEGEND PROPERTIES
     */
    case 'set-legend-position':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.orient = action.value;
      });
    case 'set-legend-background':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.fillColor = action.value;
      });
    case 'set-legend-title':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.title = action.value;
      });
    case 'set-legend-symbol':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.symbolType = action.value;
      });
    case 'set-legend-symbol-size':

      return produce(state, draftState => {
          draftState.spec.layer[0].encoding.color.legend.symbolSize = parseInt(action.value);
      });
    
    /**
     * DATA LABELS PROPERTIES
     */
    case 'set-data-labels':
      return produce(state, draftState => {
          if(action.value){
            const encoding = draftState.spec.layer[0].encoding;
            draftState.spec.layer.push({
              "mark": { "type": "text", "dx": 0, "dy": 10, fontSize: 12 },
              "encoding": {
                "x": encoding.x,
                "y": { ...encoding.y, "stack": "zero" },
                "color": { ...encoding.color, scale: {range: ["#ffffff"]}, "legend": null },
                "text": { ...encoding.y,"format": ".1f" }
              }
            });
          } else {
            draftState.spec.layer.splice(1, 1);  
          }
      });
    case 'set-data-labels-color':

      return produce(state, draftState => {
          draftState.spec.layer[1].encoding.color.scale.range = [action.value];
      });

    case 'set-data-labels-size':

      return produce(state, draftState => {
          draftState.spec.layer[1].mark.fontSize = action.value;
      });

    case 'set-data-labels-format':

      return produce(state, draftState => {
          draftState.spec.layer[1].encoding.text.format = action.value;
      });

    default:
      return state;
  }
}

export default sharedReducer;