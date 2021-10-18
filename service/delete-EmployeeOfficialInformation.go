package service

import (
	"employeeManagementSystem/database/mysql"
	"employeeManagementSystem/models"
)

func Service_DeleteDetails(employee *models.Employee) error {
	return  mysql.Db_DeleteDetails(employee)
}
