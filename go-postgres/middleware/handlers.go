package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamuditg/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error .env File")
	}
	dbCon, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = dbCon.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successful connected to DB")
	return dbCon
}

func CreateStock(writer http.ResponseWriter, request *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(request.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}
	insertId := insertStock(stock)
	res := response{
		ID:      insertId,
		Message: "stock is created",
	}
	json.NewEncoder(writer).Encode(res)
}

func insertStock(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()
	var id int64
	sqlStatement := `INSERT INTO stocks(name,price,company) VALUES ($1,$2,$3) RETURNING stockid`
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single resord: %v", id)
	return id
}

func GetStock(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
		return
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("unable to get stock %v", err)
	}
	json.NewEncoder(writer).Encode(stock)
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()
	var stock models.Stock
	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the ro. %v", err)
	}
	return stock, err
}

func GetAllStock(writer http.ResponseWriter, request *http.Request) {
	stocks, err := getAllStock()
	if err != nil {
		log.Fatalf("unable to get stock %v", err)
	}
	json.NewEncoder(writer).Encode(stocks)
}

func getAllStock() ([]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()
	var stock []models.Stock
	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unale to execute the query %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var stocks models.Stock
		err = rows.Scan(&stocks.StockID, &stocks.Name, &stocks.Price, &stocks.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		stock = append(stock, stocks)
	}
	return stock, nil
}

func UpdateStock(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to get id %v", err)
	}
	var stock models.Stock
	err = json.NewDecoder(request.Body).Decode(stock)
	if err != nil {
		log.Fatalf("unable to get decode stock %v", err)
	}

	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock updated successfull. Total rows/recodes affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(writer).Encode(res)
}

func updateStock(id int64, stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to execute query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("affected rows %v", err)
	}
	return rowsAffected
}

func DeleteStock(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to get id %v", err)
	}
	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock updated successfull. Total rows/recodes affected %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(writer).Encode(res)
}

func deleteStock(id int64) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("affected rows %v", err)
	}
	return rowsAffected
}
