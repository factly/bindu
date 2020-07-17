import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import {getValueFromNestedPath} from "../../utils/index.js";

import {SET_PIE_OUTER_RADIUS, SET_PIE_INNER_RADIUS, SET_PIE_CORNER_RADIUS, SET_PIE_PADDING_ANGLE} from "../../constants/segment.js";
function Segment(props) {
	const spec = useSelector(state => state.chart.spec);
	// const {outerRadius, innerRadius, cornerRadius, padAngle} = layer.mark;
	
	const outerRadiusObj = props.properties.find(d => d.prop === "outer_radius");
	const outerRadius = getValueFromNestedPath(spec, outerRadiusObj.path);

	const innerRadiusObj = props.properties.find(d => d.prop === "inner_radius");
	const innerRadius = getValueFromNestedPath(spec, innerRadiusObj.path);

	const cornerRadiusObj = props.properties.find(d => d.prop === "corner_radius");
	const cornerRadius = getValueFromNestedPath(spec, cornerRadiusObj.path);

	const padAngleObj = props.properties.find(d => d.prop === "pad_angle");
	const padAngle = getValueFromNestedPath(spec, padAngleObj.path);

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Outer Radius</label>
			</Col>
			<Col span={12}>
				<Input min={0} value={outerRadius} placeholder="title" type="number" onChange={(e) => dispatch({type: SET_PIE_OUTER_RADIUS, payload:{value: e.target.value, path: outerRadiusObj.path}, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Doughnut hole</label>
			</Col>
			<Col span={12}>
				<Input min={0} value={innerRadius} placeholder="title" type="number" onChange={(e) => dispatch({type: SET_PIE_INNER_RADIUS, payload:{value: e.target.value, path: innerRadiusObj.path}, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Corner Curve</label>
			</Col>
			<Col span={12}>
				<Input min={0} value={cornerRadius} placeholder="width" type="number" onChange={(e) => dispatch({type: SET_PIE_CORNER_RADIUS, payload:{value: e.target.value, path: cornerRadiusObj.path}, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Padding Angle</label>
			</Col>
			<Col span={12}>
				<Input min={0} max={0.5} step="0.025" value={padAngle} placeholder="height" type="number" onChange={(e) => dispatch({type: SET_PIE_PADDING_ANGLE, payload:{value: e.target.value, path: padAngleObj.path}, chart: "shared"})}/>
			</Col>
		</Row>
    </div>
  );
}

export default Segment;