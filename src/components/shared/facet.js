import React from 'react';
import { InputNumber, Select, Form } from 'antd';

const getResolveOptions = async (form, setResolveOptions) => {
  try {
    const schema = form.getFieldValue('$schema');
    const res = await fetch(schema);
    const jsonData = await res.json();
    setResolveOptions(jsonData.definitions.ResolveMode.enum);
  } catch (error) {
    console.error(error);
  }
};

function Facet(props) {
  const [resolveOptions, setResolveOptions] = React.useState([]);

  const columnsObj = props.properties.find((d) => d.prop === 'column');
  const spacingObj = props.properties.find((d) => d.prop === 'spacing');
  const xaxisObj = props.properties.find((d) => d.prop === 'xaxis');
  const yaxisObj = props.properties.find((d) => d.prop === 'yaxis');

  React.useEffect(() => {
    getResolveOptions(props.form, setResolveOptions);
  }, []);

  return (
    <div className="property-container">
      <Form.Item name={columnsObj.path} label="Columns">
        <InputNumber
          formatter={(value) => parseInt(value) || 1}
          parser={(value) => parseInt(value) || 1}
          placeholder="columns"
          min={1}
        />
      </Form.Item>

      <Form.Item name={spacingObj.path} label="Spacing">
        <InputNumber
          formatter={(value) => parseInt(value) || 0}
          parser={(value) => parseInt(value) || 0}
          placeholder="spacing"
          min={0}
        />
      </Form.Item>
      {xaxisObj ? (
        <Form.Item name={xaxisObj.path} label="X Axis">
          <Select>
            {resolveOptions.map((option) => (
              <Select.Option key={option} value={option}>
                {option}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>
      ) : null}
      {yaxisObj ? (
        <Form.Item name={yaxisObj.path} label="Y Axis">
          <Select>
            {resolveOptions.map((option) => (
              <Select.Option key={option} value={option}>
                {option}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>
      ) : null}
    </div>
  );
}

export default Facet;
