package models
import (
	"github.com/fatih/structs"
)
type Employee struct {
	ID              string `json:"id" structs:"id"`
	UserId          string `json:"userId" structs:"userid"`
	Salary          string `json:"salary" structs:"salary"`
	Designation     string `json:"designation" structs:"designation"`
	Department      string `json:"department" structs:"department"`
	TeamLead        string `json:"teamlead" structs:"teamLead"`
	JobType         string `json:"jobtype" structs:"jobtype"`
	HealthInsurance string `json:"healthinsurance" structs:"healthinsurance"`
	LifeInsurance   string `json:"lifeinsurance" structs:"lifeinsurance"`
}


func (e *Employee) Map() map[string]interface{} {
	return structs.Map(e)
	}

func (e *Employee) Names() []string {
	fields := structs.Fields(e)
	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}
	return names
}
