import ChartProperties from '../../../components/shared/chart_properties.js';
import Colors from '../../../components/shared/colors.js';
import Bars from '../../../components/shared/bars.js';
import Legend from '../../../components/shared/legend.js';
import LegendLabel from '../../../components/shared/legend_label.js';
import XAxis from '../../../components/shared/x_axis.js';
import YAxis from '../../../components/shared/y_axis.js';
import DataLabels from '../../../components/shared/data_labels.js';

import Spec from './default.json';

export const spec = Spec;

export const properties = [
  {
    name: 'Chart Properties',
    properties: [
      {
        prop: 'title',
        path: ['title'],
      },
      {
        prop: 'width',
        path: ['width'],
      },
      {
        prop: 'height',
        path: ['height'],
      },
      {
        prop: 'background',
        path: ['background'],
      },
    ],
    Component: ChartProperties,
  },
  {
    name: 'Colors',
    properties: [
      {
        prop: 'color',
        type: 'array',
        path: ['layer', 0, 'encoding', 'color', 'scale', 'range'],
      },
    ],
    Component: Colors,
  },
  {
    name: 'Bars',
    Component: Bars,
  },
  {
    name: 'X Axis',
    Component: XAxis,
  },
  {
    name: 'Y Axis',
    Component: YAxis,
  },
  {
    name: 'Legend',
    Component: Legend,
  },
  {
    name: 'Legend Label',
    Component: LegendLabel,
  },
  {
    name: 'Data Labels',
    Component: DataLabels,
  },
];
