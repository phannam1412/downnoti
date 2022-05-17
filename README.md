# downnoti

Check if a server is up by sending HTTP requests to server and expect status code < 400.
If server is down or request is timed-out, execute a shell script.

### Quickstart

Install 

```
go install github.com/phannam1412/downnoti
```

Usage 

```
NAME:
   downnoti - A new cli application

USAGE:
   downnoti [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --exec value, -e value      
   --interval value, -i value  interval in seconds (default: 60)
   --help, -h                  show help (default: false)
```

### EXAMPLES

Send slack message when your website https://example.com is down

```
downnoti -e "curl -X POST -H Content-type: application/json' --data '{\"text\": \"your website is down, please hurry !\"}' \
  <YOUR SLACK WEBHOOK>" \
  https://example.com
```

