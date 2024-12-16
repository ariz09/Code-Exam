Problem: Valid Parentheses
Problem Description:
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. An input string is valid if:
1.	Open brackets must be closed by the same type of brackets.
2.	Open brackets must be closed in the correct order.

Examples:
Input: ()[]{}
Output: True

Input: ([)]
Output: False

Instructions:
1.	Bracket Matching:
a.	Pay attention to matching different types of brackets correctly.
2.	Order of Closing:
a.	Ensure that open brackets are closed in the correct order.
3.	Output Format:
a.	The output should be a boolean indicating whether the input string is valid or not.
4.	Test Your Solution:
a.	Test your solution with various strings to verify its correctness.

isValid(s string) bool to check if a string containing brackets (parentheses, braces, and square brackets) is valid. 
It uses a stack to track opening brackets and ensures that each closing bracket matches the most recent opening bracket. 
If the stack is empty at the end, the string is valid; otherwise, it's invalid.
