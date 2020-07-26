## Action Resource

#### Properties

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_id`|yes|number|The application to add the action to|`32423`|
|`action_type`|yes|string|The type of the action|`"EMAIL"`|
|`emails`|no|string[]|The email addresses to be notified when he action is performed|`["bob@example.com"]`|

###### Action Types
- SMS
- EMAIL
- CUSTOM_EMAIL
- THREAD_DUMP
- HTTP_REQUEST
- CUSTOM
- RUN_SCRIPT_ON_NODES
- DIAGNOSE_BUSINESS_TRANSACTIONS
- CREATE_UPDATE_JIRA

#### Examples

###### Email
```terraform
resource "appd_action" "my_first_email_action" {
  application_id = var.application_id
  action_type = "EMAIL"
  emails = [
    "bob@example.com",
    "sandra@example.com"
  ]
}
```