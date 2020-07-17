import React from 'react';
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_LINE_WIDTH,
  SET_LINE_OPACITY,
  SET_LINE_CURVE,
  SET_LINE_DASHED,
} from '../../constants/lines.js';
const { Option } = Select;

function Lines(props) {
  const spec = useSelector((state) => state.chart.spec);
  // const { strokeWidth, opacity, interpolate, strokeDash} = spec.layer[0].mark;
  const strokeWidthObj = props.properties.find((d) => d.prop === 'strokeWidth');
  const opacityObj = props.properties.find((d) => d.prop === 'opacity');
  const interpolateObj = props.properties.find((d) => d.prop === 'interpolate');
  const strokeDashObj = props.properties.find((d) => d.prop === 'strokeDash');

  const strokeWidth = getValueFromNestedPath(spec, strokeWidthObj.path);
  const opacity = getValueFromNestedPath(spec, opacityObj.path);
  const interpolate = getValueFromNestedPath(spec, interpolateObj.path);
  const strokeDash = getValueFromNestedPath(spec, strokeDashObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Width</label>
        </Col>
        <Col span={12}>
          <Input
            value={strokeWidth}
            type="number"
            onChange={(e) =>
              dispatch({
                type: SET_LINE_WIDTH,
                payload: { value: e.target.value, path: strokeWidthObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Opacity</label>
        </Col>
        <Col span={12}>
          <Input
            value={opacity}
            min={0}
            max={1}
            step={0.05}
            type="number"
            onChange={(e) =>
              dispatch({
                type: SET_LINE_OPACITY,
                payload: { value: e.target.value, path: opacityObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Line Curve</label>
        </Col>
        <Col span={12}>
          <Select
            value={interpolate}
            onChange={(value) =>
              dispatch({
                type: SET_LINE_CURVE,
                payload: { value: value, path: interpolateObj.path },
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
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Dash Width</label>
        </Col>
        <Col span={12}>
          <Input
            value={strokeDash}
            type="number"
            onChange={(e) =>
              dispatch({
                type: SET_LINE_DASHED,
                payload: { value: e.target.value, path: strokeDashObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
    </div>
  );
}

export default Lines;
