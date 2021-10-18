package service

import (
	"employeeManagementSystem/database/mysql"
	"employeeManagementSystem/models"
)

func Service_ShowDetails(employee *models.Employee) error {
	return mysql.Db_ShowDetails(employee)
}