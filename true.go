package true

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

func GetConfiguration() {

	port := flag.String("port", "5432", "port to connect to db")
	db_url := flag.String("db_url", "http", "url to connect to www")
	jaeger_url := flag.String("jaeger_url", "http", "url to connect to www")
	sentry_url := flag.String("entry_url", "http", "url to connect to www")
	kafka_broker := flag.String("kafka_broker", "5432", "kafka connect to port")
	some_app_id := flag.Int("some_app_id", 111, "ID to connect to www")
	some_app_key := flag.Int("some_app_key", 111, "KEY to connect to www")

	flag.Parse()

	_, err := url.Parse("http")
	if err != nil {
		fmt.Printf("Ошибка: %s/n", err.Error())
		return
	}

	connetntoDB(*port, *db_url, *jaeger_url, *sentry_url, *kafka_broker, *some_app_id, *some_app_key)

	file, err := os.OpenFile(
		"C:/Users/Admin/LessonGO/lesson8/text.txt",
		os.O_RDONLY,
		0777,
	)

	if err != nil {
		panic(err)
	}

	buf := make([]byte, 35)
	n, err := file.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, string(buf))

}

func connetntoDB(port, db_url, jaeger_url, sentry_ur, kafka_broker string, some_app_id, some_app_key int) {

	fmt.Println(port, db_url, jaeger_url, sentry_ur, kafka_broker, some_app_id, some_app_key)

}
