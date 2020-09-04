# bindu-web
Bindu is a modern open source Data visualization platform built on Vega-Lite

**Releasability:** [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=alert_status)](https://sonarcloud.io/dashboard?id=factly_bindu-web)  
**Reliability:** [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=factly_bindu-web) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=bugs)](https://sonarcloud.io/dashboard?id=factly_bindu-web)  
**Security:** [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=security_rating)](https://sonarcloud.io/dashboard?id=factly_bindu-web) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=factly_bindu-web)  
**Maintainability:** [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=factly_bindu-web) [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=sqale_index)](https://sonarcloud.io/dashboard?id=factly_bindu-web) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=code_smells)](https://sonarcloud.io/dashboard?id=factly_bindu-web)  
**Other:** [![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=ncloc)](https://sonarcloud.io/dashboard?id=factly_bindu-web) [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=factly_bindu-web) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=factly_bindu-web&metric=coverage)](https://sonarcloud.io/dashboard?id=factly_bindu-web)  


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