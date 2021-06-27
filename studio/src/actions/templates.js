import axios from 'axios';
import {
  ADD_TEMPLATE,
  ADD_TEMPLATES,
  ADD_TEMPLATES_REQUEST,
  SET_TEMPLATES_LOADING,
  RESET_TEMPLATES,
  TEMPLATES_API,
} from '../constants/templates';
import { addCategoriesList } from './categories';
import { addErrorNotification, addSuccessNotification } from './notification';

export const getTemplates = (query) => {
  return (dispatch) => {
    dispatch(loadingTemplates());
    return axios
      .get(TEMPLATES_API, {
        params: query,
      })
      .then((response) => {
        dispatch(addTemplates(response.data));
        dispatch(
          addTemplatesRequest({
            data: response.data.map((item) => item.id),
            query: query,
            total: response.data.total,
          }),
        );
        dispatch(stopTemplatesLoading());
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      })
      .finally(() => {
        dispatch(stopTemplatesLoading());
      });
  };
};

export const getTemplate = (id) => {
  return (dispatch) => {
    dispatch(loadingTemplates());
    return axios
      .get(TEMPLATES_API + '/' + id)
      .then((response) => {
        dispatch(getTemplateByID(response.data));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      })
      .finally(() => {
        dispatch(stopTemplatesLoading());
      });
  };
};
export const addTemplate = (data) => {
  return (dispatch) => {
    dispatch(loadingTemplates());
    return axios
      .post(TEMPLATES_API, data)
      .then(() => {
        dispatch(resetTemplates());
        dispatch(getTemplates());
        dispatch(addSuccessNotification('Template added'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      })
      .finally(() => {
        dispatch(stopTemplatesLoading());
      });
  };
};

export const updateTemplate = (data) => {
  return (dispatch) => {
    dispatch(loadingTemplates());
    return axios
      .put(TEMPLATES_API + '/' + data.id, data)
      .then((response) => {
        dispatch(getTemplateByID(response.data));
        dispatch(stopTemplatesLoading());
        dispatch(addSuccessNotification('Template updated'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const deleteTemplate = (id) => {
  return (dispatch) => {
    dispatch(loadingTemplates());
    return axios
      .delete(TEMPLATES_API + '/' + id)
      .then(() => {
        dispatch(resetTemplates());
        dispatch(getTemplates());
        dispatch(addSuccessNotification('Template deleted'));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const addTemplates = (nodes) => (dispatch) => {
  const templates = [];
  const categories = [];
  nodes.forEach((node) => {
    templates.push(...node.templates);
    categories.push({ ...node, template_ids: node.templates.map((template) => template.id) });
  });
  console.log({ templates, categories });
  dispatch(addTemplatesList(templates));
  dispatch(addCategoriesList(categories));
};

export const loadingTemplates = () => ({
  type: SET_TEMPLATES_LOADING,
  payload: true,
});

export const stopTemplatesLoading = () => ({
  type: SET_TEMPLATES_LOADING,
  payload: false,
});

export const getTemplateByID = (data) => ({
  type: ADD_TEMPLATE,
  payload: data,
});

export const addTemplatesList = (data) => ({
  type: ADD_TEMPLATES,
  payload: data,
});

export const addTemplatesRequest = (data) => ({
  type: ADD_TEMPLATES_REQUEST,
  payload: data,
});

export const resetTemplates = () => ({
  type: RESET_TEMPLATES,
});
