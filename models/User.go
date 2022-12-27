//
// User.go
// Copyright (C) 2022 tinix <tinix@archlinux>
//
// Distributed under terms of the MIT license.
//

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
