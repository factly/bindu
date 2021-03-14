import React from 'react';
import { Input, Select, Form } from 'antd';

import { aggregateOptions } from '../../constants/x_axis';

const { Option } = Select;

function YAxis(props) {
  const [fields, setFields] = React.useState([]);

  React.useEffect(() => {
    const url = props.form.getFieldValue(['data', 'url']);
    const values = props.form.getFieldValue(['data', 'values']);
    const ext = url.split('.').pop();

    let fetchdata;
    if (ext === 'csv') {
      fetchdata = async (url) => {
        fetch(url)
          .then((response) => response.text())
          .then((csvData) => {
            setFields(csvData.split('\n')[0].split(','));
          })
          .catch((error) => {
            console.error(error);
          });
      };
    } else if (ext === 'json') {
      fetchdata = async (url) => {
        fetch(url)
          .then((response) => response.json())
          .then((jsonData) => {
            setFields(Object.keys(jsonData[0]));
          })
          .catch((error) => {
            console.error(error);
          });
      };
    }
    if (url) {
      fetchdata(url);
    } else if (values) {
      setFields(Object.keys(values[0]));
    }
  }, []);

  const titleObj = props.properties.find((d) => d.prop === 'title');
  const orientObj = props.properties.find((d) => d.prop === 'orient');
  const formatObj = props.properties.find((d) => d.prop === 'format');
  const labelColorObj = props.properties.find((d) => d.prop === 'label_color');
  const aggregateObj = props.properties.find((d) => d.prop === 'aggregate');
  const fieldObj = props.properties.find((d) => d.prop === 'field');

  return (
    <div className="property-container">
      <Form.Item name={titleObj.path} label="Title">
        <Input placeholder="Title" type="text" />
      </Form.Item>

      <Form.Item name={orientObj.path} label="Position">
        <Select>
          <Option value="left">Left</Option>
          <Option value="right">Right</Option>
        </Select>
      </Form.Item>

      <Form.Item name={formatObj.path} label="Label Format">
        <Input placeholder="Label Format" type="text" />
      </Form.Item>

      <Form.Item name={labelColorObj.path} label="Label Color">
        <Input placeholder="Label Color" type="color" />
      </Form.Item>

      <Form.Item name={aggregateObj.path} label="Aggregate">
        <Select placeholder="Label Color" defaultValue={null}>
          <Select.Option value={null}>None</Select.Option>
          {aggregateOptions.map((option) => (
            <Select.Option key={option.value} value={option.value}>
              {option.name}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item name={fieldObj.path} label="Field">
        <Select placeholder="Label Color">
          {fields.map((field) => (
            <Select.Option key={field} value={field}>
              {field}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
    </div>
  );
}

export default YAxis;
