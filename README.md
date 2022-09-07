# Running the app

Open command prompt in project directory.
Build image with command:

```shell
docker build -t btc_app .
```

To run container use:

```shell
docker run -p 80:80 btc_app
```

# Using the app

App can provide bitcoin rate in uah, subscribe email
for getting bitcoin rate and send rate to subscribed emails

### /rate

**GET**

BTC rate in UAH.
Plain integer value

### /subscribe

**POST. application/x-www-form-urlencoded**

Subscribe email for future rate emails.
Email should be sent as request body with key "email".

### /sendEmails

**POST**

Send bitcoin rate to all subscribed emails 


### More info can be found in api/swagger.yaml