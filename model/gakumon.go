package model

type Gakumon struct {
	//外部から呼び出すために先頭の文字は大文字
	GakumonId       string `binding:"required"`
	Name            string `binding:"required"`
	Comment         string `binding:"required"`
	ImageUrl        *string
	Habitat         *string
	NumberOfCredits *int
	Type            *string
}
