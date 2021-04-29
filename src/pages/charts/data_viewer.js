import React, { useState, useRef } from 'react';
import './data_viewer.css';
import { VariableSizeGrid as Grid } from 'react-window';
import classNames from 'classnames';
import { Table } from 'antd';
import { useDispatch } from 'react-redux';

function DataViewer(props) {
  const { columns, scroll } = props;
  const [tableWidth, setTableWidth] = useState(scroll.x);
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
  const dispatch = useDispatch();

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
        {({ columnIndex, rowIndex, style }) => (
          <div
            className={classNames('virtual-table-cell', {
              'virtual-table-cell-last': columnIndex === mergedColumns.length - 1,
            })}
            style={style}
          >
            {rawData[rowIndex][mergedColumns[columnIndex].dataIndex]}
          </div>
        )}
      </Grid>
    );
  };

  return (
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
  );
}

export default DataViewer;
