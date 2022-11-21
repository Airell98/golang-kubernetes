package main

import (
	"golang-kubernetes/payment/events/payment_listener"
	"golang-kubernetes/payment/events/payment_producer"
)

func main() {
	payment_producer.PaymentProducer.SetUpProducer()
	payment_listener.PaymentListener.InitiliazeMainListener()
}
