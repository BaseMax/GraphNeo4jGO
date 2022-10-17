package neo4j

import (
	"GraphNeo4jGO/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (n *Neo4j) NewTweet(t model.Tweet) (uuid [16]byte, err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()
	res, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(`MERGE (t:Tweet {uuid: $uuid, text: $text}-[:POSTED_BY]->(:User {username: $username})
									 RETURN t.uuid`,
			params{"uuid": t.UUID, "text": t.Text, "username": t.Username})
		if err != nil {
			return nil, err
		}

		res, err := result.Single()
		if err != nil {
			return nil, err
		}

		uuid, ok := res.Get("t.uuid")
		if !ok {
			return nil, ErrNoResult
		}
		if err = tx.Commit(); err != nil {
			return nil, err
		}
		return uuid.([16]byte), nil
	})
	if err != nil {
		return [16]byte{}, err
	}

	return res.([16]byte), nil
}
