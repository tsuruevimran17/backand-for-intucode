package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Title       string `json:"title"`
	CurrentWeek int    `json:"current_week"`
	TotalWeeks  int    `json:"total_weeks"`
	InFinished  bool   `json:"in_finished"`
}

type GroupForPost struct {
	Title       string `json:"title"`
	CurrentWeek int    `json:"current_week"`
	TotalWeeks  int    `json:"total_weeks"`
	InFinished  bool   `json:"in_finished"`
}

type GroupForPatch struct {
	Title       *string `json:"title"`
	CurrentWeek *int    `json:"current_week"`
	TotalWeeks  *int    `json:"total_weeks"`
	InFinished  *bool   `json:"in_finished"`
}
