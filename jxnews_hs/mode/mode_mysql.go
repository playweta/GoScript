package mode

type TM struct {
	//ID       int64  `gorm:"id"`
	Question string `gorm:"question"`
	Answer   string `gorm:"answer"`
}
