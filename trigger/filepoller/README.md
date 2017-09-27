# directory poller
This trigger provides your flogo application the ability to continuously poll a directory for capturing any changes happening inside it (like create a new file, modify existing file, etc.)

## Installation

```bash
flogo add trigger github.com/anshulsharmas/flogo-contrib/trigger/filepoller
```

## Schema
Outputs and Handlers:

```json
{
"outputs": [
    {
      "name": "filename",
      "type": "string"
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "dirName",
        "type": "string"
      }
    ]
   }
```

## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the directory poller.

### Providing the directory name to be polled
Configure the Trigger to start a flogo flow. "settings" "dirName" is the directory name it polls continuously for any changes. 
So in this case the flogo flow will be triggered whenever any change is detected in the directory "Full_Path_to_test_directory".
Changes can be write operation on the existing file, creating a new file, etc.

```json
{
  "triggers": [
    {
      "name": "Directory Poller",
      "description": "Poller",
      "settings": {
        "dirName": "Full_Path_to_test_directory"
      },
      "id": "directory_poller",
      "handlers": [
        {
          "settings": {
            "dirName": "Full_Path_to_test_directory"
          },
          "actionId": "cc"
        }
      ]
    }
  ]
}
```
