import ChartProperties from '../../../components/shared/chart_properties.js';
import Colors from '../../../components/shared/colors.js';
import Bars from '../../../components/shared/bars.js';
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
        type: 'string',
        path: ['layer', 0, 'encoding', 'color', 'value'],
      },
    ],
    Component: Colors,
  },
  {
    name: 'Bars',
    properties: [
      {
        prop: 'opacity',
        path: ['layer', 0, 'encoding', 'opacity', 'value'],
      },
      {
        prop: 'corner_radius',
        path: ['layer', 0, 'mark', 'cornerRadius'],
      },
    ],
    Component: Bars,
  },
  {
    name: 'X Axis',
    properties: [
      {
        prop: 'title',
        path: ['layer', 0, 'encoding', 'x', 'axis', 'title'],
      },
      {
        prop: 'orient',
        path: ['layer', 0, 'encoding', 'x', 'axis', 'orient'],
      },
      {
        prop: 'format',
        path: ['layer', 0, 'encoding', 'x', 'axis', 'format'],
      },
      {
        prop: 'label_color',
        path: ['layer', 0, 'encoding', 'x', 'axis', 'labelColor'],
      },
      {
        prop: 'aggregate',
        path: ['layer', 0, 'encoding', 'x', 'aggregate'],
      },
      {
        prop: 'field',
        path: ['layer', 0, 'encoding', 'x', 'field'],
      },
    ],
    Component: XAxis,
  },
  {
    name: 'Y Axis',
    properties: [
      {
        prop: 'title',
        path: ['layer', 0, 'encoding', 'y', 'axis', 'title'],
      },
      {
        prop: 'orient',
        path: ['layer', 0, 'encoding', 'y', 'axis', 'orient'],
      },
      {
        prop: 'format',
        path: ['layer', 0, 'encoding', 'y', 'axis', 'format'],
      },
      {
        prop: 'label_color',
        path: ['layer', 0, 'encoding', 'y', 'axis', 'labelColor'],
      },
      {
        prop: 'aggregate',
        path: ['layer', 0, 'encoding', 'y', 'aggregate'],
      },
      {
        prop: 'field',
        path: ['layer', 0, 'encoding', 'y', 'field'],
      },
    ],
    Component: YAxis,
  },
  {
    name: 'Data Labels',
    properties: [
      {
        prop: 'color',
        path: ['layer', 1, 'encoding', 'color', 'value'],
      },
      {
        prop: 'font_size',
        path: ['layer', 1, 'mark', 'fontSize'],
      },
      {
        prop: 'format',
        path: ['layer', 1, 'encoding', 'text', 'format'],
      },
    ],
    Component: DataLabels,
  },
];
