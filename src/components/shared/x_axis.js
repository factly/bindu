import React from "react";
import { Input, Select, Row, Col } from 'antd';

import { useDispatch, useSelector } from 'react-redux';
import _ from "lodash";

const { Option } = Select;

function XAxis() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const {title, orient, format, labelColor} = layer.encoding.x.axis;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Title</label>
			</Col>
			<Col span={12}>
				<Input value={title} placeholder="Title" type="text" onChange={(e) => dispatch({type: "set-xaxis-title", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Position</label>
			</Col>
			<Col span={12}>
				<Select value={orient} onChange = {(value) => dispatch({type: "set-xaxis-position", value: value, chart: "shared"})}>
			      <Option value="top">Top</Option>
			      <Option value="bottom">bottom</Option>
			    </Select>
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Label Format</label>
			</Col>
			<Col span={12}>
				<Input value={format} placeholder="Label Format" type="text" onChange={(e) => dispatch({type: "set-xaxis-label-format", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Label Color</label>
			</Col>
			<Col span={12}>
				<Input value={labelColor} placeholder="Label Color" type="color" onChange={(e) => dispatch({type: "set-xaxis-label-color", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
    </div>
  );
}

export default XAxis;