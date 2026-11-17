package dto

type SearchRequest struct {
	Mention string   `json:"mention"`
	Hashtag string   `json:"hashtag"`
	Text    []string `json:"text"`
}
