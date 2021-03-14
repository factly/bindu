import axios from 'axios';
import {
  GET_SPACES_SUCCESS,
  ADD_SPACE_SUCCESS,
  LOADING_SPACES,
  SET_SELECTED_SPACE,
  DELETE_SPACE_SUCCESS,
  UPDATE_SPACE_SUCCESS,
  SPACES_API,
} from '../constants/spaces';
import { addErrorNotification, addSuccessNotification } from './notifications';

export const getSpaces = () => {
  return (dispatch) => {
    dispatch(loadingSpaces());
    return axios
      .get(SPACES_API)
      .then((response) => {
        dispatch(getSpacesSuccess(response.data));
        return response.data;
      })
      .catch((error) => {
        if (error.response && error.response.data && error.response.data.errors.length > 0) {
          dispatch(addErrorNotification(error.response.data.errors[0].message));
        } else {
          dispatch(addErrorNotification(error.message));
        }
      });
  };
};

export const setSelectedSpace = (space) => {
  return (dispatch) => {
    dispatch({
      type: SET_SELECTED_SPACE,
      payload: space,
    });
    dispatch(addSuccessNotification('Space changed'));
  };
};

export const addSpace = (data) => {
  return (dispatch) => {
    dispatch(loadingSpaces());
    return axios
      .post(SPACES_API, data)
      .then((response) => {
        dispatch(addSpaceSuccess(response.data));
        dispatch(addSuccessNotification('Space added'));
      })
      .catch((error) => {
        if (error.response && error.response.data && error.response.data.errors.length > 0) {
          dispatch(addErrorNotification(error.response.data.errors[0].message));
        } else {
          dispatch(addErrorNotification(error.message));
        }
      });
  };
};

export const deleteSpace = (id) => {
  return (dispatch) => {
    dispatch(loadingSpaces());
    return axios
      .delete(SPACES_API + '/' + id)
      .then(() => {
        dispatch(deleteSpaceSuccess(id));
        dispatch(addSuccessNotification('Space deleted'));
      })
      .catch((error) => {
        if (error.response && error.response.data && error.response.data.errors.length > 0) {
          dispatch(addErrorNotification(error.response.data.errors[0].message));
        } else {
          dispatch(addErrorNotification(error.message));
        }
      });
  };
};

export const updateSpace = (data) => {
  return (dispatch) => {
    dispatch(loadingSpaces());
    return axios
      .put(SPACES_API + '/' + data.id, data)
      .then((response) => {
        dispatch(updateSpaceSuccess(response.data));
        dispatch(addSuccessNotification('Space updated'));
      })
      .catch((error) => {
        if (error.response && error.response.data && error.response.data.errors.length > 0) {
          dispatch(addErrorNotification(error.response.data.errors[0].message));
        } else {
          dispatch(addErrorNotification(error.message));
        }
      });
  };
};

export const loadingSpaces = () => ({
  type: LOADING_SPACES,
});

export const getSpacesSuccess = (spaces) => ({
  type: GET_SPACES_SUCCESS,
  payload: spaces,
});

export const addSpaceSuccess = (space) => ({
  type: ADD_SPACE_SUCCESS,
  payload: space,
});

export const updateSpaceSuccess = (data) => ({
  type: UPDATE_SPACE_SUCCESS,
  payload: data,
});

export const deleteSpaceSuccess = (id) => ({
  type: DELETE_SPACE_SUCCESS,
  payload: id,
});
