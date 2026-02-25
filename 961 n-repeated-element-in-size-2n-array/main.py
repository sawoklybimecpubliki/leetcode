from typing import List


class Solution:
    def repeatedNTimes(self, nums: List[int]) -> int:
        n = len(nums)/2
        for i in range(len(nums)):
            if nums.count(nums[i]) == n:
                return nums[i]

"""
        n = len(nums)/2
        d = defaultdict(int)
        for i in range(len(nums)):
            d[nums[i]]+=1
            if d[nums[i]] == n:
                return nums[i]
"""

nums = [2,1,2,5,3,2]
solution = Solution()
print(solution.repeatedNTimes(nums))