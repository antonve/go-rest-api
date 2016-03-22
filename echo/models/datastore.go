package models

import (
    "log"
    "gopkg.in/mgo.v2"
)

var session *mgo.Session

func init() {
    var err error
    session, err = mgo.Dial("localhost:27017")
    if err != nil {
        log.Fatalln("Couldn't connect to data store")
    }
}

type DataStore struct {
    session *mgo.Session
}

func (d *DataStore) Close() {
    d.session.Close()
}

func (d *DataStore) C(name string) *mgo.Collection {
    return d.session.DB("productdb").C(name)
}

func NewDataStore() *DataStore {
    ds := &DataStore{
        session: session.Copy(),
    }
    return ds
}
