package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Expense represents the structure of an expense record.
type Expense struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}

// Global variable for the database connection pool.
var db *sql.DB
var mu sync.Mutex

// initDB initializes the database connection.
/* func initDB() {
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MySQL database!")
} */

// createExpense adds a new expense to the database.
func createExpense(w http.ResponseWriter, r *http.Request) {
	var exp Expense
	if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Assuming a database connection exists. This part is commented out in your original code.
	// You will need to uncomment and fix the initDB function for this to work.
	// As the `db` variable is currently `nil`, this code will panic at runtime.
	result, err := db.Exec("INSERT INTO expenses (user_id, category, amount) VALUES (?, ?, ?)", exp.UserID, exp.Category, exp.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	exp.ID = int(lastID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exp)
}

// getExpensesByUser fetches all expenses for a specific user from the database.
func getExpensesByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userId"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// This code will also panic if db is nil.
	rows, err := db.Query("SELECT id, user_id, category, amount FROM expenses WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var userExpenses []Expense
	for rows.Next() {
		var exp Expense
		if err := rows.Scan(&exp.ID, &exp.UserID, &exp.Category, &exp.Amount); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userExpenses = append(userExpenses, exp)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userExpenses)
}

// getAllExpenses fetches all expenses from the database.
func getAllExpenses(w http.ResponseWriter, r *http.Request) {
	// This code will also panic if db is nil.
	rows, err := db.Query("SELECT id, user_id, category, amount FROM expenses")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allExpenses []Expense
	for rows.Next() {
		var exp Expense
		if err := rows.Scan(&exp.ID, &exp.UserID, &exp.Category, &exp.Amount); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		allExpenses = append(allExpenses, exp)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allExpenses)
}

func main() {
	// Initialize database connection
	//initDB()
	// Close the database connection when the main function exits.
	// This will cause a panic if initDB() is not called and a connection is not established.
	// defer db.Close() 

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/expenses", getAllExpenses).Methods("GET")
	r.HandleFunc("/expenses", createExpense).Methods("POST")
	r.HandleFunc("/expenses/{userId}", getExpensesByUser).Methods("GET")

	// Correct CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://3.95.198.203"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		Debug: true, // This is very helpful for debugging CORS issues
	})

	handler := c.Handler(r)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", handler)
}
