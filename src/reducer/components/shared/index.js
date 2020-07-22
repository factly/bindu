import produce from 'immer';
import { getValueFromNestedPath } from '../../../utils/index.js';
import {
  SET_AREA_LINES,
  SET_AREA_LINE_WIDTH,
  SET_AREA_LINE_OPACITY,
  SET_AREA_LINE_CURVE,
  SET_AREA_LINE_DASHED,
} from '../../../constants/area_lines.js';
import { SET_BAR_OPACITY, SET_BAR_CORNER_RADIUS } from '../../../constants/bars.js';
import {
  SET_TITLE,
  SET_WIDTH,
  SET_HEIGHT,
  SET_BACKGROUND,
} from '../../../constants/chart_properties.js';
import { SET_COLOR } from '../../../constants/colors.js';
import {
  SET_DATA_LABELS,
  SET_DATA_LABELS_COLOR,
  SET_DATA_LABELS_SIZE,
  SET_DATA_LABELS_FORMAT,
} from '../../../constants/data_labels.js';
import {
  SET_LINE_DOTS,
  SET_LINE_DOT_SHAPE,
  SET_LINE_DOT_SIZE,
  SET_LINE_DOTS_HOLLOW,
} from '../../../constants/dots.js';
import {
  SET_FACET_COLUMNS,
  SET_FACET_SPACING,
  SET_FACET_XAXIS,
  SET_FACET_YAXIS,
} from '../../../constants/facet.js';
import {
  SET_LEGEND_POSITION,
  SET_LEGEND_TITLE,
  SET_LEGEND_BACKGROUND,
  SET_LEGEND_SYMBOL,
  SET_LEGEND_SYMBOL_SIZE,
} from '../../../constants/legend.js';
import {
  SET_LEGEND_LABEL_POSITION,
  SET_LEGEND_LABEL_BASELINE,
  SET_LEGEND_LABEL_COLOR,
} from '../../../constants/legend_label.js';
import {
  SET_LINE_WIDTH,
  SET_LINE_OPACITY,
  SET_LINE_CURVE,
  SET_LINE_DASHED,
} from '../../../constants/lines.js';
import {
  SET_PIE_DATA_LABELS,
  SET_PIE_DATA_LABELS_SIZE,
  SET_PIE_DATA_LABELS_POSITION,
} from '../../../constants/pie_data_labels.js';
import {
  SET_PIE_OUTER_RADIUS,
  SET_PIE_INNER_RADIUS,
  SET_PIE_CORNER_RADIUS,
  SET_PIE_PADDING_ANGLE,
} from '../../../constants/segment.js';
import {
  SET_XAXIS_TITLE,
  SET_XAXIS_POSITION,
  SET_XAXIS_LABEL_FORMAT,
  SET_XAXIS_LABEL_COLOR,
} from '../../../constants/x_axis.js';
import {
  SET_YAXIS_TITLE,
  SET_YAXIS_POSITION,
  SET_YAXIS_LABEL_FORMAT,
  SET_YAXIS_LABEL_COLOR,
} from '../../../constants/y_axis.js';

