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
	records, err := session.Run(`
        MERGE (u:User {username: $username})
        // CREATE CONSTRAINT uniq_user IF NOT EXISTS FOR (u) REQUIRE u.username IS UNIQUE
        RETURN u.username
        `,
		params{"username": user.Username})
	if err != nil {
		return err
	}

	record, err := records.Single()
	if err != nil {
		return err
	}
	_, found := record.Get("u.username")
	if !found {
		return ErrNoResult
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

	_, err = session.Run(`
        MATCH (u:User {username: $username})
        MATCH (u)-[tr:TWEETED]->(t:Tweet)
        MATCH (u)-[lr:LIKED]->(:Tweet)
        DELETE tr, t, lr, u`,
		params{"username": user.Username})
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

	_, err = session.Run(`MATCH (u:User {username: $old} SET u.username = $new`, params{"old": old, "new": newUN})
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
	_, err = session.Run(`MATCH (u1:User {username: $username1})
                          MATCH (u2:User {username: $username2})
                          MERGE (u1)-[:FOLLOWS]->(u2)`, params{"username1": u1.Username, "username2": u2.Username})
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
	_, err = session.Run(`MATCH (:User {username: $username1})-[r:FOLLOWS]->(:User {username: $username2}
                          DELETE r`, params{"username1": u1.Username, "username2": u2.Username})
	if err != nil {
		return err
	}
	return nil

}

// return list of user follower from username
func (n *Neo4j) GetFollowers(username string) ([]string, error) {
	var (
		err       error
		usernames []string
	)
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()

	results, err := session.Run(
		`SELECT (users:User)-[:FOLLOWS]->(:User {username: $username}) RETURN users.username`,
		params{"username": username},
	)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		record := results.Record()
		u, ok := record.Get("users.username")
		if ok {
			usernames = append(usernames, u.(string))
		}
	}

	return usernames, err
}
