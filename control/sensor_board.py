import serial
import json

class SensorBoard:
    ser = None
    sensors = None
    last_sensors = None

    sensor_mapping = {
        "front" : 0,
        "front right" : 1,
        "right front" : 2,
        "right rear" : 3,
        "rear right" : 4,
        "rear" : 5,
        "rear left" : 6,
        "left rear" : 7,
        "left front" : 8,
        "front left" : 9
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

    def distance_check(direction, distance):
        """Return true if the sensor is reading closer than the given distance"""
        return self.sensor_value(direction) < distance and self.sensor_value(direction) != 0

    