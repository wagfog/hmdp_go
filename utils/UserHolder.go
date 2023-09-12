package utils

import (
	"sync"

	"github.com/wagfog/hmdp_go/dto"
)

type UserHolder struct {
	userMap map[int]dto.UserDTO
	mux     sync.RWMutex
}

// SaveUser saves the user data for the current goroutine
func (uh *UserHolder) SaveUser(user dto.UserDTO) {
	goroutineID := getGoroutineID() // get a unique identifier for the current goroutine

	uh.mux.Lock()
	defer uh.mux.Unlock()

	if uh.userMap == nil {
		uh.userMap = make(map[int]dto.UserDTO)
	}

	uh.userMap[goroutineID] = user
}

func (uh *UserHolder) GetUser() (dto.UserDTO, bool) {
	goroutineID := getGoroutineID()

	uh.mux.RLock()
	defer uh.mux.RUnlock()

	user, exists := uh.userMap[goroutineID]
	return user, exists
}

// RemoveUser removes the user data for the current goroutine
func (uh *UserHolder) RemoveUser() {
	goroutineID := getGoroutineID()

	uh.mux.Lock()
	defer uh.mux.Unlock()

	delete(uh.userMap, goroutineID)
}

// Utility function to get a unique identifier for the current goroutine
func getGoroutineID() int {
	// Implement your own logic here to generate a unique identifier
	// This can be based on the current goroutine's ID or any other mechanism you prefer

	return 0 // Placeholder value, replace it with your own logic
}
