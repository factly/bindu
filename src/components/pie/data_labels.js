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
				<Input type="checkbox" onChange={(e) => dispatch({type: "set-data-labels", value: e.target.checked, chart: "pie"})} />
			</Col>
		</Row>
		{layersLength > 1 ? 
		<React.Fragment>
			<Row gutter={[0, 12]}>
				<Col span={12}>
					<label htmlFor="">Size</label>
				</Col>
				<Col span={12}>
					<Input type="number" value={spec.layer[1].mark.fontSize} onChange={(e) => dispatch({type: "set-data-labels-size", value: e.target.value, chart: "pie"})} />
				</Col>
			</Row>
			<Row gutter={[0, 12]}>
				<Col span={12}>
					<label htmlFor="">Position( from center)</label>
				</Col>
				<Col span={12}>
					<Input type="number" value={spec.layer[1].mark.radius} onChange={(e) => dispatch({type: "set-data-labels-position", value: e.target.value, chart: "pie"})} />
				</Col>
			</Row>
		</React.Fragment> : null}
    </div>
  );
}

export default DataLabels;