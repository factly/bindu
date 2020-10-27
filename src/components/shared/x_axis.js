import React from 'react';
import { Input, Select, Form } from 'antd';

const { Option } = Select;

function XAxis(props) {
  const titleObj = props.properties.find((d) => d.prop === 'title');
  const orientObj = props.properties.find((d) => d.prop === 'orient');
  const formatObj = props.properties.find((d) => d.prop === 'format');
  const labelColorObj = props.properties.find((d) => d.prop === 'label_color');

  return (
    <div className="property-container">
      <Form.Item name={titleObj.path} lable="Title">
        <Input placeholder="Title" type="text" />
      </Form.Item>

      <Form.Item name={orientObj.path} label="Position">
        <Select>
          <Option value="top">Top</Option>
          <Option value="bottom">bottom</Option>
        </Select>
      </Form.Item>

      <Form.Item name={formatObj.path} label="Label Format">
        <Input placeholder="Label Format" type="text" />
      </Form.Item>

      <Form.Item name={labelColorObj.path} label="Label Color">
        <Input placeholder="Label Color" type="color" />
      </Form.Item>
    </div>
  );
}

export default XAxis;
