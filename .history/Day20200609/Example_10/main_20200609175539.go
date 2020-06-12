package main

import "sync"

type UserAges struct {
	ages map[string]int
	sync.Mutex
}