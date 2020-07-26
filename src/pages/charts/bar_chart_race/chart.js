import * as d3 from 'd3';
import data from './data.json';
import './chart.css';
export default function BarChartRace(ref, chartSettings) {
  chartSettings.innerWidth = chartSettings.width - chartSettings.padding * 2;
  chartSettings.innerHeight = chartSettings.height - chartSettings.padding * 2;

  const chartDataSets = data;
  let chartTransition;
  let timerStart, timerEnd;
  let currentDataSetIndex = 0;
  let elapsedTime = chartSettings.duration;

  const svg = d3.select(ref).append('svg').attr('id', 'bar-chart-race');
  const chartContainer = svg.append('g').attr('class', 'chart-container');

  const xAxisContainer = chartContainer.append('g').attr('class', 'x-axis');
  const yAxisContainer = chartContainer.append('g').attr('class', 'y-axis');

  chartContainer.append('g').attr('class', 'columns');
  chartContainer.append('text').attr('class', 'chart-title');
  chartContainer.append('text').attr('class', 'current-date');

  const xAxisScale = d3.scaleLinear().range([0, chartSettings.innerWidth]);

  const yAxisScale = d3
    .scaleBand()
    .range([0, chartSettings.innerHeight])
    .padding(chartSettings.columnPadding);

  const colorScale = d3
    .scaleOrdinal()
    .domain([0, chartSettings.maximumModelCount])
    .range(d3.schemeCategory10);

  svg.attr('width', chartSettings.width).attr('height', chartSettings.height);

  chartContainer.attr('transform', `translate(${chartSettings.padding} ${chartSettings.padding})`);

  chartContainer
    .select('.current-date')
    .attr('transform', `translate(${chartSettings.innerWidth} ${chartSettings.innerHeight})`);

  d3.select('.chart-title')
    .attr('x', chartSettings.width / 2)
    .attr('y', -chartSettings.padding / 2)
    .text(chartSettings.title);

  function draw({ dataSet, date: currentDate }, transition) {
    const { innerHeight, ticksInXAxis, titlePadding } = chartSettings;
    const dataSetDescendingOrder = dataSet
      .sort(({ value: firstValue }, { value: secondValue }) => secondValue - firstValue)
      .slice(0, chartSettings.maximumModelCount);

    chartContainer.select('.current-date').text(currentDate);

    xAxisScale.domain([0, dataSetDescendingOrder[0].value]);
    yAxisScale.domain(dataSetDescendingOrder.map(({ name }) => name));

    xAxisContainer
      .transition(transition)
      .call(d3.axisTop(xAxisScale).ticks(ticksInXAxis).tickSize(-innerHeight));

    yAxisContainer.transition(transition).call(d3.axisLeft(yAxisScale).tickSize(0));

    // The general update Pattern in d3.js

    // Data Binding
    const barGroups = chartContainer
      .select('.columns')
      .selectAll('g.column-container')
      .data(dataSetDescendingOrder, ({ name }) => name);

    // Enter selection
    const barGroupsEnter = barGroups
      .enter()
      .append('g')
      .attr('class', 'column-container')
      .attr('transform', `translate(0,${innerHeight})`);

    barGroupsEnter
      .append('rect')
      .attr('class', 'column-rect')
      .attr('width', 0)
      .attr('height', yAxisScale.step() * (1 - chartSettings.columnPadding))
      .attr('fill', (d, i) => colorScale(i));

    barGroupsEnter
      .append('text')
      .attr('class', 'column-title')
      .attr('y', (yAxisScale.step() * (1 - chartSettings.columnPadding)) / 2)
      .attr('x', -titlePadding)
      .text(({ name }) => name);

    barGroupsEnter
      .append('text')
      .attr('class', 'column-value')
      .attr('y', (yAxisScale.step() * (1 - chartSettings.columnPadding)) / 2)
      .attr('x', titlePadding)
      .text(0);

    // Update selection
    const barUpdate = barGroupsEnter.merge(barGroups);

    barUpdate
      .transition(transition)
      .attr('transform', ({ name }) => `translate(0,${yAxisScale(name)})`)
      .attr('fill', 'normal');

    barUpdate
      .select('.column-rect')
      .transition(transition)
      .attr('width', ({ value }) => xAxisScale(value));

    barUpdate
      .select('.column-title')
      .transition(transition)
      .attr('x', ({ value }) => xAxisScale(value) - titlePadding);

    barUpdate
      .select('.column-value')
      .transition(transition)
      .attr('x', ({ value }) => xAxisScale(value) + titlePadding)
      .tween('text', function ({ value }) {
        const interpolateStartValue =
          elapsedTime === chartSettings.duration ? this.currentValue || 0 : +this.innerHTML;

        const interpolate = d3.interpolate(interpolateStartValue, value);
        this.currentValue = value;

        return function (t) {
          d3.select(this).text(Math.ceil(interpolate(t)));
        };
      });

    // Exit selection
    const bodyExit = barGroups.exit();

    bodyExit
      .transition(transition)
      .attr('transform', `translate(0,${innerHeight})`)
      .on('end', function () {
        d3.select(this).attr('fill', 'none');
      });

    bodyExit.select('.column-title').transition(transition).attr('x', 0);

    bodyExit.select('.column-rect').transition(transition).attr('width', 0);

    bodyExit
      .select('.column-value')
      .transition(transition)
      .attr('x', titlePadding)
      .tween('text', function () {
        const interpolate = d3.interpolate(this.currentValue, 0);
        this.currentValue = 0;

        return function (t) {
          d3.select(this).text(Math.ceil(interpolate(t)));
        };
      });

    return this;
  }

  async function render(index = 0) {
    currentDataSetIndex = index;
    timerStart = d3.now();

    chartTransition = chartContainer
      .transition()
      .duration(elapsedTime)
      .ease(d3.easeLinear)
      .on('end', () => {
        if (index < chartDataSets.length) {
          elapsedTime = chartSettings.duration;
          render(index + 1);
        } else {
          d3.select('button').text('Play');
        }
      })
      .on('interrupt', () => {
        timerEnd = d3.now();
      });

    if (index < chartDataSets.length) {
      draw(chartDataSets[index], chartTransition);
    }

    return this;
  }

  function stop() {
    d3.select(`ref`).selectAll('*').interrupt();

    return this;
  }

  function start() {
    elapsedTime -= timerEnd - timerStart;

    render(currentDataSetIndex);

    return this;
  }

  return {
    render,
    start,
    stop,
  };
}
