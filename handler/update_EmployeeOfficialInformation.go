package handler

import (
	"employeeManagementSystem/models"
	"employeeManagementSystem/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler_UpdateDetails (w http.ResponseWriter, r *http.Request) {
	var employee *models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		fmt.Println("Error 1")
		panic(err)
	}
	err := service.Service_UpdateDetails(employee)
	if err != nil{
		fmt.Println("Error 2")
		json.NewEncoder(w).Encode("Not Done query")
	}else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(200)
		fmt.Println(err)
	}

}