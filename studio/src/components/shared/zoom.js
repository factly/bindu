import React from 'react';
import { InputNumber, Form } from 'antd';

function Zoom(props) {
  const scaleObj = props.properties.find((d) => d.prop === 'scale');
  const rotate0Obj = props.properties.find((d) => d.prop === 'rotate0');
  const rotate1Obj = props.properties.find((d) => d.prop === 'rotate1');
  const rotate2Obj = props.properties.find((d) => d.prop === 'rotate2');
  const center0Obj = props.properties.find((d) => d.prop === 'center0');
  const center1Obj = props.properties.find((d) => d.prop === 'center1');

  return (
    <div className="property-container">
      <Form.Item name={scaleObj.path} label="Scale">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={50}
          max={2000}
          placeholder="scale"
        />
      </Form.Item>

      <Form.Item name={rotate0Obj.path} label="Rotate 0">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={-180}
          max={180}
          placeholder="rotate0"
        />
      </Form.Item>

      <Form.Item name={rotate1Obj.path} label="Rotate 1">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={-90}
          max={90}
          placeholder="rotate1"
        />
      </Form.Item>

      <Form.Item name={rotate2Obj.path} label="Rotate 2">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={-90}
          max={90}
          placeholder="rotate2"
        />
      </Form.Item>

      <Form.Item name={center0Obj.path} label="Center 0">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={-180}
          max={180}
          placeholder="center0"
        />
      </Form.Item>

      <Form.Item name={center1Obj.path} label="Center 1">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={-180}
          max={180}
          placeholder="center1"
        />
      </Form.Item>
    </div>
  );
}

export default Zoom;
