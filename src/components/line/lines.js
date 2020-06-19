import React from "react";
import { Input, Select } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

const { Option } = Select;

function Dimensions() {
	const spec = useSelector(state => state.chart.spec);
	const { strokeWidth, opacity, interpolate, strokeDash} = spec.layer[0].mark;


	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Width</label>
			<Input value={strokeWidth} type="number" onChange={(e) => dispatch({type: "set-line-witdth", value: e.target.value, chart: "line"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Opacity</label>
			<Input value={opacity} type="number" onChange={(e) => dispatch({type: "set-line-opacity", value: e.target.value, chart: "line"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Line Curve</label>
			<Select value={interpolate} onChange = {(value) => dispatch({type: "set-line-curve", value: value, chart: "line"})}>
		      <Option value="linear">linear</Option>
		      <Option value="linear-closed">Linear Closed</Option>
		      <Option value="step">Step</Option>
		      <Option value="basis">Basis</Option>
		      <Option value="monotone">Monotone</Option>
		    </Select>
		</div>
		<div className="item-container">
			<label htmlFor="">Dash Width</label>
			<Input value={strokeDash} type="number" onChange={(e) => dispatch({type: "set-line-dashed", value: e.target.value, chart: "line"})} />
		</div>
    </div>
  );
}

export default Dimensions;