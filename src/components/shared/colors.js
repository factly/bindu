import React from "react";
import { Input } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

function Dimensions() {
	const spec = useSelector(state => state.chart.spec);
	const colors = spec.layer[0].encoding.color.scale.range;


	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<div className="item-container">
			<label htmlFor="">Colors</label>
			<div>
				{colors.map((d, i) => <Input type="color" value={d} key={i} onChange={(e) => dispatch({type: "set-color", index: i, value: e.target.value, chart: "shared"})}></Input>)}
			</div>
		</div>
    </div>
  );
}

export default Dimensions;