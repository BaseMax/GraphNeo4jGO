package neo4j

import (
	"GraphNeo4jGO/model"

	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (n *Neo4j) CommentOn(c model.Comment) (string, error) {
	var err error

	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	results, err := session.Run(
		`MATCH (:User {username: $poster})-[:TWEETED]->(tweet:Tweet {uuid: $tweet_uuid})
         MERGE (:User {username: $commenter})-[r:COMMENTED {uuid: $comment_uuid}]->(tweet)
         SET r.comment = $comment
         RETURN r.uuid`,
		params{"poster": c.Poster, "tweet_uuid": c.TweetID, "commenter": c.Commenter, "comment": c.Text, "comment_uuid": c.UUID})

	if err != nil {
		return "", err
	}

	result, err := results.Single()
	if err != nil {
		return "", err
	}
	uuid, ok := result.Get("r.uuid")
	if !ok {
		return "", ErrNoResult
	}

	return uuid.(string), err
}

func (n *Neo4j) DeleteComment(c model.Comment) error {
	var err error

	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	_, err = session.Run(
		`MATCH (tweet:Tweet {uuid: $uuid})<-[:TWEETED]-(:User {username: $poster})
         MATCH (commenter:User {username: $commenter})-[r:COMMENTED {uuid: $comment_uuid}]->(tweet)
         DELETE r`,
		params{"uuid": c.TweetID, "poster": c.Poster, "commenter": c.Commenter, "comment_uuid": c.UUID})

	return err
}

func (n *Neo4j) GetComments(username, tweetID string) ([]model.Comment, error) {
	var (
		err      error
		comments []model.Comment
	)

	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	results, err := session.Run(
		`MATCH (:User {username: $username})-[c:COMMENTED]->(:Tweet {uuid: $uuid})
        RETURN c`,
		params{"username": username, "uuid": tweetID},
	)

	if err != nil {
		return nil, err
	}

	signleRes, err := results.Single()
	if err != nil {
		return nil, err
	}

	if c, ok := signleRes.Get("c"); ok {
		var comment model.Comment
		node := c.(neo4j.Node)
		if err := mapstructure.Decode(node.Props, &comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, err
}
