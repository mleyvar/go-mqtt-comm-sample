package config

import "mqttComm/structure"

/*
*  FUNCTION:   GetConfigurationMQTT
*  PARAM:
*  RETURN:  Structure object Structure_MQTT
*  DESCRIPTION: Method for set values i the Structure_MQTT object. Start configuration mqtt server
 */
func GetConfigurationMQTT() structure.Structure_MQTT {

	config := structure.Structure_MQTT{}

	config.Broker = "tcp://localhost:1883" // change host and port
	config.Id = "pubtestgoid"
	config.Cleansess = false
	config.User = "userxxx"    // change user
	config.Password = "pwdxxx" // change pwd
	config.Payload = ""
	config.Qos = 0
	config.Store = ":memory:"
	config.Topic = "GO_SERVER/"
	config.Action = "sub"

	return config
}
