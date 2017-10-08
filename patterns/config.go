package patterns

import(
	"time";
  "gopkg.in/mgo.v2";
)

const(
  dbip                     = "localhost"
  dbport                   = "27017"
  Dbname                   = "Myapi"
  Users_collection         = "users"
  urldb                    = dbip+":"+dbport
  authUserName             = "admin"
  authPassword             = "password"
)

var mongoDBDialInfo = &mgo.DialInfo{
  Addrs:    []string{urldb},
  Timeout:  60 * time.Second,
  Database: Dbname,
  Username: authUserName,
  Password: authPassword,
}
