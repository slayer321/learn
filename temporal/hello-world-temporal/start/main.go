package main

import (
	"context"
	"flag"
	"fmt"
	"hello-world-temporal/app"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	var workflowName string
	flag.StringVar(&workflowName, "workflow", "", "The name of the workflow to run (greeting or even-number)")

	flag.Parse()

	switch workflowName {
	case "odd-number":
		workflowName = "PrintOddNumbers"
	case "even-number":
		workflowName = "PrintEvenNumbers"
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow-id",
		TaskQueue: app.GreetingTaskQueue,
	}

	input := []int{1, 4, 5, 6, 7, 8, 9, 10}
	we, err := c.ExecuteWorkflow(context.Background(), options, workflowName, input)
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}

	var result []int
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get workflow result", err)
	}
	printResults(result, we.GetID(), we.GetRunID())

}

func printResults(greeting []int, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%v\n\n", greeting)
}
