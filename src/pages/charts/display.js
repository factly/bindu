import React from 'react';
import * as vega from 'vega';
import { compile } from 'vega-lite';

import BarChartRace from './bar_chart_race/chart.js';

function Chart({ form, setView = () => {} }) {
  const spec = form.getFieldValue();
  const refContainer = React.useRef(null);

  const getSpec = () => {
    switch (spec.mode) {
      case 'vega':
        return spec;
      case 'vega-lite':
        return compile(spec).spec;
      default:
        return spec;
    }
  };

  const renderVega = () => {
    if (Object.keys(spec).length) {
      const spec = getSpec();
      let runtime = vega.parse(spec);
      const loader = vega.loader();
      const view = new vega.View(runtime, {
        loader,
      }).hover();

      view.logLevel(vega.Warn).renderer('svg').initialize(refContainer.current).runAsync();
      setView(view);
    }
  };

  const renderCustomChart = () => {
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

  const renderChart = () => {
    switch (spec.mode) {
      case 'vega':
        return renderVega();
      case 'vega-lite':
        return renderVega();
      case 'custom':
        return renderCustomChart();
      default:
        return spec;
    }
  };

  React.useEffect(() => {
    renderChart();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [spec]);

  return <div style={{ height: 'inherit', overflow: 'auto' }} ref={refContainer}></div>;
}

export default Chart;
