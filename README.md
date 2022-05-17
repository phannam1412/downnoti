# downnoti

Check if a server is up by sending HTTP requests to server and expect status code < 400.
If server is down or request is timed-out, execute a shell script.

### EXAMPLES

Send slack message when your website https://example.com is down

```
downnoti -e "curl -X POST -H Content-type: application/json' \
  --data '{\"text\": \"your website is down, please hurry !\"}' <YOUR SLACK WEBHOOK>" \
  https://example.com
```

