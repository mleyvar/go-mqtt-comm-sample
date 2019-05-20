package main

import (
	"mqttComm/audit"
	config "mqttComm/configuration"
	"mqttComm/structure"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
*  FUNCTION:   main
*  PARAM:
*  RETURN:
*  DESCRIPTION: Start process
 */

func main() {

	SendMessageMQTT("GO_SERVER/TEST", "Test message for MQTT localhost")

}

/*
*  FUNCTION:   SendMessageMQTT
*  PARAM:      sus= Suscription (topic) for send message   -  msg = message
*  RETURN:
*  DESCRIPTION: Method for send messages in topic mqtt
 */

func SendMessageMQTT(sus string, msg string) {

	var m structure.Structure_MQTT

	m = config.GetConfigurationMQTT()

	if m.Topic == "" {
		println("Invalid setting for -topic, must not be empty")
		return
	} else {
		m.Topic += "0001"
	}
	println("Broker: " + m.Broker)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(m.Broker)
	opts.SetClientID(m.Id)
	opts.SetUsername(m.User)
	opts.SetPassword(m.Password)
	opts.SetCleanSession(m.Cleansess)
	if m.Store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(m.Store))
	}

	client_pub := MQTT.NewClient(opts)

	if token := client_pub.Connect(); token.Wait() && token.Error() != nil {

		if !audit.CheckErr(error(token.Error()), "SendMessageMQTT - 1") {
			return
		}
	}

	println("Suscription........: " + sus)
	println("Message to published : " + msg)

	token := client_pub.Publish(sus, byte(m.Qos), false, msg)

	token.Wait()

	client_pub.Disconnect(250)
	println("Publisher Disconnected.")

}
