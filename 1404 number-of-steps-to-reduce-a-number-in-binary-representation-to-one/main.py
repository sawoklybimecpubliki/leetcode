class Solution:
    def numSteps(self, s: str) -> int:
        num = self.sqrt(s)
        c = 0
        while num != 1:
            c += 1
            if num % 2 == 0:
                num //= 2
            else:
                num += 1
        return c
    def sqrt(self, s: str) -> int:
        num = 0
        for i in range(len(s)):
            if s[i] == '1':
                num = num + 2 ** (len(s) - 1 - i)
        return num
s = '10'
sol = Solution()
print(sol.numSteps(s))