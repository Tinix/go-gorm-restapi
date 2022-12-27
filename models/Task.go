//
// Task.go
// Copyright (C) 2022 tinix <tinix@archlinux>
//
// Distributed under terms of the MIT license.
//

package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `gorm:"not null; unique_index"`
	Description string
	Done        bool `gorm:"default:false"`
	UserID      uint
}
