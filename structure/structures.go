package structure

type Structure_MQTT struct {
	Topic     string // "The topic name to/from which to publish/subscribe"
	Broker    string // "The broker URI. ex: tcp://10.10.1.1:1883"
	Password  string // "The password (optional)")
	User      string // "The User (optional)")
	Id        string // "The ClientID (optional)")
	Cleansess bool   // "Set Clean Session (default false)")
	Qos       int    // "The Quality of Service 0,1,2 (default 0)")
	Payload   string // "The message text to publish (default empty)")
	Action    string // "Action publish or subscribe (required)")
	Store     string // "The Store Directory (default use memory store)")
}
