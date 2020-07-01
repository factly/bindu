import React from "react";
import { useSelector } from 'react-redux';
import Area from "./area/index.js";
import StackedArea from "./stacked_area/index.js";
import Bar from "./bar/index.js";
import HorizontalBar from "./horizontal_bar/index.js";
import HorizontalStackBar from "./horizontal_stacked_bar/index.js";
import GroupedBar from "./grouped_bar/index.js";
import Line from "./line/index.js";
import Pie from "./pie/index.js";
import GroupedLine from "./grouped_line/index.js";
import LineBar from "./line_bar/index.js";
import DivergingBar from "./diverging_bar/index.js";
import { useParams } from "react-router-dom";

function OptionComponent(){
	let { id } = useParams();
  	const templates = useSelector(state => state.templates);
	const selectedChartName = templates.options[id].name;
	switch (selectedChartName) {
	  case 'Area Chart':
	    return <Area />
	  case 'Stacked Area Chart':
	    return <StackedArea />
	  case 'Bar Chart':
	    return <Bar />
	  case 'Grouped Bar Chart':
	    return <GroupedBar />
	  case 'Line Chart':
	    return <Line />
	  case 'Pie Chart':
	    return <Pie />
	  case 'Grouped Line Chart':
	    return <GroupedLine />
	  case 'Horizontal Bar Chart':
	    return <HorizontalBar />
	  case 'Horizontal Stack Bar Chart':
	    return <HorizontalStackBar />
	  case 'Line + Bar Chart':
	    return <LineBar />
	  case 'Diverging Bar Chart':
	    return <DivergingBar />
	  default:
	    return null
	}
}

export default OptionComponent;