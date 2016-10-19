import serial

# A Serial-based peripheral board
class Peripheral:
    ser = None

    def __init__(self, port, baud, timeout):
        self.ser = serial.Serial(port, baud, timeout=timeout)

    def send(self, command):
        tosend = (command + '\n').encode('ascii')
        self.ser.write(tosend)