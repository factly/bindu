import React, { useEffect } from 'react';

import { Collapse } from 'antd';
import ChartProperties from '../../../components/shared/chart_properties.js';
import TreeMap from '../../../components/shared/tree_map.js';
import { useDispatch } from 'react-redux';

import Spec from './default.json';
const { Panel } = Collapse;

function GroupedBarChart() {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch({ type: 'set-config', value: Spec, mode: 'vega' });
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
      name: 'Tree Map',
      properties: [
        {
          prop: 'layout',
          path: ['data', 0, 'transform', 1, 'method'],
        },
        {
          prop: 'aspect_ratio',
          path: ['data', 0, 'transform', 1, 'ratio'],
        },
      ],
      Component: TreeMap,
    },
  ];

  return (
    <div className="options-container">
      <Collapse className="option-item-collapse">
        {properties.map((d, i) => {
          return (
            <Panel className="option-item-panel" header={d.name} key={i}>
              <d.Component properties={d.properties} />
            </Panel>
          );
        })}
      </Collapse>
    </div>
  );
}

export default GroupedBarChart;
