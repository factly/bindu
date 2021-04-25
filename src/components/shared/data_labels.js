import React from 'react';
import { Input, InputNumber, Form, Checkbox } from 'antd';

function DataLabels(props) {
  const { form } = props;
  const [enable, setEnable] = React.useState(false);

  const colorObj = props.properties.find((d) => d.prop === 'color');
  const fontSizeObj = props.properties.find((d) => d.prop === 'font_size');
  const formatObj = props.properties.find((d) => d.prop === 'format');

  const handleEnable = (checked) => {
    let values = form.getFieldValue([]);
    if (checked) {
      const { encoding } = values.layer[0];
      values.layer.push({
        mark: { type: 'text', dx: 0, dy: 10, fontSize: 12 },
        encoding: {
          x: encoding.x,
          y: { ...encoding.y, stack: 'zero' },
          color: { value: '#ffffff' },
          text: { ...encoding.y, format: '.1f' },
        },
      });
    } else {
      values.layer.splice(1, 1);
    }
    setEnable(checked);
    form.setFieldsValue(values);
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
          <Form.Item name={colorObj.path} label="Color">
            <Input type="color" />
          </Form.Item>

          <Form.Item name={fontSizeObj.path} label="Size">
            <InputNumber min={0} />
          </Form.Item>

          <Form.Item name={formatObj.path} label="Format">
            <Input type="text" />
          </Form.Item>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default DataLabels;
