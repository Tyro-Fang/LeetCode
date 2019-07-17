def twoSum(nums, target) :
        nums.sort()
        nlen=len(nums)
        if nlen<2:
            return [None,None]
        for i in range(nlen-1):
            if nums[i]+nums[nlen-1]<target:
                continue
            if nums[i]+nums[i+1]>target:
                return [None,None]
            for j in range(i+1,nlen):
                if nums[i]+nums[j]==target:
                    l1=[]
                    l1.append(i)
                    l1.append(j)
                    return l1
        return [None,None]

a=[0,4,3,0]
b=0
c=twoSum(a,b)
print(c)