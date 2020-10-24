import React, { useState } from 'react';
import Display from './display.js';
import ChartOption from './options.js';
import './index.css';

import { Card, Tooltip, Button, Input, Form, Modal } from 'antd';
import { SaveOutlined, SettingOutlined, EditOutlined, UploadOutlined } from '@ant-design/icons';
import UppyUploader from '../../components/uppy';

function Chart() {
  const [form] = Form.useForm();

  const [showModal, setShowModal] = useState(false);
  const [showOptions, setShowOptions] = useState(true);
  const [chartName, setChartName] = useState('Untitled');
  const [isChartNameEditable, setChartNameEditable] = useState(false);

  const IconSize = 20;
  const actions = [
    {
      name: 'Customize',
      Component: (
        <SettingOutlined
          style={{ fontSize: IconSize }}
          onClick={() => setShowOptions(!showOptions)}
        />
      ),
    },
    {
      name: 'Save',
      Component: <SaveOutlined style={{ fontSize: IconSize }} />,
    },
    {
      name: 'Upload',
      Component: (
        <UploadOutlined style={{ fontSize: IconSize }} onClick={() => setShowModal(true)} />
      ),
    },
  ];

  const actionsList = (
    <div className="extra-actions-container">
      <ul>
        {actions.map((item) => (
          <li key={item.name}>
            <Tooltip title={item.name}>{item.Component}</Tooltip>
          </li>
        ))}
      </ul>
    </div>
  );
  let titleComponent;
  if (isChartNameEditable) {
    titleComponent = (
      <div className="chart-name-editable-container">
        <Input
          onPressEnter={() => setChartNameEditable(false)}
          value={chartName}
          onChange={(e) => setChartName(e.target.value)}
        />{' '}
        <Button
          style={{ padding: '4px 0px' }}
          size="medium"
          onClick={() => setChartNameEditable(false)}
          type="primary"
        >
          Save
        </Button>{' '}
      </div>
    );
  } else {
    titleComponent = (
      <div className="chart-name-container">
        <label className="chart-name">{chartName}</label>
        <EditOutlined style={{ fontSize: IconSize }} onClick={() => setChartNameEditable(true)} />
      </div>
    );
  }

  return (
    <>
      <Form form={form}>
        <Card
          title={titleComponent}
          extra={actionsList}
          bodyStyle={{ overflow: 'hidden', display: 'flex', padding: '0px' }}
        >
          <div className="display-container" style={{ width: '100%' }}>
            <Form.Item shouldUpdate={true}>
              {({ getFieldValue }) => {
                return (
                  <Form.Item>
                    <Display spec={getFieldValue()} />
                  </Form.Item>
                );
              }}
            </Form.Item>
          </div>
          <div className="option-container" style={{ right: showOptions ? '0' : '-250px' }}>
            <ChartOption form={form} />
          </div>
        </Card>
      </Form>
      <Modal
        title="Upload Dataset"
        visible={showModal}
        onOk={() => setShowModal(false)}
        onCancel={() => setShowModal(false)}
      >
        <UppyUploader onUpload={console.log} />
      </Modal>
    </>
  );
}

export default Chart;
