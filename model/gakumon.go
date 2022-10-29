package model

type Gakumon struct {
	//外部から呼び出すために先頭の文字は大文字
	GakumonID       string  `binding:"required" firestore:"gakumon_id"`
	Name            string  `binding:"required" firestore:"name"`
	Comment         string  `binding:"required" firestore:"comment"`
	ImageURL        *string `firestore:"image_url"`
	Habitat         *string `firestore:"habitat"`
	NumberOfCredits *int    `firestore:"number_of_credits"`
	Type            *string `firestore:"type"`
}

type GakumonID struct {
	GakumonID string `binding:"required" firestore:"gakumon_id"`
}
