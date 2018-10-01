package dao

// Session Data Access Object
type Session struct{}

// Create new session
func (s *Session) Create(token string) error {
	return nil
}

// Update existing session
func (s *Session) Update(oldToken, newToken string) error {
	return nil
}

// Destroy session
func (s *Session) Destroy(token string) error {
	return nil
}
