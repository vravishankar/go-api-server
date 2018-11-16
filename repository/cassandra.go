package repository

import (
	"log"
	"os"

	"../config"

	"github.com/gocql/gocql"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

// InitDB - Initialize the Database
func init() {
	var err error
	config := config.GetConfig()

	cluster := gocql.NewCluster(config.DB.Cluster)
	cluster.Keyspace = config.DB.Keyspace
	Session, err = cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}
	Session.SetTrace(gocql.NewTraceWriter(Session, os.Stdout))
	log.Println("Connected to Database...")
}
