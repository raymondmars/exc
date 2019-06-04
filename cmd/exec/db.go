package main

var Db = Storage{}

func init() {
	Db.LoadAll()
}
