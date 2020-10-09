import React from 'react';
import { Input, Form, InputNumber } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_PIE_OUTER_RADIUS,
  SET_PIE_INNER_RADIUS,
  SET_PIE_CORNER_RADIUS,
  SET_PIE_PADDING_ANGLE,
} from '../../constants/segment.js';
function Segment(props) {
  const spec = useSelector((state) => state.chart.spec);
  // const {outerRadius, innerRadius, cornerRadius, padAngle} = layer.mark;

  const outerRadiusObj = props.properties.find((d) => d.prop === 'outer_radius');
  const outerRadius = getValueFromNestedPath(spec, outerRadiusObj.path);

  const innerRadiusObj = props.properties.find((d) => d.prop === 'inner_radius');
  const innerRadius = getValueFromNestedPath(spec, innerRadiusObj.path);

  const cornerRadiusObj = props.properties.find((d) => d.prop === 'corner_radius');
  const cornerRadius = getValueFromNestedPath(spec, cornerRadiusObj.path);

  const padAngleObj = props.properties.find((d) => d.prop === 'pad_angle');
  const padAngle = getValueFromNestedPath(spec, padAngleObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Form.Item name={outerRadiusObj.path} label="Outer Radius">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="title"
          onChange={(value) =>
            dispatch({
              type: SET_PIE_OUTER_RADIUS,
              payload: { value, path: outerRadiusObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>

      <Form.Item name={innerRadiusObj.path} label="Doughnut hole">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="title"
          onChange={(value) =>
            dispatch({
              type: SET_PIE_INNER_RADIUS,
              payload: { value, path: innerRadiusObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>

      <Form.Item name={cornerRadiusObj.path} label="Corner Curve">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="width"
          onChange={(value) =>
            dispatch({
              type: SET_PIE_CORNER_RADIUS,
              payload: { value, path: cornerRadiusObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>

      <Form.Item name={padAngleObj.path} label="Padding Angle">
        <InputNumber
          min={0}
          max={0.5}
          step="0.025"
          placeholder="height"
          type="number"
          onChange={(value) =>
            dispatch({
              type: SET_PIE_PADDING_ANGLE,
              payload: { value: value, path: padAngleObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>
    </div>
  );
}

export default Segment;
