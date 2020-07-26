import React, { useEffect } from 'react';

import { Collapse } from 'antd';
import ChartProperties from '../../../components/shared/chart_properties.js';
import TreeMap from '../../../components/shared/tree_map.js';
import { useDispatch } from 'react-redux';

import Spec from './default.json';
const { Panel } = Collapse;

function BarChartRace() {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch({ type: 'set-config', value: Spec, mode: 'custom' });
  }, [dispatch]);

  return null;
}

export default BarChartRace;
