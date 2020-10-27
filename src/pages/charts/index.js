import React, { useState } from 'react';
import Display from './display.js';
import ChartOption from './options.js';
import './index.css';
import _ from 'lodash';

import { Card, Tooltip, Button, Input, Form } from 'antd';
import { SaveOutlined, SettingOutlined, EditOutlined } from '@ant-design/icons';

function Chart() {
  const [form] = Form.useForm();

  const [showOptions, setShowOptions] = useState(true);
  const [chartName, setChartName] = useState('Untitled');
  const [isChartNameEditable, setChartNameEditable] = useState(false);

  const actions = [
    {
      name: 'Customize',
      Icon: SettingOutlined,
      props: { onClick: () => setShowOptions(!showOptions) },
    },
    {
      name: 'Save',
      Icon: SaveOutlined,
      props: {},
    },
  ];

  const IconSize = 20;

  const actionsList = (
    <div className="extra-actions-container">
      <ul>
        {actions.map((item) => (
          <li key={item.name}>
            <Tooltip title={item.name}>
              {<item.Icon {...item.props} style={{ fontSize: IconSize }} />}
            </Tooltip>
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
  );
}

export default Chart;
