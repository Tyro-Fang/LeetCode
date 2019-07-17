def countAndSay(n):
        i=1
        r='1'
        while i<n:
            rlen=len(r)
            a=0
            res=''
            while a<rlen:
                p=1
                while a+p<rlen:
                    if r[a]==r[a+p]:
                        p+=1
                        continue
                    else:
                        res=res+str(p)+r[a]
                        #a+=p
                        break
                if a+p>=rlen:
                    res=res+str(p)+r[a]
                a+=p
            r=res
            i+=1
        return r

a=countAndSay(6)
print(a)