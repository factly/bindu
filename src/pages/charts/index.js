import React from "react";
import Display from "./display.js";
import ChartOption from "./options.js";
import "./index.css";

import { useDispatch, useSelector } from 'react-redux';

import { Card, Tooltip, Modal, Button, Input } from 'antd';
import { SaveOutlined, SettingOutlined, EditOutlined } from '@ant-design/icons';


function Chart() {
	const showOptions = useSelector(state => state.chart.showOptions);
	const chartName = useSelector(state => state.chart.chartName);
	const isChartNameEditable = useSelector(state => state.chart.isChartNameEditable);
	// const openCopyModal = useSelector(state => state.chart.openCopyModal);
	const dispatch = useDispatch();
	const actions = [{
		name: "Customize",
		Icon: SettingOutlined,
		onClick: () => dispatch({type: "set-options"})
	},
	{
		name: "Save",
		Icon: SaveOutlined
	}];

	const IconSize = 20;

	const actionsList = <div className="extra-actions-container"><ul>{actions.map(item => <li key={item.name} onClick={item.onClick}><Tooltip title={item.name}>{<item.Icon style={{fontSize: IconSize}}/>}</Tooltip></li> )}</ul></div>;
	let titleComponent;
	if (isChartNameEditable){
		titleComponent = <div className="chart-name-editable-container"><Input onPressEnter={() => dispatch({type: "edit-chart-name", value: false})} value={chartName} onChange={(e) => dispatch({type: "set-chart-name", value: e.target.value})}/> <Button style={{padding: "4px 0px"}} size="medium" onClick={() => dispatch({type: "edit-chart-name", value: false})} type="primary">Save</Button> </div>;
	} else {
		titleComponent = <div className="chart-name-container"><label className="chart-name">{chartName}</label><EditOutlined style={{fontSize: IconSize}} onClick={() => dispatch({type: "edit-chart-name", value: true})}/></div>;
	}
  return (
    <Card title={titleComponent} extra={actionsList} bodyStyle={{overflow: "hidden", display: "flex", "padding": "0px"}}>
      	<div className="display-container" style={{width: "100%"}}>
			<Display />
		</div>
		<div className="option-container" style={{right: showOptions ? "0": "-250px"}}>
  			<ChartOption />
		</div>
    </Card>
  );
}

export default Chart;