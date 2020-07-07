import React from "react";
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';


const { Option } = Select;

function LegendLabel() {
	const spec = useSelector(state => state.chart.spec);
	const layer = spec.layer[0];
	const {labelAlign, labelBaseline, labelColor} = layer.encoding.color.legend;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Align</label>
			</Col>
			<Col span={12}>
				<Select value={labelAlign} onChange = {(value) => dispatch({type: "set-legend-label-position", value: value, chart: "shared"})}>
			      <Option value="left">Left</Option>
			      <Option value="right">Right</Option>
			      <Option value="center">Center</Option>
			    </Select>
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Baseline</label>
			</Col>
			<Col span={12}>
				<Select value={labelBaseline} onChange = {(value) => dispatch({type: "set-legend-label-baseline", value: value, chart: "shared"})}>
			      <Option value="top">Top</Option>
			      <Option value="bottom">Bottom</Option>
			      <Option value="middle">Center</Option>
			    </Select>
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Color</label>
			</Col>
			<Col span={12}>
				<Input value={labelColor} type="color" onChange={(e) => dispatch({type: "set-legend-label-color", value: e.target.value, chart: "shared"})} />
			</Col>
		</Row>
    </div>
  );
}

export default LegendLabel;