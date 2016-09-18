# Ansel Robotics Platform
Ansel is a robotics platform I'm developing to be used as a base for computer vision and machine learning experiments.

## Software Components
Currently, it consists of three main subsystems:
1. Drive: Arduino board controlling a dual H-Bridge driver and two DC motors in a differential drive configuration
2. Sensors: A second Arduino board wired up with several ultrasonic ping sensors for basic boundary detection
3. Control: A Raspberry Pi 2 running some Python to tie the other subsystems together and control the high level function

## Electrical
Aside from the aforementioned Arduino and Raspberry Pi components, Ansel is powered by two 3000mAH NiMH battery packs in parallel.
The two motors get the raw the battery voltage and the rest of the components are powered by a low dropout linear regulator at 5v on a handmade board.

## Mechanical
Ansel sits atop two DC drive motors and two caster wheels. The chassis is a combination of 1/4" plexiglass sheets and 3D printed parts.