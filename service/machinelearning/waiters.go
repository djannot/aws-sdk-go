// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package machinelearning

import (
	"github.com/djannot/aws-sdk-go/private/waiter"
)

func (c *MachineLearning) WaitUntilBatchPredictionAvailable(input *DescribeBatchPredictionsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeBatchPredictions",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *MachineLearning) WaitUntilDataSourceAvailable(input *DescribeDataSourcesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeDataSources",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *MachineLearning) WaitUntilEvaluationAvailable(input *DescribeEvaluationsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeEvaluations",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *MachineLearning) WaitUntilMLModelAvailable(input *DescribeMLModelsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeMLModels",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Results[].Status",
				Expected: "COMPLETED",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Results[].Status",
				Expected: "FAILED",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
