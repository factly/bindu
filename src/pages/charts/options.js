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
import { useParams } from 'react-router-dom';

function OptionComponent(props) {
  let { id } = useParams();
  id = parseInt(id);
  let Component;
  switch (id) {
    case 0:
      Component = Area;
      break;
    case 1:
      Component = StackedArea;
      break;
    case 2:
      Component = StackedAreaProportional;
      break;
    case 3:
      Component = Bar;
      break;
    case 4:
      Component = HorizontalBar;
      break;
    case 5:
      Component = HorizontalStackBar;
      break;
    case 6:
      Component = StackedBar;
      break;
    case 7:
      Component = GroupedLine;
      break;
    case 8:
      Component = Line;
      break;
    case 9:
      Component = LineProjected;
      break;
    case 10:
      Component = Pie;
      break;
    case 11:
      Component = Donut;
      break;
    case 12:
      Component = LineBar;
      break;
    case 13:
      Component = DivergingBar;
      break;
    case 14:
      Component = GroupedBarProportional;
      break;
    case 15:
      Component = HorizontalGroupedBarProportional;
      break;
    case 16:
      Component = GridPie;
      break;
    case 17:
      Component = GridBar;
      break;
    case 18:
      Component = GridLine;
      break;
    case 19:
      Component = GroupedBar;
      break;
    default:
      return null;
  }

  return <Component {...props} />;
}

export default OptionComponent;
