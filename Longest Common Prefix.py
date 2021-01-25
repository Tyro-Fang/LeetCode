def longestCommonPrefix(strs):
    listLen=len(strs)
    if listLen<1:
        return ''
    elif listLen<2:
        return strs[0]
    else :
        result=strs[0]
        for i in range(1,listLen):
            temp=''
            for j in range (min(len(result),len(strs[i]))):
                if result[j]==strs[i][j]:
                    temp+=strs[0][j]
                else: 
                    if j==0:
                        result=''
                    break
            result=temp
            if result=='':
                break
        return result

def moreEfficentWay(strs):
    s1=min(strs)
    s2=max(strs)
    result=''
    for i,c in enumerate(s1):
        if c==s2[i]:
            result+=c
        else: break
    return result
                
print(moreEfficentWay(["dog","dogecar","dor"]))     
        

