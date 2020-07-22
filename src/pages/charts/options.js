import React from 'react';
import Area from './area/index.js';
import StackedArea from './stacked_area/index.js';
import StackedAreaProportional from './stacked_area_proportional/index.js';
import Bar from './bar/index.js';
import GridBar from './grid_bar/index.js';
import HorizontalBar from './horizontal_bar/index.js';
import HorizontalStackBar from './horizontal_stacked_bar/index.js';
import StackedBar from './stacked_bar/index.js';
import GroupedBar from './grouped_bar/index.js';
import Line from './line/index.js';
import GridLine from './grid_line/index.js';
import LineProjected from './line_projected/index.js';
import Pie from './pie/index.js';
import GridPie from './grid_pie/index.js';
import GroupedLine from './grouped_line/index.js';
import LineBar from './line_bar/index.js';
import DivergingBar from './diverging_bar/index.js';
import Donut from './donut/index.js';
import GroupedBarProportional from './grouped_bar_proportional/index.js';
import HorizontalGroupedBarProportional from './horizontal_grouped_bar_proportional/index.js';

import IndiaStates from './india_states/index.js';
import { useParams } from 'react-router-dom';

function OptionComponent() {
  let { id } = useParams();
  id = parseInt(id);
  switch (id) {
    case 0:
      return <Area />;
    case 1:
      return <StackedArea />;
    case 2:
      return <StackedAreaProportional />;
    case 3:
      return <Bar />;
    case 4:
      return <HorizontalBar />;
    case 5:
      return <HorizontalStackBar />;
    case 6:
      return <StackedBar />;
    case 7:
      return <GroupedLine />;
    case 8:
      return <Line />;
    case 9:
      return <LineProjected />;
    case 10:
      return <Pie />;
    case 11:
      return <Donut />;
    case 12:
      return <LineBar />;
    case 13:
      return <DivergingBar />;
    case 14:
      return <GroupedBarProportional />;
    case 15:
      return <HorizontalGroupedBarProportional />;
    case 16:
      return <GridPie />;
    case 17:
      return <GridBar />;
    case 18:
      return <GridLine />;
    case 19:
      return <GroupedBar />;
    case 20:
      return <IndiaStates />;
    default:
      return null;
  }
}

export default OptionComponent;
