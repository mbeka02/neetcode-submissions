
// filterAlphanumeric removes all non-alphanumeric characters from a string.
func filterAlphanumeric(s string) string {
    // Compile a regular expression that matches any character *not* in the
    // set a-z, A-Z, or 0-9. The `^` inside the square brackets means "not".
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
    // Replace all matched non-alphanumeric characters with an empty string.
	filtered := re.ReplaceAllString(s, "")
	return filtered
}

func isPalindrome(s string) bool {
filteredString:=filterAlphanumeric(s)
fmt.Println(strings.ToLower(filteredString))
runes:= []rune(strings.ToLower(filteredString))
left:=0
right:=len(runes)-1

for left < right {
if runes[left] !=runes[right]{
return false
}
left++
right--
}
return true
}
