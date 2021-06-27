export const getValueFromNestedPath = (object, path) => {
  let value = object;
  if (path && path.length) {
    path.forEach((key) => {
      value = value[key];
    });
  }
  return value;
};
