import React from 'react';
import ReactJson from 'react-json-view';

const JSONEditor = ({ value, onChange }) => {
  return (
    <ReactJson
      src={value}
      onEdit={({ updated_src }) => onChange(updated_src)}
      onDelete={() => {}}
      onAdd={() => {}}
    />
  );
};

export default JSONEditor;
