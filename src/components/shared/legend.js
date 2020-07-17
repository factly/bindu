import React from 'react';
import { Input, Select, Row, Col } from 'antd';
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
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Position</label>
        </Col>
        <Col span={12}>
          <Select
            value={orient}
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
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Title</label>
        </Col>
        <Col span={12}>
          <Input
            value={title}
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
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Background Color</label>
        </Col>
        <Col span={12}>
          <Input
            value={fillColor}
            type="color"
            onChange={(e) =>
              dispatch({
                type: SET_LEGEND_BACKGROUND,
                payload: { value: e.target.value, path: fillColorObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Symbol</label>
        </Col>
        <Col span={12}>
          <Select
            value={symbolType}
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
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Symbol Size</label>
        </Col>
        <Col span={12}>
          <Input
            value={symbolSize}
            placeholder="Symbol Size"
            type="number"
            onChange={(e) =>
              dispatch({
                type: SET_LEGEND_SYMBOL_SIZE,
                payload: { value: e.target.value, path: symbolSizeObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
    </div>
  );
}

export default Legend;
