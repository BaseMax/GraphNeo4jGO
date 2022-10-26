package neo4j

import (
	"GraphNeo4jGO/model"

	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (n *Neo4j) NewTweet(t model.Tweet) (uuid string, err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	result, err := session.Run(`
        MERGE (:User {username: $username})-[:TWEETED]->(t:Tweet {uuid: $uuid})
        ON CREATE set t.date = datetime()
        ON CREATE set t.text = $text
        RETURN t.uuid`,
		params{"uuid": t.UUID, "text": t.Text, "username": t.Poster})
	if err != nil {
		return "", err
	}

	res, err := result.Single()
	if err != nil {
		return "", err
	}

	tweetUUID, ok := res.Get("t.uuid")
	if !ok {
		return "", ErrNoResult
	}

	return tweetUUID.(string), nil
}

func (n *Neo4j) UserTweets(username string, limit, skip int) ([]model.Tweet, error) {
	var err error
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	var tweets []model.Tweet
	res, err := session.Run(`
        MATCH (:User {username: $username})-[:TWEETED]->(t:Tweet)
        RETURN t SKIP $skip LIMIT $limit`,
		params{"username": username, "limit": limit, "skip": skip})
	if err != nil {
		return nil, err
	}

	for res.Next() {
		record := res.Record()

		if value, ok := record.Get("t"); ok {
			var tweet model.Tweet
			node := value.(neo4j.Node)

			if err = mapstructure.Decode(node.Props, &tweet); err != nil {
				return nil, err
			}

			tweets = append(tweets, tweet)
		}
	}
	if err = res.Err(); err != nil {
		return nil, err
	}

	return tweets, err
}

func (n *Neo4j) GetTweet(username, uuid string) (model.Tweet, error) {
	var (
		err   error
		tweet model.Tweet
	)

	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	res, err := session.Run(`
        MATCH (:User {username: $username})-[:TWEETED]->(t:Tweet {uuid: $uuid}) RETURN t`,
		params{"uuid": uuid, "username": username})
	if err != nil {
		return model.Tweet{}, err
	}

	singleRes, err := res.Single()
	if err != nil {
		return model.Tweet{}, err
	}
	if value, ok := singleRes.Get("t"); ok {
		node := value.(neo4j.Node)
		if err = mapstructure.Decode(node.Props, &tweet); err != nil {
			return model.Tweet{}, err
		}
	}

	return tweet, err
}

func (n *Neo4j) Delete(user, uuid string) (err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	_, err = session.Run(
		`MATCH (:User {username: $username})-[:TWEETED]->(t:Tweet {uuid: $uuid})
        DELETE t`,
		params{"username": user, "uuid": uuid},
	)

	return err
}
