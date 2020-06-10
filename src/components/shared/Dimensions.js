import React from "react";
import { Input } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { createSelector } from 'reselect';

import _ from "lodash";

function Dimensions() {
	const spec = useSelector(state => state.spec);
	const {width, height} = spec;


	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Width</label>
			<Input value={width} placeholder="width" type="number" onChange={(e) => dispatch({type: "set-width", value: e.target.value, chart: "bar"})} />
		</div>
		<div className="item-container">
			<label htmlFor="">Height</label>
			<Input value={height} placeholder="height" type="number" onChange={(e) => dispatch({type: "set-height", value: e.target.value})}/>
		</div>
    </div>
  );
}

export default Dimensions;
