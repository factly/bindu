import React from 'react';
import { InputNumber, Select, Checkbox, Form } from 'antd';

import _ from 'lodash';

const { Option } = Select;

function Lines(props) {
  const { form } = props;
  const [enable, setEnable] = React.useState(false);

  const markObj = props.properties.find((d) => d.prop === 'mark');

  const handleEnable = (checked) => {
    setEnable(checked);
    if (!checked) {
      let values = form.getFieldValue();
      _.unset(values, [...markObj.path, 'line']);
      form.setFieldsValue(values);
    }
  };

  return (
    <div className="property-container">
      <Checkbox
        onChange={(e) => {
          const checked = e.target.checked;
          handleEnable(checked);
        }}
      >
        Enable
      </Checkbox>
      {enable ? (
        <React.Fragment>
          <Form.Item name={[...markObj.path, 'line', 'strokeWidth']} initialValue={4} label="Width">
            <InputNumber />
          </Form.Item>

          <Form.Item name={[...markObj.path, 'line', 'opacity']} initialValue={1} label="Opacity">
            <InputNumber min={0} max={1} step={0.05} />
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'line', 'interpolate']}
            initialValue="linear"
            label="Line Curve"
          >
            <Select>
              <Option value="linear">Linear</Option>
              <Option value="linear-closed">Linear Closed</Option>
              <Option value="step">Step</Option>
              <Option value="basis">Basis</Option>
              <Option value="monotone">Monotone</Option>
            </Select>
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'line', 'strokeDash']}
            initialValue={0}
            label="Dash Width"
          >
            <InputNumber />
          </Form.Item>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default Lines;
