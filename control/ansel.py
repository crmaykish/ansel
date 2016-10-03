# Ansel Sensor Board - Colin Maykish - September 2016
import time
import dao
import sensor_board
import motor_board
from threading import Thread

class Ansel:
    DISTANCE_LIMIT = 30
    running = True
    motor = None
    sensor = None
    sensor_thread = None
    db = None

    def __init__(self):
        # Connect to database
        self.db = dao.DAO("/home/colin/ansel-bot/db/ansel.db")
        self.db.connect()

        print(self.db.run)

        # Connect to motor board
        self.motor = motor_board.MotorBoard("/dev/ttyUSB0", 115200, timeout=10)

        # Connect to sensor board
        self.sensor = sensor_board.SensorBoard("/dev/ttyUSB1", 9600, timeout=10)

        # Start the sensor board thread
        self.sensor_thread = Thread(target=self.sensor_loop)
        self.sensor_thread.start()

    def sensor_loop(self):
        # Continuous read sensor values
        while self.running:
            self.sensor.read_sensors()

    def control_loop(self):
        print("Starting Ansel control loop...")

        # Don't start driving until the sensor readings are coming in
        while(self.sensor.sensors_ready() is False):
            continue

        print("Starting to drive.")
        while self.running:
            self.update_movement()

            # Only update the DB when the sensor values change
            if (self.sensor.sensors != self.sensor.last_sensors):
                self.db.save_sensors(self.sensor.sensors)
                self.sensor.last_sensors = self.sensor.sensors.copy()

            time.sleep(self.motor.sleep_time())

    def update_movement(self):
        if (self.sensor.distance_check("front", self.DISTANCE_LIMIT) or self.sensor.distance_check("front left", self.DISTANCE_LIMIT) or self.sensor.distance_check("front right", self.DISTANCE_LIMIT)):
            self.motor.stop_movement()

            if (self.sensor.sensor_value("front left") > self.sensor.sensor_value):
                self.motor.set_movement("left", 190)
            else:
                self.motor.set_movement("right", 190)
        else:
            self.motor.set_movement("forward", 255)

    def stop(self):
        self.running = False
        self.db.disconnect()

# Setup and run Ansel
try:
    ansel = Ansel()
    ansel.control_loop()
except KeyboardInterrupt:
    print("Stopping Ansel...")
finally:
    ansel.stop()