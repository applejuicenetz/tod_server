#!/bin/sh

CONF_FILE=/app/config.yaml

sed -i "s/tcp_enabled.*/tcp_enabled: $TCP_ENABLED/" $CONF_FILE
sed -i "s/udp_enabled.*/udp_enabled: $UDP_ENABLED/" $CONF_FILE
sed -i "s/level.*/level: '$LOG_LEVEL'/" $CONF_FILE
sed -i "s/format.*/format: '$LOG_FORMAT'/" $CONF_FILE

/app/tod_server -config $CONF_FILE
