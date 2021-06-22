package token

import "time"

type Maker interface {
	// create new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// check input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
