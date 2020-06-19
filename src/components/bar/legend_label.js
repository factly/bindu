import React from "react";
import { Input, Select } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

const { Option } = Select;

function LegendLabel() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const {labelAlign, labelBaseline, labelColor} = layer.encoding.color.legend;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Align</label>
			<Select value={labelAlign} onChange = {(value) => dispatch({type: "set-legend-label-position", value: value, chart: "bar"})}>
		      <Option value="left">Left</Option>
		      <Option value="right">Right</Option>
		      <Option value="center">Center</Option>
		    </Select>
		</div>
		<div className="item-container">
			<label htmlFor="">Baseline</label>
			<Select value={labelBaseline} onChange = {(value) => dispatch({type: "set-legend-label-baseline", value: value, chart: "bar"})}>
		      <Option value="top">Top</Option>
		      <Option value="bottom">Bottom</Option>
		      <Option value="middle">Center</Option>
		    </Select>
		</div>
		<div className="item-container">
			<label htmlFor="">Color</label>
			<Input value={labelColor} type="color" onChange={(e) => dispatch({type: "set-legend-label-color", value: e.target.value, chart: "bar"})} />
		</div>
    </div>
  );
}

export default LegendLabel;