import time
import sqlite3

run = None

def dao_init():
    conn = sqlite3.connect("/home/colin/ansel-bot/db/ansel.db")
    c = conn.cursor()
    c.execute("SELECT MAX(run)+1 from sensors")
    global run
    run = int(c.fetchone()[0])

def save_sensor(sensor):
    conn = sqlite3.connect("/home/colin/ansel-bot/db/ansel.db")
    c = conn.cursor()
    c.execute("INSERT INTO sensors VALUES (?, datetime('now'), ?, ?)", (run, sensor['pos'], sensor['val']))
    conn.commit()
    conn.close()

def save_motor(motor, speed):
    conn = sqlite3.connect("/home/colin/ansel-bot/db/ansel.db")
    c = conn.cursor()
    c.execute("INSERT INTO motors VALUES (datetime('now'), ?, ?)", (motor, speed))
    conn.commit()
    conn.close()