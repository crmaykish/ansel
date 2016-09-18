# Ansel Sensor Board - Colin Maykish - September 2016
import serial
import time
from threading import Thread

# TODO: Serial handshake to confirm two way communication
# TODO: Call/response system vs async communication?
# TODO: Drive board resetting after start? Doesn't work without serial monitor open

sensors = None

# Send a command to a serial device
def command(board, command):
    tosend = (command + '\n').encode('ascii')
    board.write(tosend)

def sensor_check(direction, distance):
    return sensors[direction] < distance and sensors[direction] != 0

def sensor_loop():
    sensor_board = serial.Serial("/dev/ttyUSB1", 115200, timeout=10)
    print("Connected to Sensor Board")
    while True:
        reading = sensor_board.readline().decode("utf-8").split(",")
        global sensors
        sensors = {'right': int(reading[0]), 'front': int(reading[1]), 'left': int(reading[2]), 'rear': int(reading[3])}

def control_loop():
    drive_board = serial.Serial("/dev/ttyUSB0", 115200, timeout=10)
    print("Connected to Drive Board")
    
    # Don't start driving until the sensor readings are coming in
    while(sensors is None):
        continue
    
    print("Starting to drive")
    while True:
        if sensor_check("front", 30) or sensor_check("left", 20) or sensor_check("right", 20):
            # Too close to something
            if sensors["left"] == 0:
                # Left is wide open
                command(drive_board, "d,left,190")
            elif sensors["right"] == 0:
                # Right is wide open
                command(drive_board, "d,right,190")
            elif sensors["left"] > sensors["right"]:
                # Left has more room than right
                command(drive_board, "d,left,190")
            else:
                # Right has more room then left (or they're the same)
                command(drive_board, "d,right,190")
        else:
            # Nothing blocking, move forward
            command(drive_board, "d,forward,255")

        # 25 updates per second
        time.sleep(0.04)

# Start the sensor board thread
sensor_thread = Thread(target=sensor_loop)
sensor_thread.start()

# Start the main drive control loop
control_loop()