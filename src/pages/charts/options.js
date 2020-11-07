import React, { useEffect } from 'react';
import { Collapse } from 'antd';
import { useParams } from 'react-router-dom';

import Categories from '../../components/categories';
import Tags from '../../components/tags';
import * as Area from './area/index.js';
import * as StackedArea from './stacked_area/index.js';
import * as StackedAreaProportional from './stacked_area_proportional/index.js';
import * as Bar from './bar/index.js';
import * as GridBar from './grid_bar/index.js';
import * as HorizontalBar from './horizontal_bar/index.js';
import * as HorizontalStackBar from './horizontal_stacked_bar/index.js';
import * as StackedBar from './stacked_bar/index.js';
import * as GroupedBar from './grouped_bar/index.js';
import * as Line from './line/index.js';
import * as GridLine from './grid_line/index.js';
import * as LineProjected from './line_projected/index.js';
import * as Pie from './pie/index.js';
import * as GridPie from './grid_pie/index.js';
import * as GroupedLine from './grouped_line/index.js';
import * as LineBar from './line_bar/index.js';
import * as DivergingBar from './diverging_bar/index.js';
import * as Donut from './donut/index.js';
import * as GroupedBarProportional from './grouped_bar_proportional/index.js';
import * as HorizontalGroupedBarProportional from './horizontal_grouped_bar_proportional/index.js';

const { Panel } = Collapse;

function OptionComponent(props) {
  const { form } = props;

  let { id } = useParams();
  id = parseInt(id);
  let component;

  useEffect(() => {
    form.setFieldsValue(component.spec);
  }, [id, component]);

  switch (id) {
    case 0:
      component = Area;
      break;
    case 1:
      component = StackedArea;
      break;
    case 2:
      component = StackedAreaProportional;
      break;
    case 3:
      component = Bar;
      break;
    case 4:
      component = HorizontalBar;
      break;
    case 5:
      component = HorizontalStackBar;
      break;
    case 6:
      component = StackedBar;
      break;
    case 7:
      component = GroupedLine;
      break;
    case 8:
      component = Line;
      break;
    case 9:
      component = LineProjected;
      break;
    case 10:
      component = Pie;
      break;
    case 11:
      component = Donut;
      break;
    case 12:
      component = LineBar;
      break;
    case 13:
      component = DivergingBar;
      break;
    case 14:
      component = GroupedBarProportional;
      break;
    case 15:
      component = HorizontalGroupedBarProportional;
      break;
    case 16:
      component = GridPie;
      break;
    case 17:
      component = GridBar;
      break;
    case 18:
      component = GridLine;
      break;
    case 19:
      component = GroupedBar;
      break;
    default:
      return null;
  }

  return (
    <Collapse className="option-item-collapse">
      <Panel className="option-item-panel" header={'Tags'} key={'tags'}>
        <Tags form={form} />
      </Panel>
      <Panel className="option-item-panel" header={'Categories'} key={'categories'}>
        <Categories form={form} />
      </Panel>
      {component.properties.map((d, i) => {
        return (
          <Panel className="option-item-panel" header={d.name} key={i}>
            <d.Component properties={d.properties} form={form} />
          </Panel>
        );
      })}
    </Collapse>
  );
}

export default OptionComponent;
