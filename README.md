# cogs

A simple server for executing adhoc scripts or running scripts on a schedule all in a single file

## Features

- [x] Run scripts on demand or cron schedule
- [x] Webhook, email and Slack notifications
- [x] Script execution history
- [x] Light brand configuration
- [x] Secret management
- [x] Permissions per project
- [x] Web UI
- [x] REST API (Basic)

## Quick Start (non docker)

Download the latest release from the [releases page](https://blank)

Start cogs with the following command

```bash
./cogs -data=/path/to/data
```

Once Cogs has launched navigate to http://localhost:8080 and login with the default username and password - **make sure to change these after logging in**

```
Username: admin@localhost.com
Password: ChangeMe1234!
```

## Docker Quick Start

TODO

## Best Practices

- **Change the default admin password**
- **Do not run cogs as root as script execution can leak into your environment**
- **Do not run cogs on a public network as it does not have MFA for the web UI**

## Configuration

You can configure cogs with either command line flags or environment variables.  The following table shows the available options

When both a flag and environment variable are set, the environment variable takes precedence.

**If you do not supply values for the smtp configuration options your email notifications will fail to send**

| Flag           | Environment Variable        | Description                                             | Default               |
|----------------|-----------------------------|---------------------------------------------------------|-----------------------|
| -data          | COGS_DATA                   | The path to the data directory where all data is stored | ./data                |
| -port          | COGS_PORT                   | The port to run the server on                           | 8080                  |
| -callback      | COGS_CALLBACK_URL           | The base URL for notifications to link back to          | http://localhost:8080 |
| -smtp-host     | COGS_SMTP_HOST              | The SMTP host to use for sending emails                 |                       |
| -smtp-port     | COGS_SMTP_PORT              | The SMTP port to use for sending emails                 | 25                    |
| -smtp-username | COGS_SMTP_USERNAME          | The SMTP username to use for sending emails             |                       |
| -smtp-password | COGS_SMTP_PASSWORD          | The SMTP password to use for sending emails             |                       |
| -smtp-from     | COGS_SMTP_FROM              | The SMTP from address to use for sending emails         | cogs@localhost.com    |
| -level         | COGS_LOG_LEVEL              | The log Level (DEBUG, INFO, WARN, ERROR)                | INFO                  |
| -format        | COGS_LOG_FORMAT             | The log format                                          | text                  |
| -brand         | COGS_BRAND                  | The brand name for login and headers                    | Cogs                  |
| -retention     | COGS_HISTORY_RETENTION_DAYS | The number of days to keep job history                  | 30                    |

## Notifications

### Slack

Slack Notifications use the [Slack Incoming Webhook](https://api.slack.com/messaging/webhooks) integration.  You will need to create a webhook and supply the URL to cogs.

### Email

Email notifications use SMTP to send emails.  You will need to supply the SMTP host, port, username and password to cogs either through flags or environment variables.

### Webhook

Webhooks notifications are always sent as `POST` requests and include the data schema below.

```json
{
  "project_name": "Your Project Name",
  "script_name": "Your Script Name",
  "history_link": "http://localhost:8080/projects/3/5/history/17",
  "success": true,
  "triggered_by": "admin@localhost.com",
  "trigger": "webUI",
  "duration_seconds": 3,
  "run_id": 17,
  "created_at": "2023-07-21T22:07:48.587324101-05:00"
}
```

## Using Secrets

Secrets are scoped to the project level and therefore are created inside each project.  Scripts reference secrets in their project with environment variables like below

```bash
# A secret created with the name "password" would be referenced as $COGS_SECRET_PASSWORD
echo $COGS_SECRET_PASSWORD
```

## Using Schedules

Each job may have a schedule set and either enabled or disabled.  Schedules that are configured but disabled do not run.

A schedule must be specified in the [cron format](https://en.wikipedia.org/wiki/Cron#CRON_expression) and may be set to run at any interval.

**Note:** Scheduled scripts will not have any input values available to them

### Special cases

- If a job is running and the schedule is set to run the job will run a second time
- If a job is set to run frequently and a previous run does not finish the new run will still start

## Specifying Script Inputs

Script inputs are written in raw JSON currently, the below example shows both a drop-down (select) implementation and a free form implementation

When `strict_options` is set to true you must provide an array of options

Arguments are accessed by scripts as environment variables.  The name of the argument is converted to uppercase and prefixed with `COGS_`


```json
[
  {
    "name": "Name",
    "description": "Freeform input example",
    "strict_options": false
  },
  {
    "name": "Greeting",
    "description": "Dropdown input example",
    "strict_options": true,
    "options": [
      "hello",
      "hi",
      "howdy",
      "sup"
    ]
  }
]
```

When the above input is used in a script the following environment variables will be available for the script to use

```bash
echo $COGS_NAME
echo $COGS_GREETING
```

## Backups and Restoration

Cogs data is stored in a sqlite3 database and can be backed up by simply creating a copy of that file and storing it in a different location.  The file is `data.db` and can be found at the location of `-data` or `COGS_DATA`

Restoring a backup is as simple as stopping the service and then replacing the `data.db` file with the backup file.

## REST API

The REST API is available at `/api/v1/` and requires an administrator API token to interact with.

### Authentication

All requests to the API must have the header `x-api-key` set to a valid API token with administrator privileges.

Available endpoints are

- `GET /api/v1/projects` | Lists all projects
- `GET /api/v1/projects/{project_id}` | Lists all scripts within the specified project
- `GET /api/v1/scripts/{script_id}` | Gets the specified script
- `PUT /api/v1/scripts/{script_id}` | Updates the specified script
- `POST /api/v1/run/{script_id}` | Runs the specified script
- `GET /api/v1/history/{script_id}` | Lists execution history for the specified script
- `GET /api/v1/history/{script_id}/{history_id}` | Gets the specified execution history

## Development

### Running cogs locally for development

```bash
go run cmd/server/main.go -data=tmp -port=8080
```

### Adding Database Schemas

```bash
go run -mod=mod entgo.io/ent/cmd/ent new NameOfObject
```

Generate the updated models

```bash
go generate ./ent
```