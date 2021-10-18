package models

type Employee struct {
	ID              string `json:"id"`
	UserId          string `json:"userid"`
	Salary          string `json:"salary"`
	Designation     string `json:"designation"`
	Department      string `json:"department"`
	TeamLead        string `json:"teamlead"`
	JobType         string `json:"jobtype"`
	HealthInsurance string `json:"healthinsurance"`
	LifeInsurance   string `json:"lifeinsurance"`
}

/*// Map function returns map values.
func (e *Employee) Map() map[string]interface{} {
	return structs.Map(e)
}

// Names function returns field names.
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
}*/
