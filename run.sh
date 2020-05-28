#!/usr/bin/env bashio

export RMQ_HOST=$(bashio::config 'rmq_host')
export RMQ_PORT=$(bashio::config 'rmq_port')
export RMQ_LOGIN=$(bashio::config 'rmq_login')
export RMQ_PASSWORD=$(bashio::config 'rmq_password')
export RMQ_EXCHANGE=$(bashio::config 'rmq_exchange')
export RMQ_QUEUE=$(bashio::config 'rmq_queue')
export RMQ_ROUTING_KEY=$(bashio::config 'rmq_routing_key')
export RMQ_MSG_EXPIRATION=$(bashio::config 'rmq_message_expiration')

bashio::log.info "Start Alexa RMQ Consumer"

arm_alexa_rmq_consumer