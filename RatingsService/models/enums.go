package models

type SagaStatus string

const (
	PENDING  SagaStatus = "PENDING"
	ACCEPTED SagaStatus = "ACCEPTED"
	REJECTED SagaStatus = "REJECTED"
)
