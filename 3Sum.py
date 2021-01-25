def threeSum(nums):
    result=[]
    dictN=dict()
    for i in nums:
        if  i not in dictN:
            dictN[i]=0
        dictN[i]+=1
    if 0 in dictN and dictN[0]>2:
        result=[[0,0,0]]
    pos=[p for p in dictN if p >0]
    neg=[n for n in dictN if n <0]
    for p in pos:
        for n in neg:
            i=0-p-n
            if i not in dictN:
                continue
            if i==0 :
                result.append([n,0,p])
            elif i==p and dictN[i]>1:
                result.append([n,i,p])
            elif i==n and dictN[i]>1:
                result.append([n,i,p])
            elif i>p :
                result.append([n, p, i])
            elif i<n :
                result.append([i,n,p])
    return result
    

nums=[0,0,0]
print(threeSum(nums))
        