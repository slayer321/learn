# Time O(N) space O(2N) 



# def trap(self, height: List[int]) -> int:
#     left = []
#     max_left = 0
#     n = len(height)
#     for i in range(n):
#         max_left = max(max_left, height[i])
#         left.append(max_left)

#     right = []
#     max_right = 0
#     for i in range(n-1, -1, -1):
#         max_right = max(max_right, height[i])
#         right.append(max_right)
#     right.reverse()

#     water = []
#     for i in range(n):
#         val = min(left[i], right[i]) - height[i]
#         water.append(val)

#     res = 0
#     for i in range(len(water)):
#         res += water[i]

#     return res





