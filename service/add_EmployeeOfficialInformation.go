package service

import (
	"employeeManagementSystem/database/mysql"
	"employeeManagementSystem/models"
)

func Service_AddDetails(employee *models.Employee) error{
	return  mysql.Db_addDetails(employee)
}