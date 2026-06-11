package adapters

import (
	"github.com/tcero76/marketplace/bff-service/dto/demo"
	"github.com/tcero76/marketplace/postgres/model"
)

func ToPostDTO(post model.PostsDemo) demo.PostDTO {
	result := demo.PostDTO{
		ID:        post.ID,
		Texto:     post.Content,
		ReplyToID: post.ReplyToID,
		ProductId: post.ProductId,
	}

	if post.ReplyTo != nil {
		reply := ToPostDTO(*post.ReplyTo)
		result.ReplyTo = &reply
	}

	if len(post.Replies) > 0 {
		replies := make([]demo.PostDTO, 0, len(post.Replies))

		for _, r := range post.Replies {
			replies = append(replies, ToPostDTO(r))
		}

		result.Replies = replies
	}

	return result
}

func ToPostDTOS(posts []model.PostsDemo) []demo.PostDTO {
	result := make([]demo.PostDTO, 0, len(posts))
	for _, post := range posts {
		result = append(result, ToPostDTO(post))
	}
	return result
}

func ToPosteoModel(posteoDTO *demo.PostDTO) model.PostsDemo {
	return model.PostsDemo{
		ID:        posteoDTO.ID,
		Content:   posteoDTO.Texto,
		ReplyToID: posteoDTO.ReplyToID,
		ProductId: posteoDTO.ProductId,
	}
}
