const int L_FORWARD = 3;
const int L_REVERSE = 5;
const int R_FORWARD = 6;
const int R_REVERSE = 9;

void setup() {
    Serial.begin(115200);
}

void loop() {
    left(255);
    delay(10000);
    stop();
    delay(500);
    right(255);
    delay(10000);
    stop();
    delay(500);
}

void forward(int speed) {
    analogWrite(L_REVERSE, 0);
    analogWrite(R_REVERSE, 0);
    analogWrite(L_FORWARD, speed);
    analogWrite(R_FORWARD, speed);
}

void reverse(int speed) {
    analogWrite(L_FORWARD, 0);
    analogWrite(R_FORWARD, 0);
    analogWrite(L_REVERSE, speed);
    analogWrite(R_REVERSE, speed);
}

void left(int speed) {
    analogWrite(L_FORWARD, 0);
    analogWrite(R_FORWARD, speed);
    analogWrite(L_REVERSE, speed);
    analogWrite(R_REVERSE, 0);
}

void right(int speed) {
    analogWrite(L_FORWARD, speed);
    analogWrite(R_FORWARD, 0);
    analogWrite(L_REVERSE, 0);
    analogWrite(R_REVERSE, speed);
}

void stop() {
    analogWrite(L_FORWARD, 0);
    analogWrite(L_REVERSE, 0);
    analogWrite(R_FORWARD, 0);
    analogWrite(R_REVERSE, 0);
}
