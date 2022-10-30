package model

type QandA struct {
	GakumonID string `json:"gakumon_id" binding:"required"`
	Question  string `json:"question" binding:"required"`
	Answer    string `json:"answer" binding:"required"`
}

type Question struct {
	//GakumonID string `binding:"required"`
	Question string `json:"question" binding:"required"`
}

type Answer struct {
	//GakumonID string `binding:"required"`
	Answer string `json:"answer" binding:"required"`
}

type GakumonByAnswer struct {
	IsCollect bool     `json:"is_correct"`
	User      *Gakumon `json:"user"`
}
