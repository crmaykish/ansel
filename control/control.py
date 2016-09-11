import serial
import time

def command(command):
    tosend = (command + '\n').encode('ascii');
    ser.write(tosend)

print("Ansel control is active")

ser = serial.Serial("/dev/ttyUSB0", 115200, timeout=10)
print("Connected to serial")

# TODO: serial handshake to confirm two way communication

print("Starting Ansel control loop")
while True:
    command("d,forward,150")
    time.sleep(5);
    command("stop,")
    time.sleep(1)
    command("d,reverse,150")
    time.sleep(5)
    command("stop, ")
    time.sleep(1)
