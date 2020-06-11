const templateReducer = (state = {}, action) => {
	
  switch (action.type) {
    case 'set-chart':
    	return {...state, selectedOption: action.index};
    default:
      return state
  }
}

export default templateReducer;