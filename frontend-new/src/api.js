import axios from 'axios';

// Base URLs for your microservices, assuming an API Gateway is in place
// Use service names as hostnames for inter-container communication
const USER_SERVICE_URL = 'http://user-service:5000';
const EXPENSE_SERVICE_URL = 'http://expense-service:8080';
const ANALYTICS_SERVICE_URL = 'http://analytics-service:8090';


export const registerUser = (username, password) => {
  return axios.post(`${USER_SERVICE_URL}/register`, { username, password });
};

export const loginUser = (username, password) => {
  return axios.post(`${USER_SERVICE_URL}/login`, { username, password });
};

export const getExpenses = (userId) => {
  return axios.get(`${EXPENSE_SERVICE_URL}/${userId}`);
};

export const addExpense = (expense) => {
  return axios.post(`${EXPENSE_SERVICE_URL}/add`, expense);
};

export const getAnalyticsSummary = (userId) => {
  return axios.get(`${ANALYTICS_SERVICE_URL}/summary/${userId}`);
};
