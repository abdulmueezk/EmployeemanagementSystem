package mysql

import (
	"employeeManagementSystem/models"
	"testing"
)

func TestDb_ShowDetails(t *testing.T) {
	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add employee details in db",
			args:    args{employee: &models.Employee{UserId: "31601", Salary: "20000", Department: "backend", Designation: "Intern" , TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: "0", LifeInsurance: "0"}},
			wantErr: false,
		},
		{
			name:    "Failed - add employee details in db",
			args:    args{employee: &models.Employee{ TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: "0", LifeInsurance: "0"}},
			wantErr: true,
		},
		{
			name:    "Failed - add employee details in db",
			args:    args{employee: &models.Employee{ TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: "0", LifeInsurance: "0"}},
			wantErr: true,
		},
		{
			name:    "Failed - add employee details in db",
			args:    args{employee: &models.Employee{ }},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Db_ShowDetails(tt.args.employee); (err != nil) != tt.wantErr {
				t.Errorf("Db_ShowDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
