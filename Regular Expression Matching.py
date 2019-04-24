def isMatch(s, p):
    flag = False
    sloc = 0
    ploc = 0
    for i in range(len(p)):
        if ploc + 1 < len(p):
            if p[ploc + 1] == '*':
                if (p[ploc] != s[sloc]) and (p[ploc] != '.'):
                    ploc = ploc + 2
                else:
                    while (p[ploc] == s[sloc]) or p[ploc] == '.':
                        sloc = sloc + 1
                        if sloc >= len(s):
                            flag = True
                            break
        else:
            if p[ploc] != s[sloc]:
                falg = False
        if ploc >= len(p):
            break
    if sloc >= len(s):
        flag = True
    return flag
print(isMatch('aa', 'a*'))
