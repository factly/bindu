import React from "react";
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';


const { Option } = Select;

function Dimensions() {
	const spec = useSelector(state => state.chart.spec);
	const { strokeWidth, opacity, interpolate, strokeDash} = spec.layer[0].mark;


	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Width</label>
			</Col>
			<Col span={12}>
				<Input value={strokeWidth} type="number" onChange={(e) => dispatch({type: "set-line-witdth", value: e.target.value, chart: "line"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Opacity</label>
			</Col>
			<Col span={12}>
				<Input value={opacity} type="number" onChange={(e) => dispatch({type: "set-line-opacity", value: e.target.value, chart: "line"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Line Curve</label>
			</Col>
			<Col span={12}>
				<Select value={interpolate} onChange = {(value) => dispatch({type: "set-line-curve", value: value, chart: "line"})}>
			      <Option value="linear">linear</Option>
			      <Option value="linear-closed">Linear Closed</Option>
			      <Option value="step">Step</Option>
			      <Option value="basis">Basis</Option>
			      <Option value="monotone">Monotone</Option>
			    </Select>
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Dash Width</label>
			</Col>
			<Col span={12}>
				<Input value={strokeDash} type="number" onChange={(e) => dispatch({type: "set-line-dashed", value: e.target.value, chart: "line"})} />
			</Col>
		</Row>
    </div>
  );
}

export default Dimensions;