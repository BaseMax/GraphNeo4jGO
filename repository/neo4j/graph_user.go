package neo4j

import (
	"GraphNeo4jGO/model"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type params map[string]any

var ErrNoResult = errors.New("cant get results")

func (n *Neo4j) CreateUser(user model.GraphUser) (err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		defer tx.Close()
		records, err := tx.Run(`MERGE (u:User {username: $username}) RETURN u.username`,
			params{"username": user.Username})
		if err != nil {
			return nil, err
		}

		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		_, found := record.Get("u.username")
		if !found {
			return nil, ErrNoResult
		}

		//if err = tx.Commit(); err != nil {
		//	return nil, err
		//}
		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (n *Neo4j) DeleteUser(user model.GraphUser) (err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err = tx.Run(`MATCH (u:User {username: $username})
								DELETE u`,
			params{"username": user.Username})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (n *Neo4j) UpdateUser(old, newUN string) (err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err = tx.Run(`MATCH (u:User {username: $old} SET u.username = $new`, params{"old": old, "new": newUN})
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}

// FollowUser adds a FOLLOWS relationship from u1 to u2
func (n *Neo4j) FollowUser(u1 model.GraphUser, u2 model.GraphUser) (err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err = tx.Run(`MATCH (u1:User {username: $username1})
                         MATCH (u2:User {username: $username2})
                         MERGE (u1)-[:FOLLOWS]->(u2)`, params{"username1": u1.Username, "username2": u2.Username})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}

// UnFollowUser removes FOLLOWS relationship from u1 to u2
func (n *Neo4j) UnFollowUser(u1 model.GraphUser, u2 model.GraphUser) (err error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err = tx.Run(`MATCH (:User {username: $username1})-[r:FOLLOWS]->(:User {username: $username2})
						 DELETE r`, params{"username1": u1.Username, "username2": u2.Username})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}
