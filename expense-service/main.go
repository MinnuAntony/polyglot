package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "sync"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

type Expense struct {
    ID       int     `json:"id"`
    UserID   int     `json:"user_id"`
    Category string  `json:"category"`
    Amount   float64 `json:"amount"`
}

var expenses []Expense
var counter int
var mu sync.Mutex

// Create a new expense
func createExpense(w http.ResponseWriter, r *http.Request) {
    var exp Expense
    json.NewDecoder(r.Body).Decode(&exp)
    mu.Lock()
    counter++
    exp.ID = counter
    expenses = append(expenses, exp)
    mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(exp)
}

// Get expenses for a specific user
func getExpensesByUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, _ := strconv.Atoi(params["userId"])
    var userExpenses []Expense
    for _, e := range expenses {
        if e.UserID == userID {
            userExpenses = append(userExpenses, e)
        }
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userExpenses)
}

// Get all expenses
func getAllExpenses(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(expenses)
}

func main() {
    r := mux.NewRouter()

    // Routes
    r.HandleFunc("/expenses", getAllExpenses).Methods("GET")         // Get all expenses
    r.HandleFunc("/expenses", createExpense).Methods("POST")         // Create expense
    r.HandleFunc("/expenses/{userId}", getExpensesByUser).Methods("GET") // Get expenses by user

    // Enable CORS
    handler := cors.Default().Handler(r)

    http.ListenAndServe(":8080", handler)
}

