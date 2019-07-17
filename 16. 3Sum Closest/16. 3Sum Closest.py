import sys
def threeSumClosest(nums, target):
    nlen=len(nums)#取得长度循环使用，避免重复获取长度
    nums.sort()#对数组进行排序，当数字重复时可以不必重复计算
    result,diff=0,sys.maxsize
    for i in range(nlen-2):#从第一个元素开始计算差值，在最后两个数据无需进入循环，因为前面方案已经包括
        if i>1 and nums[i]==nums[i-1]:#跳过除第一个元素以外的重复元素，因为之前方案中已经计算完成
            continue
        left,right=i+1,nlen-1#从每个元素的右方开始作为剩下两个元素取值的集合，进行
        while left<right:
            total=nums[i]+nums[left]+nums[right]
            res=abs(total-target)
            if not res:
                return total
            if res<diff:
                diff=res
                result=total
            if total<target:
                left+=1
            else:
                right-=1
    return result
        

a=[-1, 2, 1, -4]
b=-1
print(threeSumClosest(a,b))