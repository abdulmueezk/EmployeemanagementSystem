package mysql

import (
	"employeeManagementSystem/models"
	"fmt"
)
func Db_addDetails(employee *models.Employee) error{
	//query := "insert into employeeOfficial (salary,designation,department,teamLead,jobType,healthInsurance,lifeInsurance) values ( '"+employee.Salary+"' ,'"+employee.Designation+"' ,'"+employee.Department+"','"+employee.TeamLead+"','"+employee.JobType+"','"+employee.JobType+"','"+employee.HealthInsurance+"','"+employee.LifeInsurance+"')"
	db := Sqlclient()
	result, err := db.Query("INSERT INTO employeeOfficial VALUES ( '"+employee.Salary+"' ,'"+employee.Designation+"' ,'"+employee.Department+"','"+employee.TeamLead+"','"+employee.JobType+"','"+employee.HealthInsurance+"','"+employee.LifeInsurance+"','"+employee.UserId +"')")
		if err != nil {
			return err
		}

		defer result.Close()
		//fmt.Println(result) //print employee values in the terminal
	/*for result.Next() {
		err := result.Scan(&employee.Salary, &employee.Designation, &employee.Department, &employee.TeamLead, &employee.JobType, &employee.HealthInsurance, &employee.LifeInsurance, &employee.UserId)
		if err != nil {
			panic(err.Error())
		}
	}*/
	fmt.Println(result)
	return nil
}
