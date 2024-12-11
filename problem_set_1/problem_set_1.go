package main

import "fmt"

// Helper function to check if a string is a palindrome
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Helper function to reverse a string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Helper function to remove duplicate values from a slice of int slices
func removeDuplicates(slice [][]int) [][]int {
	seen := make(map[string]bool)
	result := [][]int{}
	for _, v := range slice {
		key := fmt.Sprintf("%d_%d", v[0], v[1])
		if !seen[key] {
			result = append(result, v)
			seen[key] = true
		}
	}
	return result
}

// Main function to find all palindrome pairs
func palindromePairs(words []string) [][]int {
	wordDict := make(map[string]int)
	for i, word := range words {
		wordDict[word] = i
	}

	result := [][]int{}

	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			prefix, suffix := word[:j], word[j:]

			// If the suffix is a palindrome, check if the reversed prefix exists
			if isPalindrome(suffix) {
				reversePrefix := reverse(prefix)
				if idx, found := wordDict[reversePrefix]; found && idx != i {
					result = append(result, []int{i, idx})
				}
			}

			// If the prefix is a palindrome, check if the reversed suffix exists
			if j != len(word) && isPalindrome(prefix) {
				reverseSuffix := reverse(suffix)
				if idx, found := wordDict[reverseSuffix]; found && idx != i {
					result = append(result, []int{wordDict[reverseSuffix], i})
				}
			}
		}
	}
	return removeDuplicates(result) // Add this line to remove duplicates
}

// Main function for testing
func main() {
	inputWords := []string{"bat", "tab", "cat"}
	result := palindromePairs(inputWords)
	fmt.Println(result) // Output: [[0 1] [1 0]]
}
