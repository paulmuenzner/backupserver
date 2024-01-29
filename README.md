<a name="readme-top"></a>


<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- [![Golang][golang-shield]][golang-url] -->
[![Go Report Card](https://goreportcard.com/badge/github.com/paulmuenzner/backupserver)](https://goreportcard.com/report/github.com/paulmuenzner/backupserver)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/650472443a824243884952184a6732dd)](https://app.codacy.com/gh/paulmuenzner/backupserver/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Issues][issues-shield]][issues-url]
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/paulmuenzner/backupserver)
[![GNU License][license-shield]][license-url]
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/paulmuenzner/backupserver)
![GitHub top language](https://img.shields.io/github/languages/top/paulmuenzner/backupserver)
 <!-- [![paulmuenzner.com][website-shield]][website-url] -->
[![paulmuenzner github][github-shield]][github-url] 
[![Contributors][contributors-shield]][contributors-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/paulmuenzner/backupserver">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Golang Backup Server</h3>

  <p align="center">
    Circular Buffer Backups For MongoDB Using AWS S3
    <br />
    <a href="#about-the-project"><strong>EXPLORE DOCS</strong></a>
    <br />
    <br />
    <a href="#configuration">High Flexibility</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Report Bug</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This Golang-based server is designed to automate recurring backups of MongoDB databases, offering flexibility and ease of use. The server converts each collection into CSV file format, providing human-readable backups that are easily managed. Backups are then uploaded to AWS S3, ensuring secure and scalable storage.


### Features
-   **Recurring Backups**: Define automatic, scheduled backups using cron jobs.
-   **Circular Buffer**: Implement a circular buffer to manage and optimize backup storage.
-   **CSV Format**: Each MongoDB collection is saved as a CSV file, offering simplicity and human readability. 
-   **Configuration Flexibility**: Easily modify several parameters, such as for cron jobs and adjust the number of kept backups thanks to a flexible configuration system.
-   **Limit File Size**: Define max and limit size of csv files. Large collections are splitted into multiple numbered files.
-   **Dependency Injection (DI) setup**: This Golang webserver boasts a robust architecture designed for flexibility, reduced coupling and testibiliy through a dedicated Dependency Injection (DI) setup. The core functionalities of database communications, and AWS operations and sending email notifications are seamlessly integrated, providing a cohesive and modular solution.
-   **AWS S3 Integration**: Backups are securely uploaded to AWS S3 for reliable and scalable storage. Multipart uploads are applied automatically for large csv files improving throughput by uploading a number of parts in parallel.
-   **S3 Pagination**: Pagination implemented to handle large object lists with AWS S3.
-   **Local Backups**: Store backups optionally on your local machine; even with circular buffer functionality.
-   **Robust Error Handling Mechanism**: Any encountered errors are diligently logged to the designated log folder and simultaneously dispatched via email notifications.



### Advantage CSV Backups

Choosing to make backups in CSV format offers simplicity, portability, and human readability. CSV files are plain text, making them easy to understand, edit, and share across various platforms. They don't rely on database-specific tools (eg. MongoDB Compass), providing independence and ease of use. Additionally, CSV allows for straightforward analysis, version control, and transparency into data structure. This format is database-agnostic, facilitating compatibility and reducing dependencies. While MongoDB backups have their advantages, CSV backups are often preferred for their versatility and accessibility.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Tech Stack <a name="tech-stack"></a>

This project is basically built with and for:

*   [![Aws][aws-shield]][aws-url]
*   [![Golang][golang-shield]][golang-url]
*   [![MongoDB][mongodb-shield]][mongodb-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Prior to launching the program, clone the repo, install go dependencies and ensure that all configurations are set. 


### Prerequisites 
-   Make sure MongoDB is installed and available.
-   Make sure a properly configured [AWS S3 Bucket](https://aws.amazon.com/s3/?nc1=h_ls) is ready.


### Installation

-   Clone the repo
   ```sh
   git clone https://github.com/paulmuenzner/backupserver.git
   ```
-   Install go dependencies by running
   ```sh
   go get
   ```

### Environment file (.env)
Before running the program, you need to set up the required environment variables by creating a .env file in the root directory of the project. This file holds sensitive information and configurations needed for the proper functioning of the application.

#### Mandatory Environment Variables

AWS S3 & MongoDB Configuration:

If your application involves interactions with AWS S3, you must provide the following key-value pairs in the .env file:

-   BUCKET_NAME: The name of your AWS S3 bucket.
-   AWS_REGION: The AWS region where your S3 bucket is located.
-   AWS_ACCESS_KEY_ID: Your AWS access key ID.
-   AWS_SECRET_ACCESS_KEY: Your AWS secret access key.
-   MONGO_URI: MongoDB URI (Uniform Resource Identifier) 

#### Optional Environment Variables

Email Notification Configuration:

If you intend to use email notifications (configured with SendEmailNotifications in the config file), include the following additional variables in your .env file:

-   EMAIL_PROVIDER_PASSWORD: Password for the email provider.
-   EMAIL_PROVIDER_USERNAME: Username for the email provider.
-   EMAIL_PROVIDER_SMTP_PORT: SMTP port for the email provider.
-   EMAIL_PROVIDER_HOST: Hostname of the email provider.
-   EMAIL_ADDRESS_SENDER_BACKUP: Sender email address for backup notifications.
-   EMAIL_ADDRESS_RECEIVER_BACKUP: Receiver email address for backup notifications.

#### Important Note

Make sure to keep your '.env' file secure and do not share it publicly.

The program relies on these configurations to run successfully. Without the correct values in the .env file, certain features may not work as expected.

#### Template

Here's an example .env template in code format. Replace "your-..." placeholders with your actual values. Ensure that this file is kept secure, and sensitive information is not shared publicly. Users should fill in the appropriate values for their specific configurations.
```sh
# AWS S3 Configuration
BUCKET_NAME=your-s3-bucket-name
AWS_REGION=your-aws-region
AWS_ACCESS_KEY_ID=your-access-key-id
AWS_SECRET_ACCESS_KEY=your-secret-access-key

# Database Configuration
MONGO_URI=mongodb://localhost:27017

# Email Notification Configuration (Optional)
EMAIL_PROVIDER_PASSWORD=your-email-provider-password
EMAIL_PROVIDER_USERNAME=your-email-provider-username
EMAIL_PROVIDER_SMTP_PORT=your-smtp-port
EMAIL_PROVIDER_HOST=your-email-provider-host
EMAIL_ADDRESS_SENDER_BACKUP=your-sender-email-address
EMAIL_ADDRESS_RECEIVER_BACKUP=your-receiver-email-address

   ```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Configuration
<a name="configuration"></a>
The following configurations can be modified in the config file located at => /config/base_config.go

| Key                               |  Description |  Type |  Example 
|:-----                             |:---------    |:---------  |:---------    
| DeleteLogsAfterDays               | Errors are logged to 'log/'-folder. Log file names are assigned by day. All logs generated during one day are collected in a designated backup file. This parameter indicates after how many days log files will be deleted automatically. | int|   5   
| NameDatabase                      |  Configure your database name you like to backup. It must be 100% identical to the MongoDB database name. | string| "MyProjectDB"      
| FolderNameBackup                  | Determine the folder name where your backup is stored in the cloud; inside the S3 bucket. | string|"mydbbackup"
| FileNameMetaData                  |File name for meta data file containing information on each created backup file | string| "meta_data.csv"  
| IntervalBackup                    |Cron-like syntax format to define the recurring schedule of your automatic backup  |string|"@every 6h"
| MaxFileSizeInBytes                | The maximum size of a backup file. Be aware of the max upload size permitted by AWS S3. Of the configured file size is not sufficient, a new backup file is created with the same name plus an added sequential numbering at the end           |int64| 2 * 1024 * 1024 * 1024
| SendEmailNotifications            |Decide whether you want to send email notifications or not. Emails are send in both cases error and successfully completed backup. |bool| false
| EmailProviderUserNameEnv          |Name of .env key. The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.|string| "EMAIL_PROVIDER_USERNAME"
| EmailProviderPasswordEnv          |Name of .env key. The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.|string| "EMAIL_PROVIDER_PASSWORD"
| EmailProviderSmtpPortEnv                       | Name of .env key. The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.          |string| "EMAIL_PROVIDER_SMTP_PORT"
| EmailProviderHostEnv              |Name of .env key. The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| string|"EMAIL_PROVIDER_HOST"
| EmailAddressSenderEnv             |Name of .env key. The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| string|"EMAIL_ADDRESS_SENDER_BACKUP"
| EmailAddressReceiverEnv           |Name of .env key. The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| string|"EMAIL_ADDRESS_RECEIVER_BACKUP"
| IsCircularBufferActivatedS3       |Decide whether you like to implement a circular buffer or not. If 'false', all backups on S3 are stored without deleting them, which might increase costs depending on backup interval and database size. | bool|true
| MaxBackupsS3                      |Configuration for circular buffer. If IsCircularBufferActivatedS3 is set to true, circular buffer deletes backups older than latest number of MaxBackupsS3 in S3. In this example, the 12 newest backups are stored on S3 only - older backups will be deleted. | int|12
| UseLocalBackupStorage             |Decide if you like to store backups on your local machine (where this program is running on), too.|bool| true 
| IsCircularBufferActivatedLocally  | Same function as 'IsCircularBufferActivatedS3' but for local backup storage. |bool| true
| MaxBackupsLocally                 | Same as 'MaxBackupsS3' but for local backup storage. If MaxBackupsLocally is set to true, circular buffer deletes backups older than latest number of MaxBackupsLocally locally.|int| 10
| S3BucketEnv                   |Name of .env key to configure bucket name. The value behind this .env key is placed in your .env file. Needed, to configure AWS S3. Check your S3 AWS dashboard for this value. The bucket with the exact same name must be ready in your AWS account.|string| "BUCKET_NAME"
| S3RegionEnv                   |Name of .env key to configure S3 region. The value behind this .env key is placed in your .env file. The region with the exact same name is mentioned in your AWS account.| string|"AWS_REGION"
| S3AccessKeyEnv                |Name of .env key to add S3 access key. The value behind this .env key is placed in your .env file. The access key is available in your AWS account.| string|"AWS_ACCESS_KEY_ID"
| S3SecretKeyEnv                |Name of .env key to add S3 secret key. The value behind this .env key is placed in your .env file. The secret key is available in your AWS account.| string|"AWS_SECRET_ACCESS_KEY"    
| MongoURIEnv                |Name of .env key to define a MongoDB URI (Uniform Resource Identifier). The value behind this .env key is placed in your .env file. |string| "MONGO_URI"



### Run program

Run program by: `go run main.go` or use live-reloader such as [air](https://github.com/cosmtrek/air) with `air`


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

-   [+] Add optional circular buffer feature for S3 
-   [+] Add optional circular buffer feature for local storage
-   [+] Add optional email notification feature
-   [-] Add gzip compression feature for entire backup files
-   [-] Extend testing
-   [-] Add option to also backup SQL databases besides MongoDB
-   [-] Add option to backup multiple databases
-   [-] Add option to upload backups to MS Azure


See the [open issues](https://github.com/paulmuenzner/backupserver/issues) to report bugs or request fatures.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Contributions are more than welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for
more info.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the GNU General Public License v2.0. See [LICENSE](LICENSE.txt) for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Paul Münzner: [https://paulmuenzner.com](https://paulmuenzner.com) 

Project Link: [https://github.com/paulmuenzner/backupserver](https://github.com/paulmuenzner/backupserver)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

*   [AWS S3 Upload Size](https://docs.aws.amazon.com/AmazonS3/latest/userguide/upload-objects.html)
*   [MongoDB Go Docs](https://www.mongodb.com/docs/drivers/go/current/quick-start/)
*   [AWS SDK for Go V2 Docs][aws-url]
*   [Gomail Docs](https://pkg.go.dev/gopkg.in/gomail.v2?utm_source=godoc)
*   [Testing](https://pkg.go.dev/testing) & [assert](https://pkg.go.dev/github.com/stretchr/testify/assert)
*   [Cron](https://pkg.go.dev/github.com/robfig/cron/v3)


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[golang-shield]: https://img.shields.io/badge/golang-black.svg?logo=go&logoColor=ffffff&colorB=00ADD8
[golang-url]: https://go.dev/
[aws-shield]: https://img.shields.io/badge/aws_s3-black.svg?logo=amazons3&logoColor=ffffff&colorB=569A31
[aws-url]: https://aws.github.io/aws-sdk-go-v2/docs/
[mongodb-shield]: https://img.shields.io/badge/mongodb-black.svg?logo=mongodb&logoColor=ffffff&colorB=47A248
[mongodb-url]: https://go.dev/
[github-shield]: https://img.shields.io/badge/paulmuenzner-black.svg?logo=github&logoColor=ffffff&colorB=000000
[github-url]: https://github.com/paulmuenzner
[contributors-shield]: https://img.shields.io/github/contributors/paulmuenzner/backupserver.svg
[contributors-url]: https://github.com/paulmuenzner/backupserver/graphs/contributors
[issues-shield]: https://img.shields.io/github/issues/paulmuenzner/backupserver.svg
[issues-url]: https://github.com/paulmuenzner/backupserver/issues
[license-shield]: https://img.shields.io/github/license/paulmuenzner/backupserver.svg
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[website-shield]: https://img.shields.io/badge/www-paulmuenzner.com-blue
[website-url]: https://paulmuenzner.com
[product-screenshot]: images/screenshot.png
