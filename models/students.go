package models

import "gorm.io/gorm"

type PaymentStatus string
type StudyStatus string

const (
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusUnpaid  PaymentStatus = "unpaid"
	PaymentStatusPartial PaymentStatus = "partial"
)

const (
	StudyStatusLearning  StudyStatus = "learning"
	StudyStatusJobSearch StudyStatus = "job_search"
	StudyStatusOffer     StudyStatus = "offer"
	StudyStatusWorking   StudyStatus = "working"
)

type Student struct {
	gorm.Model
	FullName      string        `json:"full_name"`
	Email         string        `json:"email"`
	Telegram      string        `json:"telegram"`
	GroupID       uint          `json:"group_id"`
	Group         Group         `gorm:"foreignKey:GroupID" json:"group"`
	TuitionTotal  int           `json:"tuition_total"`
	TuitionPaid   int           `json:"tuition_paid"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	StudyStatus   StudyStatus   `json:"study_status"`
}

type StudentForPost struct {
	FullName      string        `json:"full_name"`
	Email         string        `json:"email"`
	Telegram      string        `json:"telegram"`
	GroupID       uint          `json:"group_id"`
	TuitionTotal  int           `json:"tuition_total"`
	TuitionPaid   int           `json:"tuition_paid"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	StudyStatus   StudyStatus   `json:"study_status"`
}

type StudentForPatch struct {
	FullName      *string        `json:"full_name"`
	Email         *string        `json:"email"`
	Telegram      *string        `json:"telegram"`
	GroupID       *uint          `json:"group_id"`
	TuitionTotal  *int           `json:"tuition_total"`
	TuitionPaid   *int           `json:"tuition_paid"`
	PaymentStatus *PaymentStatus `json:"payment_status"`
	StudyStatus   *StudyStatus   `json:"study_status"`
}
