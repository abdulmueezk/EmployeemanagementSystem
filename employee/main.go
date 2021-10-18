package main

import (
	"employeeManagementSystem/handler"
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/adddofficialdetail", handler.Handler_AddDetails)
	http.HandleFunc("/showdofficialdetail", handler.Handler_ShowDetails)
	http.HandleFunc("/updatedofficialdetail", handler.Handler_UpdateDetails)
	http.HandleFunc("/deletedofficialdetail", handler.Handler_DeleteDetails)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
