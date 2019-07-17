def search(nums,target) :
    nlen=len(nums)
    begindex=0
    endindex=nlen-1
    midindex=0
    while begindex<endindex:
        midindex=(begindex+endindex)//2
        if nums[midindex]<nums[begindex]:
            endindex=midindex
        else:
            begindex=midindex+1
    return -1
a=[4,5,6,7,0,1,2]
b=0
c=search(a,b)