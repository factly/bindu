import React, { useState, useRef } from 'react';
import './data_viewer.css';
import { VariableSizeGrid as Grid } from 'react-window';
import classNames from 'classnames';
import { Button, Input, InputNumber, Table } from 'antd';

function DataViewer(props) {
  const { columns, scroll, onDataChange } = props;
  const tableWidth = scroll.x;
  const [isEditable, setIsEditable] = useState(false);
  const mergedColumns = columns.map((column) => {
    if (column.width) {
      return column;
    }

    return {
      ...column,
      width: Math.max(Math.floor((tableWidth - 4) / columns.length), 150),
    };
  });

  const gridRef = useRef();

  const [connectObject] = useState(() => {
    const obj = {};
    Object.defineProperty(obj, 'scrollLeft', {
      get: () => null,
      set: (scrollLeft) => {
        if (gridRef.current) {
          gridRef.current.scrollTo({
            scrollLeft,
          });
        }
      },
    });
    return obj;
  });

  const renderVirtualList = (rawData, { scrollbarSize, ref, onScroll }) => {
    ref.current = connectObject;
    return (
      <Grid
        ref={gridRef}
        className="virtual-grid"
        columnCount={mergedColumns.length}
        columnWidth={(index) => {
          const { width } = mergedColumns[index];
          return index === mergedColumns.length - 1 ? width - scrollbarSize - 1 : width;
        }}
        height={scroll.y}
        rowCount={rawData.length}
        rowHeight={() => 54}
        width={tableWidth}
        onScroll={({ scrollLeft }) => {
          onScroll({ scrollLeft });
        }}
      >
        {({ columnIndex, rowIndex, style }) => {
          const value = rawData[rowIndex][mergedColumns[columnIndex].dataIndex];
          let InputComponent = Input;
          let onChange = (event) =>
            onDataChange(rowIndex, mergedColumns[columnIndex].dataIndex, event.target.value);
          if (typeof value === 'number') {
            InputComponent = InputNumber;
            onChange = (value) =>
              onDataChange(rowIndex, mergedColumns[columnIndex].dataIndex, value);
          }
          return (
            <div
              className={classNames('virtual-table-cell', {
                'virtual-table-cell-last': columnIndex === mergedColumns.length - 1,
              })}
              style={style}
            >
              {isEditable ? <InputComponent value={value} onChange={onChange} /> : value}
            </div>
          );
        }}
      </Grid>
    );
  };

  return (
    <>
      {!isEditable ? (
        <Button onClick={() => setIsEditable(true)}>Edit Data</Button>
      ) : (
        <Button onClick={() => setIsEditable(false)}>Done</Button>
      )}
      <Table
        {...props}
        className="virtual-table"
        columns={mergedColumns}
        pagination={false}
        showHeader={true}
        components={{
          body: renderVirtualList,
        }}
      />
    </>
  );
}

export default DataViewer;
