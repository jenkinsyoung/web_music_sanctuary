classes = [
    ('classical', 'guitar_form'), 
    ('concert', 'guitar_form'), 
    ('dreadnought', 'guitar_form'), 
    ('es', 'guitar_form'), 
    ('explorer', 'guitar_form'), 
    ('flying-v', 'guitar_form'), 
    ('h-pickup', 'pickup_config'), 
    ('hh-pickup', 'pickup_config'), 
    ('hs-pickup', 'pickup_config'), 
    ('hsh-pickup', 'pickup_config'), 
    ('les-paul', 'guitar_form'), 
    ('mustang', 'guitar_form'), 
    ('parlor', 'guitar_form'), 
    ('sg', 'guitar_form'), 
    ('ss-pickup', 'pickup_config'), 
    ('ssh-pickup', 'pickup_config'), 
    ('sss-pickup', 'pickup_config'), 
    ('stratocaster', 'guitar_form'), 
    ('superstrat', 'guitar_form'), 
    ('telecaster', 'guitar_form'), 
    ('triple-o', 'guitar_form')
]
source = 'src/yolov9/runs/detect/exp4/labels/image.txt'

def parsepred() -> dict:
    res = {}
    detections = {}
    
    with open(source) as f:
        for line in f:
            classid, coef1, coef2, coef3, coef4 = line.split()
            try:
                detections[int(classid)] = max(detections[int(classid)], (coef1, coef2, coef3, coef4), key=lambda x: x[1])
            except Exception:
                detections[int(classid)] = (coef1, coef2, coef3, coef4)
                
    with open('src/yolov9/runs/detect/exp/labels/image.txt', 'w') as f:
        f.write('')
    
    for d in detections:
        try:
            res[classes[d][1]] = max((d, detections[d][1]), res[classes[d][1]], key= lambda x: x[1])
        except Exception:
            res[classes[d][1]] = (d, detections[d][1])
        
    for k, v in res.items(): # example: {'guitar_form': (18, '0.51875'), 'pickup_config': (7, '0.626042')}
        res[k] = classes[v[0]][0]
    
    return res
