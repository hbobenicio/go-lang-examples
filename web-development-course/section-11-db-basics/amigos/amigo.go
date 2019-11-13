package amigos

// Amigo models the amigo entity
type Amigo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// IsValid checks is an amigo is valid for insertion
func IsValid(a *Amigo) bool {
	if a.Name == "" {
		return false
	}
	return true
}
