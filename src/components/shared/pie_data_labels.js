import React from "react";
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import {SET_PIE_DATA_LABELS, SET_PIE_DATA_LABELS_SIZE, SET_PIE_DATA_LABELS_POSITION} from "../../constants/pie_data_labels.js";

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
				<Input type="checkbox" onChange={(e) => dispatch({type: SET_PIE_DATA_LABELS, value: e.target.checked, chart: "shared"})} />
			</Col>
		</Row>
		{layersLength > 1 ? 
		<React.Fragment>
			<Row gutter={[0, 12]}>
				<Col span={12}>
					<label htmlFor="">Size</label>
				</Col>
				<Col span={12}>
					<Input type="number" value={spec.layer[1].mark.fontSize} onChange={(e) => dispatch({type: SET_PIE_DATA_LABELS_SIZE, value: e.target.value, chart: "shared"})} />
				</Col>
			</Row>
			<Row gutter={[0, 12]}>
				<Col span={12}>
					<label htmlFor="">Position( from center)</label>
				</Col>
				<Col span={12}>
					<Input type="number" value={spec.layer[1].mark.radius} onChange={(e) => dispatch({type: SET_PIE_DATA_LABELS_POSITION, value: e.target.value, chart: "shared"})} />
				</Col>
			</Row>
		</React.Fragment> : null}
    </div>
  );
}

export default DataLabels;