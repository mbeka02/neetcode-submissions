func characterReplacement(s string, k int) int {
if len(s)==0{
return 0
}
  // frequency of each letter in current window
    var freq [26]int
    left := 0
    maxLen := 0

    // current max frequency in the window (helps decide feasibility)
    maxCount := 0

    for right := 0; right < len(s); right++ {
        ch := s[right] - 'A'
        freq[ch]++
        if freq[ch] > maxCount {
            maxCount = freq[ch]
        }

        // If window size - maxCount > k, window can't be uniform with at most k changes
        for (right-left+1)-maxCount > k {
            rem := s[left] - 'A'
            freq[rem]--
            left++
            // Recompute maxCount to be safe (O(26) each time we shrink)
            maxCount = 0
            for i := 0; i < 26; i++ {
                if freq[i] > maxCount {
                    maxCount = freq[i]
                }
            }
        }

        // Update answer
        curLen := right - left + 1
        if curLen > maxLen {
            maxLen = curLen
        }
    }

    return maxLen
}
