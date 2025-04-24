'''
Given n non-negative integers representing an elevation
map where the width of each bar is 1, compute how much
water it can trap after raining.
'''

class Solution:
    def trap(self, height: List[int]) -> int:
        n = len(height)
        left_walls = [0] * n
        max_wall = 0
        for i, h in enumerate(height):
            left_walls[i] = max_wall
            max_wall = max(max_wall, h)
        right_walls = [0] * n
        max_wall = 0
        for i in range(n-1,-1,-1):
            right_walls[i] = max_wall
            max_wall = max(max_wall, height[i])
        total_water = 0
        for i, h in enumerate(height):
            water = min(left_walls[i], right_walls[i]) - h
            if water > 0:
                total_water += water
        return total_water