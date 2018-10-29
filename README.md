# docker-go-hook

Webhook for sending notifications from docker hub builds to slack

### Setting up Slack webhook

https://api.slack.com/incoming-webhooks


### Deploy to GCP App service

Quickstart for Go 1.11 in the App Engine Standard Environment
https://cloud.google.com/appengine/docs/standard/go111/quickstart

1. Create app
```
$ gcloud app create
```

2. Pull this project

3. Add you slack webhook to app.yaml
```yaml
runtime: go111

instance_class: F1

env_variables:
  SLACK_HOOK: "YOUR_SLACK_WEBHOOK"
#  PORT: 80 Optional parameter
```

4. Deploy app to cloud
```
$ gcloud app deploy
```

### Add webhook to Docker hub

https://docs.docker.com/docker-hub/webhooks/


### Altermatives

You can build this application and run on any enviroment supported by Go as standalone webservice.
You can modify this repository to be deployed on Heroku.


### Maintainer
Stefan Monko || smonko@simianlabs.io  
`Simian Labs` - (https://github.com/simianlabs)  
http://simianlabs.io || sl@simianlabs.io