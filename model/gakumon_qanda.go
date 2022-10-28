package model

type GakumonQandA struct {
	GakumonId string `binding:"required"`
	Question  string `binding:"required"`
	Answer    string `binding:"required"`
}

type GakumonQuestion struct {
	GakumonId string `binding:"required"`
	Question  string `binding:"required"`
}

type GakumonAnswer struct {
	GakumonId string `binding:"required"`
	Answer    string `binding:"required"`
}
