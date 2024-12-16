Problem: Longest Increasing Subsequence
Given an unsorted array of integers, find the length of the longest increasing subsequence.

For context. The "Longest Increasing Subsequence" is a common problem in computer science and dynamic programming. 
In the context of an array of integers, the goal is to find the length of the longest subsequence of a given array such 
that all elements of the subsequence are sorted in increasing order.

Here's a more detailed explanation:
Subsequence: A subsequence is a sequence of elements that appear in the same order as they are in the original sequence,
but not necessarily consecutively.
Increasing Subsequence: An increasing subsequence is a subsequence in which the elements are in strictly increasing order.

The "Longest Increasing Subsequence" problem asks for the length of the longest increasing subsequence in a given array.
For example, given the array: [10, 9, 2, 5, 3, 7, 101, 18]

One possible increasing subsequence is [2,5,7,18], and the length of this subsequence is 4.
The goal is to find the length of the longest increasing subsequence for a given array.

Solving this problem efficiently often involves dynamic programming techniques, where you build up a solution
for each subproblem and use those solutions to solve the overall problem.

Example:
Input: [10,9,2,5,3,7,101,18]
Output: 4

Edge Case:

If the input array nums is empty, return 0 since there's no subsequence.
Initialization:

A DP array dp is created with the same length as nums, where each element is initially set to 
1. This represents the minimum subsequence length (a subsequence of a single element).
Filling the DP Array:

The program iterates through the array starting from the second element. 
For each element nums[i], it checks all previous elements (nums[j] where j < i).
If nums[i] is greater than nums[j], it updates dp[i] to the maximum of its current value and dp[j] + 1,
meaning it extends the subsequence ending at j by including nums[i].

Finding the Result:
After processing all elements, the length of the longest increasing subsequence is the maximum value in the dp array.
