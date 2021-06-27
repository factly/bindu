import axios from 'axios';

function createAxiosAuthMiddleware() {
  return ({ getState }) => (next) => (action) => {
    axios.defaults.headers.common['X-Space'] = getState().spaces.selected;
    axios.defaults.baseURL = process.env.REACT_APP_API_URL;
    return next(action);
  };
}

const axiosAuth = createAxiosAuthMiddleware();

export default axiosAuth;
