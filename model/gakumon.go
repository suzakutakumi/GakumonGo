package model

type Gakumon struct {
	//外部から呼び出すために先頭の文字は大文字
	GakumonId       string  `binding:"required" firestore:"gakumon_id"`
	Name            string  `binding:"required" firestore:"name"`
	Comment         string  `binding:"required" firestore:"comment"`
	ImageUrl        *string `firestore:"image_url"`
	Habitat         *string `firestore:"habitat"`
	NumberOfCredits *int    `firestore:"number_of_credits"`
	Type            *string `firestore:"type"`
}

type GakumonId struct {
	GakumonId string `binding:"required" firestore:"gakumon_id"`
}
