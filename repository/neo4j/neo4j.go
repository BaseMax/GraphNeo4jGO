package neo4j

import (
	"GraphNeo4jGO/config"
	"context"

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

func (n *Neo4j) Ping(ctx context.Context) error {
	res := make(chan struct{}, 1)

	go func() {
		session := n.driver.NewSession(neo4j.SessionConfig{})
		_, err := session.ReadTransaction(pingTx)
		if err != nil {
			return
		}
		res <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-res:
            // log.Println("connection ok")
			return nil
		}
	}

	return nil
}

func pingTx(tx neo4j.Transaction) (any, error) {
	_, err := tx.Run(`RETURN 1`, params{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
