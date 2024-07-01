package model

import "time"

type BaseModel struct {
	CreatedBy uint `json:"created_by"`
	// CreatedAt time.Time  `json:"created_at,omitempty" gorm:"type:datetime"`
	// UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"type:datetime"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index" gorm:"type:datetime"`
	IsDeleted  bool      `json:"is_deleted"`
	TglCreated time.Time `gorm:"column:created_at;type:datetime;<-:false" json:"created_at" search:"created_at" swaggerignore:"true"`
	TglUpdated time.Time `gorm:"column:modified_at;type:datetime;<-:false" json:"modified_at" search:"modified_at" swaggerignore:"true"`
}

type MetaData struct {
	TotalData int64 `gorm:"column:count" json:"total_data,omitempty"`
	Page      int   `json:"page,omitempty"`
	Limit     int   `json:"limit,omitempty"` // offset
	PrevPage  *int  `json:"prev_page,omitempty"`
	NextPage  *int  `json:"next_page,omitempty"`
	LastPage  int   `json:"last_page,omitempty"`
	Extra     any   `json:"extra,omitempty"`
}

type PaginationFilter struct {
	Page  *int `query:"page"`
	Limit *int `query:"limit"`
}
