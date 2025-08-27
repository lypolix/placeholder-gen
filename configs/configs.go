package configs

import (
    "os"
    "fmt"

    "github.com/joho/godotenv"
)

type ConfI interface {
    GetPort() string
}

type Conf struct {
    port string
}

func (c *Conf) GetPort() string {
    return c.port
}

func New(confPath string) (*Conf, error) {
    err := godotenv.Load(confPath)
    if err != nil {
        return nil, err
    }
   
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        return nil, fmt.Errorf("environment variable SERVER_PORT is not set")
    }
    return &Conf{port: port}, nil
}
