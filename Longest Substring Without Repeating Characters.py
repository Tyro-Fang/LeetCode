import time
def findstr(x):
    start = time.clock()    
    l1=list(x)
    dict1={}
    if len(l1)==0:
        target=0
    else:
        target=1
    j=0
    for i in range(len(l1)):
        if l1[i] in dict1:
            cm=len(''.join((list(dict1.keys()))))
            if target<cm:
                target=cm
            l=int(dict1[l1[i]])
            for k in range(j,l):
                dict1.pop(l1[k])
            dict1[l1[i]]=i
            j=l+1            
        else:
            dict1[l1[i]]=i
    cm=len(''.join((list(dict1.keys()))))
    if target<cm:
        target=cm
    end = time.clock()
    print(end-start)   
    return target
def find2(s):
    start = time.clock()
    s_letter = set([w for w in s])
    s_dict = dict.fromkeys(s_letter,-1) 
    tmp_len = 0
    record_len = 0
    start_point = -1
    for i,le in enumerate(s):
        tmp_point = s_dict[le]
        if tmp_point <= start_point:
            s_dict[le] = i
            tmp_len += 1
        else:
            start_point = tmp_point
            s_dict[le] = i
            if tmp_len > record_len:
                record_len = tmp_len
            tmp_len = i-tmp_point
    if tmp_len > record_len:
        record_len = tmp_len
    end = time.clock()
    print(end-start)
    return record_len

print(findstr("abcabcbbdsacascacvewvwecadscascdavijdfsanabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnvishanvn adsjivnsanvoadsn"))
print(find2("abcabcbbdsacascacvewvwecadscascdavijdfsanabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnabcabcbbdsacascacvewvwecadscascdavijdfsanvishanvn adsjivnsanvoadsnvishanvn adsjivnsanvoadsn"))

