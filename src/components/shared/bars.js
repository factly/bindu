import React from 'react';
import { InputNumber, Form } from 'antd';

function Bars(props) {
  const opacityObj = props.properties.find((d) => d.prop === 'opacity');
  const cornerRadiusObj = props.properties.find((d) => d.prop === 'corner_radius');

  return (
    <div className="property-container">
      <Form.Item name={opacityObj.path} label="Bar Opacity">
        <InputNumber min={0} max={1} step={0.05} placeholder="Opacity" />
      </Form.Item>

      <Form.Item name={cornerRadiusObj.path} label="Corner Radius">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="Corner Radius"
        />
      </Form.Item>
    </div>
  );
}

export default Bars;
