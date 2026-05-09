func lengthOfLongestSubstring(s string) int {
    occurrences := make(map[byte]int)
    maxLength := 0
    l := 0

    for r := 0; r < len(s); r++ {
        key := s[r]

        // Shrink window from the left until duplicate is removed
        for occurrences[key] > 0 {
            occurrences[s[l]]--
            l++
        }

        occurrences[key]++
        maxLength = max(maxLength, r-l+1)
    }

    return maxLength
}