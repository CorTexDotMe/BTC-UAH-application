# Testing the app

As an example I provide command prompt curl commands to test the app on Windows.

File C:\Windows\System32\drivers\etc\hosts has to be changed. One line to add:

```shell
127.0.0.1 gses2.app
```
### /rate
```shell
curl --location --request GET "gses2.app/api/rate"
```

### /subscribe
To subscribe 123example@gmail.com

```shell
curl --location --request POST "gses2.app/api/subscribe" \
--header "Content-Type: application/x-www-form-urlencoded" \
--data-urlencode "email=123example@gmail.com"
```
Same command in one line
```shell
curl --location --request POST "gses2.app/api/subscribe" --header "Content-Type: application/x-www-form-urlencoded" --data-urlencode "email=123example@gmail.com"
```
### /sendEmails
```shell
curl --location --request POST "gses2.app/api/sendEmails"
```

