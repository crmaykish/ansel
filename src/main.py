from flask import Flask, render_template
from flask_socketio import SocketIO, emit
import logging
import control

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

socketio.start_background_task(target=control.start)

if __name__ == '__main__':
    print("Starting web application")
    socketio.run(app, host="0.0.0.0", debug=False)
