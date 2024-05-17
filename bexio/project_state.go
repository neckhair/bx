package bexio

type ProjectState int

const (
	ProjectStateOpen     ProjectState = 1
	ProjectStateActive   ProjectState = 2
	ProjectStateArchived ProjectState = 3
)

var StateNames = map[ProjectState]string{
	ProjectStateOpen:     "Open",
	ProjectStateActive:   "Active",
	ProjectStateArchived: "Archived",
}

func GetProjectStateNameFromId(id ProjectState) string {
	return StateNames[id]
}
