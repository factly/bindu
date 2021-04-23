import React from 'react';
import { Input, Select, Form } from 'antd';
import { getInterpolateOptions } from './area_lines';

function Lines(props) {
  const [interpolateOptions, setInterpolateOptions] = React.useState([]);

  const strokeWidthObj = props.properties.find((d) => d.prop === 'strokeWidth');
  const opacityObj = props.properties.find((d) => d.prop === 'opacity');
  const interpolateObj = props.properties.find((d) => d.prop === 'interpolate');
  const strokeDashObj = props.properties.find((d) => d.prop === 'strokeDash');

  React.useEffect(() => {
    getInterpolateOptions(props.form, setInterpolateOptions);
  }, []);

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
          {interpolateOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item name={strokeDashObj.path} label="Dash Width">
        <Input type="number" />
      </Form.Item>
    </div>
  );
}

export default Lines;
