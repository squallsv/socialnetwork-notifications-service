package DbConnection

import (
	"github.com/gocql/gocql"
	"fmt"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "notifications"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
