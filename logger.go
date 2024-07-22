package Logger

import "fmt"

func Info(message string, params interface{}) {
	fmt.Println("Log info message: ", message, " params: ", params)
}

func Error(message string, params interface{}) {
	fmt.Println("Log error message: ", message, " params: ", params)
}

func Warning(message string, params interface{}) {
	fmt.Println("Log warning message: ", message, " params: ", params)
}

func Fatal(message string, params interface{}) {
	fmt.Println("Log fatal: message", message, " params: ", params)
}
