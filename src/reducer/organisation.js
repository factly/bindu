import {
  SET_SELECTED_ORGANISATION,
  GET_ORGANISATIONS_SUCCESS,
  LOADING_ORGANISATIONS,
} from '../constants/organisation';

const initialState = {
  details: {},
  loading: true,
  selected: 0,
};

export default function organisationsReducer(state = initialState, action = {}) {
  if (!action.payload) {
    return state;
  }
  switch (action.type) {
    case LOADING_ORGANISATIONS:
      return {
        ...state,
        loading: true,
      };
    case GET_ORGANISATIONS_SUCCESS:
      let organisation_details = {};

      action.payload.forEach((element) => {
        organisation_details[element.id] = element;
      });

      const defaultOrganisation =
        Object.keys(organisation_details).length > 0
          ? organisation_details[Object.keys(organisation_details)[0]].id
          : 0;

      return {
        ...state,
        details: organisation_details,
        loading: false,
        selected: organisation_details[state.selected] ? state.selected : defaultOrganisation,
      };
    case SET_SELECTED_ORGANISATION:
      return {
        ...state,
        selected: action.payload,
      };
    default:
      return state;
  }
}
