# Ansel Sensor Board - Colin Maykish - September 2016
from threading import Thread
import time
import ultrasonic
import motor
import control_loop

class Control(control_loop.ControlLoop):
    motor = None
    ultrasonic = None

    def __init__(self):
        self.motor = motor.Motor("/dev/ttyUSB1", 115200, timeout=10)
        self.ultrasonic = ultrasonic.Ultrasonic("/dev/ttyUSB0", 9600, timeout=10)

    def start(self):
        # Start the ultrasonic sensor loop in its own thread
        Thread(target=self.ultrasonic.start).start()

        print("Started sensor thread")

        # Don't start driving until the sensor readings are coming in
        while not self.ultrasonic.sensors_ready():
            continue

        print("Starting main control loop...")
        super().start()

    def iterate(self):
        print("control loop")
        time.sleep(1)