const sharedReducer = (state = {}, action) => {
  const payload = action.payload;
  switch (action.type) {
    /**
     * GENERAL PROPERTIES
     */

    case SET_WIDTH:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const widthObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        widthObj[path[path.length - 1]] = parseFloat(value);
      });
    case SET_HEIGHT:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const heightObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        heightObj[path[path.length - 1]] = value;
      });
    case SET_TITLE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const titleObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        titleObj[path[path.length - 1]] = value;
      });
    case SET_BACKGROUND:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const backgroundObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        backgroundObj[path[path.length - 1]] = value;
      });
    /**
     * COLOR PROPERTIES
     */
    case SET_COLOR:
      return produce(state, (draftState) => {
        const { value, path, type, index } = payload;
        const colorObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        if (type === 'string') {
          colorObj[path[path.length - 1]] = value;
        } else {
          colorObj[path[path.length - 1]][index] = value;
        }
      });

    /**
     * X AXIS PROPERTIES
     */
    case SET_XAXIS_TITLE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const xaxisTitleObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        xaxisTitleObj[path[path.length - 1]] = value;
      });

    case SET_XAXIS_POSITION:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const xaxisPositionObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        xaxisPositionObj[path[path.length - 1]] = value;
      });

    case SET_XAXIS_LABEL_FORMAT:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const labelFormatObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        labelFormatObj[path[path.length - 1]] = value;
      });

    case SET_XAXIS_LABEL_COLOR:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const labelColorObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        labelColorObj[path[path.length - 1]] = value;
      });

    /**
     * Y AXIS PROPERTIES
     */
    case SET_YAXIS_TITLE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const titleObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        titleObj[path[path.length - 1]] = value;
      });

    case SET_YAXIS_POSITION:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const positionObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        positionObj[path[path.length - 1]] = value;
      });

    case SET_YAXIS_LABEL_FORMAT:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const labelFormatObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        labelFormatObj[path[path.length - 1]] = value;
      });

    case SET_YAXIS_LABEL_COLOR:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const labelColorObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        labelColorObj[path[path.length - 1]] = value;
      });

    /**
     * LEGEND LABEL AXIS PROPERTIES
     */
    case SET_LEGEND_LABEL_POSITION:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const positionObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        positionObj[path[path.length - 1]] = value;
      });
    case SET_LEGEND_LABEL_BASELINE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const labelBaselineObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        labelBaselineObj[path[path.length - 1]] = value;
      });
    case SET_LEGEND_LABEL_COLOR:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const labelColorObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        labelColorObj[path[path.length - 1]] = value;
      });

    /**
     * LEGEND PROPERTIES
     */
    case SET_LEGEND_POSITION:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const positionObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        positionObj[path[path.length - 1]] = value;
      });
    case SET_LEGEND_BACKGROUND:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const legendBgColorObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        legendBgColorObj[path[path.length - 1]] = value;
      });
    case SET_LEGEND_TITLE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const titleObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        titleObj[path[path.length - 1]] = value;
      });
    case SET_LEGEND_SYMBOL:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const legendSymbolObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        legendSymbolObj[path[path.length - 1]] = value;
      });
    case SET_LEGEND_SYMBOL_SIZE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const symbolSizeObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        symbolSizeObj[path[path.length - 1]] = parseInt(value);
      });

    /**
     * DATA LABELS PROPERTIES
     */
    case SET_DATA_LABELS:
      return produce(state, (draftState) => {
        if (action.value) {
          const encoding = draftState.spec.layer[0].encoding;
          draftState.spec.layer.push({
            mark: { type: 'text', dx: 0, dy: 10, fontSize: 12 },
            encoding: {
              x: encoding.x,
              y: { ...encoding.y, stack: 'zero' },
              color: { value: '#ffffff' },
              text: { ...encoding.y, format: '.1f' },
            },
          });
        } else {
          draftState.spec.layer.splice(1, 1);
        }
      });
    case SET_DATA_LABELS_COLOR:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const colorObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        colorObj[path[path.length - 1]] = value;
      });

    case SET_DATA_LABELS_SIZE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const sizeObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        sizeObj[path[path.length - 1]] = parseInt(value);
      });

    case SET_DATA_LABELS_FORMAT:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const formatObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        formatObj[path[path.length - 1]] = value;
      });

    /**
     * DOTS PROPERTIES
     */

    case SET_LINE_DOTS:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        if (value) {
          markObj.point = {
            shape: 'circle',
            size: 40,
            filled: true,
          };
        } else {
          delete markObj['point'];
        }
      });
    case SET_LINE_DOT_SHAPE:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        markObj.point.shape = value;
      });
    case SET_LINE_DOT_SIZE:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        markObj.point.size = parseInt(value);
      });
    case SET_LINE_DOTS_HOLLOW:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        markObj.point.filled = !value;
      });

    /**
     * LINE PROPERTIES
     */
    case SET_LINE_WIDTH:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const widthObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        widthObj[path[path.length - 1]] = parseInt(value);
      });
    case SET_LINE_OPACITY:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const opacityObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        opacityObj[path[path.length - 1]] = parseFloat(value);
      });
    case SET_LINE_CURVE:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const curveObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        curveObj[path[path.length - 1]] = value;
      });
    case SET_LINE_DASHED:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const dashedObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        dashedObj[path[path.length - 1]] = parseInt(value);
      });

    /**
     * FACET PROPERTIES
     */

    case SET_FACET_COLUMNS:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const columnsObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        columnsObj[path[path.length - 1]] = parseInt(value);
      });
    case SET_FACET_SPACING:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const spacingObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        spacingObj[path[path.length - 1]] = parseInt(value);
      });
    case SET_FACET_XAXIS:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const xaxisObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        xaxisObj[path[path.length - 1]] = value;
      });
    case SET_FACET_YAXIS:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const yaxisObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        yaxisObj[path[path.length - 1]] = value;
      });

    /**
     * AREA LINE PROPERTIES
     */
    case SET_AREA_LINES:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        if (value) {
          markObj.line = {
            interpolate: 'linear',
            strokeWidth: 4,
            opacity: 1,
            strokeDash: 0,
          };
        } else {
          delete markObj['line'];
        }
      });
    case SET_AREA_LINE_WIDTH:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        markObj.line.strokeWidth = parseInt(value);
      });
    case SET_AREA_LINE_OPACITY:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        markObj.line.opacity = parseFloat(value);
      });
    case SET_AREA_LINE_CURVE:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        markObj.line.interpolate = value;
      });
    case SET_AREA_LINE_DASHED:
      return produce(state, (draftState) => {
        const { value, path } = payload;
        const markObj = getValueFromNestedPath(draftState.spec, path);
        markObj.line.strokeDash = parseInt(value);
      });

    /**
     * BARS PROPERTIES
     *
     */
    case SET_BAR_OPACITY:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const opacityObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        opacityObj[path[path.length - 1]] = parseFloat(value);
      });

    case SET_BAR_CORNER_RADIUS:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const cornerRadiusObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        cornerRadiusObj[path[path.length - 1]] = parseInt(value);
      });

    /**
     * PIE DATA LABELS
     */
    case SET_PIE_DATA_LABELS:
      return produce(state, (draftState) => {
        if (action.value) {
          const layer = draftState.spec.layer[0];
          const { encoding, mark } = layer;
          draftState.spec.layer.push({
            mark: { type: 'text', fontSize: 12, radius: mark.outerRadius + 10 },
            encoding: {
              theta: { ...encoding.theta, stack: true },
              text: { field: encoding.color.field, type: encoding.color.type },
              color: { ...encoding.color, legend: null },
            },
          });
        } else {
          draftState.spec.layer.splice(1, 1);
        }
      });
    case SET_PIE_DATA_LABELS_SIZE:
      return produce(state, (draftState) => {
        draftState.spec.layer[1].mark.fontSize = parseInt(action.value);
      });
    case SET_PIE_DATA_LABELS_POSITION:
      return produce(state, (draftState) => {
        draftState.spec.layer[1].mark.radius = parseFloat(action.value);
      });

    /**
     * SEGMENT PROPERTIES
     */

    case SET_PIE_OUTER_RADIUS:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const outerRadiusObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        outerRadiusObj[path[path.length - 1]] = parseInt(value);
      });
    case SET_PIE_INNER_RADIUS:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const innerRadiusObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        innerRadiusObj[path[path.length - 1]] = parseInt(value);
      });
    case SET_PIE_CORNER_RADIUS:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const cornerRadiusObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        cornerRadiusObj[path[path.length - 1]] = parseInt(value);
      });
    case SET_PIE_PADDING_ANGLE:
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const paddingAngleObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        paddingAngleObj[path[path.length - 1]] = parseFloat(value);
      });
    /**
     * REGIONS LAYER PROPERTIES
     */
    case 'set-map-stroke-width':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const paddingAngleObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        paddingAngleObj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-stroke-color':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const paddingAngleObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        paddingAngleObj[path[path.length - 1]] = value;
      });

    case 'set-map-stroke-opacity':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const paddingAngleObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        paddingAngleObj[path[path.length - 1]] = parseFloat(value);
      });

    case 'set-map-color-scheme':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const paddingAngleObj = getValueFromNestedPath(
          draftState.spec,
          path.slice(0, path.length - 1),
        );
        paddingAngleObj[path[path.length - 1]] = value;
      });

    /**
     * GRATICULE PROPERTIES
     */

    case 'set-map-graticule-long':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const longObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        longObj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-graticule-lat':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const latObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        latObj[path[path.length - 1]] = value;
      });

    case 'set-map-graticule-width':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const widthObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        widthObj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-graticule-dash':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const dashObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        dashObj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-graticule-opacity':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const opacityObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        opacityObj[path[path.length - 1]] = parseFloat(value);
      });

    case 'set-map-graticule-color':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const colorObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        colorObj[path[path.length - 1]] = value;
      });

    /**
     * ZOOM PROPERTIES
     */

    case 'set-map-zoom-scale':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const scaleObj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        scaleObj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-zoom-rotate0':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const rotate0Obj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        rotate0Obj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-zoom-rotate1':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const rotate1Obj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        rotate1Obj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-zoom-rotate2':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const rotate2Obj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        rotate2Obj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-zoom-center0':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const center0Obj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        center0Obj[path[path.length - 1]] = parseInt(value);
      });

    case 'set-map-zoom-center1':
      return produce(state, (draftState) => {
        const { path, value } = payload;
        const center1Obj = getValueFromNestedPath(draftState.spec, path.slice(0, path.length - 1));
        center1Obj[path[path.length - 1]] = parseInt(value);
      });
    default:
      return state;
  }
};

export default sharedReducer;
