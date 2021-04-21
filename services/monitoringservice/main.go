package main

import (
	"github.com/VladyslavLukyanenko/MonitoringService/configs"
	"github.com/VladyslavLukyanenko/MonitoringService/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	configs.ReadConfig()
	initDatabase()
	initAMQP()
}
func initDatabase() {

}

//type TestPayload struct {
// test string
//}
//
//type Test struct {
//	payload *TestPayload
//}
////
////func BindQueueToFunction(name string, handler func (amqp.Delivery, *interface{}), data *interface{}) {
////
////}
//func (*Test) GetRoutingKey() string {
//	return "test"
//}
//
//func (*Test) RouteFunction(message amqp.Delivery, p *interface{}) {
//}
//
//func (typez *Test) GetDataStructure() *interface{} {
//	var data interface{}
// 	data = typez.payload
//	return &data
//}
func initAMQP() {
	InitAMQP()
	BindQueueToFunction("monitor-add-task", routes.MonitorAddTask)
	BindQueueToFunction("monitor-remove-task", routes.MonitorRemoveTask)
	//BindQueueToFunction_(arg)
	//BindQueueToFunction("monitor-remove-task", routes.MonitorRemoveTask, *interface{})
}