package service

import "sync"

type Service struct {
	mutex sync.Mutex
}
