def convert(s,numrows):
    if len(s)<=numrows:
            return s
    if numrows==1:
            return s
    result=''
    for nrow in range(numrows):
        loc=nrow
        result+=s[nrow]
        for i in range(len(s)):
            if i%2==0:
                ldown=(numrows-nrow-1)*2
                if ldown>0:
                    loc+=ldown
                    if loc>=len(s):
                        break
                    result+=s[loc]
            else:
                lup=nrow*2
                if lup>0:
                    loc+=lup
                    if loc>=len(s):
                        break
                    result+=s[loc]
    return result


def conv(s,numRows):
    if numRows <= 1:
            return s
    ret = [''] * numRows
    grp = 2 * numRows - 2

    for i, c in enumerate(s):
        cur = i % grp
        if cur < numRows:
            ret[cur] += c
        else:
            ret[grp - cur] += c

    return ''.join(ret)
print(conv('PAYPALISHIRING',3))
            

            
            

        
    
    







