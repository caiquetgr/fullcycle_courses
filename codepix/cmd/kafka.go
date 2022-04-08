/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/caiquetgr/fullcycle_courses/codepix/application/kafka"
	"github.com/caiquetgr/fullcycle_courses/codepix/infrastructure/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {

		deliveryChannel := make(chan ckafka.Event)
		database := db.ConnectDB(os.Getenv("env"))
		producer := kafka.NewKafkaProducer()

		//kafka.Publish("oi kafka", "teste", producer, deliveryChannel)
		go kafka.DeliveryReport(deliveryChannel)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChannel)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
