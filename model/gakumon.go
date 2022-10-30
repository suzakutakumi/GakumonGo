package model

type Gakumon struct {
	//外部から呼び出すために先頭の文字は大文字
	GakumonID       string  `json:"gakumon_id" binding:"required" firestore:"gakumon_id"`
	Name            string  `json:"name" binding:"required" firestore:"name"`
	Comment         string  `json:"comment" binding:"required" firestore:"comment"`
	ImageURL        *string `json:"image_url" firestore:"image_url"`
	Habitat         *string `json:"habitat" firestore:"habitat"`
	NumberOfCredits *int    `json:"number_of_credits" firestore:"number_of_credits"`
	Type            *string `json:"type" firestore:"type"`
}

type GakumonID struct {
	GakumonID string `json:"gakumon_id" binding:"required" firestore:"gakumon_id"`
}
