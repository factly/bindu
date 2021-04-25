import React from 'react';
import { Input, Select, Form } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import _ from 'lodash';

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

export const getSortOrderOptions = async (form, setSortOrderOptions) => {
  try {
    const schema = form.getFieldValue('$schema');
    const res = await fetch(schema);
    const jsonData = await res.json();
    setSortOrderOptions(jsonData.definitions.SortOrder.enum);
    // can also contain null. improve the fetching using definitions.SortField.properties.order
  } catch (error) {
    console.error(error);
  }
};

export const getTypeOptions = async (form, setTypeOptions) => {
  try {
    const schema = form.getFieldValue('$schema');
    const res = await fetch(schema);
    const jsonData = await res.json();
    setTypeOptions(jsonData.definitions.Type.enum);
    // can also contain null. improve the fetching using definitions.SortField.properties.order
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

export const getTimeUnitOptions = (form, setFields) => {
  setFields([
    'year',
    'quarter',
    'month',
    'date',
    'week',
    'day',
    'dayofyear',
    'hours',
    'minutes',
    'seconds',
    'milliseconds',
  ]);
};

function Axis(props) {
  const [fields, setFields] = React.useState([]);
  const [sortOrderOptions, setSortOrderOptions] = React.useState([]);
  const [aggregateOptions, setAggregateOptions] = React.useState([]);
  const [typeOptions, setTypeOptions] = React.useState([]);
  const [timeUnitOptions, setTimeUnitOptions] = React.useState([]);
  const positionOptions = props.axis === 'x' ? ['top', 'bottom'] : ['left', 'right'];

  React.useEffect(() => {
    getFields(props.form, setFields);
    getAggregateOptions(props.form, setAggregateOptions);
    getSortOrderOptions(props.form, setSortOrderOptions);
    getTypeOptions(props.form, setTypeOptions);
    getTimeUnitOptions(props.form, setTimeUnitOptions);
  }, []);

  const titleObj = props.properties.find((d) => d.prop === 'title');
  const orientObj = props.properties.find((d) => d.prop === 'orient');
  const formatObj = props.properties.find((d) => d.prop === 'format');
  const labelColorObj = props.properties.find((d) => d.prop === 'label_color');
  const aggregateObj = props.properties.find((d) => d.prop === 'aggregate');
  const fieldObj = props.properties.find((d) => d.prop === 'field');
  const typeObj = props.properties.find((d) => d.prop === 'type');
  const sortObj = props.properties.find((d) => d.prop === 'sort');

  const type = props.form.getFieldValue(typeObj.path);
  const timeUnitPath = [...typeObj.path.slice(0, -1), 'timeUnit'];
  const onTimeUnitChange = (selectedValues) => {
    const timeUnitValue = selectedValues
      .sort()
      .map((i) => timeUnitOptions[i])
      .join('');
    let values = props.form.getFieldValue([]);
    _.set(values, timeUnitPath, timeUnitValue);
    props.form.setFieldsValue(values);
    console.log({ timeUnitValue });
  };

  const removeField = (path) => {
    let values = props.form.getFieldValue([]);
    _.unset(values, path);
    props.form.setFieldsValue(values);
  };

  const onTypeChange = () => {
    removeField(sortObj.path);
  };

  return (
    <div className="property-container">
      <Form.Item name={titleObj.path} label="Title">
        <Input placeholder="Title" type="text" />
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

      <Form.Item name={orientObj.path} label="Position">
        <Select>
          {positionOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>

      <Form.Item name={formatObj.path} label="Label Format">
        <Input placeholder="Label Format" type="text" />
      </Form.Item>

      <Form.Item name={labelColorObj.path} label="Label Color">
        <Input placeholder="Label Color" type="color" />
      </Form.Item>

      <Form.Item
        name={typeObj.path}
        label={
          <div>
            Type{' '}
            <InfoCircleOutlined
              onClick={() =>
                window.open('https://vega.github.io/vega-lite/docs/type.html', '_blank')
              }
            />
          </div>
        }
      >
        <Select showSearch placeholder="Type" onChange={onTypeChange}>
          {typeOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
            </Select.Option>
          ))}
        </Select>
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
        <Select placeholder="Aggregate">
          <Select.Option value={null}>None</Select.Option>
          {aggregateOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>

      {type === 'temporal' && (
        <Form.Item label="Time unit">
          <Select mode="multiple" onChange={onTimeUnitChange}>
            {timeUnitOptions.map((field, index) => (
              <Select.Option key={field} value={index}>
                {field}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>
      )}

      <Form.Item label="Sort">
        {type === 'ordinal' || type === 'nominal' ? (
          <>
            <Form.Item name={[...sortObj.path, 'order']} label="order">
              <Select>
                <Select.Option value={null}>None</Select.Option>
                {sortOrderOptions.map((field) => (
                  <Select.Option key={field} value={field}>
                    {field}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>

            <Form.Item name={[...sortObj.path, 'field']} label="By field">
              <Select>
                {fields.map((field) => (
                  <Select.Option key={field} value={field}>
                    {field}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>

            <Form.Item name={[...sortObj.path, 'op']} label="Aggregate operation">
              <Select>
                <Select.Option value={null}>None</Select.Option>
                {aggregateOptions.map((option) => (
                  <Select.Option key={option} value={option}>
                    {option}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </>
        ) : type === 'quantitative' || type === 'temporal' ? (
          <Form.Item name={sortObj.path} label="order">
            <Select>
              <Select.Option value={null}>None</Select.Option>
              {sortOrderOptions.map((field) => (
                <Select.Option key={field} value={field}>
                  {field}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
        ) : (
          'Not applicable for this chart'
        )}
      </Form.Item>
    </div>
  );
}

export default Axis;
