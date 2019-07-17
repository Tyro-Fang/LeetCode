def firstMissingPositive(nums ) :
        if not nums:
            return 1
        
        for i in range(1,len(nums)+2):
            if i not in nums:
                return i

a=[0,2,2,1,1,1,1,3]
b=firstMissingPositive(a)