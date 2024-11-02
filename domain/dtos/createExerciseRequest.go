package dtos

type CreateExerciseRequest struct {
	Name   string `json:"name"`
	Metric string `json:"metric"`
}
