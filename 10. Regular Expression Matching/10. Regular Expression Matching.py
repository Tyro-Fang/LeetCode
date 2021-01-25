def isMatch(s, p):
    #利用动规算法，需要s与p完全匹配
    #记录两个字符串长度
    slen=len(s)
    plen=len(p)
    #建立动规结果表
    T = [
            [False for _ in range(len(s)+1)]
            for _ in range(len(p)+1)
        ]
    #T=[[False]*(slen+1)]*(plen+1)
    
    #1.两个空字符串肯定相等
    T[0][0]=True
    #2.空字符串与非空字符串肯定不相等,故，T[0][i]=false,但是T[j][0]与空字符串匹配就必须有*
    for i in range(1,plen+1):
        T[i][0]=T[i-2][0] and i>1 and p[i-1]=='*'
    #3.对剩余字符串进行对比
    for i in range(1,plen+1):
        for j in range(1,slen+1):
            #3.1 s与p一一匹配情况
            if p[i-1]==s[j-1] or p[i-1]=='.':
                T[i][j]=T[i-1][j-1]
            #3.2 p是*的情况
            elif p[i-1]=='*':
                #1.不需要*即可完成匹配
                #解决'aaa*'与'aaa'的匹配
                T[i][j]=T[i-1][j]|T[i][j]
                #再解决'a'与'aa*'的匹配
                T[i][j]=T[i-2][j] and i>1
                #2.需要*来完成匹配
                if i>1 and p[i-2]==s[j-1] or p[i-2]=='.' :
                    T[i][j]=T[i][j]|T[i][j-1]
    return T[-1][-1]

print(isMatch('aaa','a'))
