import React from 'react';
import { Input, Form } from 'antd';

function Colors(props) {
  const colorObj = props.properties.find((d) => d.prop === 'color');

  let colors = props.form.getFieldValue(colorObj.path);

  if (colorObj.type === 'string') {
    colors = [colors];
  }

  const getName = (index) => {
    if (colorObj.type === 'string') {
      return colorObj.path;
    }

    return [...colorObj.path, index];
  };

  return (
    <div className="property-container">
      {colors &&
        colors.map((d, i) => (
          <Form.Item name={getName(i)} label="Colors">
            <Input type="color" key={i}></Input>
          </Form.Item>
        ))}
    </div>
  );
}

export default Colors;
