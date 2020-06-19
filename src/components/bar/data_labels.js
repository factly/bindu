import React from "react";
import { Input } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

function DataLabels() {
	const spec = useSelector(state => state.chart.spec);
	const layersLength = spec.layer.length;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Enable</label>
			<Input type="checkbox" onChange={(e) => dispatch({type: "set-data-labels", value: e.target.checked, chart: "bar"})} />
		</div>
		{layersLength > 1 ? 
		<React.Fragment>
			<div className="item-container">
				<label htmlFor="">Color</label>
				<Input type="color" value={spec.layer[1].encoding.color.scale.range[0]} onChange={(e) => dispatch({type: "set-data-labels-color", value: e.target.value, chart: "bar"})} />
			</div>
			<div className="item-container">
				<label htmlFor="">Size</label>
				<Input type="number" value={spec.layer[1].mark.fontSize} onChange={(e) => dispatch({type: "set-data-labels-size", value: e.target.value, chart: "bar"})} />
			</div>
			<div className="item-container">
				<label htmlFor="">Format</label>
				<Input type="text" value={spec.layer[1].encoding.text.format} onChange={(e) => dispatch({type: "set-data-labels-format", value: e.target.value, chart: "bar"})} />
			</div>
		</React.Fragment> : null}
    </div>
  );
}

export default DataLabels;