package main

var Db = Storage{}

//to load all commands
func init() {
	Db.LoadAll()
}
