import React from 'react';
import { Input, Form, InputNumber } from 'antd';

function Dimensions(props) {
  const titleObj = props.properties.find((d) => d.prop === 'title');
  const widthObj = props.properties.find((d) => d.prop === 'width');
  const heightObj = props.properties.find((d) => d.prop === 'height');
  const backgroundObj = props.properties.find((d) => d.prop === 'background');

  return (
    <div className="property-container">
      <Form.Item name={titleObj.path} label="Title">
        <Input placeholder="title" type="text" />
      </Form.Item>

      <Form.Item name={widthObj.path} label="Width">
        <InputNumber placeholder="width" />
      </Form.Item>
      <Form.Item name={heightObj.path} label="Height">
        <InputNumber placeholder="height" />
      </Form.Item>
      <Form.Item name={backgroundObj.path} label="Background">
        <Input type="color" />
      </Form.Item>
    </div>
  );
}

export default Dimensions;
