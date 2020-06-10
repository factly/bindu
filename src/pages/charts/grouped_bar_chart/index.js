import React, { useEffect } from "react";
import { useSelector } from 'react-redux';

import { Collapse } from "antd";
import Dimensions from "../../../components/shared/Dimensions.js";
import { useDispatch } from 'react-redux';

import Spec from "./default.json";
const { Panel } = Collapse;


function GroupedBarChart() {
  // const counter = useSelector(state => state.counter)
	const dispatch = useDispatch();
	useEffect(() => {
		dispatch({type: "set-config", value: Spec});
	}, []);

  return (
    <div className="options-container">
		<Collapse
          className="option-item-collapse"
        >
          <Panel className="option-item-panel" header={"Chart Properties"} key="1">
            <Dimensions />
          </Panel>
        </Collapse>
    </div>
  );
}

export default GroupedBarChart;
