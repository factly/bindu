import React from 'react';
import { Input, Form, InputNumber, Select, Checkbox } from 'antd';
import { getAutosizeOptions } from '../../utils/options';
import _ from 'lodash';

function Dimensions(props) {
  const titleObj = props.properties.find((d) => d.prop === 'title');
  const widthObj = props.properties.find((d) => d.prop === 'width');
  const heightObj = props.properties.find((d) => d.prop === 'height');
  const autosizeObj = props.properties.find((d) => d.prop === 'width');
  const backgroundObj = props.properties.find((d) => d.prop === 'background');

  const [autosizeOptions, setAutosizeOptions] = React.useState([]);
  const [autoWidth, setAutoWidth] = React.useState(false);
  const [autoHeight, setAutoHeight] = React.useState(false);

  const widthValue = props.form.getFieldValue(widthObj.path);
  const heightValue = props.form.getFieldValue(heightObj.path);
  React.useEffect(() => {
    getAutosizeOptions(props.form, setAutosizeOptions);
  }, []);

  React.useEffect(() => {
    if (widthValue === 'container') {
      setAutoWidth(true);
    }
    if (heightValue === 'container') {
      setAutoHeight(true);
    }
  }, [widthValue, heightValue]);

  const onAutoHeightChange = (event) => {
    const checked = event.target.checked;
    changeFormValues(checked, heightObj.path);
    setAutoHeight(checked);
  };

  const onAutoWidthChange = (event) => {
    const checked = event.target.checked;
    changeFormValues(checked, widthObj.path);
    setAutoWidth(checked);
  };

  const changeFormValues = (checked, path) => {
    const { form } = props;
    const values = form.getFieldValue([]);
    if (checked) {
      _.set(values, path, 'container');
    } else {
      _.set(values, path, 500);
    }
    form.setFieldsValue(values);
  };

  return (
    <div className="property-container">
      <Form.Item name={titleObj.path} label="Title">
        <Input placeholder="title" type="text" />
      </Form.Item>

      <Checkbox checked={autoWidth} onChange={onAutoWidthChange}>
        Auto
      </Checkbox>
      <Form.Item name={widthObj.path} label="Width">
        {autoWidth ? (
          <Input defaultValue="container" disabled />
        ) : (
          <InputNumber
            disabled={autoWidth}
            initialvalue={500}
            formatter={(value) => parseInt(value) || 0}
            parser={(value) => parseInt(value) || 0}
            placeholder="width"
            min={0}
          />
        )}
      </Form.Item>

      <Checkbox checked={autoHeight} onChange={onAutoHeightChange}>
        Auto
      </Checkbox>
      <Form.Item name={heightObj.path} label="Height">
        {autoHeight ? (
          <Input defaultValue="container" disabled />
        ) : (
          <InputNumber
            formatter={(value) => parseInt(value) || 0}
            parser={(value) => parseInt(value) || 0}
            placeholder="height"
            min={0}
          />
        )}
      </Form.Item>

      <Form.Item name={[...autosizeObj.path.slice(0, -1), 'autosize']} label="Autosize">
        <Select initialvalue={autosizeOptions[0]}>
          {autosizeOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item name={backgroundObj.path} label="Background">
        <Input type="color" />
      </Form.Item>
    </div>
  );
}

export default Dimensions;
