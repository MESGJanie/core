## mesg service stop

Stop a service

### Synopsis

By stoping a service, your node will not process any other actions from this service.
/!\ This action will slash your stake if you didn't respect the duration

```
mesg service stop [flags]
```

### Examples

```
mesg-cli service stop ethereum
```

### Options

```
  -c, --confirm   Confirm
  -h, --help      help for stop
```

### SEE ALSO

* [mesg service](mesg_service.md)	 - Manage the services you are running

