def findls(s):
    l=list(s)
    lr=l[::-1]
    maxlen=len(l)
    al=maxlen
    for i in range(maxlen):
        times=len(l)-maxlen+1
        for j in range(times):
            l1=l[j:j+maxlen]
            if l1==lr[times-1-j:times-1-j+maxlen]:
                return "".join(l1)
        maxlen=maxlen-1
    return ""


def find2(s):
    l=list(s)
    maxl=0
    mb=0
    for i in range(len(l)):
        maxlen=2
        mbegin=i
        l1=l[i:i+maxlen]
        l2=l1[::-1]
        while l1==l2:
            maxlen=maxlen+2
            l1=l[i-1:i-1+maxlen]
            l2=l1[::-1]
        if maxl<maxlen:
            mb=mbegin
            maxl=maxlen
        maxlen=3
        
        l1=l[i:i+maxlen]
        l2=l1[::-1]
        while l1==l2:
            maxlen=maxlen+2
            l1=l[i-1:i-1+maxlen]
            l2=l1[::-1]
        if maxl<maxlen:
            mb=mbegin
            maxl=maxlen
        if  maxlen+mbegin>=len(l):
            break
        
    return l[mb:mb+maxl]


s="tembwxtvddsttolegryohdlhxhycymqybzfzcbkzdwcekzfapmpfyeivfoeeqoqdhylziqpnyzuzeeutpqpalbddlliuehvkhqevgjdkskvphidcjmpcmetzwqkzcnxjcjywhfzplntbkuddmbcovearburjqyirbladcrhfkfdfgsmyhdsfmmxmslwkymkgaguilxghmfgaldcogtfnbqakctqtqakupwrxkmbjpmzqngwldmaugzizgwmediyzxevspxdwruyzrmnhchtxlgtb"
s1="babad"
print(find2(s1))
