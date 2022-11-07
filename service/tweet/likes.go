package tweet

import "GraphNeo4jGO/DTO"

// LikeTweet implements service.Tweet
func (t *TweetService) LikeTweet(liker, poster, uuid string) (DTO.TweetResponse, error) {
	if err := t.validate.Var(uuid, "uuid4"); err != nil {
		return DTO.TweetResponse{}, err
	}

	if err := t.repo.TweetGraph().LikeTweet(liker, poster, uuid); err != nil {
		return DTO.TweetResponse{}, err
	}

	return DTO.TweetResponse{
		Status:  DTO.StatusCreated,
		TweetID: uuid,
	}, nil
	//panic("unimplemented")
}

// UnLike implements service.Tweet
func (t *TweetService) UnLike(liker, poster, uuid string) (DTO.TweetResponse, error) {
	if err := t.validate.Var(uuid, "uuid4"); err != nil {
		return DTO.TweetResponse{}, err
	}

	if err := t.repo.TweetGraph().UnLikeTweet(liker, poster, uuid); err != nil {
		return DTO.TweetResponse{}, err
	}

	return DTO.TweetResponse{
		Status:  DTO.StatusDeleted,
		TweetID: uuid,
	}, nil
	//panic("unimplemented")
}
