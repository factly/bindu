import deepEqual from 'deep-equal';
import {
  ADD_TEMPLATE,
  ADD_TEMPLATES,
  ADD_TEMPLATES_REQUEST,
  SET_TEMPLATES_LOADING,
  RESET_TEMPLATES,
} from '../constants/templates';
import * as donut from '../pages/charts/donut/index';
import * as diverging_bar from '../pages/charts/diverging_bar';
import * as area from '../pages/charts/area/index';
import * as bar from '../pages/charts/bar/index';
import * as grid_bar from '../pages/charts/grid_bar/index';
import * as grid_line from '../pages/charts/grid_line/index';
import * as grid_pie from '../pages/charts/grid_pie/index';
import * as grouped_bar from '../pages/charts/grouped_bar/index';
import * as grouped_bar_proportional from '../pages/charts/grouped_bar_proportional/index';
const chart1 = {
  created_at: '2021-03-18T18:25:27.27127Z',
  created_by_id: 1,
  deleted_at: null,
  id: 1,
  medium: null,
  medium_id: null,
  properties: grid_bar.properties,
  spec: grid_bar.spec,
  slug: 'grid_bar-chart',
  space_id: 1,
  title: 'grid_bar Chart Local',
  updated_at: '2021-03-18T18:25:27.27127Z',
  updated_by_id: 1,
};
const chart2 = {
  created_at: '2021-03-18T18:25:27.27127Z',
  created_by_id: 1,
  deleted_at: null,
  id: 2,
  medium: null,
  medium_id: null,
  properties: grid_line.properties,
  spec: grid_line.spec,
  slug: 'grid_line-chart',
  space_id: 1,
  title: 'grid_line Chart Local',
  updated_at: '2021-03-18T18:25:27.27127Z',
  updated_by_id: 1,
};
const chart3 = {
  created_at: '2021-03-18T18:25:27.27127Z',
  created_by_id: 1,
  deleted_at: null,
  id: 3,
  medium: null,
  medium_id: null,
  properties: grid_pie.properties,
  spec: grid_pie.spec,
  slug: 'grid_pie-chart',
  space_id: 1,
  title: 'grid_pie Chart Local',
  updated_at: '2021-03-18T18:25:27.27127Z',
  updated_by_id: 1,
};
const chart4 = {
  created_at: '2021-03-18T18:25:27.27127Z',
  created_by_id: 1,
  deleted_at: null,
  id: 4,
  medium: null,
  medium_id: null,
  properties: grouped_bar.properties,
  spec: grouped_bar.spec,
  slug: 'grouped_bar-chart',
  space_id: 1,
  title: 'grouped_bar Chart Local',
  updated_at: '2021-03-18T18:25:27.27127Z',
  updated_by_id: 1,
};
const chart5 = {
  created_at: '2021-03-18T18:25:27.27127Z',
  created_by_id: 1,
  deleted_at: null,
  id: 5,
  medium: null,
  medium_id: null,
  properties: grouped_bar_proportional.properties,
  spec: grouped_bar_proportional.spec,
  slug: 'grouped_bar_proportional-chart',
  space_id: 1,
  title: 'grouped_bar_proportional Chart Local',
  updated_at: '2021-03-18T18:25:27.27127Z',
  updated_by_id: 1,
};
const initialState = {
  req: [],
  details: {
    1: chart1,
    2: chart2,
    3: chart3,
    4: chart4,
    5: chart5,
  },
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
