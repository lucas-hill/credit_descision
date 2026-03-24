package model

import "time"

type ApplicationStatus string

const (
	StatusPending  ApplicationStatus = "PENDING"
	StatusApproved ApplicationStatus = "APPROVED"
	StatusDenied   ApplicationStatus = "DENIED"
	StatusReview   ApplicationStatus = "MANUAL_REVIEW"
)

const (
	ApplicationSKMetadata = "METADATA"
)

type Application struct {
	// DynamoDB keys
	PK string `json:"-" dynamodbav:"PK"`
	SK string `json:"-" dynamodbav:"SK"`

	// Application data
	ID        string            `json:"id" dynamodbav:"ID"`
	Status    ApplicationStatus `json:"status" dynamodbav:"Status"`
	CreatedAt time.Time         `json:"created_at" dynamodbav:"CreatedAt"`
	UpdatedAt time.Time         `json:"updated_at" dynamodbav:"UpdatedAt"`

	// Applicant info
	FirstName   string `json:"first_name" dynamodbav:"FirstName"`
	LastName    string `json:"last_name" dynamodbav:"LastName"`
	Email       string `json:"email" dynamodbav:"Email"`
	SSNLastFour string `json:"ssn_last_four" dynamodbav:"SSNLastFour"`

	// Financial info
	AnnualIncome     float64 `json:"annual_income" dynamodbav:"AnnualIncome"`
	MonthlyDebt      float64 `json:"monthly_debt" dynamodbav:"MonthlyDebt"`
	EmploymentStatus string  `json:"employment_status" dynamodbav:"EmploymentStatus"`

	// Loan details
	RequestedAmount float64 `json:"requested_amount" dynamodbav:"RequestedAmount"`
	LoanPurpose     string  `json:"loan_purpose" dynamodbav:"LoanPurpose"`
}

func BuildPK(id string) string {
	return "APP#" + id
}
