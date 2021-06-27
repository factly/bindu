import React from 'react';
import * as vega from 'vega';
import { compile } from 'vega-lite';

import BarChartRace from './bar_chart_race/chart.js';

function Chart({ spec, mode, setView = () => {} }) {
  const refContainer = React.useRef(null);
  const getSpec = (spec, new_mode) => {
    switch (new_mode) {
      case 'vega':
        return spec;
      case 'vega-lite':
        return compile(spec).spec;
      default:
        return spec;
    }
  };

  const renderVega = (spec, new_mode) => {
    if (Object.keys(spec).length) {
      try {
        const runtimeSpec = getSpec(spec, new_mode);
        let runtime = vega.parse(runtimeSpec);
        const loader = vega.loader();
        const view = new vega.View(runtime, {
          loader,
        }).hover();

        view.logLevel(vega.Warn).renderer('svg').initialize(refContainer.current).runAsync();
        setView(view);
      } catch (error) {
        refContainer.current.innerHTML = error.message;
      }
    }
  };

  const renderCustomChart = (spec, new_mode) => {
    const type = spec.type;
    switch (type) {
      case 'bar-chart-race':
        var customChart = new BarChartRace(refContainer.current, spec);
        customChart.render();
        break;
      default:
        return null;
    }
  };

  const renderChart = (spec, new_mode) => {
    switch (new_mode) {
      case 'vega':
        return renderVega(spec, new_mode);
      case 'vega-lite':
        return renderVega(spec, new_mode);
      case 'custom':
        return renderCustomChart(spec, new_mode);
      default:
        return spec;
    }
  };

  React.useEffect(() => {
    renderChart(spec, mode);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [spec, mode]);

  return <div style={{ height: 'inherit', overflow: 'auto' }} ref={refContainer}></div>;
}

export default Chart;
