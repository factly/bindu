import React from 'react';
import * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
import MonacoEditor from 'react-monaco-editor';
import { debounce } from 'vega';

const MONACOEditor = ({ value, onChange }) => {
  const SCHEMA = {
    vega: 'https://vega.github.io/schema/vega/v5.json',
    'vega-lite': 'https://vega.github.io/schema/vega-lite/v4.json',
  };

  const addVegaSchemaURL = (value) => {
    let spec = JSON.parse(value);
    if (spec.$schema !== undefined) {
      return;
    }
    if (spec.$schema === undefined) {
      spec = {
        $schema: SCHEMA['vega'],
        ...spec,
      };
    }
    if (confirm('Adding schema URL will format the specification too.')) {
      onChange(JSON.stringify(spec, null, 4));
    }
  };

  const addVegaLiteSchemaURL = (value) => {
    let spec = JSON.parse(value);
    if (spec.$schema !== undefined) {
      return;
    }
    if (spec.$schema === undefined) {
      spec = {
        $schema: SCHEMA['vega-lite'],
        ...spec,
      };
    }
    if (confirm('Adding schema URL will format the specification too.')) {
      onChange(JSON.stringify(spec, null, 4));
    }
  };

  const editorDidMount = (editor) => {
    editor.addAction({
      contextMenuGroupId: 'vega',
      contextMenuOrder: 0,
      id: 'ADD_VEGA_SCHEMA',
      label: 'Add Vega schema URL',
      run: () => addVegaSchemaURL(editor.getModel().getValue()),
    });

    editor.addAction({
      contextMenuGroupId: 'vega',
      contextMenuOrder: 1,
      id: 'ADD_VEGA_LITE_SCHEMA',
      label: 'Add Vega-Lite schema URL',
      run: () => addVegaLiteSchemaURL(editor.getModel().getValue()),
    });

    // editor.addAction({
    //   contextMenuGroupId: 'vega',
    //   contextMenuOrder: 2,
    //   id: 'CLEAR_EDITOR',
    //   label: 'Clear Spec',
    //   run: this.onClear.bind(this),
    // });

    // editor.addAction({
    //   contextMenuGroupId: 'vega',
    //   contextMenuOrder: 3,
    //   id: 'MERGE_CONFIG',
    //   label: 'Merge Config Into Spec',
    //   run: this.handleMergeConfig.bind(this),
    // });

    // editor.addAction({
    //   contextMenuGroupId: 'vega',
    //   contextMenuOrder: 4,
    //   id: 'EXTRACT_CONFIG',
    //   label: 'Extract Config From Spec',
    //   run: this.handleExtractConfig.bind(this),
    // });
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
