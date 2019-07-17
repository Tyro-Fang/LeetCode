def allSolution():
    path = 'D:\\ChangZhi\\A.in'
    nums=getNums(path)
    result=[]
    for i in range(len(nums)):
        num=nums[i]
        num.sort()
        count=getCount(num)
        #print(count)
        result.append("Case #"+str(i+1)+": "+str(count)+"\n")
    f=open('D:\\ChangZhi\\a.out','w')
    f.writelines(result)
    f.close()

def getCount(num):
    count = 0
    for i in range(len(num)-2):
        a = num[i]
        begin = i + 1
        end = begin + 1
        while end <len(num):
            if a*num[begin]<num[end]:
                break
            elif a*num[begin]>num[end]:
                if end < len(num) - 1:
                    end+=1
                    while (num[end]==num[end-1]) and (end<len(num)-1):
                        end+=1
                else:break
            else:
                count+=1
                begin += 1
                end += 1
                if begin<len(num)-2:
                    while (num[begin]==num[begin-1]) and (begin<len(num)-1):
                        begin+=1
                if end < len(num)-1 :
                    while (num[end]==num[end-1]) and (end<len(num)-1):
                        end+=1
            if end>len(num)-1:
                break
        if end ==len(num)-1 and a!=1:
            break
    return count

def getNums(path):
    f=open(path,'r')
    r=f.readlines()
    nums=[]
    for i in range(1,len(r)):
        if i%2!=0:
            continue
        line=r[i].split('\n')
        num=list(map(int,line[0].split(' ')))
        nums.append(num)
    f.close()
    return nums

allSolution()