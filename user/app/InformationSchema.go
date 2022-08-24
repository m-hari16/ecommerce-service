package app

import "os"

func getServiceName() string{
     return os.Getenv("SERVICE_NAME")
}