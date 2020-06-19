import React from "react";
import { Input, Select } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

const { Option } = Select;

function Legend() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const {title, fillColor, symbolType, symbolSize, orient} = layer.encoding.color.legend;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Position</label>
			<Select value={orient} onChange = {(value) => dispatch({type: "set-legend-position", value: value, chart: "bar"})}>
		      <Option value="left">Left</Option>
		      <Option value="right">Right</Option>
		      <Option value="top">Top</Option>
		      <Option value="bottom">Bottom</Option>
		      <Option value="top-left">Top Left</Option>
		      <Option value="top-right">Top Right</Option>
		      <Option value="bottom-left">Bottom Left</Option>
		      <Option value="bottom-right">Bottom Right</Option>
		    </Select>
		</div>
		<div className="item-container">
			<label htmlFor="">Title</label>
			<Input value={title} placeholder="Title" type="text" onChange={(e) => dispatch({type: "set-legend-title", value: e.target.value, chart: "bar"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Background Color</label>
			<Input value={fillColor} type="color" onChange={(e) => dispatch({type: "set-legend-background", value: e.target.value, chart: "bar"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Symbol</label>
			<Select value={symbolType} onChange = {(value) => dispatch({type: "set-legend-symbol", value: value, chart: "bar"})}>
		      <Option value="circle">Circle</Option>
		      <Option value="square">Square</Option>
		      <Option value="cross">Cross</Option>
		      <Option value="diamond">Diamond</Option>
		      <Option value="triangle-up">Triangle Up</Option>
		      <Option value="triangle-down">Triangle Down</Option>
		      <Option value="triangle-right">Triangle Right</Option>
		      <Option value="triangle-left">Triangle Left</Option>
		    </Select>
		</div>
		<div className="item-container">
			<label htmlFor="">Symbol Size</label>
			<Input value={symbolSize} placeholder="Symbol Size" type="number" onChange={(e) => dispatch({type: "set-legend-symbol-size", value: e.target.value, chart: "bar"})} />
		</div>
    </div>
  );
}

export default Legend;