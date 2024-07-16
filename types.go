package main

type ClientCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	Next     *string
	Previous *string
}
