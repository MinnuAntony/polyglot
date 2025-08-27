import axios from 'axios';

const USER_SERVICE_URL = '/user-service';
const EXPENSE_SERVICE_URL = '/expense-service';
const ANALYTICS_SERVICE_URL = '/analytics-service';

// Register user with safe error handling
export const registerUser = async (username, password) => {
  try {
    const res = await axios.post(`${USER_SERVICE_URL}/register`, { username, password });
    console.log("Register response:", res);
    return res.data; // return data directly
  } catch (error) {
    console.error("Register error:", error);
    if (error.response && error.response.data) {
      throw new Error(error.response.data.detail || 'Registration failed');
    } else {
      throw new Error(error.message || 'Registration failed');
    }
  }
};

// Login user
export const loginUser = async (username, password) => {
  try {
    const res = await axios.post(`${USER_SERVICE_URL}/login`, { username, password });
    return res.data;
  } catch (error) {
    if (error.response && error.response.data) {
      throw new Error(error.response.data.detail || 'Login failed');
    } else {
      throw new Error(error.message || 'Login failed');
    }
  }
};

// Expenses
export const getExpenses = async (userId) => {
  const res = await axios.get(`${EXPENSE_SERVICE_URL}/${userId}`);
  return res.data;
};

export const addExpense = async (expense) => {
  const res = await axios.post(`${EXPENSE_SERVICE_URL}/add`, expense);
  return res.data;
};

// Analytics
export const getAnalyticsSummary = async (userId) => {
  const res = await axios.get(`${ANALYTICS_SERVICE_URL}/summary/${userId}`);
  return res.data;
};

