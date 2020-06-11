import React from "react";
import { Input, Select } from 'antd';

import { useDispatch, useSelector } from 'react-redux';
import _ from "lodash";

const { Option } = Select;

function XAxis() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const {title, orient, format, labelColor} = layer.encoding.x.axis;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Title</label>
			<Input value={title} placeholder="Title" type="text" onChange={(e) => dispatch({type: "set-xaxis-title", value: e.target.value, chart: "bar"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Position</label>
			<Select value={orient} onChange = {(value) => dispatch({type: "set-xaxis-position", value: value, chart: "bar"})}>
		      <Option value="top">Top</Option>
		      <Option value="bottom">bottom</Option>
		    </Select>
		</div>
		<div className="item-container">
			<label htmlFor="">Label Format</label>
			<Input value={format} placeholder="Label Format" type="text" onChange={(e) => dispatch({type: "set-xaxis-label-format", value: e.target.value, chart: "bar"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Label Color</label>
			<Input value={labelColor} placeholder="Label Color" type="color" onChange={(e) => dispatch({type: "set-xaxis-label-color", value: e.target.value, chart: "bar"})} />
		</div>
    </div>
  );
}

export default XAxis;