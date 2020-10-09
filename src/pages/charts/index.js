import React, { useState } from 'react';
import Display from './display.js';
import ChartOption from './options.js';
import './index.css';
import _ from 'lodash';

import { useDispatch, useSelector } from 'react-redux';

import { Card, Tooltip, Button, Input } from 'antd';
import { SaveOutlined, SettingOutlined, EditOutlined } from '@ant-design/icons';

function Chart() {
  const [spec, setSpec] = useState({});
  const showOptions = useSelector((state) => state.chart.showOptions);
  const chartName = useSelector((state) => state.chart.chartName);
  const isChartNameEditable = useSelector((state) => state.chart.isChartNameEditable);
  // const openCopyModal = useSelector(state => state.chart.openCopyModal);
  const dispatch = useDispatch();
  const actions = [
    {
      name: 'Customize',
      Icon: SettingOutlined,
      onClick: () => dispatch({ type: 'set-options' }),
    },
    {
      name: 'Save',
      Icon: SaveOutlined,
    },
  ];

  const IconSize = 20;

  const actionsList = (
    <div className="extra-actions-container">
      <ul>
        {actions.map((item) => (
          <li key={item.name} onClick={item.onClick}>
            <Tooltip title={item.name}>{<item.Icon style={{ fontSize: IconSize }} />}</Tooltip>
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
          onPressEnter={() => dispatch({ type: 'edit-chart-name', value: false })}
          value={chartName}
          onChange={(e) => dispatch({ type: 'set-chart-name', value: e.target.value })}
        />{' '}
        <Button
          style={{ padding: '4px 0px' }}
          size="medium"
          onClick={() => dispatch({ type: 'edit-chart-name', value: false })}
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
        <EditOutlined
          style={{ fontSize: IconSize }}
          onClick={() => dispatch({ type: 'edit-chart-name', value: true })}
        />
      </div>
    );
  }

  const changeSpec = (values) => {
    setSpec(values);
  };

  const addLayer = (func) => {
    console.log('addlayer');
    const layer = spec.layer[0];
    const { encoding, mark } = layer;
    const newLayer = func({ encoding, mark });
    console.log({ newLayer });
    const newSpec = { ...spec };
    newSpec.layer.push(newLayer);
    setSpec(newSpec);
  };

  const removeLayer = () => {
    const newSpec = { ...spec };
    newSpec.layer.splice(1, 1);
    setSpec(newSpec);
  };

  const addField = (path, field, value) => {
    const newSpec = { ...spec };
    _.set(newSpec, [...path, field], value);
    console.log(newSpec, path, field, value);
    setSpec(newSpec);
  };

  const removeField = (path, field) => {
    const newSpec = { ...spec };
    _.unset(newSpec, [...path, field]);
    setSpec(newSpec);
  };

  return (
    <Card
      title={titleComponent}
      extra={actionsList}
      bodyStyle={{ overflow: 'hidden', display: 'flex', padding: '0px' }}
    >
      <div className="display-container" style={{ width: '100%' }}>
        <Display spec={spec} />
      </div>
      <div className="option-container" style={{ right: showOptions ? '0' : '-250px' }}>
        <ChartOption
          onSpecChange={changeSpec}
          addLayer={addLayer}
          removeLayer={removeLayer}
          addField={addField}
          removeField={removeField}
        />
      </div>
    </Card>
  );
}

export default Chart;
