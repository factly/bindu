import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

function Colors() {
	const spec = useSelector(state => state.chart.spec);
	let colors = [];	
	if (spec.layer[0].encoding.color.hasOwnProperty("field")){
		colors = spec.layer[0].encoding.color.scale.range;
	} else {
		colors = [spec.layer[0].encoding.color.value];
	}

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Colors</label>
			</Col>
			<Col span={12}>
				{colors.map((d, i) => <Input type="color" value={d} key={i} onChange={(e) => dispatch({type: "set-color", index: i, value: e.target.value, chart: "shared"})}></Input>)}
			</Col>
		</Row>
    </div>
  );
}

export default Colors;