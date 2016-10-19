# Generic control loop interface
class ControlLoop:
    running = False

    def start(self):
        if not self.running:
            self.running = True
            self._loop()
        else:
            print("Loop is already running")

    def stop(self):
        if self.running:
            self.running = False
        else:
            print("Loop is not running")

    def _loop(self):
       while self.running:
           self.iterate()

    def iterate(self):
        raise NotImplementedError(self.__class__.__name__ + " must implement iterate() method.")