import React from 'react';
import { InputNumber, Form } from 'antd';

function Graticule(props) {
  const longSepObj = props.properties.find((d) => d.prop === 'long-sep');
  const latSepObj = props.properties.find((d) => d.prop === 'lat-sep');
  const colorObj = props.properties.find((d) => d.prop === 'color');
  const widthObj = props.properties.find((d) => d.prop === 'width');
  const opacityObj = props.properties.find((d) => d.prop === 'opacity');
  const strokeDashObj = props.properties.find((d) => d.prop === 'dash');

  return (
    <div className="property-container">
      <Form.Item name={longSepObj.path} label="Longitude Separation">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={0}
          placeholder="Long Sep"
        />
      </Form.Item>
      <Form.Item name={latSepObj.path} label="Latitude Separation">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={0}
          placeholder="Lat Sep"
        />
      </Form.Item>
      <Form.Item name={widthObj.path} label="Width">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={0}
          placeholder="width"
        />
      </Form.Item>
      <Form.Item name={strokeDashObj.path} label="Dashed">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={0}
          placeholder="dash"
        />
      </Form.Item>
      <Form.Item name={opacityObj.path} label="Opacity">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          min={0}
          max={1}
          step={0.05}
          placeholder="opacity"
        />
      </Form.Item>
      <Form.Item name={colorObj.path} label="Color">
        <InputNumber type="color" />
      </Form.Item>
    </div>
  );
}

export default Graticule;
