import React, { useEffect } from 'react';

import { Collapse } from 'antd';
import ChartProperties from '../../../components/shared/chart_properties.js';
import RegionsLayer from '../../../components/shared/regions_layer.js';
import GraticuleLayer from '../../../components/shared/graticule.js';
import ZoomLayer from '../../../components/shared/zoom.js';
import { useDispatch } from 'react-redux';

import Spec from './default.json';
const { Panel } = Collapse;

function IndiaStates() {
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
      name: 'Regions Layer',
      properties: [
        {
          prop: 'stroke-color',
          path: ['marks', 1, 'encode', 'update', 'stroke', 'value'],
        },
        {
          prop: 'stroke-width',
          path: ['marks', 1, 'encode', 'update', 'strokeWidth', 'value'],
        },
        {
          prop: 'stroke-opacity',
          path: ['marks', 1, 'encode', 'update', 'strokeOpacity', 'value'],
        },
      ],
      Component: RegionsLayer,
    },
    {
      name: 'Highlight Regions Layer',
      properties: [
        {
          prop: 'stroke-color',
          path: ['marks', 1, 'encode', 'hover', 'stroke', 'value'],
        },
        {
          prop: 'stroke-width',
          path: ['marks', 1, 'encode', 'hover', 'strokeWidth', 'value'],
        },
        {
          prop: 'color-scheme',
          path: ['scales', 0, 'range', 'scheme'],
        },
      ],
      Component: RegionsLayer,
    },
    {
      name: 'Graticule Layer',
      properties: [
        {
          prop: 'long-sep',
          path: ['data', 2, 'transform', 0, 'step', 0],
        },
        {
          prop: 'lat-sep',
          path: ['data', 2, 'transform', 0, 'step', 1],
        },
        {
          prop: 'color',
          path: ['marks', 0, 'encode', 'enter', 'stroke', 'value'],
        },
        {
          prop: 'width',
          path: ['marks', 0, 'encode', 'enter', 'strokeWidth', 'value'],
        },
        {
          prop: 'opacity',
          path: ['marks', 0, 'encode', 'enter', 'strokeOpacity', 'value'],
        },
        {
          prop: 'dash',
          path: ['marks', 0, 'encode', 'enter', 'strokeDash', 'value'],
        },
      ],
      Component: GraticuleLayer,
    },
    {
      name: 'Zoom',
      properties: [
        {
          prop: 'scale',
          path: ['signals', 2, 'value'],
        },
        {
          prop: 'rotate0',
          path: ['signals', 8, 'value'],
        },
        {
          prop: 'rotate1',
          path: ['projections', 0, 'rotate', 1],
        },
        {
          prop: 'rotate2',
          path: ['projections', 0, 'rotate', 2],
        },
        {
          prop: 'center0',
          path: ['projections', 0, 'center', 0],
        },
        {
          prop: 'center1',
          path: ['signals', 9, 'value'],
        },
      ],
      Component: ZoomLayer,
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

export default IndiaStates;
