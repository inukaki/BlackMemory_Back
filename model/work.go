package model

// 仕事データの構造体
type Work struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Date    string `json:"date"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
	Hours   string `json:"hours"`
	Content string `json:"content"`
	//CASCADE:親テーブルのレコードが削除されたら、子テーブルのレコードも削除する
	User   User `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	UserID uint `json:"user_id" gorm:"not null"`
}

type WorkResponse struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Date    string `json:"date"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
	Hours   string `json:"hours"`
	Content string `json:"content"`
}
