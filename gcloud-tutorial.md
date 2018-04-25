

// create an app using gcloud cli

gcloud app create

// run app in cgloud terminal

goapp serve app.yaml

// example app.yaml
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app


// deploy app to live
goapp deploy -application name-of-gcloud-project -version 1

url is metal-dimension-115204.appspot.com


// verifying current working project on gcloud
$ gcloud config list 


// configuring gcloud to current project
$ gcloud init

// explicitly set a project 
$ gcloud config set project <PROJ-ID>

// list current accounts projects
$ gcloud projects list


// gcloud golang tutorial
https://cloud.google.com/appengine/docs/standard/go/building-app/handling-forms

https://cloud.google.com/sdk/gcloud/reference/app/deploy

https://cloud.google.com/appengine/docs/standard/go/config/appref

https://cloud.google.com/appengine/docs/standard/go/building-app/creating-your-application

https://cloud.google.com/appengine/docs/standard/go/runtime




