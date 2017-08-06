# TOD server
This is a TOD server conforming [RFC 868](https://tools.ietf.org/html/rfc868).

## Run
Launching the binary is enough. Since it binds to port 37, it must be run as root.

To override defaults, run with `-config` flag:
```
# ./tod_server -config conf/config.yaml
```
## Configuration
The configuration file looks like this:
```
tcp_enabled: true
tcp_host: '0.0.0.0'
tcp_port: 37

udp_enabled: true
udp_host: '0.0.0.0'
udp_port: 37

logging:
  stdout:
    level: 'INFO'
    format: 'plain'
  file:
    level: 'INFO'
    format: 'plain'
    dir: '/var/log/'
```
* `tcp_enabled`: enable or disable TCP. Both UDP and TCP can be used, but at least one must be enabled.
* `tcp_host`: ip address.
* `tcp_port`: shouldn't be changed since 37 is the port defined by RFC.
* `udp_enabled`: enable or disable UDP. Both UDP and TCP can be used, but at least one must be enabled.
* `udp_host`: ip address.
* `udp_port`: shouldn't be changed since 37 is the port defined by RFC.
* `logging`: defines logging appenders. By default only logs to STDOUT.



## Docker image
[![](https://images.microbadger.com/badges/version/elpadrinoiv/tod_server.svg)](https://microbadger.com/images/elpadrinoiv/tod_server "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/image/elpadrinoiv/tod_server.svg)](https://microbadger.com/images/elpadrinoiv/tod_server "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/commit/elpadrinoiv/tod_server.svg)](https://microbadger.com/images/elpadrinoiv/tod_server "Get your own commit badge on microbadger.com")
[![](https://images.microbadger.com/badges/license/elpadrinoiv/tod_server.svg)](https://microbadger.com/images/elpadrinoiv/tod_server "Get your own license badge on microbadger.com")

### Configuration
There are a few things that can be changed with environment variables:
* `TCP_ENABLED`: default is true
* `UDP_ENABLED`: default is true
* `LOG_LEVEL`: default is INFO
* `LOG_FORMAT`: plain or json. Default is plain
