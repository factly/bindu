import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

function Dimensions() {
	const spec = useSelector(state => state.chart.spec);
	const {title, width, height, background} = spec;


	const dispatch = useDispatch();

  return (
    <div className="property-container">
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Title</label>
			</Col>
			<Col span={12}>
				<Input value={title} placeholder="title" type="text" onChange={(e) => dispatch({type: "set-title", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Width</label>
			</Col>
			<Col span={12}>
				<Input value={width} placeholder="width" type="number" onChange={(e) => dispatch({type: "set-width", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Height</label>
			</Col>
			<Col span={12}>
				<Input value={height} placeholder="height" type="number" onChange={(e) => dispatch({type: "set-height", value: e.target.value, chart: "shared"})}/>
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Background</label>
			</Col>
			<Col span={12}>
				<Input value={background} type="color" onChange={(e) => dispatch({type: "set-background", value: e.target.value, chart: "shared"})}/>
			</Col>
		</Row>
    </div>
  );
}

export default Dimensions;