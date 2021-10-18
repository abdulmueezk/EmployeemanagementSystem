package mysql

import (
	"employeeManagementSystem/models"
	"fmt"
)

func Db_DeleteDetails(employee *models.Employee) error{
	var abcd error
	db := Sqlclient()
	query := "UPDATE employeeOfficial SET jobType='"+employee.JobType+"' where userid='"+employee.UserId +"'"
	querychk := "select * from employeeOfficial where userid ='" +employee.UserId + "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		panic(err)
	}
	defer resultchk.Close()
	var emp models.Employee
	for resultchk.Next() {
		err := resultchk.Scan(&emp.Salary, &emp.Designation, &emp.Department, &emp.TeamLead, &emp.JobType, &emp.HealthInsurance, &emp.LifeInsurance, &emp.UserId)
		if err != nil {
			panic(err.Error())
		}
	}


	if employee.UserId == emp.UserId {
		result, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		fmt.Println(result)
	} else {
		var x int
		err := fmt.Errorf("math: square root of negative number %g", x)
		abcd = err
		fmt.Println("Sorry Invalid Entry")
	}
	return abcd
}

