import React from "react";
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from "../../utils/index.js";
import {SET_LINE_DOTS, SET_LINE_DOT_SHAPE, SET_LINE_DOT_SIZE, SET_LINE_DOTS_HOLLOW} from "../../constants/dots.js";

const { Option } = Select;

function Dots(props) {
	const spec = useSelector(state => state.chart.spec);
	const markObj = props.properties.find(d => d.prop === "mark");
	const mark = getValueFromNestedPath(spec, markObj.path);

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Enable</label>
			</Col>
			<Col span={12}>
				<Input type="checkbox" onChange={(e) => dispatch({type: SET_LINE_DOTS, payload: {value: e.target.checked, path: markObj.path}, chart: "shared"})} />
			</Col>
		</Row>
		{
			mark.hasOwnProperty("point") ? 
				<React.Fragment>
					<Row gutter={[0, 12]}>
						<Col span={12}>
							<label htmlFor="">Symbol</label>
						</Col>
						<Col span={12}>
							<Select value={mark.point.shape} onChange = {(value) => dispatch({type: SET_LINE_DOT_SHAPE, payload: {value: value, path: markObj.path}, chart: "shared"})}>
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
							<Input value={mark.point.size} placeholder="Symbol Size" type="number" onChange={(e) => dispatch({type: SET_LINE_DOT_SIZE, payload: {value: e.target.value, path: markObj.path}, chart: "shared"})} />
						</Col>
					</Row>
					<Row gutter={[0, 12]}>
						<Col span={12}>
							<label htmlFor="">Hollow</label>
						</Col>
						<Col span={12}>
							<Input value={mark.point.filled} type="checkbox" onChange={(e) => dispatch({type: SET_LINE_DOTS_HOLLOW, payload: {value: e.target.checked, path: markObj.path}, chart: "shared"})} />
						</Col>
					</Row>
				</React.Fragment>
			: null
		}
    </div>
  );
}

export default Dots;