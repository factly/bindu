import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import {getValueFromNestedPath} from "../../utils/index.js";

import {SET_BAR_OPACITY, SET_BAR_CORNER_RADIUS} from "../../constants/bars.js";
function Bars(props) {
	const spec = useSelector(state => state.chart.spec);
	const opacityObj = props.properties.find(d => d.prop === "opacity");
	const opacity = getValueFromNestedPath(spec, opacityObj.path);

	const cornerRadiusObj = props.properties.find(d => d.prop === "corner_radius");
	const cornerRadius = getValueFromNestedPath(spec, cornerRadiusObj.path);

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Bar Opacity</label>
			</Col>
			<Col span={12}>
				<Input value={opacity} min={0} max={1} step={0.05} placeholder="Opacity" type="number" onChange={(e) => dispatch({type: SET_BAR_OPACITY, payload: {value: e.target.value, path: opacityObj.path}, chart: "shared"})} />
			</Col>
		</Row>
		<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Corner radius</label>
			</Col>
			<Col span={12}>
				<Input min={0} value={cornerRadius} placeholder="Corner Radius" type="number" onChange={(e) => dispatch({type: SET_BAR_CORNER_RADIUS, payload: {value: e.target.value, path: cornerRadiusObj.path}, chart: "shared"})} />
			</Col>
		</Row>
    </div>
  );
}

export default Bars;