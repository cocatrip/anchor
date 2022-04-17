# Required Variable
Anchor will look for these variable inside config.yaml and exit if one or more variable doesn't exist
- APPLICATION_NAME
- BUSINESS_NAME
- TESTING_TAG
- SERVER_NAME
- TEMPLATE

# Usage

## Define variable inside config.yaml
```yaml
// config.yaml
APPLICATION_NAME: ad1-lead-main
BUSINESS_NAME: administrasi
TESTING_TAG: uat
SERVER_NAME: https://adira.co.id
TEMPLATE: maven

jenkins:
  example: This is an example
```

## Using variable
```groovy
// Jenkinsfile
pipeline {
  environment {
    ...
    String SonarProjectKey = '[[.Global.APPLICATION_NAME]]'
    String SonarHost = '[[.Jenkins.example]]'
    ...
  }
  ...
}

// Result
pipeline {
  environment {
    ...
    String SonarProjectKey = 'ad1-lead-main'
    String SonarHost = 'This is an example'
    ...
  }
  ...
}
```

## Enabling command completion

```
anchor completion [shell] --help
```
