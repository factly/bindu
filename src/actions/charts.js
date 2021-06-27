import axios from 'axios';
import {
  ADD_CHART,
  ADD_CHARTS,
  ADD_CHARTS_REQUEST,
  SET_CHARTS_LOADING,
  RESET_CHARTS,
  CHARTS_API,
} from '../constants/charts';
import { addErrorNotification, addSuccessNotification } from './notification';

export const getCharts = (query) => {
  return (dispatch) => {
    dispatch(loadingCharts());
    return axios
      .get(CHARTS_API, {
        params: { limit: 20 },
      })
      .then((response) => {
        dispatch(addChartsList(response.data.nodes));
        dispatch(
          addChartsRequest({
            data: response.data.nodes.map((item) => item.id),
            query: query,
            total: response.data.total,
          }),
        );
        dispatch(stopChartsLoading());
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const getChart = (id) => {
  return (dispatch) => {
    dispatch(loadingCharts());
    return axios
      .get(CHARTS_API + '/' + id)
      .then((response) => {
        dispatch(getChartByID(response.data));
        dispatch(stopChartsLoading());
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const addChart = (data) => {
  return (dispatch) => {
    dispatch(loadingCharts());
    return axios
      .post(CHARTS_API, data)
      .then(() => {
        dispatch(resetCharts());
        dispatch(addSuccessNotification('Chart added'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const updateChart = (data) => {
  return (dispatch) => {
    dispatch(loadingCharts());
    return axios
      .put(CHARTS_API + '/' + data.id, data)
      .then((response) => {
        dispatch(getChartByID(response.data));
        dispatch(stopChartsLoading());
        dispatch(addSuccessNotification('Chart updated'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const deleteChart = (id) => {
  return (dispatch) => {
    dispatch(loadingCharts());
    return axios
      .delete(CHARTS_API + '/' + id)
      .then(() => {
        dispatch(resetCharts());
        dispatch(addSuccessNotification('Chart deleted'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const addCharts = (charts) => {
  return (dispatch) => {
    dispatch(addChartsList(charts));
  };
};

export const loadingCharts = () => ({
  type: SET_CHARTS_LOADING,
  payload: true,
});

export const stopChartsLoading = () => ({
  type: SET_CHARTS_LOADING,
  payload: false,
});

export const getChartByID = (data) => ({
  type: ADD_CHART,
  payload: data,
});

export const addChartsList = (data) => ({
  type: ADD_CHARTS,
  payload: data,
});

export const addChartsRequest = (data) => ({
  type: ADD_CHARTS_REQUEST,
  payload: data,
});

export const resetCharts = () => ({
  type: RESET_CHARTS,
});
