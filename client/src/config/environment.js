const isDev = import.meta.env.DEV;
const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || (isDev ? 'http://localhost:8080' : '');

export const config = {
    apiBaseUrl: apiBaseUrl + '/api',
    isDev,
};