import React, { useEffect } from 'react';

import { Collapse, Form } from 'antd';
import ChartProperties from '../../../components/shared/chart_properties.js';
import Colors from '../../../components/shared/colors.js';
import Legend from '../../../components/shared/legend.js';
import LegendLabel from '../../../components/shared/legend_label.js';

import Segment from '../../../components/shared/segment.js';
import Facet from '../../../components/shared/facet.js';

import { useDispatch } from 'react-redux';

import Spec from './default.json';
const { Panel } = Collapse;

function PieChart(props) {
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
          type: 'array',
          path: ['encoding', 'color', 'scale', 'range'],
        },
      ],
      Component: Colors,
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
      ],
      Component: Facet,
    },
    {
      name: 'Segment',
      properties: [
        {
          prop: 'inner_radius',
          path: ['mark', 'innerRadius'],
        },
        {
          prop: 'corner_radius',
          path: ['mark', 'cornerRadius'],
        },
        {
          prop: 'pad_angle',
          path: ['mark', 'padAngle'],
        },
        {
          prop: 'outer_radius',
          path: ['mark', 'outerRadius'],
        },
      ],
      Component: Segment,
    },
    {
      name: 'Legend',
      properties: [
        {
          prop: 'title',
          path: ['encoding', 'color', 'legend', 'title'],
        },
        {
          prop: 'fill_color',
          path: ['encoding', 'color', 'legend', 'fillColor'],
        },
        {
          prop: 'symbol_type',
          path: ['encoding', 'color', 'legend', 'symbolType'],
        },
        {
          prop: 'symbol_size',
          path: ['encoding', 'color', 'legend', 'symbolSize'],
        },
        {
          prop: 'orient',
          path: ['encoding', 'color', 'legend', 'orient'],
        },
      ],
      Component: Legend,
    },
    {
      name: 'Legend Label',
      properties: [
        {
          prop: 'label_align',
          path: ['encoding', 'color', 'legend', 'labelAlign'],
        },
        {
          prop: 'label_baseline',
          path: ['encoding', 'color', 'legend', 'labelBaseline'],
        },
        {
          prop: 'label_color',
          path: ['encoding', 'color', 'legend', 'labelColor'],
        },
      ],
      Component: LegendLabel,
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

export default PieChart;
