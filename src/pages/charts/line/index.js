import ChartProperties from '../../../components/shared/chart_properties.js';
import Colors from '../../../components/shared/colors.js';
import Legend from '../../../components/shared/legend.js';
import LegendLabel from '../../../components/shared/legend_label.js';
import XAxis from '../../../components/shared/x_axis.js';
import YAxis from '../../../components/shared/y_axis.js';

import Lines from '../../../components/shared/lines.js';
import Dots from '../../../components/shared/dots.js';

import DataLabels from '../../../components/shared/data_labels.js';

import Spec from './default.json';

export const spec = Spec;

export const properties = [
  {
    name: 'Chart Properties',
    component: ChartProperties,
  },
  {
    name: 'Colors',
    component: Colors,
  },
  {
    name: 'Lines',
    component: Lines,
  },
  {
    name: 'Dots',
    component: Dots,
  },
  {
    name: 'Data Labels',
    component: DataLabels,
  },
  {
    name: 'X Axis',
    component: XAxis,
  },
  {
    name: 'Y Axis',
    component: YAxis,
  },
  {
    name: 'Legend',
    component: Legend,
  },
  {
    name: 'Legend Label',
    component: LegendLabel,
  },
  {
    name: 'Data Labels',
    component: DataLabels,
  },
];
