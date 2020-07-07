import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';


function Segment() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const {outerRadius, innerRadius, cornerRadius, padAngle} = layer.mark;


	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Outer Radius</label>
			</Col>
			<Col span={12}>
				<Input value={outerRadius} placeholder="title" type="number" onChange={(e) => dispatch({type: "set-outer-radius", value: e.target.value, chart: "pie"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Doughnut hole</label>
			</Col>
			<Col span={12}>
				<Input value={innerRadius} placeholder="title" type="number" onChange={(e) => dispatch({type: "set-inner-radius", value: e.target.value, chart: "pie"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Corner Curve</label>
			</Col>
			<Col span={12}>
				<Input value={cornerRadius} placeholder="width" type="number" onChange={(e) => dispatch({type: "set-corner-radius", value: e.target.value, chart: "pie"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Padding Angle</label>
			</Col>
			<Col span={12}>
				<Input min="0" max="0.5" step="0.025" value={padAngle} placeholder="height" type="number" onChange={(e) => dispatch({type: "set-padding-angle", value: e.target.value, chart: "pie"})}/>
			</Col>
		</Row>
    </div>
  );
}

export default Segment;