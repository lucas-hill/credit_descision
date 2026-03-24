package model

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateApplicationRequest struct {
	FirstName        string  `json:"first_name" validate:"required,min=1,max=100"`
	LastName         string  `json:"last_name" validate:"required,min=1,max=100"`
	Email            string  `json:"email" validate:"required,email"`
	SSNLastFour      string  `json:"ssn_last_four" validate:"required,len=4,numeric"`
	AnnualIncome     float64 `json:"annual_income" validate:"required,gt=0"`
	MonthlyDebt      float64 `json:"monthly_debt" validate:"gte=0"`
	EmploymentStatus string  `json:"employment_status" validate:"required,oneof=EMPLOYED SELF_EMPLOYED UNEMPLOYED RETIRED"`
	RequestedAmount  float64 `json:"requested_amount" validate:"required,gt=0,lte=1000000"`
	LoanPurpose      string  `json:"loan_purpose" validate:"required,oneof=HOME AUTO PERSONAL EDUCATION BUSINESS"`
}

func (r *CreateApplicationRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		var messages []string
		for _, e := range err.(validator.ValidationErrors) {
			messages = append(messages, formatValidationError(e))
		}
		return fmt.Errorf("%s", strings.Join(messages, "; "))
	}
	return nil
}

func formatValidationError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", e.Field())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", e.Field(), e.Param())
	case "gte":
		return fmt.Sprintf("%s must be %s or greater", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("%s must be %s or less", e.Field(), e.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", e.Field(), e.Param())
	case "numeric":
		return fmt.Sprintf("%s must be numeric", e.Field())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s failed validation: %s", e.Field(), e.Tag())
	}
}
