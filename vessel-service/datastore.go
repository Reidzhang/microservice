package main

import (
	"gopkg.in/mgo.v2"
)

// CreateSession creates mongodb instance
func CreateSession(host string) (*mgo.session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}