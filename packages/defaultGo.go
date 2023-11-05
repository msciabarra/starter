package main

func Main(obj map[string]interface{}) map[string]interface{} {
	msg := make(map[string]interface{})
	msg["message"] = "Hello, default go Deployer!"
	return msg
}
