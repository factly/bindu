let binduChartsURL = 'http://localhost:7002/charts/';
let embedElements = document.getElementsByClassName('factly-embed');

var scripts = [
  'https://cdn.jsdelivr.net/npm/vega@5',
  'https://cdn.jsdelivr.net/npm/vega-lite@4',
  'https://cdn.jsdelivr.net/npm/vega-embed@6',
];

// RECURSIVE LOAD SCRIPTS
const load_scripts = (urls, final_callback, index = 0) => {
  if (typeof urls[index + 1] === 'undefined') {
    load_script(urls[index], final_callback);
  } else {
    load_script(urls[index], function () {
      load_scripts(urls, final_callback, index + 1);
    });
  }
};

// LOAD SCRIPT
const load_script = (url, callback) => {
  var script = document.createElement('script');
  script.type = 'text/javascript';
  if (script.readyState) {
    // IE
    script.onreadystatechange = function () {
      if (script.readyState === 'loaded' || script.readyState === 'complete') {
        script.onreadystatechange = null;
        callback();
      }
    };
  } // Others
  else {
    script.onload = function () {
      callback();
    };
  }
  script.src = url;
  document.getElementsByTagName('head')[0].appendChild(script);
};

// embedjs function
var main = function () {
  for (let i = 0; i < embedElements.length; i++) {
    let chartID = embedElements[i].getAttribute('data-src');
    fetch(binduChartsURL + chartID)
      .then((res) => {
        return res.json();
      })
      .then((json) => {
        embedElements[i].setAttribute('id', 'embed' + i);
        vegaEmbed('#embed' + i, json).catch(console.error);
      })
      .catch((err) => {
        console.log(err);
      });
  }
};

load_scripts(scripts, main);