class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if not nums or target == None:
            return -1
        nLen = len(nums)
        left = 0
        right = nLen - 1
        while left + 1 < right:
            mid = left + (right - left) // 2
            if nums[mid] ==  target:
                return mid
            if nums[mid] > nums[left]:
                if nums[left] <= target < nums[mid]:
                    right = mid - 1
                else:
                    left = mid + 1
            else:
                if nums[mid] < target <= nums[right]:
                    left = mid + 1
                else:
                    right = mid - 1
        if nums[left] == target:
            return left
        if nums[right] == target:
            return right
        return -1