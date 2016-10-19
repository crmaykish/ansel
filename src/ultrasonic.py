import peripheral, control_loop, time

class Ultrasonic(peripheral.Peripheral, control_loop.ControlLoop):
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
        super().__init__(port, baud, timeout)
        print("Connected to Sensor Board")

    def sensors_ready(self):
        return self.sensors is not None and len(self.sensors) == 10

    def read_sensors(self):
        try:
            reading = self.ser.readline().decode("utf-8")
            kv = reading.split(":")

            if self.sensors is None:
                self.sensors = {}

            self.sensors[int(kv[0])] = int(kv[1])
        except UnicodeDecodeError:
            print("Serial data is not valid unicode, likely an incomplete byte stream.")
        except IndexError:
            print("Not a valid key/value pair.")
        except ValueError:
            print("Value error, probably just bad serial data.")

    def sensor_value(self, sensor):
        return self.sensors[self.sensor_mapping[sensor]]

    def distance_check(self, direction, distance):
        """Return true if the sensor is reading closer than the given distance"""
        return self.sensor_value(direction) < distance and self.sensor_value(direction) != 0

    def iterate(self):
        self.read_sensors()