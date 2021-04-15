package main

/*develop a simple REST API in Golang,
by building a CRUD(Create/Read/Update/Delete)
service for managing orders*/
import (
 	"encoding/json" //go type to json and vice-versa
	"log"       //formatting and printing log messages
	"net/http" //provide  http server and client implementations
	"strconv" //converting string into other datatype
	"time"   //provide time values

	"github.com/gorilla/mux"//provide methods to route incoming req to their  handler methods
 	"github.com/swaggo/http-swagger"
)

//Order represents the model for an order
type Order struct {
	OrderID      string    `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemID"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var orders []Order //to store  ours orders
var prevOrderID = 0 //store orderid

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)//converts the body of incoming Http req and save in the field in order
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)// convert into string
	orders = append(orders, order)
	w.Header().Set("Content-Type", "application/json")
	//set the content type header to  json return a json content as response
	json.NewEncoder(w).Encode(order) //order variable to a json as the response
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]//param give orderid from req
	for _, order := range orders {
		if order.OrderID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func main() {
	router := mux.NewRouter()//handles a certain Api req.
	// Create
	router.HandleFunc("/orders", createOrder).Methods("POST")
	// Read
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	// Read-all
	router.HandleFunc("/orders", getOrders).Methods("GET")
	// Update
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	// Delete
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))//starts an Http server

}

/*{"orderedAt":"2019-11-09T21:21:46+00:00","customerName":"Tom Jerry","items":[{"itemId":"123","description":"IPhone 10X","quantity":1}]}*/