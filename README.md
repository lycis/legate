# Legate
Legate is a webservice forwarder that is intertwined with [Consul](https://github.com/hashicorp/consul).
It forwards any HTTP requests via HTTP redirection to a service that is registered by Consul.

## Usage
Simply start legate with a call to its command:
```
legate
```
Any configuration will be loaded from a file called `legate.yml` in the working directory.

## Configuration
Legate is configured by a file called `legate.yml` in its working directory. Example:
```
consul:
    address: consulserver:8500
    datacenter: dc01
port: 8080
bind: 0.0.0.0
```

* _consul_: This section directly refers to the configuration of the Consul API. See [here](https://godoc.org/github.com/hashicorp/consul/api#Config)
* _port_: Port to listen on for requests (default: 8080)
* _bind_: Address to bind to (default: 0.0.0.0)
