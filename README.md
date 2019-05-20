go-mqtt-comm-sample



This sample use  github.com/eclipse/paho.mqtt.golang   PAHO libary.

Install paho librari with:

go get github.com/eclipse/paho.mqtt.golang


For testing services mqtt you can use this commands:

1.- Install mosquitto client in you system


Suscribe (Listener)

mosquitto_sub -h localhost -p 1883 -t "TOPIC/#" -q 1  -u "USERXXX" -P "PWDXXXX"

Sample:

mosquitto_sub -h localhost -p 1883 -t "GO_SERVER/#" -q 1  -u "myuser" -P "mypwd"



Publish message:

mosquitto_pub -h localhost -p 1883 -t "TOPIC" -m "MESSAGE" -u "USERXXXXX" -P "PWDXXXX"

Sample

mosquitto_pub -h localhost -p 1883 -t "GO_SERVER/MACADDRESS_DEVICE" -m "SWITC_ON" -u "myuser" -P "mypwd"


Sample commands for this program:

@ONE  =  Simulates the execution of the ONE process

@TWO  =  Simulates the execution of the TWO process


@EXIT  =  Exit listener and program



