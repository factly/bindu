import React from "react";
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';


const { Option } = Select;

function Legend() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const {title, fillColor, symbolType, symbolSize, orient} = layer.encoding.color.legend;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Position</label>
			</Col>
			<Col span={12}>
				<Select value={orient} onChange = {(value) => dispatch({type: "set-legend-position", value: value, chart: "shared"})}>
			      <Option value="left">Left</Option>
			      <Option value="right">Right</Option>
			      <Option value="top">Top</Option>
			      <Option value="bottom">Bottom</Option>
			      <Option value="top-left">Top Left</Option>
			      <Option value="top-right">Top Right</Option>
			      <Option value="bottom-left">Bottom Left</Option>
			      <Option value="bottom-right">Bottom Right</Option>
			    </Select>
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Title</label>
			</Col>
			<Col span={12}>
				<Input value={title} placeholder="Title" type="text" onChange={(e) => dispatch({type: "set-legend-title", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Background Color</label>
			</Col>
			<Col span={12}>
				<Input value={fillColor} type="color" onChange={(e) => dispatch({type: "set-legend-background", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Symbol</label>
			</Col>
			<Col span={12}>
				<Select value={symbolType} onChange = {(value) => dispatch({type: "set-legend-symbol", value: value, chart: "shared"})}>
			      <Option value="circle">Circle</Option>
			      <Option value="square">Square</Option>
			      <Option value="cross">Cross</Option>
			      <Option value="diamond">Diamond</Option>
			      <Option value="triangle-up">Triangle Up</Option>
			      <Option value="triangle-down">Triangle Down</Option>
			      <Option value="triangle-right">Triangle Right</Option>
			      <Option value="triangle-left">Triangle Left</Option>
			    </Select>
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Symbol Size</label>
			</Col>
			<Col span={12}>
				<Input value={symbolSize} placeholder="Symbol Size" type="number" onChange={(e) => dispatch({type: "set-legend-symbol-size", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
    </div>
  );
}

export default Legend;