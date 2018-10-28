package main

import "gopkg.in/mgo.v2"

func CreateSession(host string) (*mgo.Session, error) {
	session, error := mgo.Dial(host)
	if error != nil {
		return nil, error
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
