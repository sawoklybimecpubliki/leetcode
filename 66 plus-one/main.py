from typing import List

class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        i = len(digits) - 1
        digits[i] += 1
        while digits[i] == 10 and i >= 0:
            digits[i] = 0
            if i > 0:
                digits[i-1] += 1
            else:
                digits.insert(0, 1)
            i-=1
        return digits


digits = [8,9,9,9]
solution = Solution()
print("\n", solution.plusOne(digits))
