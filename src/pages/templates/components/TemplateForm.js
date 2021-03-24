import React from 'react';
import { Button, Form, Input, Space } from 'antd';
import { maker, checker } from '../../../utils/slug';
import ReactJson from 'react-json-view';

import MediaSelector from '../../../components/MediaSelector';

import * as Bar from '../../charts/bar';

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

const JSONEditor = ({ value, onChange }) => {
  return (
    <ReactJson
      src={value}
      onEdit={({ updated_src }) => onChange(updated_src)}
      onDelete={() => {}}
      onAdd={() => {}}
    />
  );
};

const TemplateForm = ({ onSubmit, data = {} }) => {
  const [form] = Form.useForm();

  const onReset = () => {
    form.resetFields();
  };

  const onTitleChange = (string) => {
    form.setFieldsValue({
      slug: maker(string),
    });
  };

  return (
    <Form
      {...layout}
      form={form}
      initialValues={{ ...data }}
      name="create-chart"
      onValuesChange={({ values }) => console.log({ values })}
      onFinish={(values) => {
        onSubmit(values);
        onReset();
      }}
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
      <Form.Item name="schema" label="Spec" initialValue={Bar.spec}>
        <JSONEditor />
      </Form.Item>
      <Form.Item name="properties" label="Properties" initialValue={Bar.properties}>
        <JSONEditor />
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
