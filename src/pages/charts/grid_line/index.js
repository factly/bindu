import ChartProperties from '../../../components/shared/chart_properties.js';
import Colors from '../../../components/shared/colors.js';
import XAxis from '../../../components/shared/x_axis.js';
import YAxis from '../../../components/shared/y_axis.js';

import Lines from '../../../components/shared/lines.js';
import Dots from '../../../components/shared/dots.js';
import Facet from '../../../components/shared/facet.js';

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
    Component: 'ChartProperties',
  },
  {
    name: 'Colors',
    properties: [
      {
        prop: 'color',
        type: 'array',
        path: ['encoding', 'color', 'scale', 'range'],
      },
    ],
    Component: 'Colors',
  },
  {
    name: 'Grid',
    properties: [
      {
        prop: 'column',
        path: ['encoding', 'facet', 'columns'],
      },
      {
        prop: 'spacing',
        path: ['encoding', 'facet', 'spacing'],
      },
      {
        prop: 'xaxis',
        path: ['resolve', 'axis', 'x'],
      },
      {
        prop: 'yaxis',
        path: ['resolve', 'axis', 'y'],
      },
    ],
    Component: 'Facet',
  },
  {
    name: 'Lines',
    properties: [
      {
        prop: 'strokeWidth',
        path: ['mark', 'strokeWidth'],
      },
      {
        prop: 'opacity',
        path: ['mark', 'opacity'],
      },
      {
        prop: 'interpolate',
        path: ['mark', 'interpolate'],
      },
      {
        prop: 'strokeDash',
        path: ['mark', 'strokeDash'],
      },
    ],
    Component: 'Lines',
  },
  {
    name: 'Dots',
    properties: [
      {
        prop: 'mark',
        path: ['mark'],
      },
    ],
    Component: 'Dots',
  },
  {
    name: 'X Axis',
    properties: [
      {
        prop: 'title',
        path: ['encoding', 'x', 'axis', 'title'],
      },
      {
        prop: 'orient',
        path: ['encoding', 'x', 'axis', 'orient'],
      },
      {
        prop: 'format',
        path: ['encoding', 'x', 'axis', 'format'],
      },
      {
        prop: 'label_color',
        path: ['encoding', 'x', 'axis', 'labelColor'],
      },
      {
        prop: 'aggregate',
        path: ['encoding', 'x', 'aggregate'],
      },
      {
        prop: 'field',
        path: ['encoding', 'x', 'field'],
      },
      {
        prop: 'sort',
        path: ['encoding', 'x', 'sort'],
      },
      {
        prop: 'type',
        path: ['encoding', 'x', 'type'],
      },
    ],
    Component: 'XAxis',
  },
  {
    name: 'Y Axis',
    properties: [
      {
        prop: 'title',
        path: ['encoding', 'y', 'axis', 'title'],
      },
      {
        prop: 'orient',
        path: ['encoding', 'y', 'axis', 'orient'],
      },
      {
        prop: 'format',
        path: ['encoding', 'y', 'axis', 'format'],
      },
      {
        prop: 'label_color',
        path: ['encoding', 'y', 'axis', 'labelColor'],
      },
      {
        prop: 'aggregate',
        path: ['encoding', 'y', 'aggregate'],
      },
      {
        prop: 'field',
        path: ['encoding', 'y', 'field'],
      },
      {
        prop: 'sort',
        path: ['encoding', 'y', 'sort'],
      },
      {
        prop: 'type',
        path: ['encoding', 'y', 'type'],
      },
    ],
    Component: 'YAxis',
  },
];
