import React from 'react';
import { Input, Select, Form } from 'antd';

const { Option } = Select;

function Lines(props) {
  // const { strokeWidth, opacity, interpolate, strokeDash} = spec.layer[0].mark;
  const strokeWidthObj = props.properties.find((d) => d.prop === 'strokeWidth');
  const opacityObj = props.properties.find((d) => d.prop === 'opacity');
  const interpolateObj = props.properties.find((d) => d.prop === 'interpolate');
  const strokeDashObj = props.properties.find((d) => d.prop === 'strokeDash');

  return (
    <div className="property-container">
      <Form.Item name={strokeWidthObj.path} label="Width">
        <Input min={0} type="number" />
      </Form.Item>

      <Form.Item name={opacityObj.path} label="Opacity">
        <Input min={0} max={1} step={0.05} type="number" />
      </Form.Item>

      <Form.Item name={interpolateObj.path} label="Line Curve">
        <Select>
          <Option value="linear">Linear</Option>
          <Option value="linear-closed">Linear Closed</Option>
          <Option value="step">Step</Option>
          <Option value="basis">Basis</Option>
          <Option value="monotone">Monotone</Option>
        </Select>
      </Form.Item>

      <Form.Item name={strokeDashObj.path} label="Dash Width">
        <Input type="number" />
      </Form.Item>
    </div>
  );
}

export default Lines;
