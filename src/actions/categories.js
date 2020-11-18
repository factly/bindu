import axios from 'axios';
import {
  ADD_CATEGORY,
  ADD_CATEGORIES,
  ADD_CATEGORIES_REQUEST,
  SET_CATEGORIES_LOADING,
  RESET_CATEGORIES,
  CATEGORIES_API,
} from '../constants/categories';
import { addErrorNotification, addSuccessNotification } from './notification';

export const getCategories = (query) => {
  return (dispatch) => {
    dispatch(loadingCategories());
    return axios
      .get(CATEGORIES_API, {
        params: query,
      })
      .then((response) => {
        dispatch(addCategoriesList(response.data.nodes));
        dispatch(
          addCategoriesRequest({
            data: response.data.nodes.map((item) => item.id),
            query: query,
            total: response.data.total,
          }),
        );
        dispatch(stopCategoriesLoading());
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const getCategory = (id) => {
  return (dispatch) => {
    dispatch(loadingCategories());
    return axios
      .get(CATEGORIES_API + '/' + id)
      .then((response) => {
        dispatch(getCategoryByID(response));
        dispatch(stopCategoriesLoading());
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const addCategory = (data) => {
  return (dispatch) => {
    dispatch(loadingCategories());
    return axios
      .post(CATEGORIES_API, data)
      .then(() => {
        dispatch(resetCategories());
        dispatch(addSuccessNotification('Category added'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const updateCategory = (data) => {
  return (dispatch) => {
    dispatch(loadingCategories());
    return axios
      .put(CATEGORIES_API + '/' + data.id, data)
      .then((response) => {
        dispatch(getCategoryByID(response.data));
        dispatch(stopCategoriesLoading());
        dispatch(addSuccessNotification('Category updated'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const deleteCategory = (id) => {
  return (dispatch) => {
    dispatch(loadingCategories());
    return axios
      .delete(CATEGORIES_API + '/' + id)
      .then(() => {
        dispatch(resetCategories());
        dispatch(addSuccessNotification('Category deleted'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const addCategories = (categories) => {
  return (dispatch) => {
    dispatch(addCategoriesList(categories));
  };
};

export const loadingCategories = () => ({
  type: SET_CATEGORIES_LOADING,
  payload: true,
});

export const stopCategoriesLoading = () => ({
  type: SET_CATEGORIES_LOADING,
  payload: false,
});

export const getCategoryByID = (data) => ({
  type: ADD_CATEGORY,
  payload: data,
});

export const addCategoriesList = (data) => ({
  type: ADD_CATEGORIES,
  payload: data,
});

export const addCategoriesRequest = (data) => ({
  type: ADD_CATEGORIES_REQUEST,
  payload: data,
});

export const resetCategories = () => ({
  type: RESET_CATEGORIES,
});
