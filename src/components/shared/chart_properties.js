import React from "react";
import { Input } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

function Dimensions() {
	const spec = useSelector(state => state.chart.spec);
	const {title, width, height, background} = spec;


	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Title</label>
			<Input value={title} placeholder="title" type="text" onChange={(e) => dispatch({type: "set-title", value: e.target.value, chart: "shared"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Width</label>
			<Input value={width} placeholder="width" type="number" onChange={(e) => dispatch({type: "set-width", value: e.target.value, chart: "shared"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Height</label>
			<Input value={height} placeholder="height" type="number" onChange={(e) => dispatch({type: "set-height", value: e.target.value, chart: "shared"})}/>
		</div>
		<div className="item-container">
			<label htmlFor="">Background</label>
			<Input value={background} type="color" onChange={(e) => dispatch({type: "set-background", value: e.target.value, chart: "shared"})}/>
		</div>
    </div>
  );
}

export default Dimensions;