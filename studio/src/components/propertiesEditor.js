import React from 'react';
import MonacoEditor from 'react-monaco-editor';
import { debounce } from 'vega';

const MONACOEditor = ({ value, onChange }) => {
  const addNewProperty = (value) => {
    let spec = JSON.parse(value);
    if (spec.$schema === undefined) {
      spec = {
        $schema: 'PropertiesSchema',
        ...spec,
      };
    }

    if (spec.$components === undefined) {
      spec.$components = [];
    }

    spec.$components = [...spec.$components, { Component: '', name: '', properties: [] }];
    if (confirm('Adding schema URL will format the specification too.')) {
      onChange(JSON.stringify(spec, null, 4));
    }
  };

  const editorDidMount = (editor) => {
    editor.addAction({
      contextMenuGroupId: 'vega',
      contextMenuOrder: 0,
      id: 'ADD_NEW_COMPONENT',
      label: 'Add new component',
      run: () => addNewProperty(editor.getModel().getValue()),
    });
  };

  return (
    <MonacoEditor
      language="json"
      options={{
        autoClosingBrackets: 'never',
        autoClosingQuotes: 'never',
        cursorBlinking: 'smooth',
        folding: true,
        lineNumbersMinChars: 4,
        minimap: { enabled: false },
        scrollBeyondLastLine: false,
        wordWrap: 'on',
      }}
      height="640"
      value={value}
      onChange={debounce(300, onChange)}
      editorDidMount={editorDidMount}
    />
  );
};

export default MONACOEditor;
