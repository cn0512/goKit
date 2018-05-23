package ykmq

/*
	Use AMQP.RabbitMQ
*/

type IMQ interface {
	Create(url string, s2c bool)
	PushMsg(buf []byte)
	PopMsg()
}
