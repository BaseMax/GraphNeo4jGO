package neo4j

import (
	"GraphNeo4jGO/config"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4j struct {
	driver neo4j.Driver
}

func New(cfg config.Neo4j) (*Neo4j, error) {
	driver, err := neo4j.NewDriver(cfg.URI, neo4j.BasicAuth(cfg.Username, cfg.Password, cfg.Realm))
	if err != nil {
		return nil, err
	}

	return &Neo4j{
		driver: driver,
	}, nil
}

func (n *Neo4j) Close() error {
	if err := n.driver.Close(); err != nil {
		return err
	}
	return nil
}
