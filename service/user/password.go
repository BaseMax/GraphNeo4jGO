package user

import "golang.org/x/crypto/bcrypt"

func hashPassword(pass string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func comparePassword(old, new string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(old), []byte(new))
	return err == nil
}
