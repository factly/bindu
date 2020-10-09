import React from 'react';
import { Input, Select, Form } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_FACET_COLUMNS,
  SET_FACET_SPACING,
  SET_FACET_XAXIS,
  SET_FACET_YAXIS,
} from '../../constants/facet.js';

const { Option } = Select;
function Facet(props) {
  const spec = useSelector((state) => state.chart.spec);

  const columnsObj = props.properties.find((d) => d.prop === 'column');
  const spacingObj = props.properties.find((d) => d.prop === 'spacing');
  const xaxisObj = props.properties.find((d) => d.prop === 'xaxis');
  const yaxisObj = props.properties.find((d) => d.prop === 'yaxis');

  const columns = getValueFromNestedPath(spec, columnsObj.path);
  const spacing = getValueFromNestedPath(spec, spacingObj.path);

  let xaxis;
  if (xaxisObj) {
    xaxis = getValueFromNestedPath(spec, xaxisObj.path);
  }
  let yaxis;
  if (yaxisObj) {
    yaxis = getValueFromNestedPath(spec, yaxisObj.path);
  }
  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Form.Item name={columnsObj.path} label="Columns">
        <Input
          placeholder="columns"
          min={0}
          type="number"
          onChange={(e) =>
            dispatch({
              type: SET_FACET_COLUMNS,
              payload: { value: e.target.value, path: columnsObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>

      <Form.Item name={spacingObj.path} label="Spacing">
        <Input
          placeholder="columns"
          min={0}
          type="number"
          onChange={(e) =>
            dispatch({
              type: SET_FACET_SPACING,
              payload: { value: e.target.value, path: spacingObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>
      {xaxisObj ? (
        <Form.Item name={xaxisObj.path} label="X Axis">
          <Select
            onChange={(value) =>
              dispatch({
                type: SET_FACET_XAXIS,
                payload: { value: value, path: xaxisObj.path },
                chart: 'shared',
              })
            }
          >
            <Option value="shared">Shared</Option>
            <Option value="independent">Independent</Option>
          </Select>
        </Form.Item>
      ) : null}
      {yaxisObj ? (
        <Form.Item name={yaxisObj.path} label="Y Axis">
          <Select
            onChange={(value) =>
              dispatch({
                type: SET_FACET_YAXIS,
                payload: { value: value, path: yaxisObj.path },
                chart: 'shared',
              })
            }
          >
            <Option value="shared">Shared</Option>
            <Option value="independent">Independent</Option>
          </Select>
        </Form.Item>
      ) : null}
    </div>
  );
}

export default Facet;
