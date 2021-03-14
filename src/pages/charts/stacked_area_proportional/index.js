import ChartProperties from '../../../components/shared/chart_properties.js';
import Colors from '../../../components/shared/colors.js';
import Legend from '../../../components/shared/legend.js';
import LegendLabel from '../../../components/shared/legend_label.js';
import XAxis from '../../../components/shared/x_axis.js';
import YAxis from '../../../components/shared/y_axis.js';
import DataLabels from '../../../components/shared/data_labels.js';

import Dots from '../../../components/shared/dots.js';
import Line from '../../../components/shared/area_lines.js';

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
    name: 'Line',
    properties: [
      {
        prop: 'mark',
        path: ['layer', 0, 'mark'],
      },
    ],
    Component: Line,
  },
  {
    name: 'Dots',
    properties: [
      {
        prop: 'mark',
        path: ['layer', 0, 'mark'],
      },
    ],
    Component: Dots,
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
    name: 'Legend',
    properties: [
      {
        prop: 'title',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'title'],
      },
      {
        prop: 'fill_color',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'fillColor'],
      },
      {
        prop: 'symbol_type',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'symbolType'],
      },
      {
        prop: 'symbol_size',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'symbolSize'],
      },
      {
        prop: 'orient',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'orient'],
      },
    ],
    Component: Legend,
  },
  {
    name: 'Legend Label',
    properties: [
      {
        prop: 'label_align',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'labelAlign'],
      },
      {
        prop: 'label_baseline',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'labelBaseline'],
      },
      {
        prop: 'label_color',
        path: ['layer', 0, 'encoding', 'color', 'legend', 'labelColor'],
      },
    ],
    Component: LegendLabel,
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
