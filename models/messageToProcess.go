package models

// MessageToProcess : mensaje a procesar

type MessageToProcess struct {
	LoginID         string
	Name            string
	Email           string
	Recommendations []Agents
}
