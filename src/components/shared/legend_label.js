import React from 'react';
import { Input, Select, Form } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_LEGEND_LABEL_POSITION,
  SET_LEGEND_LABEL_BASELINE,
  SET_LEGEND_LABEL_COLOR,
} from '../../constants/legend_label.js';
const { Option } = Select;

function LegendLabel(props) {
  const spec = useSelector((state) => state.chart.spec);
  // const layer = spec.layer[0];
  // const {labelAlign, labelBaseline, labelColor} = layer.encoding.color.legend;
  const labelAlignObj = props.properties.find((d) => d.prop === 'label_align');
  const labelAlign = getValueFromNestedPath(spec, labelAlignObj.path);

  const labelBaselineObj = props.properties.find((d) => d.prop === 'label_baseline');
  const labelBaseline = getValueFromNestedPath(spec, labelBaselineObj.path);

  const labelColorObj = props.properties.find((d) => d.prop === 'label_color');
  const labelColor = getValueFromNestedPath(spec, labelColorObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Form.Item name={labelAlignObj.path} label="Align">
        <Select
          onChange={(value) =>
            dispatch({
              type: SET_LEGEND_LABEL_POSITION,
              payload: { value: value, path: labelAlignObj.path },
              chart: 'shared',
            })
          }
        >
          <Option value="left">Left</Option>
          <Option value="right">Right</Option>
          <Option value="center">Center</Option>
        </Select>
      </Form.Item>

      <Form.Item name={labelBaselineObj.path} label="Baseline">
        <Select
          onChange={(value) =>
            dispatch({
              type: SET_LEGEND_LABEL_BASELINE,
              payload: { value: value, path: labelBaselineObj.path },
              chart: 'shared',
            })
          }
        >
          <Option value="top">Top</Option>
          <Option value="bottom">Bottom</Option>
          <Option value="middle">Center</Option>
        </Select>
      </Form.Item>

      <Form.Item name={labelColorObj.path} label="Color">
        <Input
          type="color"
          onChange={(e) =>
            dispatch({
              type: SET_LEGEND_LABEL_COLOR,
              payload: { value: e.target.value, path: labelColorObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>
    </div>
  );
}

export default LegendLabel;
