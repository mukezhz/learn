"""
    Given an array of integers, return indices of the two numbers such that they add up to a specific target.
    Example:
    Given nums = [2, 11, 7, 15], target = 9,
    Because nums[0] + nums[1] = 2 + 7 = 9,
    reutrn [0, 1]
"""

"""
    Solution 1: Brute Force
    Time Complexity: O(n^2)
    Space Complexity: O(1)
    
    Solution 2: Two-pass Hash Table
    Time Complexity: O(n)
    Space Complexity: O(n)
    
    Solution 3: One-pass Hash Table
    Time Complexity: O(n)
    Space Complexity: O(n)
    
    Solution 4: Two-pointer
    Time Complexity: O(nlogn)
    Space Complexity: O(1)
    
    Solution 5: Binary Search
    Time Complexity: O(nlogn)
    Space Complexity: O(1)
    
Solution: using hash table
2
7
11
15

step 1: 2
9 - 2 = 7
{2: 0}

step 2: 11
9 - 11 = -2
{2: 0, 11: 1}

step 3: 7
9 - 7 = 2 which already exits in hash table
{2: 0, 11: 1, 7: 2}

end

    
"""


a = [2, 11, 7, 15]
target = 9

def two_sum(nums: list[int], target: int)-> list[int]:
    h: dict[int, int] = {}
    for i, num in enumerate(nums):
        n = target - num
        if num > target:
            continue
        if n in h:
            return [h[n], i]
        h[num] = i
    return []

print(two_sum(a, target))