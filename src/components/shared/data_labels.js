import React from 'react';
import { Input, InputNumber, Form, Checkbox } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_DATA_LABELS,
  SET_DATA_LABELS_COLOR,
  SET_DATA_LABELS_SIZE,
  SET_DATA_LABELS_FORMAT,
} from '../../constants/data_labels.js';
import { SET_LEGEND_LABEL_BASELINE } from '../../constants/legend_label.js';

function DataLabels(props) {
  const [enable, setEnable] = React.useState(false);

  const spec = useSelector((state) => state.chart.spec);
  const layersLength = spec.layer.length;

  const dispatch = useDispatch();

  const colorObj = props.properties.find((d) => d.prop === 'color');
  const fontSizeObj = props.properties.find((d) => d.prop === 'font_size');
  const formatObj = props.properties.find((d) => d.prop === 'format');

  let color;
  let fontSize;
  let format;
  if (layersLength > 1 && colorObj.path) {
    color = getValueFromNestedPath(spec, colorObj.path);
  }
  if (layersLength > 1 && fontSizeObj.path) {
    fontSize = getValueFromNestedPath(spec, fontSizeObj.path);
  }
  if (layersLength > 1 && formatObj.path) {
    format = getValueFromNestedPath(spec, formatObj.path);
  }

  const handleEnable = (checked) => {
    if (checked) {
      props.addLayer(({ encoding }) => ({
        mark: { type: 'text', dx: 0, dy: 10, fontSize: 12 },
        encoding: {
          x: encoding.x,
          y: { ...encoding.y, stack: 'zero' },
          color: { value: '#ffffff' },
          text: { ...encoding.y, format: '.1f' },
        },
      }));
    } else {
      props.removeLayer();
    }
    setEnable(checked);
  };

  return (
    <div className="property-container">
      <Checkbox
        onChange={(e) => {
          const checked = e.target.checked;
          handleEnable(checked);
          dispatch({ type: SET_DATA_LABELS, value: e.target.checked, chart: 'shared' });
        }}
      >
        Enable
      </Checkbox>
      {enable ? (
        <React.Fragment>
          <Form.Item name={colorObj.path} label="Color">
            <Input
              type="color"
              value={color}
              onChange={(e) =>
                dispatch({
                  type: SET_DATA_LABELS_COLOR,
                  payload: { value: e.target.value, path: colorObj.path },
                  chart: 'shared',
                })
              }
            />
          </Form.Item>

          <Form.Item name={fontSizeObj.path} label="Size">
            <InputNumber
              value={fontSize}
              onChange={(value) =>
                dispatch({
                  type: SET_DATA_LABELS_SIZE,
                  payload: { value: value, path: fontSizeObj.path },
                  chart: 'shared',
                })
              }
            />
          </Form.Item>

          <Form.Item name={formatObj.path} label="Format">
            <Input
              type="text"
              value={format}
              onChange={(e) =>
                dispatch({
                  type: SET_DATA_LABELS_FORMAT,
                  payload: { value: e.target.value, path: formatObj.path },
                  chart: 'shared',
                })
              }
            />
          </Form.Item>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default DataLabels;
