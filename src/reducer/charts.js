import {
  ADD_CHART,
  ADD_CHARTS,
  ADD_CHARTS_REQUEST,
  SET_CHARTS_LOADING,
  RESET_CHARTS,
} from '../constants/charts';
import deepEqual from 'deep-equal';

const initialState = {
  req: [],
  details: {},
  loading: true,
};

export default function chartsReducer(state = initialState, action = {}) {
  switch (action.type) {
    case RESET_CHARTS:
      return {
        ...state,
        req: [],
        details: {},
        loading: true,
      };
    case SET_CHARTS_LOADING:
      return {
        ...state,
        loading: action.payload,
      };
    case ADD_CHARTS_REQUEST:
      return {
        ...state,
        req: state.req
          .filter((value) => !deepEqual(value.query, action.payload.query))
          .concat(action.payload),
      };
    case ADD_CHARTS:
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
    case ADD_CHART:
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
}
