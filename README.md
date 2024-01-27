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
    An awesome README template to jumpstart your projects!
    <br />
    <a href="https://github.com/paulmuenzner/golang-backup-server"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/paulmuenzner/golang-backup-server">View Demo</a>
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
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

There are many great README templates available on GitHub; however, I didn't find one that really suited my needs so I created this enhanced one. I want to create a README template so amazing that it'll be the last one you ever need -- I think this is it.

Here's why:
* Your time should be focused on creating something amazing. A project that solves a problem and helps others
* You shouldn't be doing the same tasks over and over like creating a README from scratch
* You should implement DRY principles to the rest of your life :smile:

Of course, no one template will serve all projects since your needs may be different. So I'll be adding more in the near future. You may also suggest changes by forking this repo and creating a pull request or opening an issue. Thanks to all the people have contributed to expanding this template!

Use the `BLANK_README.md` to get started.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built For/With

This project is basically built with and applies:

* [![Aws][aws-shield]][aws-url]
* [![Golang][golang-shield]][golang-url]
* [![MongoDB][mongodb-shield]][mongodb-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Before starting the program, make sure to make all configurations ==> <a href="#configuration">Configuration</a>.

### Prerequisites & Installation

- Make sure MongoDB is installed and available locally.
- Clone the repo
   ```sh
   git clone https://github.com/paulmuenzner/golang-backup-server.git
   ```
- Install go dependencies by running
   ```sh
   go get
   ```
- Run program by: `go run main.go` or use live-reloader such as [air](https://github.com/cosmtrek/air) with `air`




<!-- USAGE EXAMPLES -->
## Features

- Upload and backup your entire MongoDB database in csv file format to AWS S3
- No limit for MongoDB database size. Collections which are too large for S3 are split
- Pagination implemented to handle large object lists with AWS S3


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONFIGURATION -->
## Configuration
<a name="configuration"></a>
The following configurations can be made in the config file => /config/base_config.go

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