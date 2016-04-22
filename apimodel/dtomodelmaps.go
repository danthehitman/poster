package apimodel

import "model"

func PostFromPostDto(postDto PostDto) model.Post{
	post := model.Post{
		Description: postDto.Description,
		OwnerId: postDto.OwnerId,
		Title: postDto.Title,
		Body: postDto.Body,
		Uuid: postDto.Uuid,
	}
	return post
}

func PostDtoFromPost(post model.Post) PostDto{
	postDto := PostDto{
		Description: post.Description,
		Title: post.Title,
		Body: post.Body,
		OwnerId: post.OwnerId,
		Uuid: post.Uuid,
	}
	return postDto
}
