package auth

import "fmt"

// capital letter means public function
// LoginWithCredentials is a public function
func LoginWithCredentials(username string, password string) {
	fmt.Println("Logging in with credentials", username, password)
}
