const specReducer = (state = {}, action) => {
	
  switch (action.chart) {
    case 'bar':
      return action.value
    case 'set-width':
    	// barSetWidth();
      return {...state, width: parseInt(action.value)}
    case 'set-height':
      return {...state, height: parseInt(action.value)}
    default:
      return state
  }
}

export default specReducer;