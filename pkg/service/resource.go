package service

import "sync"

type ResourceService struct {
	// store interface
	mutex sync.Mutex
}
