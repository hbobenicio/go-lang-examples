package repo

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
)

// DBName contém o nome do banco
const DBName = "section-15-mongodb"

// NewDBSession inicializa a sessão com o banco MongoDB
func NewDBSession() (*mgo.Session, error) {
	log.Println("db session: initializing...")

	// mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	host := "localhost"
	user := "root"
	password := "root"
	port := "27017"
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)

	// TODO analisar o uso do mgo.DialInfo e mgo.DialWithInfo(info)
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("repo: initdb: %v", err)
	}

	session.SetPoolLimit(10)
	session.SetPoolTimeout(2 * time.Minute)
	session.SetSyncTimeout(30 * time.Second)

	log.Println("db session: initiallized successfully!")

	return session, nil
}
