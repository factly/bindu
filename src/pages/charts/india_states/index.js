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
    Component: 'ChartProperties',
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
    Component: 'RegionsLayer',
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
    Component: 'RegionsLayer',
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
    Component: 'GraticuleLayer',
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
    Component: 'ZoomLayer',
  },
];
