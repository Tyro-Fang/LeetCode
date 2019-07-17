def strStr(haystack: str, needle: str) -> int:
        hlen=len(haystack)
        nlen=len(needle)
        if hlen<nlen:
            return -1
        if hlen==0 and nlen==0:
            return 0
        for i in range(hlen):
            if i+nlen>hlen:
                break
            a=i
            b=0
            for j in range(nlen):
                if a>=hlen or haystack[a]!=needle[j]  :
                    break
                else:
                    a+=1
                b+=1
            if b==nlen:
                return i
        return -1

a="mississippi"
b="issipi"
c=strStr(a,b)