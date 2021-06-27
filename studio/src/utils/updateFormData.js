const updateFormData = async (specData, path) => {
  const minoURL = fetch(
    'http://127.0.0.1:3020/s3/params?' +
      new URLSearchParams({
        filename: path,
        type: 'application/json',
      }),
  )
    .then((res) => res.json())
    .then((data) => {
      var jsonse = JSON.stringify(specData.values);
      var blob = new Blob([jsonse], { type: 'application/json' });

      const chartFormData = new FormData();

      Object.keys(data.fields).forEach((key) => {
        chartFormData.append(key, data.fields[key]);
      });

      chartFormData.append('file', blob);

      const url = fetch('http://localhost:9000/dega', {
        method: data.method,
        body: chartFormData,
      })
        .then((re) => {
          return re;
        })
        .then((da) => {
          return 'http://localhost:9000/dega/' + path;
        });
      return url;
    });
  return minoURL;
};

export default updateFormData;
