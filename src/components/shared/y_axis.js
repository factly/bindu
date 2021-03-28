import React from 'react';
import { Input, Select, Form } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import _ from 'lodash';

import {
  getAggregateOptions,
  getSortOrderOptions,
  getFields,
  getTypeOptions,
  getTimeUnitOptions,
} from './x_axis';

const { Option } = Select;

function YAxis(props) {
  const [fields, setFields] = React.useState([]);
  const [aggregateOptions, setAggregateOptions] = React.useState([]);
  const [sortOrderOptions, setSortOrderOptions] = React.useState([]);
  const [typeOptions, setTypeOptions] = React.useState([]);
  const [timeUnitOptions, setTimeUnitOptions] = React.useState([]);

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
        <Select showSearch placeholder="Aggregate" defaultValue={null}>
          <Select.Option value={null}>None</Select.Option>
          {aggregateOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
            </Select.Option>
          ))}
        </Select>
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
        <Select showSearch placeholder="Type" defaultValue={null}>
          <Select.Option value={null}>None</Select.Option>
          {typeOptions.map((option) => (
            <Select.Option key={option} value={option}>
              {option}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>

      {type === 'temporal' && (
        <Form.Item label="Time unit">
          <Select mode="multiple" onChange={onTimeUnitChange}>
            <Select.Option value={null}>None</Select.Option>
            {timeUnitOptions.map((field, index) => (
              <Select.Option key={field} value={index}>
                {field}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>
      )}

      <Form.Item name={fieldObj.path} label="Field">
        <Select placeholder="Label Color">
          {fields.map((field) => (
            <Select.Option key={field} value={field}>
              {field}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>

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
              <Select defaultValue={null}>
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

export default YAxis;
