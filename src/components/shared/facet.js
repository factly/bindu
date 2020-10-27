import React from 'react';
import { Input, Select, Form } from 'antd';

const { Option } = Select;

function Facet(props) {
  const columnsObj = props.properties.find((d) => d.prop === 'column');
  const spacingObj = props.properties.find((d) => d.prop === 'spacing');
  const xaxisObj = props.properties.find((d) => d.prop === 'xaxis');
  const yaxisObj = props.properties.find((d) => d.prop === 'yaxis');

  return (
    <div className="property-container">
      <Form.Item name={columnsObj.path} label="Columns">
        <Input placeholder="columns" min={0} type="number" />
      </Form.Item>

      <Form.Item name={spacingObj.path} label="Spacing">
        <Input placeholder="columns" min={0} type="number" />
      </Form.Item>
      {xaxisObj ? (
        <Form.Item name={xaxisObj.path} label="X Axis">
          <Select>
            <Option value="shared">Shared</Option>
            <Option value="independent">Independent</Option>
          </Select>
        </Form.Item>
      ) : null}
      {yaxisObj ? (
        <Form.Item name={yaxisObj.path} label="Y Axis">
          <Select>
            <Option value="shared">Shared</Option>
            <Option value="independent">Independent</Option>
          </Select>
        </Form.Item>
      ) : null}
    </div>
  );
}

export default Facet;
