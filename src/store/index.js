import { createStore } from 'redux';
import rootReducer from '../reducer/index.js';

const defaultStore = {
  chart: {
    spec: {},
    showOptions: true,
    chartName: 'Untitled',
  },
  templates: {
    options: [
      {
        id: 0,
        name: 'Area Chart',
        icon: 'images/area.vl.png',
      },
      {
        id: 1,
        name: 'Stacked Area Chart',
        icon: 'images/area_density_stacked.vl.png',
      },
      {
        id: 2,
        name: 'Stacked Area Chart(Proportional)',
        icon: 'images/area_density_stacked_fold.vl.png',
      },
      {
        id: 3,
        name: 'Bar Chart',
        icon: 'images/bar.vl.png',
      },
      {
        id: 4,
        name: 'Horizontal Bar Chart',
        icon: 'images/bar_aggregate.vl.png',
      },
      {
        id: 5,
        name: 'Horizontal Stack Bar Chart',
        icon: 'images/area.vl.png',
      },
      {
        id: 6,
        name: 'Grouped Bar Chart',
        icon: 'images/facet_grid_bar.vl.png',
      },
      {
        id: 7,
        name: 'Grouped Line Chart',
        icon: 'images/line_color_label.vl.png',
      },
      {
        id: 8,
        name: 'Line Chart',
        icon: 'images/line.vl.png',
      },
      {
        id: 9,
        name: 'Line Chart (Projected)',
        icon: 'images/line_dashed_part.vl.png',
      },
      {
        id: 10,
        name: 'Pie Chart',
        icon: 'images/arc_pie.vl.png',
      },
      {
        id: 11,
        name: 'Donut Chart',
        icon: 'images/arc_donut.vl.png',
      },
      {
        id: 12,
        name: 'Line + Bar Chart',
        icon: 'images/layer_bar_line.vl.png',
      },
      {
        id: 13,
        name: 'Diverging Bar Chart',
        icon: 'images/bar_diverging_stack_population_pyramid.vl.png',
      },
      {
        id: 14,
        name: 'Grouped Bar Chart(Proportional)',
        icon: 'images/facet_grid_bar.vl.png',
      },
      {
        id: 15,
        name: 'Horizontal Grouped Bar Chart(Proportional)',
        icon: 'images/facet_grid_bar.vl.png',
      },
      {
        id: 16,
        name: 'Grid of Pie Chart',
        icon: 'images/area.vl.png',
      },
      {
        id: 17,
        name: 'Grid of Bar Chart',
        icon: 'images/area.vl.png',
      },
      {
        id: 18,
        name: 'Grid of Line Chart',
        icon: 'images/area.vl.png',
      },
      {
        id: 19,
        name: 'Grouped Bar',
        icon: 'images/facet_grid_bar.vl.png',
      },
      {
        id: 20,
        name: 'India States',
        icon: 'images/india_states.png',
      },
      {
        id: 21,
        name: 'Tree Map',
        icon: 'images/treemap.vg.png',
      },
      {
        id: 22,
        name: 'Bar Chart Race',
        icon: 'images/bar_aggregate.vl.png',
      },
    ],
  },
};

const store = createStore(rootReducer, defaultStore);
window.store = store;

export default store;
