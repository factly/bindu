import React from 'react';
import { Input, Select, Form, InputNumber } from 'antd';

const { Option } = Select;

function Legend(props) {
  // const {title, fillColor, symbolType, symbolSize, orient} = layer.encoding.color.legend;
  const titleObj = props.properties.find((d) => d.prop === 'title');
  const fillColorObj = props.properties.find((d) => d.prop === 'fill_color');
  const symbolTypeObj = props.properties.find((d) => d.prop === 'symbol_type');
  const symbolSizeObj = props.properties.find((d) => d.prop === 'symbol_size');
  const orientObj = props.properties.find((d) => d.prop === 'orient');

  return (
    <div className="property-container">
      <Form.Item name={orientObj.path} label="Position">
        <Select>
          <Option value="left">Left</Option>
          <Option value="right">Right</Option>
          <Option value="top">Top</Option>
          <Option value="bottom">Bottom</Option>
          <Option value="top-left">Top Left</Option>
          <Option value="top-right">Top Right</Option>
          <Option value="bottom-left">Bottom Left</Option>
          <Option value="bottom-right">Bottom Right</Option>
        </Select>
      </Form.Item>

      <Form.Item name={titleObj.path} label="Title">
        <Input placeholder="Title" type="text" />
      </Form.Item>

      <Form.Item name={fillColorObj.path} label="Background Color">
        <Input type="color" />
      </Form.Item>

      <Form.Item name={symbolTypeObj.path} labe="Symbol">
        <Select>
          <Option value="circle">Circle</Option>
          <Option value="square">Square</Option>
          <Option value="cross">Cross</Option>
          <Option value="diamond">Diamond</Option>
          <Option value="triangle-up">Triangle Up</Option>
          <Option value="triangle-down">Triangle Down</Option>
          <Option value="triangle-right">Triangle Right</Option>
          <Option value="triangle-left">Triangle Left</Option>
        </Select>
      </Form.Item>

      <Form.Item name={symbolSizeObj.path} label="Symbol Size">
        <InputNumber placeholder="Symbol Size" min={0} />
      </Form.Item>
    </div>
  );
}

export default Legend;
