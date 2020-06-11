import React from "react";
import { Input } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

function Bars() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const opacity = layer.encoding.opacity.value;
	const cornerRadius = layer.mark.cornerRadius;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Bar Opacity</label>
			<Input value={opacity} placeholder="Opacity" type="number" onChange={(e) => dispatch({type: "set-opacity", value: e.target.value, chart: "bar"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Corner radius</label>
			<Input value={cornerRadius} placeholder="Corner Radius" type="number" onChange={(e) => dispatch({type: "set-corner-radius", value: e.target.value, chart: "bar"})} />
		</div>
    </div>
  );
}

export default Bars;