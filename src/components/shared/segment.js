import React from 'react';
import { Form, InputNumber } from 'antd';

function Segment(props) {
  // const {outerRadius, innerRadius, cornerRadius, padAngle} = layer.mark;
  const outerRadiusObj = props.properties.find((d) => d.prop === 'outer_radius');
  const innerRadiusObj = props.properties.find((d) => d.prop === 'inner_radius');
  const cornerRadiusObj = props.properties.find((d) => d.prop === 'corner_radius');
  const padAngleObj = props.properties.find((d) => d.prop === 'pad_angle');

  return (
    <div className="property-container">
      <Form.Item name={outerRadiusObj.path} label="Outer Radius">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="title"
        />
      </Form.Item>

      <Form.Item name={innerRadiusObj.path} label="Doughnut hole">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="title"
        />
      </Form.Item>

      <Form.Item name={cornerRadiusObj.path} label="Corner Curve">
        <InputNumber
          min={0}
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="width"
        />
      </Form.Item>

      <Form.Item name={padAngleObj.path} label="Padding Angle">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={0}
          max={0.5}
          step="0.025"
          placeholder="height"
          type="number"
        />
      </Form.Item>
    </div>
  );
}

export default Segment;
