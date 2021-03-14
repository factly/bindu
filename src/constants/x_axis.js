export const SET_XAXIS_TITLE = 'set-xaxis-title';
export const SET_XAXIS_POSITION = 'set-xaxis-position';
export const SET_XAXIS_LABEL_FORMAT = 'set-xaxis-label-format';
export const SET_XAXIS_LABEL_COLOR = 'set-xaxis-label-color';
export const aggregateOptions = [
  { value: 'count', name: 'Count' },
  { value: 'valid', name: 'Valid' },
  { value: 'values', name: 'Values' },
  { value: 'missing', name: 'Missing' },
  { value: 'distinct', name: 'Distinct' },
  { value: 'sum', name: 'Sum' },
  { value: 'product', name: 'Product' },
  { value: 'mean', name: 'Mean' },
  { value: 'average', name: 'Average' },
  { value: 'variance', name: 'Variance' },
  { value: 'variancep', name: 'Population Variance' },
  { value: 'stdev', name: 'Standard Deviation' },
  { value: 'stdevp', name: 'Population Standard Deviation' },
  { value: 'stderr', name: 'Standard Error' },
  { value: 'median', name: 'Median' },
  { value: 'q1', name: 'Lower Quartile Boundary' },
  { value: 'q3', name: 'Upper Quartile Boundary' },
  {
    value: 'ci0',
    name: 'Lower boundary of the bootstrapped 95% confidence interval of the mean field value',
  },
  {
    value: 'ci1',
    name: 'Upper boundary of the bootstrapped 95% confidence interval of the mean field value.',
  },
  { value: 'min', name: 'Minimum' },
  { value: 'max', name: 'Maximum' },
  { value: 'argmin', name: 'Argmin (Data object containing the minimum field value)' },
  { value: 'argmax', name: 'Argmax (Data object containing the maximum field value)' },
];
