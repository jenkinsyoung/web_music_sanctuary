from http.server import BaseHTTPRequestHandler
from http.server import HTTPServer
import utilities.encode as b64
from utilities.predictionparse import parsepred
from yolov9 import detect

import sys
import json
import cgi

class Server(BaseHTTPRequestHandler):
    def _set_headers(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        
    def do_HEAD(self):
        self._set_headers()
        
    # GET sends back a Hello world message
    def do_GET(self):
        self._set_headers()
        self.wfile.write(json.dumps({'hello': 'world', 'received': 'ok'}))
        
    # POST echoes the message adding a JSON field
    def do_POST(self):
        ctype, pdict = cgi.parse_header(self.headers.get('content-type'))
        
        # refuse to receive non-json content
        if ctype != 'application/json':
            self.send_response(400)
            self.end_headers()
            return
            
        # read the message and convert it into a python dictionary
        length = int(self.headers.get('content-length'))
        message = json.loads(self.rfile.read(length))
        
        imagebytes = message["image"]
        b64.decode_image(imagebytes)
        opt = {
            "device": 0,
            "save_txt": True,
            "weights": "src/yolov9/gelan-c.pt",
            "source": 'src/yolov9/data/images/image.png',
            "exist_ok": True
        }
        detect.main(opt)
        try:
            del message["image"]
        except KeyError:
            pass
        responce = parsepred()
        # send the message back
        self._set_headers()
        self.wfile.write(json.dumps(responce).encode())
        
def run(server_class=HTTPServer, handler_class=Server, port=8008):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    
    print (f'Starting httpd on port {port}...')
    httpd.serve_forever()
    

if __name__ == "__main__":
    if len(sys.argv) == 2:
        run(port=int(sys.argv[1]))
    else:
        run()