package main

import (
	"mqttComm/audit"
	config "mqttComm/configuration"
	"mqttComm/structure"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
*  FUNCTION:   main
*  PARAM:
*  RETURN:
*  DESCRIPTION: Start process
 */

func main() {

	process()
}

/*
*  FUNCTION:   process
*  PARAM:
*  RETURN:
*  DESCRIPTION: Method for read messages in topic mqtt
 */
func process() {
	var m structure.Structure_MQTT

	m = config.GetConfigurationMQTT()

	if m.Topic == "" {
		println("Invalid setting for -topic, must not be empty")
		return

	}

	// add # command for listener all message in this suscription (topic)
	m.Topic = m.Topic + "#"

	println("Topic for suscription listener: " + m.Topic)

	opts := MQTT.NewClientOptions()
	//opts.ClientID="111110"    // change this client id, random number for example
	opts.AddBroker(m.Broker)
	opts.SetClientID(m.Id)
	opts.SetUsername(m.User)
	opts.SetPassword(m.Password)
	opts.SetCleanSession(m.Cleansess)
	if m.Store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(m.Store))
	}

	receiveCount := 0
	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		audit.CheckErr(error(token.Error()), "process - 2")
	}

	if token := client.Subscribe(m.Topic, byte(m.Qos), nil); token.Wait() && token.Error() != nil {
		audit.CheckErr(error(token.Error()), "process - 3")

	}

	for {

		println("Start listener.....")
		incoming := <-choke

		println("*******************************************************")
		println("Topic Recive: " + incoming[0])
		println("Message Recive: " + incoming[1])
		println("*******************************************************")

		if strings.Index(incoming[1], "@ONE") > -1 {

			println("Start process ONE")
		} else if strings.Index(incoming[1], "@TWO") > -1 {

			println("Start process TWO")

		} else if strings.Index(incoming[1], "@EXIT") > -1 {
			return

		}

		receiveCount++

	}
	client.Disconnect(250)
	println("Subscriber Disconnected")

}
