import React from "react";
import { Input, Select } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

const { Option } = Select;

function Dimensions() {
	const spec = useSelector(state => state.chart.spec);
	const mark = spec.layer[0].mark;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Enable</label>
			<Input type="checkbox" onChange={(e) => dispatch({type: "set-line-dots", value: e.target.checked, chart: "line"})} />
		</div>
		{
			mark.hasOwnProperty("point") ? 
				<React.Fragment>
					<div className="item-container">
						<label htmlFor="">Symbol</label>
						<Select value={mark.point.shape} onChange = {(value) => dispatch({type: "set-line-dot-shape", value: value, chart: "line"})}>
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
						<Input value={mark.point.size} placeholder="Symbol Size" type="number" onChange={(e) => dispatch({type: "set-line-dot-size", value: e.target.value, chart: "line"})} />
					</div>
					<div className="item-container">
						<label htmlFor="">Hollow</label>
						<Input value={mark.point.filled} type="checkbox" onChange={(e) => dispatch({type: "set-line-dots-hollow", value: e.target.checked, chart: "line"})} />
					</div>
				</React.Fragment>
			: null
		}
    </div>
  );
}

export default Dimensions;