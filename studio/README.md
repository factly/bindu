# bindu-studio

Bindu is a modern open source Data visualization platform built on Vega-Lite

**Releasability:** [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=alert_status)](https://sonarcloud.io/dashboard?id=factly_bindu-studio)  
**Reliability:** [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=factly_bindu-studio) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=bugs)](https://sonarcloud.io/dashboard?id=factly_bindu-studio)  
**Security:** [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=security_rating)](https://sonarcloud.io/dashboard?id=factly_bindu-studio) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=factly_bindu-studio)  
**Maintainability:** [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=factly_bindu-studio) [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=sqale_index)](https://sonarcloud.io/dashboard?id=factly_bindu-studio) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=code_smells)](https://sonarcloud.io/dashboard?id=factly_bindu-studio)  
**Other:** [![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=ncloc)](https://sonarcloud.io/dashboard?id=factly_bindu-studio) [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=factly_bindu-studio) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-studio&metric=coverage)](https://sonarcloud.io/dashboard?id=factly_bindu-studio)

### Folder Structure

- actions: Contains all the api calls of different components.
- components: Contains all the options which are included in the chart to customize the chart.
  Eg <ChartProperties properties= [
  {
  prop: 'title', //name of the property
  path: ['title'], //path to update the property in the spec
  },
  {
  prop: 'width',
  path: ['width'],
  },
  {
  prop: 'height',
  path: ['height'],
  },
  {
  prop: 'background',
  path: ['background'],
  },
  ]
  />
- constants - Contains the consts for all the components.
- pages :
  - charts :
    - [chart_name]:
      index.js: Specifies the options to customize the chart using components folder.
      default.json: Specifies the spec of the chart
    - index.js: Container for the chart and options

# How to add a chart:

- Add a new folder in pages/charts.
- Add default.json: specifying the spec of the chart
- Add index.js; specifying the options to customize the chart
- Add the chart in the redux store: src/reducer/templates.js
- Add a switch case for the new chart in src/pages/charts/options.js

## Usage Instructions

To run the server locally, you must first install the dependencies and then launch a local web server

1. Install the dependencies:

```
$ npm install
```

2. Start the server:

```
$ npm start
```
