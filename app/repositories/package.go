package repositories

import (
	"labix.org/v2/mgo"
)

const (
	DBBitoolaPolls = "bitoola-polls"

	ColEmails = "emails"
	ColUsers  = "users"
)

var (
	sess *mgo.Session
)

func getSession() *mgo.Session {
	if sess == nil {
		var err error
		sess, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		sess.SetMode(mgo.Monotonic, true)
	}
	return sess.Copy()
}
