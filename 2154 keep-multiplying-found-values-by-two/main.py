from typing import List

class Solution:
    def findFinalValue(self, nums: List[int], original: int) -> int:
        while original in nums:
            original = original*2
        return original

nums = [2,7,9]
original = 4
sol = Solution()
print(sol.findFinalValue(nums, original))
