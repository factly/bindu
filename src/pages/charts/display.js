import React from "react";
import { connect } from "react-redux";

import * as vega from "vega";
import { compile } from "vega-lite";
import { deepEqual } from "vega-lite/build/src/util";

const mapStateToProps = state => {
  return { spec: state.chart.spec, templates: state.templates };
};

class Chart extends React.Component {

  componentDidMount() {
    this.renderVega();
  }

  componentDidUpdate(prevProps) {
    if (!deepEqual(prevProps.spec, this.props.spec)) {
      this.renderVega();
    }
  }

  getSpec = () => {
    let { spec } = this.props;

    return compile(spec).spec;
  };

  renderVega = () => {
    if(Object.keys(this.props.spec).length){
      const spec = this.getSpec();
      let runtime = vega.parse(spec);
      const loader = vega.loader();
      this.view = new vega.View(runtime, {
        loader,
      }).hover();

      this.view
        .logLevel(vega.Warn)
        .renderer("svg")
        .initialize(this.refs.chart)
        .runAsync();
    }
  };

  download = () => {
    this.view
      .toImageURL("png")
      .then(function (url) {
        var link = document.createElement("a");
        link.setAttribute("href", url);
        link.setAttribute("target", "_blank");
        link.setAttribute("download", "vega-export.png");
        link.dispatchEvent(new MouseEvent("click"));
      })
      .catch(function (error) {
        /* error handling */
      });
  };

  openInEditor = (e) => {
    const url = "http://localhost:8080";
    const editor = window.open(url);
    const wait = 10000;
    const step = 250;
    const data = {
      spec: JSON.stringify(this.state.spec, null, 4),
      config: {},
      mode: "vega-lite",
      renderer: this.state.renderer,
    };
    const { origin } = new URL(url);
    // eslint-disable-next-line no-bitwise
    let count = ~~(wait / step);

    function listen(evt) {
      if (evt.source === editor) {
        count = 0;
        window.removeEventListener("message", listen, false);
      }
    }
    window.addEventListener("message", listen, false);

    // send message
    // periodically resend until ack received or timeout
    function send() {
      if (count <= 0) {
        return;
      }
      editor.postMessage(data, origin);
      setTimeout(send, step);
      count -= 1;
    }
    setTimeout(send, step);
  };

  render() {
    return (
        <div className="chart-container" ref="chart"></div>
    );
  }
}

export default connect(
  mapStateToProps
)(Chart);

/*<div className="download-container">
          <button type="button" onClick={this.download}>
            Download
          </button>
          <button type="button" onClick={this.openInEditor}>
            Open in Editor
          </button>
        </div>*/
// function Chart(props){
// 	const chartEl = useRef(null);

// 	const options = ChartConfigs[props.path].options;
// 	const spec = ChartConfigs[props.path].spec;

// 	function renderVega() {
// 		const { renderer } = spec;

// 	    let runtime = vega.parse(compile(spec).spec);
// 	    const loader = vega.loader();
// 	    const view = new vega.View(runtime, {
// 	      loader,
// 	    }).hover();
// 		debugger;
// 	    view.logLevel(vega.Warn).renderer(renderer).initialize(chartEl).runAsync();
// 	}

// 	useEffect(() => {

// 	    renderVega();

// 	  });

// 	return (
// 		<div ref={chartEl}></div>
// 	);
// }

// export default Chart;