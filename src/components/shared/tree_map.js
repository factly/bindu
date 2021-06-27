import React from 'react';
import { InputNumber, Select, Form } from 'antd';

const { Option } = Select;

function TreeMap(props) {
  const layoutObj = props.properties.find((d) => d.prop === 'layout');
  const aspectRatioObj = props.properties.find((d) => d.prop === 'aspect_ratio');

  return (
    <div className="property-container">
      <Form.Item name={layoutObj.path} label="Layout Mode">
        <Select>
          <Option value="squarify">Squarify</Option>
          <Option value="resquarify">Resquarify</Option>
          <Option value="binary">Binary</Option>
          <Option value="slice">Slice</Option>
          <Option value="dice">Dice</Option>
          <Option value="slicedice">Slicedice</Option>
        </Select>
      </Form.Item>

      <Form.Item name={aspectRatioObj.path} label="Aspect Ratio">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="Aspect Ratio"
        />
      </Form.Item>
    </div>
  );
}

export default TreeMap;
