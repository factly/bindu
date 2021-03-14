import React from 'react';
import { Input, Select, Form } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';

export const getAggregateOptions = async (form, setAggregateOptions) => {
  try {
    const schema = form.getFieldValue('$schema');
    const res = await fetch(schema);
    const jsonData = await res.json();
    setAggregateOptions(jsonData.definitions.AggregateOp.enum);
  } catch (error) {
    console.error(error);
  }
};

export const getFields = (form, setFields) => {
  const url = form.getFieldValue(['data', 'url']);
  const values = form.getFieldValue(['data', 'values']);
  if (url) {
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
    fetchdata(url);
  } else if (values) {
    setFields(Object.keys(values[0]));
  }
};

const { Option } = Select;

function XAxis(props) {
  const [fields, setFields] = React.useState([]);
  const [aggregateOptions, setAggregateOptions] = React.useState([]);

  React.useEffect(() => {
    getFields(props.form, setFields);
    getAggregateOptions(props.form, setAggregateOptions);
  }, []);

  const titleObj = props.properties.find((d) => d.prop === 'title');
  const orientObj = props.properties.find((d) => d.prop === 'orient');
  const formatObj = props.properties.find((d) => d.prop === 'format');
  const labelColorObj = props.properties.find((d) => d.prop === 'label_color');
  const aggregateObj = props.properties.find((d) => d.prop === 'aggregate');
  const fieldObj = props.properties.find((d) => d.prop === 'field');

  return (
    <div className="property-container">
      <Form.Item name={titleObj.path} lable="Title">
        <Input placeholder="Title" type="text" />
      </Form.Item>

      <Form.Item name={orientObj.path} label="Position">
        <Select>
          <Option value="top">Top</Option>
          <Option value="bottom">bottom</Option>
        </Select>
      </Form.Item>

      <Form.Item name={formatObj.path} label="Label Format">
        <Input placeholder="Label Format" type="text" />
      </Form.Item>

      <Form.Item name={labelColorObj.path} label="Label Color">
        <Input placeholder="Label Color" type="color" />
      </Form.Item>

      <Form.Item
        name={aggregateObj.path}
        label={
          <div>
            Aggregate{' '}
            <InfoCircleOutlined
              onClick={() =>
                window.open('https://vega.github.io/vega-lite/docs/aggregate.html#ops', '_blank')
              }
            />
          </div>
        }
      >
        <Select placeholder="Label Color" defaultValue={null}>
          <Select.Option value={null}>None</Select.Option>
          {aggregateOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
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

export default XAxis;
