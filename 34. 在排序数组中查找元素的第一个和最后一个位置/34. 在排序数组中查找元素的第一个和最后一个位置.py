from typing import List
class Solution:
    def searchRange(self,nums: List[int], target: int):
        nlen=len(nums)
        i,j=0,nlen-1
        if nlen==0 or nums[0]>target or nums[nlen-1]<target:
            return [-1,-1]
        while i<j:#find min val
            m=(i+j)//2
            if nums[m]>=target:
                j=m
            elif nums[m]<target:
                i=m+1
        if nums[i]!=target:
            return [-1,-1]
        k,j=i,l-1
        while k+1<j:#find max val
            m=(k+j)//2
            if nums[m]<=target:
                k=m
            elif nums[m]>target:
                j=m-1
        if k+1<nlen and nums[k+1]==target:
            return [i,k+1]
        else:
            return[i,k]
    def __init__(self,x,k):
        a=self.searchRange(x,k)
        print(a)
i=[5,7,7,8,8,10]
j=8
a=Solution(i,j)
print(a)
