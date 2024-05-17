package bexio

type Project struct {
	ID        int          `json:"id"`
	UUID      string       `json:"uuid"`
	Name      string       `json:"name"`
	StartDate Timestamp    `json:"start_date"`
	EndDate   Timestamp    `json:"end_date"`
	ContactID int          `json:"contact_id"`
	StateID   ProjectState `json:"pr_state_id"`
}

func (p *Project) StateName() string {
	return GetProjectStateNameFromId(p.StateID)
}
