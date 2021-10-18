package service
import (
	"employeeManagementSystem/database/mysql"
	"employeeManagementSystem/models"
)

func Service_UpdateDetails(employee *models.Employee) error {
	return mysql.Db_UpdateDetails(employee)
}