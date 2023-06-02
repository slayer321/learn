package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func PrintOddNumbers(ctx workflow.Context, numberList []int) ([]int, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result []int
	err := workflow.ExecuteActivity(ctx, OddNumber, numberList).Get(ctx, &result)

	return result, err
}

func PrintEvenNumbers(ctx workflow.Context, numberList []int) ([]int, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result []int
	err := workflow.ExecuteActivity(ctx, EvenNumber, numberList).Get(ctx, &result)

	return result, err
}
