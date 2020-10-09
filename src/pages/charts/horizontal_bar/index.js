import React, { useEffect } from 'react';

import { Collapse, Form } from 'antd';
import ChartProperties from '../../../components/shared/chart_properties.js';
import Colors from '../../../components/shared/colors.js';
import Bars from '../../../components/shared/bars.js';
import XAxis from '../../../components/shared/x_axis.js';
import YAxis from '../../../components/shared/y_axis.js';
import DataLabels from '../../../components/shared/data_labels.js';
import { useDispatch } from 'react-redux';

import Spec from './default.json';
const { Panel } = Collapse;

function GroupedBarChart(props) {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch({ type: 'set-config', value: Spec });
    props.onSpecChange(Spec);
  }, [dispatch]);

  const properties = [
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

  return (
    <Form
      initialValues={Spec}
      onValuesChange={(changedValues, allValues) => props.onSpecChange(allValues)}
    >
      <Collapse className="option-item-collapse">
        {properties.map((d, i) => {
          return (
            <Panel className="option-item-panel" header={d.name} key={i}>
              <d.Component properties={d.properties} {...props} />
            </Panel>
          );
        })}
      </Collapse>
    </Form>
  );
}

export default GroupedBarChart;
