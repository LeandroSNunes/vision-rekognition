# vision-rekognition
This repository lets you compare image content analysis performed by Google Vísion and AWS Rekognition

# How use it?
## Configurations
Clone this project in your machine

### Rekognition
Get your credentials in AWS console and create the config.yml file in root path

> config.yml
```yml
AWS_REGION: us-east-1
AWS_ACCESS_KEY_ID: AAAAAAAAAAA
AWS_SECRET_ACCESS_KEY: AAAAAAAAAA

```

### Vision
Set up a project in Google Cloud and enable the Vision API, after obtaining a credential file, save it to the root path with name `vision-credentials.json`

### Images
Add sample images for comparison in image folder

### Exec
At this point you should have something like this in the project

```
images/
  01.jpg
  02.jpg
config.yml
vision-credentials.json
```

So, execute `make run`

```shel
$ make run
go build -o vrekog . && ./vrekog
02.jpg -- Vision: Medical: 2.00 Racy: 5.00 Spoof: 1.00 Violence: 1.00 Adult: 3.00
02.jpg -- REKOG:
------------------------
01.jpg -- Vision: Adult: 2.00 Medical: 1.00 Racy: 4.00 Spoof: 1.00 Violence: 1.00
01.jpg -- REKOG: Revealing Clothes: 75.38 Suggestive: 75.38
------------------------
03.jpg -- REKOG: Graphic Female Nudity: 99.88 Explicit Nudity: 99.88 Sexual Activity: 94.61
03.jpg -- Vision: Adult: 5.00 Medical: 1.00 Racy: 5.00 Spoof: 1.00 Violence: 1.00
```

