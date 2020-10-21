import React from 'react';
import { Input, InputNumber, Select, Checkbox, Form } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import _ from 'lodash';

import {
  SET_AREA_LINES,
  SET_AREA_LINE_WIDTH,
  SET_AREA_LINE_OPACITY,
  SET_AREA_LINE_CURVE,
  SET_AREA_LINE_DASHED,
} from '../../constants/area_lines.js';
const { Option } = Select;

function Lines(props) {
  const { form } = props;
  const [enable, setEnable] = React.useState(false);

  const markObj = props.properties.find((d) => d.prop === 'mark');

  const dispatch = useDispatch();

  const handleEnable = (checked) => {
    setEnable(checked);
    if (!checked) {
      let values = form.getFieldValue();
      _.unset(values, [...markObj.path, 'line']);
      form.setFieldsValue(values);
    }
  };

  return (
    <div className="property-container">
      <Checkbox
        onChange={(e) => {
          const checked = e.target.checked;
          handleEnable(checked);
          dispatch({
            type: SET_AREA_LINES,
            payload: { value: checked, path: markObj.path },
            chart: 'shared',
          });
        }}
      >
        Enable
      </Checkbox>
      {enable ? (
        <React.Fragment>
          <Form.Item name={[...markObj.path, 'line', 'strokeWidth']} initialValue={4} label="Width">
            <InputNumber
              onChange={(value) =>
                dispatch({
                  type: SET_AREA_LINE_WIDTH,
                  payload: { value: value, path: markObj.path },
                  chart: 'shared',
                })
              }
            />
          </Form.Item>

          <Form.Item name={[...markObj.path, 'line', 'opacity']} initialValue={1} label="Opacity">
            <InputNumber
              min={0}
              max={1}
              step={0.05}
              onChange={(value) =>
                dispatch({
                  type: SET_AREA_LINE_OPACITY,
                  payload: { value: value, path: markObj.path },
                  chart: 'shared',
                })
              }
            />
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'line', 'interpolate']}
            initialValue="linear"
            label="Line Curve"
          >
            <Select
              onChange={(value) =>
                dispatch({
                  type: SET_AREA_LINE_CURVE,
                  payload: { value: value, path: markObj.path },
                  chart: 'shared',
                })
              }
            >
              <Option value="linear">Linear</Option>
              <Option value="linear-closed">Linear Closed</Option>
              <Option value="step">Step</Option>
              <Option value="basis">Basis</Option>
              <Option value="monotone">Monotone</Option>
            </Select>
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'line', 'strokeDash']}
            initialValue={0}
            label="Dash Width"
          >
            <InputNumber
              onChange={(value) =>
                dispatch({
                  type: SET_AREA_LINE_DASHED,
                  payload: { value: value, path: markObj.path },
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

export default Lines;
