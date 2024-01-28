<a name="readme-top"></a>


<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Golang][golang-shield]][golang-url]
[![Issues][issues-shield]][issues-url]
[![GNU License][license-shield]][license-url]
[![paulmuenzner.com][website-shield]][website-url]
[![paulmuenzner github][github-shield]][github-url]
[![Contributors][contributors-shield]][contributors-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/paulmuenzner/golang-backup-server">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Golang Backup Server</h3>

  <p align="center">
    Circular Buffer Backups For MondoDB Using AWS S3 Storage
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
- Recurring Backups: Define automatic, scheduled backups using cron jobs.
- CSV Format: Each MongoDB collection is saved as a CSV file, offering simplicity and human readability. 
- Define max and limit size of csv files. Large collections are splitted into multiple numbered files.
- AWS S3 Integration: Backups are securely uploaded to AWS S3 for reliable and scalable storage. Multipart uploads are applied automatically for large csv files improving throughput by uploading a number of parts in parallel.
- Pagination implemented to handle large object lists with AWS S3.
- Configuration Flexibility: Easily modify several parameters, such as for cron jobs and adjust the number of kept backups through a flexible configuration system.
- Circular Buffer: Implement a circular buffer to manage and optimize backup storage.
- Store backups optionally on your local machine.


### Advantage CSV Backups

Choosing to make backups in CSV format offers simplicity, portability, and human readability. CSV files are plain text, making them easy to understand, edit, and share across various platforms. They don't rely on database-specific tools (eg. MongoDB Compass), providing independence and ease of use. Additionally, CSV allows for straightforward analysis, version control, and transparency into data structure. This format is database-agnostic, facilitating compatibility and reducing dependencies. While MongoDB backups have their advantages, CSV backups are often preferred for their versatility and accessibility.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Tech Stack <a name="tech-stack"></a>

This project is basically built with and applies:

* [![Aws][aws-shield]][aws-url]
* [![Golang][golang-shield]][golang-url]
* [![MongoDB][mongodb-shield]][mongodb-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Prior to launching the program, clone the repo, install go dependencies and ensure that all configurations are set. 


### Prerequisites 
- Make sure MongoDB is installed and available locally.


### Installation

- Clone the repo
   ```sh
   git clone https://github.com/paulmuenzner/golang-backup-server.git
   ```
- Install go dependencies by running
   ```sh
   go get
   ```

### Environment file (.env)
Before running the program, you need to set up the required environment variables by creating a .env file in the root directory of the project. This file holds sensitive information and configurations needed for the proper functioning of the application.

#### Mandatory Environment Variables

AWS S3 Configuration:

If your application involves interactions with AWS S3, you must provide the following key-value pairs in the .env file:

- BUCKET_NAME: The name of your AWS S3 bucket.
- AWS_REGION: The AWS region where your S3 bucket is located.
- AWS_ACCESS_KEY_ID: Your AWS access key ID.
- AWS_SECRET_ACCESS_KEY: Your AWS secret access key.

#### Optional Environment Variables

Email Notification Configuration:

If you intend to use email notifications (configured with SendEmailNotifications in the config file), include the following additional variables in your .env file:

- EMAIL_PROVIDER_PASSWORD: Password for the email provider.
- EMAIL_PROVIDER_USERNAME: Username for the email provider.
- SMTP_PORT_ENV: SMTP port for the email provider.
- EMAIL_PROVIDER_HOST: Hostname of the email provider.
- EMAIL_ADDRESS_SENDER_BACKUP: Sender email address for backup notifications.
- EMAIL_ADDRESS_RECEIVER_BACKUP: Receiver email address for backup notifications.

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

# Email Notification Configuration (Optional)
EMAIL_PROVIDER_PASSWORD=your-email-provider-password
EMAIL_PROVIDER_USERNAME=your-email-provider-username
SMTP_PORT_ENV=your-smtp-port
EMAIL_PROVIDER_HOST=your-email-provider-host
EMAIL_ADDRESS_SENDER_BACKUP=your-sender-email-address
EMAIL_ADDRESS_RECEIVER_BACKUP=your-receiver-email-address

   ```


### Configuration
<a name="configuration"></a>
The following configurations can be modified in the config file located at => /config/base_config.go

| Key                               |  Description |  Example 
|:-----                             |:---------    |:---------    
| DeleteLogsAfterDays               | Errors are logged to 'log/'-folder. Backup file names are assigned by day. All logs generated in one day are collected in a designated backup file. This parameter (of type int) indicates after how many days log files are stored before they are deleted automatically. |    5   
| NameDatabase                      |  Configure your database name (of type string) you like to backup. It must be 100% identical to the MongoDB database name. |  "MyProjectDB"      
| FolderNameBackup                  | Determine the folder name where your backup is stored in the cloud; inside the S3 bucket. | "mydbbackup"     
| FileNameMetaData                  |File name (of type string) for meta data file containing information on each created backup file |  "meta_data.csv"         
| IntervalBackup                    |Cron-like syntax (of type string) format to define the recurring schedule of your automatic backup  |"@every 6h"
| MaxFileSizeInBytes                | The maximum size (of type int64) of a backup file. Be aware of the max upload size permitted by AWS S3. Of the configured file size is not sufficient, a new backup file is created with the same name plus an added sequential numbering at the end           | 2 * 1024 * 1024 * 1024
| SendEmailNotifications            |Decide whether you want to send email notifications or not. Emails are send in both cases error and successfully completed backup. Type is bool    | false
| EmailProviderUserNameEnv          |Name of .env key (of type string). The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| "EMAIL_PROVIDER_USERNAME"
| EmailProviderPasswordEnv          |Name of .env key (of type string). The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| "EMAIL_PROVIDER_PASSWORD"
| SmtpPortEnv                       | Name of .env key (of type string). The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.          | "SMTP_PORT_ENV"
| HostEmailProviderEnv              |Name of .env key (of type string). The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| "EMAIL_PROVIDER_HOST"
| EmailAddressSenderEnv             |Name of .env key (of type string). The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| "EMAIL_ADDRESS_SENDER_BACKUP"
| EmailAddressReceiverEnv           |Name of .env key (of type string). The value behind this .env key is placed in your .env file. Needed, if you want to send transactional email notifications. Ask your provider for this value.| "EMAIL_ADDRESS_RECEIVER_BACKUP"
| IsCircularBufferActivatedS3       |Decide whether you like to implement a circular buffer or not. If 'false', all backups on S3 are stored without deleting them, which might increase costs depending on backup interval and database size. Type is bool. | true
| MaxBackupsS3                      |Configuration for circular buffer (of type int). If IsCircularBufferActivatedS3 is set to true, circular buffer deletes backups older than latest number of MaxBackupsS3 in S3. In this example, the 12 newest backups are stored on S3 only - older ones are deleted. | 12
| UseLocalBackupStorage             |Decide (with type bool) if you like to store backups on your local machine (where this program is running), too. Type is bool.| true 
| IsCircularBufferActivatedLocally  | Same function as 'IsCircularBufferActivatedS3' but for local backup storage. | true
| MaxBackupsLocally                 | Same as 'MaxBackupsS3' but for local backup storage. If MaxBackupsLocally is set to true, circular buffer deletes backups older than latest number of MaxBackupsLocally locally.| 10
| S3BucketEnvProd                   |Name of .env key (of type string) to configure bucket name. The value behind this .env key is placed in your .env file. Needed, to configure AWS S3. Check your S3 AWS dashboard for this value. The bucket with the exact same name must be created in your AWS account.| "BUCKET_NAME"
| S3RegionEnvProd                   |Name of .env key (of type string) to configure S3 region. The value behind this .env key is placed in your .env file. The region with the exact same name is mentioned in your AWS account.| "AWS_REGION"
| S3AccessKeyEnvProd                |Name of .env key (of type string) to add S3 access key. The value behind this .env key is placed in your .env file. The access key is available in your AWS account.| "AWS_ACCESS_KEY_ID"
| S3SecretKeyEnvProd                |Name of .env key (of type string) to add S3 secret key. The value behind this .env key is placed in your .env file. The secret key is available in your AWS account.| "AWS_SECRET_ACCESS_KEY"    

### Run program

- Run program by: `go run main.go` or use live-reloader such as [air](https://github.com/cosmtrek/air) with `air`


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Add optional circular buffer feature for S3 
- [x] Add optional circular buffer feature for local storage
- [x] Add optional email email notification feature
- [ ] Add gzip compression feature for entire backup files
- [ ] Extend testing
- [ ] Add option to also backup SQL databases besides MongoDB
- [ ] Add option to backup multiple databases
- [ ] Add option to upload backups to MS Azure


See the [open issues](https://github.com/paulmuenzner/golang-backup-server/issues) to report bugs or request fatures.

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

Project Link: [https://github.com/paulmuenzner/golang-backup-server](https://github.com/paulmuenzner/golang-backup-server)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [AWS S3 Upload Size](https://docs.aws.amazon.com/AmazonS3/latest/userguide/upload-objects.html)
* [MongoDB Go Docs](https://www.mongodb.com/docs/drivers/go/current/quick-start/)
* [AWS SDK for Go V2 Docs][aws-url]
* [Gomail Docs](https://pkg.go.dev/gopkg.in/gomail.v2?utm_source=godoc)
* [Testing](https://pkg.go.dev/testing) & [assert](https://pkg.go.dev/github.com/stretchr/testify/assert)
* [Cron](https://pkg.go.dev/github.com/robfig/cron/v3)


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[golang-shield]: https://img.shields.io/badge/golang-black.svg?style=for-the-badge&logo=go&logoColor=ffffff&colorB=00ADD8
[golang-url]: https://go.dev/
[aws-shield]: https://img.shields.io/badge/aws_s3-black.svg?style=for-the-badge&logo=amazons3&logoColor=ffffff&colorB=569A31
[aws-url]: https://aws.github.io/aws-sdk-go-v2/docs/
[mongodb-shield]: https://img.shields.io/badge/mongodb-black.svg?style=for-the-badge&logo=mongodb&logoColor=ffffff&colorB=47A248
[mongodb-url]: https://go.dev/
[github-shield]: https://img.shields.io/badge/paulmuenzner-black.svg?style=for-the-badge&logo=github&logoColor=ffffff&colorB=000000
[github-url]: https://github.com/paulmuenzner
[contributors-shield]: https://img.shields.io/github/contributors/paulmuenzner/golang-backup-server.svg?style=for-the-badge
[contributors-url]: https://github.com/paulmuenzner/golang-backup-server/graphs/contributors
[forks-shield]: https://img.shields.io/badge/FORKS-blue?style=for-the-badge
[forks-url]: https://github.com/paulmuenzner/golang-backup-server/network/members
[issues-shield]: https://img.shields.io/github/issues/paulmuenzner/golang-backup-server.svg?style=for-the-badge
[issues-url]: https://github.com/paulmuenzner/golang-backup-server/issues
[license-shield]: https://img.shields.io/github/license/paulmuenzner/golang-backup-server.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[website-shield]: https://img.shields.io/badge/www-paulmuenzner.com-blue?style=for-the-badge
[website-url]: https://paulmuenzner.com
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 