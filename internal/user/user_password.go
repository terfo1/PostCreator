package user

import (
	"github.com/alexedwards/argon2id"
	"log"
)

func (p *password) Set(plaintextPassword string) error {
	hash, err := argon2id.CreateHash(plaintextPassword, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}
func (p *password) Matches(plaintextPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(plaintextPassword, p.hash)
	if err != nil {
		log.Fatal(err)
	}
	return match, nil
}
