export const getAutosizeOptions = async (form, setAutosizeOptions) => {
  try {
    const schema = form.getFieldValue('$schema');
    const mode = form.getFieldValue('mode');
    const res = await fetch(schema);
    const jsonData = await res.json();

    if (mode === 'vega-lite') {
      setAutosizeOptions(jsonData.definitions.AutosizeType.enum);
    } else if (mode === 'vega') {
      // TODO: handle `oneOf` case here
      setAutosizeOptions(jsonData.definitions.autosize.oneOf[0].enum);
    }
  } catch (error) {
    console.error(error);
  }
};
