import React from 'react';
import { InputNumber, Select, Checkbox, Form } from 'antd';

import _ from 'lodash';

export const getInterpolateOptions = async (form, setInterpolateOptions) => {
  try {
    const schema = form.getFieldValue('$schema');
    const res = await fetch(schema);
    const jsonData = await res.json();
    setInterpolateOptions(jsonData.definitions.Interpolate.enum);
  } catch (error) {
    console.error(error);
  }
};

function Lines(props) {
  const { form } = props;
  const [enable, setEnable] = React.useState(false);
  const [interpolateOptions, setInterpolateOptions] = React.useState([]);

  React.useEffect(() => {
    getInterpolateOptions(props.form, setInterpolateOptions);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

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
            <InputNumber
              formatter={(value) => parseInt(value) || 0}
              parser={(value) => parseInt(value) || 0}
            />
          </Form.Item>

          <Form.Item name={[...markObj.path, 'line', 'opacity']} initialValue={1} label="Opacity">
            <InputNumber
              formatter={(value) => parseInt(value) || 0}
              parser={(value) => parseInt(value) || 0}
              min={0}
              max={1}
              step={0.05}
            />
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'line', 'interpolate']}
            initialValue="linear"
            label="Line Curve"
          >
            <Select>
              {interpolateOptions.map((option) => (
                <Select.Option key={option} value={option}>
                  {option}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'line', 'strokeDash']}
            initialValue={0}
            label="Dash Width"
          >
            <InputNumber
              formatter={(value) => parseInt(value) || 0}
              parser={(value) => parseInt(value) || 0}
            />
          </Form.Item>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default Lines;
