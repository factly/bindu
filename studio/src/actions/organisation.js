import axios from 'axios';
import {
  GET_ORGANISATIONS_SUCCESS,
  LOADING_ORGANISATIONS,
  API_GET_ORGANISATIONS,
  SET_SELECTED_ORGANISATION,
} from '../constants/organisation';
import { addErrorNotification, addSuccessNotification } from './notification';

export const getOrganisations = () => {
  return (dispatch) => {
    dispatch(loadingOrganisations());
    return axios
      .get(API_GET_ORGANISATIONS)
      .then((response) => {
        dispatch(getOrganisationsSuccess(response.data.nodes));
      })
      .catch((error) => {
        dispatch(addErrorNotification(error.message));
      });
  };
};

export const setSelectedOrganisation = (organisation) => {
  return (dispatch) => {
    dispatch({
      type: SET_SELECTED_ORGANISATION,
      payload: organisation,
    });
    dispatch(addSuccessNotification('Organisation changed'));
  };
};

export const loadingOrganisations = () => ({
  type: LOADING_ORGANISATIONS,
});

export const getOrganisationsSuccess = (organisations) => ({
  type: GET_ORGANISATIONS_SUCCESS,
  payload: organisations,
});
