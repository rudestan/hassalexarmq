package rmq

import (
    "os"

    "github.com/spf13/cast"
)

const (
    RmqHost       = "localhost"
    RmqPort       = 5672
    RmqLogin      = "guest"
    RmqPassword   = "guest"
    RmqExchange   = "sync_exchange"
    RmqQueue      = "sync_queue"
    RmqRoutingKey = "sync.json"
    RmqMessageExp = "50000" // 50 sec.
)

const (
    EnvRmqHost = "RMQ_HOST"
    EnvRmqPort = "RMQ_PORT"
    EnvRmqLogin = "RMQ_LOGIN"
    EnvRmqPassword = "RMQ_PASSWORD"
    EnvRmqExchange = "RMQ_EXCHANGE"
    EnvRmqQueue = "RMQ_QUEUE"
    EnvRmqRoutingKey = "RMQ_ROUTING_KEY"
    EnvRmqMsgExpiration = "RMQ_MSG_EXPIRATION"
)

type Config struct {
    Host       string
    Port       int
    Login      string
    Password   string
    Exchange   string
    Queue      string
    RoutingKey string
    Expiration string
}

func NewConfigFromEnv() *Config {
    config := Config{
        Host:       getEnvVar(EnvRmqHost, RmqHost),
        Port:       cast.ToInt(getEnvVar(EnvRmqPort, cast.ToString(RmqPort))),
        Login:      getEnvVar(EnvRmqLogin, RmqLogin),
        Password:   getEnvVar(EnvRmqPassword, RmqPassword),
        Exchange:   getEnvVar(EnvRmqExchange, RmqExchange),
        Queue:      getEnvVar(EnvRmqQueue, RmqQueue),
        RoutingKey: getEnvVar(EnvRmqRoutingKey, RmqRoutingKey),
        Expiration: getEnvVar(EnvRmqMsgExpiration, RmqMessageExp),
    }

    return &config
}

func getEnvVar(name string, defaultVal string) string  {
    if value, exists := os.LookupEnv(name); exists {
        return value
    }

    return defaultVal
}
