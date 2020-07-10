import $ from "jquery";
export const mergeDeep = (target, source) => {
  return $.extend(true, target, source)
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
