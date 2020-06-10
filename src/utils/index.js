export const mergeDeep = (target, source) => {
  const isObject = (obj) => obj && typeof obj === "object";

  if (!isObject(target) || !isObject(source)) {
    return source;
  }

  Object.keys(source).forEach((key) => {
    const targetValue = target[key];
    const sourceValue = source[key];

    if (Array.isArray(targetValue) && Array.isArray(sourceValue)) {
      target[key] = targetValue.concat(sourceValue);
    } else if (isObject(targetValue) && isObject(sourceValue)) {
      target[key] = mergeDeep(Object.assign({}, targetValue), sourceValue);
    } else {
      target[key] = sourceValue;
    }
  });

  return target;
};

export const getRandomId = (len) => {
  var ans = "";
  var arr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
  for (var i = len; i > 0; i--) {
    ans += arr[Math.floor(Math.random() * arr.length)];
  }
  return ans;
};

export const setId = (settings, generatedIds) => {
  generatedIds = generatedIds || {};
  settings.forEach(function (element, index) {
    let id = getRandomId(8);
    while (id in generatedIds) {
      id = getRandomId(8);
    }
    generatedIds[id] = true;
    element.id = id;
    if (element.type == "object") {
      setId(element.items, generatedIds);
    }
  });
  return settings;
};
