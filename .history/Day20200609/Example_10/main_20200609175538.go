package main

type UserAges struct {
	ages map[string]int
	sync.Mu
}