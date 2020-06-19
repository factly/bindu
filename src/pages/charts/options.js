import React from "react";
import { useSelector } from 'react-redux';
import GroupedBarChartOptions from "./grouped_bar/index.js";
import LineChartOptions from "./line/index.js";
import PieChartOptions from "./pie/index.js";
import { useParams } from "react-router-dom";

function OptionComponent(){
	let { id } = useParams();
  	const templates = useSelector(state => state.templates);
	const selectedChartName = templates.options[id].name;
	switch (selectedChartName) {
	  case 'Grouped Bar Chart':
	    return <GroupedBarChartOptions />
	  case 'Line Chart':
	    return <LineChartOptions />
	  case 'Pie Chart':
	    return <PieChartOptions />
	  default:
	    return null
	}
}

export default OptionComponent;