classes = [
    ('classical', 'guitar_form', 'acoustic'), 
    ('concert', 'guitar_form', 'acoustic'), 
    ('dreadnought', 'guitar_form', 'acoustic'), 
    ('es', 'guitar_form', 'electric'), 
    ('explorer', 'guitar_form', 'electric'), 
    ('flying-v', 'guitar_form', 'electric'), 
    ('h-pickup', 'pickup_config'), 
    ('hh-pickup', 'pickup_config'), 
    ('hs-pickup', 'pickup_config'), 
    ('hsh-pickup', 'pickup_config'), 
    ('les-paul', 'guitar_form', 'electric'), 
    ('mustang', 'guitar_form', 'electric'), 
    ('parlor', 'guitar_form', 'electric'), 
    ('sg', 'guitar_form', 'electric'), 
    ('ss-pickup', 'pickup_config'), 
    ('ssh-pickup', 'pickup_config'), 
    ('sss-pickup', 'pickup_config'), 
    ('stratocaster', 'guitar_form', 'electric'), 
    ('superstrat', 'guitar_form', 'electric'), 
    ('telecaster', 'guitar_form', 'electric'), 
    ('triple-o', 'guitar_form', 'acoustic')
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
            
    res['category'] = classes[res['guitar_form'][0]][2]
        
    for k, v in res.items(): # example: {'guitar_form': (18, '0.51875'), 'pickup_config': (7, '0.626042')}
        if k == 'category': continue
        res[k] = classes[v[0]][0]
        
    try:
        res['pickup_config']
    except KeyError:
        res['pickup_config'] = None
    
    return res
