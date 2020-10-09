import React from 'react';
import { Input, Select, Form, InputNumber } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_LEGEND_POSITION,
  SET_LEGEND_TITLE,
  SET_LEGEND_BACKGROUND,
  SET_LEGEND_SYMBOL,
  SET_LEGEND_SYMBOL_SIZE,
} from '../../constants/legend.js';

const { Option } = Select;

function Legend(props) {
  const spec = useSelector((state) => state.chart.spec);
  // const {title, fillColor, symbolType, symbolSize, orient} = layer.encoding.color.legend;
  const titleObj = props.properties.find((d) => d.prop === 'title');
  const title = getValueFromNestedPath(spec, titleObj.path);

  const fillColorObj = props.properties.find((d) => d.prop === 'fill_color');
  const fillColor = getValueFromNestedPath(spec, fillColorObj.path);

  const symbolTypeObj = props.properties.find((d) => d.prop === 'symbol_type');
  const symbolType = getValueFromNestedPath(spec, symbolTypeObj.path);

  const symbolSizeObj = props.properties.find((d) => d.prop === 'symbol_size');
  const symbolSize = getValueFromNestedPath(spec, symbolSizeObj.path);

  const orientObj = props.properties.find((d) => d.prop === 'orient');
  const orient = getValueFromNestedPath(spec, orientObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Form.Item name={orientObj.path} label="Position">
        <Select
          onChange={(value) =>
            dispatch({
              type: SET_LEGEND_POSITION,
              payload: { value: value, path: orientObj.path },
              chart: 'shared',
            })
          }
        >
          <Option value="left">Left</Option>
          <Option value="right">Right</Option>
          <Option value="top">Top</Option>
          <Option value="bottom">Bottom</Option>
          <Option value="top-left">Top Left</Option>
          <Option value="top-right">Top Right</Option>
          <Option value="bottom-left">Bottom Left</Option>
          <Option value="bottom-right">Bottom Right</Option>
        </Select>
      </Form.Item>

      <Form.Item name={titleObj.path} label="Title">
        <Input
          placeholder="Title"
          type="text"
          onChange={(e) =>
            dispatch({
              type: SET_LEGEND_TITLE,
              payload: { value: e.target.value, path: titleObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>

      <Form.Item name={fillColorObj.path} label="Background Color">
        <Input
          type="color"
          onChange={(e) =>
            dispatch({
              type: SET_LEGEND_BACKGROUND,
              payload: { value: e.target.value, path: fillColorObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>

      <Form.Item name={symbolTypeObj.path} labe="Symbol">
        <Select
          onChange={(value) =>
            dispatch({
              type: SET_LEGEND_SYMBOL,
              payload: { value: value, path: symbolTypeObj.path },
              chart: 'shared',
            })
          }
        >
          <Option value="circle">Circle</Option>
          <Option value="square">Square</Option>
          <Option value="cross">Cross</Option>
          <Option value="diamond">Diamond</Option>
          <Option value="triangle-up">Triangle Up</Option>
          <Option value="triangle-down">Triangle Down</Option>
          <Option value="triangle-right">Triangle Right</Option>
          <Option value="triangle-left">Triangle Left</Option>
        </Select>
      </Form.Item>

      <Form.Item name={symbolSizeObj.path} label="Symbol Size">
        <InputNumber
          placeholder="Symbol Size"
          onChange={(value) =>
            dispatch({
              type: SET_LEGEND_SYMBOL_SIZE,
              payload: { value, path: symbolSizeObj.path },
              chart: 'shared',
            })
          }
        />
      </Form.Item>
    </div>
  );
}

export default Legend;
