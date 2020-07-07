import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';


function Bars() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const opacity = layer.encoding.opacity.value;
	const cornerRadius = layer.mark.cornerRadius;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Bar Opacity</label>
			</Col>
			<Col span={12}>
				<Input value={opacity} placeholder="Opacity" type="number" onChange={(e) => dispatch({type: "set-opacity", value: e.target.value, chart: "bar"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Corner radius</label>
			</Col>
			<Col span={12}>
				<Input value={cornerRadius} placeholder="Corner Radius" type="number" onChange={(e) => dispatch({type: "set-corner-radius", value: e.target.value, chart: "bar"})} />
			</Col>
		</Row>
    </div>
  );
}

export default Bars;