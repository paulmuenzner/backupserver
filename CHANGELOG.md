# Changelog

## [1.0.6] - 2024-02-03

-   Decrease potential risk of users or processes gaining unauthorized access to files by using more restrictive permissions 0644 -> 0600 and 0777 -> 0750
-   Omit comparisons to bool constants
-   Update .golangci.yml

## [1.0.5] - 2024-02-02

-   Add database connection test using real database connection (no mock)
-   Standardization interface names
-   Removing clear text error logging of aws config data

## [1.0.4] - 2024-02-01

-   Change to human readable US format for dates in body of email notifications.
-   Formatting HTML email body.
-   Clean up .env file

## [1.0.3] - 2024-02-01

-   Extend email notifications for failed backups. Finally, any backup related error will can trigger an email notification.

## [1.0.2] - 2024-01-31

-   Standardization env keys

## [1.0.1] - 2024-01-29

-   Improvement Dependency Injection Setup for MongoDB client and methods
-   Revision README.md due to improved MongoDB client and methods


## [1.0.0] - 2024-01-28

-   Initial release
