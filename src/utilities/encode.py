from base64 import b64decode

def decode_image(image : str):
    with open("src/yolov9/data/images/horses.png", "wb") as fh:
        fh.write(b64decode(image))
