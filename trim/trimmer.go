package trim

import "regexp"

/*GetTrimmed generates and returns a shorter link from longURL
It returns an error if trimming fails for some reason*/
func GetTrimmed(longURL string) (string, error) {

}

/*isValidURL checks if a url is valid or not.
Returns true if it is, else false*/
func isValidURL(url string) bool {
	matcher := regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	return matcher.MatchString(url)
}