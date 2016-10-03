import serial
import json

class SensorBoard:
    ser = None
    sensors = None
    last_sensors = None

    sensor_mapping = {
        "front" : 5,
        "front right" : 6,
        "right front" : 7,
        "right rear" : 8,
        "rear right" : 9,
        "rear" : 0,
        "rear left" : 1,
        "left rear" : 2,
        "left front" : 3,
        "front left" : 4
    }

    def __init__(self, port, baud, timeout):
        self.ser = serial.Serial(port, baud, timeout=timeout)
        print("Connected to Sensor Board.")

    def sensors_ready(self):
        return self.sensors is not None

    def read_sensors(self):
        try:
            reading = self.ser.readline().decode("utf-8")
            json_dict = json.loads(reading)
            self.sensors = json_dict['sensors']
        except UnicodeDecodeError:
            print("Serial data is not valid unicode, likely an incomplete byte stream.")
        except ValueError:
            print("Error parsing JSON data.")

    def sensor_value(self, sensor):
        return self.sensors[self.sensor_mapping[sensor]]['val']

    def distance_check(self, direction, distance):
        """Return true if the sensor is reading closer than the given distance"""
        return self.sensor_value(direction) < distance and self.sensor_value(direction) != 0

    