#include <NewPing.h>

#define MAX_DISTANCE 200

NewPing right = NewPing(5, 6, MAX_DISTANCE);
NewPing front = NewPing(7, 8, MAX_DISTANCE);
NewPing left = NewPing(9, 10, MAX_DISTANCE);
NewPing rear = NewPing(11, 12, MAX_DISTANCE);

int distances[4] = {0, 0, 0, 0};

void setup()
{
    Serial.begin(115200);
}

void loop()
{
    distances[0] = right.ping_cm();
    delay(2);
    distances[1] = front.ping_cm();
    delay(2);
    distances[2] = left.ping_cm();
    delay(2);
    distances[3] = rear.ping_cm();
    delay(2);

    for (int i = 0; i < 4; i++) {
        Serial.print(distances[i]);
        if (i != 3) {
            Serial.print(",");
        }
    }

    Serial.println("");
}