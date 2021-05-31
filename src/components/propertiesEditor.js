import React from 'react';
import * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
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

    if (spec.$properties === undefined) {
      spec.$properties = [];
    }

    spec.$properties = [...spec.$properties, { Component: '', name: '', props: [] }];
    if (confirm('Adding schema URL will format the specification too.')) {
      onChange(JSON.stringify(spec, null, 4));
    }
  };

  const editorDidMount = (editor) => {
    editor.addAction({
      contextMenuGroupId: 'vega',
      contextMenuOrder: 0,
      id: 'ADD_NEW_PROPERTY',
      label: 'Add new property',
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
      height="340"
      value={value}
      onChange={debounce(300, onChange)}
      editorDidMount={editorDidMount}
    />
  );
};

export default MONACOEditor;
