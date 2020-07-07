import React, { useEffect } from "react";

import { Collapse } from "antd";
import ChartProperties from "../../../components/shared/chart_properties.js";
import Colors from "../../../components/shared/colors.js";
import Legend from "../../../components/shared/legend.js";
import LegendLabel from "../../../components/shared/legend_label.js";

import Segment from "../../../components/pie/segment.js";
import DataLabels from "../../../components/pie/data_labels.js";

import { useDispatch } from 'react-redux';

import Spec from "./default.json";
const { Panel } = Collapse;

function PieChar() {
	const dispatch = useDispatch();
	useEffect(() => {
		dispatch({type: "set-config", value: Spec});
	}, [dispatch]);

  const properties = [{
    name: "Chart Properties",
    component: <ChartProperties /> 
  },{
    name: "Segment",
    component: <Segment /> 
  },{
    name: "Colors",
    component: <Colors /> 
  },{
    name: "Legend",
    component: <Legend /> 
  },{
    name: "Legend Label",
    component: <LegendLabel /> 
  },{
    name: "Data Labels",
    component: <DataLabels /> 
  }];

  return (
    <div className="options-container">
    		<Collapse
          className="option-item-collapse"
        >
        {
          properties.map((d, i) => {
            return (
              <Panel className="option-item-panel" header={d.name} key={i}>
                {d.component}
              </Panel>
            )
          })
        }
        </Collapse>
    </div>
  );
}

export default PieChar;
