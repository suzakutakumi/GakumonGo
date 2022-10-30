package model

type QandA struct {
	GakumonID string `binding:"required"`
	Question  string `binding:"required"`
	Answer    string `binding:"required"`
}

type Question struct {
	//GakumonID string `binding:"required"`
	Question string `binding:"required"`
}

type Answer struct {
	//GakumonID string `binding:"required"`
	Answer string `binding:"required"`
}

type GakumonByAnswer struct {
	IsCollect bool
	User      *Gakumon
}
