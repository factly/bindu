import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import _ from "lodash";

function DataLabels() {
	const spec = useSelector(state => state.chart.spec);
	const layersLength = spec.layer.length;

	const dispatch = useDispatch();

  return (
    <div className="property-container">
    	<Row gutter={[0, 12]}>
			<Col span={12}>
				<label htmlFor="">Enable</label>
			</Col>
			<Col span={12}>
				<Input type="checkbox" onChange={(e) => dispatch({type: "set-data-labels", value: e.target.checked, chart: "shared"})} />
			</Col>
		</Row>
		{layersLength > 1 ? 
		<React.Fragment>
			<Row gutter={[0, 12]}>
				<Col span={12}>
					<label htmlFor="">Color</label>
				</Col>
				<Col span={12}>
					<Input type="color" value={spec.layer[1].encoding.color.scale.range[0]} onChange={(e) => dispatch({type: "set-data-labels-color", value: e.target.value, chart: "shared"})} />
				</Col>
			</Row>
			<Row gutter={[0, 12]}>
				<Col span={12}>
					<label htmlFor="">Size</label>
				</Col>
				<Col span={12}>
					<Input type="number" value={spec.layer[1].mark.fontSize} onChange={(e) => dispatch({type: "set-data-labels-size", value: e.target.value, chart: "shared"})} />
				</Col>
			</Row>
			<Row gutter={[0, 12]}>
				<Col span={12}>
					<label htmlFor="">Format</label>
				</Col>
				<Col span={12}>
					<Input type="text" value={spec.layer[1].encoding.text.format} onChange={(e) => dispatch({type: "set-data-labels-format", value: e.target.value, chart: "shared"})} />
				</Col>
			</Row>
			<div className="item-container">
			</div>
		</React.Fragment> : null}
    </div>
  );
}

export default DataLabels;