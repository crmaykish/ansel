import time
import sqlite3

class DAO:
    db_file = None
    connection = None
    run = None

    def __init__(self, db_file):
        self.db_file = db_file

    def connect(self):
        print("Connecting to database...")
        self.connection = sqlite3.connect(self.db_file)
        self.run = self.get_run()

    def disconnect(self):
        print("Disconnecting from database...")
        self.connection.commit()
        self.connection.close()

    def get_run(self):
        c = self.connection.cursor()
        c.execute("SELECT MAX(run)+1 FROM sensors")
        r = c.fetchone()[0]
        if r is None:
            return 0
        else:
            return int(r)

    def save_sensor(self, sensor):
        c = self.connection.cursor()
        c.execute("INSERT INTO sensors VALUES (?, datetime('now'), ?, ?)", (run, sensor['pos'], sensor['val']))
        self.connection.commit()

    def save_sensors(self, sensors):
        if sensors is None:
            return

        d = []
        for s in sensors:
            d.append((self.run, s['pos'], s['val']))
        
        c = self.connection.cursor()
        c.executemany("INSERT INTO sensors VALUES (?, datetime('now'), ?, ?)", d)
        self.connection.commit()

    def save_motor(self, direction, speed):
        c = self.connection.cursor()
        c.execute("INSERT INTO motors VALUES (datetime('now'), ?, ?)", (motor, speed))
        self.connection.commit()