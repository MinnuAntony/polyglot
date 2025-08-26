import React, { useState, useEffect } from 'react';
import { getExpenses, addExpense } from '../api';
import { useNavigate } from 'react-router-dom';

const Expenses = ({ userId }) => {
  const [expenses, setExpenses] = useState([]);
  const [newExpense, setNewExpense] = useState({ description: '', amount: '', category: '' });
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    if (!userId) {
      navigate('/login');
      return;
    }
    fetchExpenses();
  }, [userId, navigate]);

  const fetchExpenses = async () => {
    try {
      const response = await getExpenses(userId);
      setExpenses(response.data);
    } catch (error) {
      setMessage("Error fetching expenses.");
    }
  };

  const handleAddExpense = async (e) => {
    e.preventDefault();
    if (!userId) {
      alert('Please log in first.');
      return;
    }
    const expenseData = {
      userId: userId,
      description: newExpense.description,
      amount: parseFloat(newExpense.amount),
      category: newExpense.category,
    };
    try {
      await addExpense(expenseData);
      setNewExpense({ description: '', amount: '', category: '' });
      fetchExpenses(); // Refresh the list
      setMessage('Expense added successfully!');
    } catch (error) {
      setMessage('Error adding expense.');
    }
  };

  return (
    <div className="expenses-container">
      <h2>Your Expenses</h2>
      <form onSubmit={handleAddExpense} className="add-expense-form">
        <input
          type="text"
          placeholder="Description"
          value={newExpense.description}
          onChange={(e) => setNewExpense({ ...newExpense, description: e.target.value })}
          required
        />
        <input
          type="number"
          placeholder="Amount"
          value={newExpense.amount}
          onChange={(e) => setNewExpense({ ...newExpense, amount: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Category"
          value={newExpense.category}
          onChange={(e) => setNewExpense({ ...newExpense, category: e.target.value })}
          required
        />
        <button type="submit">Add Expense</button>
      </form>

      {expenses.length > 0 ? (
        <ul className="expense-list">
          {expenses.map((exp, index) => (
            <li key={index}>
              {exp.description} - ${exp.amount.toFixed(2)} ({exp.category})
            </li>
          ))}
        </ul>
      ) : (
        <p>No expenses found. Add one above!</p>
      )}
      {message && <p>{message}</p>}
    </div>
  );
};

export default Expenses;
