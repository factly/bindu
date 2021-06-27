import React, { useState } from 'react';
import './data_viewer.css';
import { Button } from 'antd';
import ReactDataGrid from 'react-data-grid';

const EmptyRowsView = () => {
  const message = 'No data to show';
  return (
    <div className="data-view-empty-message">
      <h3>{message}</h3>
    </div>
  );
};

function DataViewer(props) {
  const { columns, dataSource, tableWidth, tableHeight, onDataChange, tabIndex } = props;
  const [isEditable, setIsEditable] = useState(false);

  const mergedColumns = columns.map((column) => {
    return {
      ...column,
      editable: isEditable,
      resizable: true,
      width: Math.max(Math.floor((tableWidth - 4) / columns.length), 150),
    };
  });

  return (
    <>
      {!isEditable ? (
        <Button onClick={() => setIsEditable(true)}>Edit Data</Button>
      ) : (
        <Button onClick={() => setIsEditable(false)}>Done</Button>
      )}

      <ReactDataGrid
        columns={mergedColumns}
        rowGetter={(i) => dataSource[i]}
        rowsCount={dataSource.length}
        onGridRowsUpdated={(...params) => onDataChange(...params, tabIndex)}
        enableCellSelect={true}
        minHeight={tableHeight}
        minWidth={tableWidth}
        emptyRowsView={EmptyRowsView}
      />
    </>
  );
}

export default DataViewer;
