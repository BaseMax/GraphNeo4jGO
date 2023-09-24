package neo4j

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

func (n *Neo4j) LikeTweet(liker, poster, tweetID string) error {
	var err error

	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	_, err = session.Run(`MATCH (:User {username: $poster})-[:TWEETED]->(tweet:Tweet {uuid: $uuid})
                          MATCH (liker:User {username: $liker})
                          MERGE (liker)-[:LIKED]->(tweet)`,
		params{"poster": poster, "liker": liker, "uuid": tweetID})
	if err != nil {
		return err
	}

	return err
}

func (n *Neo4j) UnLikeTweet(liker, poster, tweetID string) error {
	var err error

	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	_, err = session.Run(`MATCH (:User {username: $poster})-[:TWEETED]->(tweet:Tweet {uuid: $uuid})
                          MATCH (liker:User {username: $liker})-[r:LIKED]->(tweet)
                          DELETE r`,
		params{"poster": poster, "liker": liker, "uuid": tweetID})

	return err
}
