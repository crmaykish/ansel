import peripheral

class Motor(peripheral.Peripheral):
    # Rate to run motor updates (Hertz)
    UPDATE_RATE = 20

    def __init__(self, port, baud, timeout):
        super().__init__(port, baud, timeout)
        print("Connected to Motor Board")
    
    def sleep_time(self):
        return 1 / self.UPDATE_RATE

    def set_movement(self, direction, speed):
        if (speed >= 0 and speed <= 255):
            self.send("d," + direction + "," + str(speed))

    def stop_movement(self):
        self.send("stop,")