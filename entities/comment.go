package entities

type Comment struct {
	ID uint `json:"id"`
	Text string `json:"text"`
	NewsID uint `json:"news_id"`
	News News
}