def romanToInt(s):
    dictR={'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
    lenOfS=len(s)
    i=lenOfS-1
    result=0
    while i>=0:
        result=result+dictR[s[i]]
        if i>0:
            if dictR[s[i]]>dictR[s[i-1]]:
                result-=dictR[s[i-1]]
                i-=2
            else: i-=1
        else: i-=1
    return result


print(romanToInt('XXVII'))

    
    
        

