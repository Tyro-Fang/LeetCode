def fourSum(nums,target):
    nums=sorted(nums)
    nlen=len(nums)
    result=[]
    for i in range(nlen-2):
        if i>0 and nums[i]==nums[i-1]:
            continue
        three_sum=target-nums[i]
        if nums[i]+nums[i+1]*3>target:#添加判断，当所选数字与有可能出现的最小数字和大于目标时，之后的结果肯定大于目标，无需再进行判断
            break
        if nums[i]+3*nums[nlen-1]<target:#添加判断，当所选数字与有可能出现的最大数字都小于目标时，当前数字肯定无法达成目标，进入下一数字判断
            continue
        for j in range(i+1,nlen-1):
            if j>i+1 and nums[j]==nums[j-1]:
                continue
            if nums[i]+3*nums[j]>target:
                break
            left,right=j+1,nlen-1
            while left<right:
                if 2*nums[left]+nums[j]>three_sum:
                    break
                if left>j+1 and nums[left]==nums[left-1]:
                    left+=1
                    continue
                if right<nlen-1 and nums[right]==nums[right+1]:
                    right-=1
                    continue
                total=nums[left]+nums[right]+nums[j]
                if total==three_sum:
                    a=[nums[i],nums[j],nums[left],nums[right]]
                    result.append(a)
                    left+=1
                    right-=1
                elif total>three_sum:
                    right-=1
                else:
                    left+=1
    return result
a=[-1,0,-5,-2,-2,-4,0,1,-2]
b=-9
print(fourSum(a,b))




                 
