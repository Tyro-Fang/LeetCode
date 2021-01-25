def removeDuplicates(nums):
        length=len(nums)
        if length==0:
            return 0
        elif length==1:
            return 1
        else:
            loc=1
            compare=nums[0]
            for i in range(1,length):
                if nums[i]!=compare:
                    nums[loc]=nums[i]
                    compare=nums[i]
                    loc+=1
            return loc
a=[1,1,2]
removeDuplicates(a)