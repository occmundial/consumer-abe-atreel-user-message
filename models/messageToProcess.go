package models

// MessageToProcess : mensaje a procesar

type MessageToProcess struct {
	LoginID         string   `json:"loginID"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Recommendations []Agents `json:"recommendations"`
	AbSide          string   `json:"abSide"`
	AbTestName      string   `json:"abTestName"`
}
