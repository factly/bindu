import React, { useState } from 'react';
import Display from './display.js';
import ChartOption from './options.js';
import { saveAs } from 'file-saver';
import './index.css';

import { Card, Tooltip, Button, Input, Form, Modal, Dropdown, Menu } from 'antd';
import {
  SaveOutlined,
  SettingOutlined,
  EditOutlined,
  UploadOutlined,
  MenuOutlined,
} from '@ant-design/icons';
import UppyUploader from '../../components/uppy';
import { b64toBlob } from '../../utils/file';

function Chart() {
  const [form] = Form.useForm();

  const [showModal, setShowModal] = useState(false);
  const [showOptions, setShowOptions] = useState(true);
  const [chartName, setChartName] = useState('Untitled');
  const [isChartNameEditable, setChartNameEditable] = useState(false);
  const [view, setView] = useState(null);

  const onDataUpload = (dataDetails) => {
    form.setFieldsValue({ data: { url: dataDetails.url.raw } });
  };

  const downloadSampleData = () => {
    const url = form.getFieldValue(['data', 'url']);
    const values = form.getFieldValue(['data', 'values']);
    if (url) {
      saveAs(url, url.split('/').pop());
    } else if (values) {
      const blob = new Blob([JSON.stringify(values)], { type: 'application/json;charset=utf-8' });
      saveAs(blob, 'sample.json');
    }
  };

  const downloadImage = async (e) => {
    const ext = e.key;
    const data = await view?.toImageURL(ext, 1);
    const blob = b64toBlob(data.split(',')[1], 'image/' + ext);
    saveAs(blob, `${chartName}.${ext}`);
  };

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
    {
      name: 'Options',
      Component: (
        <Dropdown
          overlay={
            <div style={{ boxShadow: '0px 0px 6px 1px #999' }}>
              <Menu onClick={downloadImage}>
                <Menu.Item key="svg">Download Image (SVG)</Menu.Item>
                <Menu.Item key="png">Download Image (PNG)</Menu.Item>
              </Menu>
              <Menu.Divider />
              <Menu onClick={downloadSampleData}>
                <Menu.Item>Download Sample</Menu.Item>
              </Menu>
            </div>
          }
        >
          <Button style={{ marginBottom: 5 }}>
            <MenuOutlined />
          </Button>
        </Dropdown>
      ),
    },
  ];

  const actionsList = (
    <div className="extra-actions-container">
      <ul style={{ display: 'flex', alignItems: 'center' }}>
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
                    <Display spec={getFieldValue()} setView={setView} />
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
        okText="Done"
      >
        <UppyUploader onUpload={onDataUpload} />
      </Modal>
    </>
  );
}

export default Chart;
