package model

type GakumonQandA struct {
	GakumonID string `binding:"required"`
	Question  string `binding:"required"`
	Answer    string `binding:"required"`
}

type GakumonQuestion struct {
	GakumonID string `binding:"required"`
	Question  string `binding:"required"`
}

type GakumonAnswer struct {
	GakumonID string `binding:"required"`
	Answer    string `binding:"required"`
}
