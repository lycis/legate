# Legate
Legate is a webservice forwarder that is intertwined with [Consul](https://github.com/hashicorp/consul).
It forwards any HTTP requests via HTTP redirection to a service that is registered by Consul.

## Usage
Simply start legate with a call to its command `legate <consul address>'. For example this call
may be like this:
```
legate consulserver:8500 
```

Any configuration will be dervied from command line parameters unless `-config` is given to set a
configuration file.

## Configuration
Usually `legate` is configured by providing matching command line parameters:

```
-bind string
        address and port to bind to (e.g. 127.0.0.1:80)  (default ":8080")
-config string
        configuration file
-dc string
        consul datacenter (default "dc01")
```

### File
If Legate is configured by file the given file has to be a YAML file with the follwing structure:

 Example:
```
consul:
    address: consulserver:8500
    datacenter: dc01
bind: 0.0.0.0:8080
```

* _consul_: This section directly refers to the configuration of the Consul API. See [here](https://godoc.org/github.com/hashicorp/consul/api#Config)
* _bind_: Address and port to bind to (default: 0.0.0.0:8080)
