package model

type Grade struct {
	Name      string `json:"name"`
	SubjectID int    `json:"subjectID"`
	Grade     int    `json:"grade"`
}
