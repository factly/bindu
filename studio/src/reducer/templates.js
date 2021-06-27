import deepEqual from 'deep-equal';
import {
  ADD_TEMPLATE,
  ADD_TEMPLATES,
  ADD_TEMPLATES_REQUEST,
  SET_TEMPLATES_LOADING,
  RESET_TEMPLATES,
} from '../constants/templates';

const initialState = {
  req: [],
  details: {},
  loading: true,
};

const templateReducer = (state = initialState, action) => {
  switch (action.type) {
    case RESET_TEMPLATES:
      return {
        ...state,
        req: [],
        details: {},
        loading: true,
      };
    case SET_TEMPLATES_LOADING:
      return {
        ...state,
        loading: action.payload,
      };
    case ADD_TEMPLATES_REQUEST:
      return {
        ...state,
        req: state.req
          .filter((value) => !deepEqual(value.query, action.payload.query))
          .concat(action.payload),
      };
    case ADD_TEMPLATES:
      if (action.payload.length === 0) {
        return state;
      }
      return {
        ...state,
        details: {
          ...state.details,
          ...action.payload.reduce((obj, item) => Object.assign(obj, { [item.id]: item }), {}),
        },
      };
    case ADD_TEMPLATE:
      return {
        ...state,
        details: {
          ...state.details,
          [action.payload.id]: action.payload,
        },
      };
    default:
      return state;
  }
};

export default templateReducer;
