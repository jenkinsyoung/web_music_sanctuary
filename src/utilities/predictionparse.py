def parsepred() -> dict:
    res = {}
    
    with open('src/yolov9/runs/detect/exp/labels/image.txt') as f:
        for line in f:
            classid, coef1, coef2, coef3, coef4, coef5 = line.split()
            try:
                res[int(classid)].append((coef1, coef2, coef3, coef4, coef5))
            except Exception:
                res[int(classid)] = [(coef1, coef2, coef3, coef4, coef5)]
                
    with open('src/yolov9/runs/detect/exp/labels/image.txt', 'w') as f:
        f.write('')

    return res
