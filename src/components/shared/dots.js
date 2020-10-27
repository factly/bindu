import React from 'react';
import { Select, Form, Checkbox, InputNumber } from 'antd';

const { Option } = Select;

function Dots(props) {
  const { form } = props;
  const [enable, setEnable] = React.useState(false);

  const markObj = props.properties.find((d) => d.prop === 'mark');

  const handleEnable = (checked) => {
    if (!checked) {
      let values = form.getFieldValue();
      _.unset(values, [...markObj.path, 'point']);
      form.setFieldsValue(values);
    }
    setEnable(checked);
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
          <Form.Item
            name={[...markObj.path, 'point', 'shape']}
            initialValue={'circle'}
            label="Symbol"
          >
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

          <Form.Item
            name={[...markObj.path, 'point', 'size']}
            initialValue={40}
            label="Symbol Size"
          >
            <InputNumber placeholder="Symbol Size" />
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'point', 'filled']}
            initialValue={true}
            valuePropName="checked"
            label="Filled"
          >
            <Checkbox />
          </Form.Item>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default Dots;
