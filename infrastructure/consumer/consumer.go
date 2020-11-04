package consumer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/iwanjunaid/basesvc/domain/model"

	"github.com/iwanjunaid/basesvc/adapter/controller"
	"github.com/iwanjunaid/basesvc/registry"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerImpl struct {
	kc            *kafka.Consumer
	appController controller.AppController
}

func NewConsumer(kc *kafka.Consumer, db *sql.DB) *ConsumerImpl {
	registry := registry.NewRegistry(db)
	appController := registry.NewAppController()

	return &ConsumerImpl{
		kc:            kc,
		appController: appController,
	}
}

func (c *ConsumerImpl) Listen(topic string) {
	err := c.kc.Subscribe(topic, nil)
	if err != nil {
		panic(err)
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false

		case ev := <-c.kc.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.kc.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.kc.Unassign()
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				var author *model.Author
				if err := json.Unmarshal(e.Value, &author); err != nil {
					panic(err)
				}
				c.appController.Author.InsertAuthor(author)
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.kc.Close()
}
