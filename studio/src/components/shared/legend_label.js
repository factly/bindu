import React from 'react';
import { Input, Select, Form } from 'antd';

const { Option } = Select;

function LegendLabel(props) {
  // const {labelAlign, labelBaseline, labelColor} = spec.layer[0].encoding.color.legend;
  const labelAlignObj = props.properties.find((d) => d.prop === 'label_align');
  const labelBaselineObj = props.properties.find((d) => d.prop === 'label_baseline');
  const labelColorObj = props.properties.find((d) => d.prop === 'label_color');

  return (
    <div className="property-container">
      <Form.Item name={labelAlignObj.path} label="Align">
        <Select>
          <Option value="left">Left</Option>
          <Option value="right">Right</Option>
          <Option value="center">Center</Option>
        </Select>
      </Form.Item>

      <Form.Item name={labelBaselineObj.path} label="Baseline">
        <Select>
          <Option value="top">Top</Option>
          <Option value="bottom">Bottom</Option>
          <Option value="middle">Center</Option>
        </Select>
      </Form.Item>

      <Form.Item name={labelColorObj.path} label="Color">
        <Input type="color" />
      </Form.Item>
    </div>
  );
}

export default LegendLabel;
