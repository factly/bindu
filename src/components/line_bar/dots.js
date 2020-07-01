import React from "react";
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

const { Option } = Select;

function Dimensions() {
	const spec = useSelector(state => state.chart.spec);
	const mark = spec.layer[1].mark;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Enable</label>
			</Col>
			<Col span={12}>
				<Input type="checkbox" onChange={(e) => dispatch({type: "set-line-dots", value: e.target.checked, chart: "line_bar"})} />
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
							<Select value={mark.point.shape} onChange = {(value) => dispatch({type: "set-line-dot-shape", value: value, chart: "line_bar"})}>
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
							<Input value={mark.point.size} placeholder="Symbol Size" type="number" onChange={(e) => dispatch({type: "set-line-dot-size", value: e.target.value, chart: "line_bar"})} />
						</Col>
					</Row>
					<Row gutter={[0, 12]}>
						<Col span={12}>
							<label htmlFor="">Hollow</label>
						</Col>
						<Col span={12}>
							<Input value={mark.point.filled} type="checkbox" onChange={(e) => dispatch({type: "set-line-dots-hollow", value: e.target.checked, chart: "line_bar"})} />
						</Col>
					</Row>
				</React.Fragment>
			: null
		}
    </div>
  );
}

export default Dimensions;