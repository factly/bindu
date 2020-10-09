import React from 'react';
import { InputNumber, Row, Col, Form } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import { SET_BAR_OPACITY, SET_BAR_CORNER_RADIUS } from '../../constants/bars.js';
function Bars(props) {
  const spec = useSelector((state) => state.chart.spec);
  const opacityObj = props.properties.find((d) => d.prop === 'opacity');
  const opacity = getValueFromNestedPath(spec, opacityObj.path);

  const cornerRadiusObj = props.properties.find((d) => d.prop === 'corner_radius');
  const cornerRadius = getValueFromNestedPath(spec, cornerRadiusObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Form.Item name={opacityObj.path} label="Bar Opacity">
        <InputNumber
          min={0}
          max={1}
          step={0.05}
          placeholder="Opacity"
          onChange={(value) =>
            dispatch({
              type: SET_BAR_OPACITY,
              payload: { value, path: opacityObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>

      <Form.Item name={cornerRadiusObj.path} label="Corner Radius">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="Corner Radius"
          onChange={(value) =>
            dispatch({
              type: SET_BAR_CORNER_RADIUS,
              payload: { value, path: cornerRadiusObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>
    </div>
  );
}

export default Bars;
