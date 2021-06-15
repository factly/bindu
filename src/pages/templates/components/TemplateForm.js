import React from 'react';
import { Button, Form, Input, Space, Select } from 'antd';
import { maker, checker } from '../../../utils/slug';
import MediaSelector from '../../../components/MediaSelector';
import SpecEditor from '../../../components/specEditor';
import PropertiesEditor from '../../../components/propertiesEditor';

import Categories from '../../../components/categories';

const { Option } = Select;

const layout = {
  labelCol: {
    span: 5,
  },
  wrapperCol: {
    span: 14,
  },
};
const tailLayout = {
  wrapperCol: {
    offset: 5,
    span: 14,
  },
};

const jsonChecker = (value) => {
  try {
    JSON.parse(value);
    return true;
  } catch {
    return false;
  }
};

const TemplateForm = ({ onSubmit, data = {}, onChange, onModeChange }) => {
  const [form] = Form.useForm();
  const onReset = () => {
    form.resetFields();
  };

  const onTitleChange = (string) => {
    form.setFieldsValue({
      slug: maker(string),
    });
  };

  const onFinish = (values) => {
    const { categories, spec, properties, ...rest } = values;
    onSubmit({
      ...rest,
      spec: JSON.parse(spec),
      properties: JSON.parse(properties),
      category_id: categories,
    });
    onReset();
  };

  return (
    <Form
      {...layout}
      form={form}
      initialValues={{
        ...data,
        spec: data?.spec ? JSON.stringify(data.spec) : '{}',
        properties: data?.properties ? JSON.stringify(data.properties) : '{}',
      }}
      onValuesChange={onChange}
      name="create-chart"
      onFinish={onFinish}
    >
      <Form.Item
        name="title"
        label="Chart Name"
        rules={[
          {
            required: true,
            message: 'Please enter chart name!',
          },
          { min: 3, message: 'Name must be minimum 3 characters.' },
          { max: 50, message: 'Name must be maximum 50 characters.' },
        ]}
      >
        <Input onChange={(e) => onTitleChange(e.target.value)} />
      </Form.Item>
      <Form.Item
        name="slug"
        label="Slug"
        rules={[
          {
            required: true,
            message: 'Please input the slug!',
          },
          {
            pattern: checker,
            message: 'Please enter valid slug!',
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Categories form={form} required label="Category" />
      <Form.Item
        name="mode"
        label="Mode"
        rules={[
          {
            required: true,
            message: 'Please add mode!',
          },
        ]}
      >
        <Select defaultValue="vega-lite" onChange={onModeChange}>
          <Option value="vega-lite">Vega Lite</Option>
          <Option value="vega">Vega</Option>
        </Select>
      </Form.Item>
      <Form.Item
        name="spec"
        label="Spec"
        rules={[
          {
            required: true,
            message: 'Please add spec!',
          },
        ]}
      >
        <SpecEditor />
      </Form.Item>
      <Form.Item
        name="properties"
        label="Properties"
        rules={[
          {
            required: true,
            message: 'Please add properties!',
          },
        ]}
      >
        <PropertiesEditor />
      </Form.Item>
      <Form.Item label="Featured Image" name="medium_id">
        <MediaSelector />
      </Form.Item>
      <Form.Item {...tailLayout}>
        <Space>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
          <Button htmlType="button" onClick={onReset}>
            Reset
          </Button>
        </Space>
      </Form.Item>
    </Form>
  );
};

export default TemplateForm;
