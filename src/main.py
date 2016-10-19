from flask import Flask, render_template
from flask_socketio import SocketIO, emit
import logging
import control
import time

# Flask app setup
app = Flask(__name__)
app.config['SECRET_KEY'] = 'secret'
socketio = SocketIO(app, async_mode='threading')

# Logging config
log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

# Ansel control logic
control = control.Control()

@app.route('/')
def index():
    return render_template('index.html')

@socketio.on('connect')
def connected():
    print("Client connected.")

def ui_loop():
    global control
    global socketio

    while True:
        socketio.emit("telemetry", control.ultrasonic.jsonify())
        time.sleep(0.10)

socketio.start_background_task(target=control.start)
socketio.start_background_task(target=ui_loop)

if __name__ == '__main__':
    print("Starting web application")
    socketio.run(app, host="0.0.0.0", debug=False)
