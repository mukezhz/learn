a = [1,2,3,4,5,1]
h: dict[int, bool] = {}

def contain_duplicate(a: list[int], h: dict[int, bool]):
    for i in a:
        if i in h:
            return True
        else:
            h[i] = True
    else:
        return False
    
print(contain_duplicate(a, h))