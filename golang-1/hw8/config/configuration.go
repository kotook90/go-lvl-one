package configuration

import (
	"fmt"
	"godotenv"
	"net/url"
	"os"
	"strconv"
)

type Configuration struct {
	Port        int
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}

func (c *Configuration) Load() {

	err := godotenv.Load("/home/anton/Projects/golangLvlOne/hw8/config/conf.env")
	if err != nil {
		fmt.Println("Не удалось згрузить файл конфигурации", err)
		return
	}
	_, err = strconv.Atoi(os.Getenv("port"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.Port, _ = strconv.Atoi(os.Getenv("port"))

	_, err = url.Parse(os.Getenv("db_url"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.DbUrl = os.Getenv("db_url")

	_, err = url.Parse(os.Getenv("jaeger_url"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.JaegerUrl = os.Getenv("jaeger_url")

	_, err = url.Parse(os.Getenv("sentry_url"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.SentryUrl = os.Getenv("sentry_url")

	c.KafkaBroker = os.Getenv("kafka_broker")

	c.SomeAppId = os.Getenv("some_app_id")

	c.SomeAppKey = os.Getenv("some_app_key")
}

func (c Configuration) End() string {

	fmt.Printf("port = %v\ndb_url = %v\njaeger_url = %v\nsentry_url = %v\nkafka_broker = %v\n"+
		"some_app_id = %v\nsome_app_key = %v", c.Port, c.DbUrl, c.JaegerUrl, c.SentryUrl, c.KafkaBroker, c.SomeAppId, c.SomeAppKey)
	return c.End()
}
