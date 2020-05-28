package rmq

import (
    "flag"
    "os"

    "github.com/spf13/cast"
)

const (
    RmqPort       = 5672
    RmqLogin      = "guest"
    RmqPassword   = "guest"
    RmqExchange   = "sync_exchange"
    RmqQueue      = "sync_queue"
    RmqRoutingKey = "sync.json"
    RmqMessageExp = "50" // sec.
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
    expiration := cast.ToString(cast.ToInt(getEnvVar(EnvRmqMsgExpiration, RmqMessageExp)) * 1000)
    config := Config{
        Host:       getEnvVar(EnvRmqHost, ""),
        Port:       cast.ToInt(getEnvVar(EnvRmqPort, cast.ToString(RmqPort))),
        Login:      getEnvVar(EnvRmqLogin, RmqLogin),
        Password:   getEnvVar(EnvRmqPassword, RmqPassword),
        Exchange:   getEnvVar(EnvRmqExchange, RmqExchange),
        Queue:      getEnvVar(EnvRmqQueue, RmqQueue),
        RoutingKey: getEnvVar(EnvRmqRoutingKey, RmqRoutingKey),
        Expiration: expiration,
    }

    return &config
}

func NewConfigFromFlags() *Config {
    config := new(Config)

    flag.StringVar(&config.Host, "h", "", "required - RabbitMQ Host")
    flag.IntVar(&config.Port, "p", RmqPort, "optional - RabbitMQ Port")
    flag.StringVar(&config.Login, "l", RmqLogin, "optional - Login")
    flag.StringVar(&config.Password, "s", RmqPassword, "optional - Password")
    flag.StringVar(&config.Exchange, "e", RmqExchange, "optional - Exchange name")
    flag.StringVar(&config.Queue, "q", RmqQueue, "optional - Queue name")
    flag.StringVar(&config.RoutingKey, "r", RmqRoutingKey, "optional - Routing key")
    flag.StringVar(&config.Expiration, "t", RmqMessageExp, "optional - Message expiration time in seconds")

    return config
}

func getEnvVar(name string, defaultVal string) string  {
    value := os.Getenv(name)

    if value != "" {
        return value
    }

    return defaultVal
}
