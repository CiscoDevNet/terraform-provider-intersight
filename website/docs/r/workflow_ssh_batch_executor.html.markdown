---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_ssh_batch_executor"
description: |-
        Intersight allows generic tasks to be created by taking the executor request
        body and a response parser specification in the form of content.Grammar object.
        SSH Batch associates the list of SSH requests to be executed as part of single
        task execution. Each SSH request takes the command to execute and a response parser
        specification based off text to extract fields of interest.

---

# Resource: intersight_workflow_ssh_batch_executor
Intersight allows generic tasks to be created by taking the executor request
body and a response parser specification in the form of content.Grammar object.
SSH Batch associates the list of SSH requests to be executed as part of single
task execution. Each SSH request takes the command to execute and a response parser
specification based off text to extract fields of interest.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `batch`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workflow.AnsiblePlaySession](#workflowAnsiblePlaySession)
[workflow.CliCommand](#workflowCliCommand)
[workflow.FileOperations](#workflowFileOperations)
[workflow.SshSession](#workflowSshSession)
[workflow.WebApi](#workflowWebApi)
[workflow.XmlApi](#workflowXmlApi)
  + `asset_target_moid`:(string)(ReadOnly) Asset target defines the remote target endpoints which are either managed byIntersight or their service APIs are invoked from Intersight. Generic API executorservice Jasmine can invoke these remote APIs via its executors. Asset targets can beaccessed directly for cloud targets or via an associated Intersight Assist. Prerequisiteto use asset targets is to persist the target details.Asset target MoRef can be given as task input of type TargetDataType. Fusion determinesand populates the target context with the Assist if associated with. It is setinternally at the API level.In case of an associated assist, it is used by Assist to fetch the target detailsand form the API request to send to endpoints. In case of cloud asset targets, Jasminefetched the target details from DB, forms the API request and sends it to the target. 
  + `body`:(string) The optional request body that is sent as part of this API request.The request body can contain a golang template that can be populated with task inputparameters and previous API output parameters. 
  + `content_type`:(string) Intersight Orchestrator, with the support of response parser specification,can extract the values from API responses and map them to task output parameters.The value extraction is supported for response content types XML, JSON and Text.The type of the content that gets passed as payload and response in thisAPI. The supported values are json, xml, text. 
  + `description`:(string) A description that task designer can add to individual API requests that explain what the API call is about. 
  + `error_content_type`:(string) Intersight Orchestrator, with the support of response parser specification,can extract the values from API responses and map them to task output parameters.The value extraction is supported for response content types XML, JSON and Text.Optional input to specify the content type in case of error API response. Thisshould be used if the content type of error response is different from that ofthe success response. If not specified, contentType input value is used to parsethe error response. 
  + `label`:(string) A user friendly label that task designers have given to the batch API request. 
  + `name`:(string) A reference name for this API request within the batch API request.This name shall be used to map the API output parameters to subsequentAPI input parameters within a batch API task. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `outcomes`:(JSON as string) All the possible outcomes of this API are captured here. Outcomes propertyis a collection property of type workflow.Outcome objects.The outcomes can be mapped to the message to be shown. The outcomes areevaluated in the order they are given. At the end of the outcomes list,an catchall success/fail outcome can be added with condition as 'true'.This is an optionalproperty and if not specified the task will be marked as success. 
  + `response_spec`:(JSON as string) The optional grammar specification for parsing the response to extract therequired values.The specification should have extraction specification specified forall the API output parameters. 
  + `skip_on_condition`:(string) The skip expression, if provided, allows the batch API executor to skip theapi execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed. 
  + `start_delay`:(int) The delay in seconds after which the API needs to be executed.By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution.Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out. 
  + `timeout`:(int) The duration in seconds by which the API response is expected from the API target.If the end point does not respond for the API request within this timeoutduration, the task will be marked as failed. 
* `constraints`:(HashMap) - Enter the constraints on when this task should match against the task definition. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `target_data_type`:(JSON as string) List of property constraints that helps to narrow down task implementations based on target device input. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) A detailed description about the batch APIs. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the batch API task. 
* `outcomes`:(JSON as string) All the possible outcomes of this task are captured here. Outcomes propertyis a collection property of type workflow.Outcome objects.The outcomes can be mapped to the message to be shown. The outcomes areevaluated in the order they are given. At the end of the outcomes list,an catchall success/fail outcome can be added with condition as 'true'.This is an optionalproperty and if not specified the task will be marked as success. 
* `output`:(JSON as string) Intersight Orchestrator allows the extraction of required values from APIresponses using the API response grammar. These extracted values can be mappedto task output parameters defined in task definition.The mapping of API output parameters to the task output parameters is providedas JSON in this property. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `retry_from_failed_api`:(bool) When an execution of a nth API in the Batch fails,Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry.By default the value is set to false. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_on_condition`:(string) The skip expression, if provided, allows the batch API executor to skip thetask execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `task_definition`:(HashMap) - A reference to a workflowTaskDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `ui_rendering_data`:(JSON as string) This will hold the data needed for task to be rendered in the user interface. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_workflow_ssh_batch_executor` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_ssh_batch_executor.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [workflow.AnsiblePlaySession](#argument-reference)
This models a single Ansible playbook execution session on the Ansible Control node. While execution of the respective Ansible task, the below provided property values are used to construct the SSH Batch executor, which gets executed on the endpoint.
* `command_line_arguments`:(string) The command line arguments for running the Ansible playbook against the given endpoint. Escape character backslash needs to be used when the command line arguments contain double quotes in them. 
* `host_inventory`:(string) The path of the host inventory file that resides on the Ansible Endpoint target or the comma separated list of hosts on which the Ansible playbook is to be run. Make sure to suffix a comma when the list of hosts is provided as input, even if the list has only one value. 
* `playbook_path`:(string) The path of the Ansible playbook that resides on the Ansible Endpoint target. 
* `ssh_op_timeout`:(string) SSH operation timeout value in seconds. Value provided should be string representation of an interger. 

### [workflow.CliCommand](#argument-reference)
This models a single CLI command that can be executed on the end point.
* `command`:(string) The command to run on the device connector. 
* `end_prompt`:(string) The regex string that identifies the end of the command response. 
* `expect_prompts`:(Array)
This complex property has following sub-properties:
  + `expect`:(string) The regex of the expect prompt of the interactive command. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `send`:(string) The answer string to the expect prompt. 
* `expected_exit_codes`:
                (Array of schema.TypeInt) -
* `skip_status_check`:(bool) Skips the execution status check of the terminal command. One use case for this is while exiting theterminal session from esxi host. 
* `terminal_end`:(bool) If this flag is set, it marks the end of the terminal session where the previous commands were executed. 
* `terminal_start`:(bool) If this flag is set, the execution of this command initiates a terminal session in which the subsequentCLI commands are executed until a command with terminalEnd flag is encountered or the end of the batch. 
* `type`:(string) The type of the command - can be interactive or non-interactive.* `NonInteractive` - The CLI command is not an interactive command.* `Interactive` - The CLI command is executed in interactive mode and the command must provide the expects andanswers. 

### [workflow.FileOperations](#argument-reference)
This models a single File Operation request within a batch of requests that get
executed within a single workflow task.
* `file_download`:(HashMap) - File operation to download a given file from Intersight storage services such asAWS or Minio bucket to a specified path on one or more Intersight connected devices. 
This complex property has following sub-properties:
  + `destination_path`:(string) Path on the Intersight connected device to which file needs to be downloaded. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `source_bucket`:(string) Source bucket name hosting the file. 
  + `source_file`:(string) Name of the file to be downloaded from bucket to endpoint devices. 
* `file_template`:(HashMap) - Populates data driven template file with input values to generate textual output. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `template_file_path`:(string) Path of the template file on the connected device. 
  + `template_values`:(JSON as string) Input values to render text output file from template file. 
* `operation_type`:(string) File operation type to be executed on the connected endpoint.* `FileDownload` - The API is executed in a remote device connected to the Intersightthrough its device connector. This operation is to download the filefrom specified storage bucket to the specific path on the device.* `FileTemplatize` - Populates data driven template file with input values to generate textual output.Inputs - the path of the template file on the device and json values to populate.An error will be returned if the file does not exists or if there is an error whileexecuting the template. 

### [workflow.SshSession](#argument-reference)
This models a single SSH session from Intersight connected endpoint to a remote
server. Multiple SSH operations can be run sequentially over a single SSH session.
* `capture_complete_response`:(string) Flag to allow capturing entire command response as batch API output. 
* `expected_exit_codes`:(JSON as string) Optional array of integer values to specify the expected exit codes of a SSH command execution. SSH commandexecution is marked success upon receiving any of the expected exit code from command execution. If not set, successexit code of 0 is expected from command execution. 
* `file_transfer_to_remote`:(HashMap) - Message to transfer a file from Intersight connected device to remote server. 
This complex property has following sub-properties:
  + `destination_file_path`:(string) Destination file path on the target server. 
  + `encrypted_aes_key`:(string) The secure properties that have large text content as value can be encrypted using AES key. In these cases, the AES key needs to be encrypted using the device connector public key and passed as the value for this property.The secure properties that are encrypted using the AES key are mapped against the property name with prefix 'AES' in SecureProperties dictionary. 
  + `encryption_key`:(string) The public key that was used to encrypt the values present in SecureProperties dictionary.If the given public key is not same as device connector's public key, an error reponse with appropriate error message is thrown back. 
  + `file_mode`:(int) File permission to set on the transferred file. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `secure_properties`:(JSON as string) A dictionary of encrypted secure values mapped against the secure property name. The values that are encrypted using AES key must be mapped against the secure property name with a 'AES' prefixDevice connector expects the message body to be a golang template and the template can use the secure property names as placeholders. 
  + `source_file_path`:(string) Source file path on the Intersight connected device. 
* `message_type`:(string) The type of SSH message to be sent to the remote server.* `ExecuteCommand` - Execute a SSH command on the remote server.* `NewSession` - Open a new SSH connection to the remote server.* `FileTransfer` - Transfer a file from Intersight connected device to the remote server.* `CloseSession` - Close the SSH connection to the remote server. 
* `ssh_command`:(JSON as string) SSH command to execute on the remote server. 
* `ssh_configuration`:(HashMap) - Carries the SSH session details for opening a new connection. 
This complex property has following sub-properties:
  + `encrypted_aes_key`:(string) The secure properties that have large text content as value can be encrypted using AES key. In these cases, the AES key needs to be encrypted using the device connector public key and passed as the value for this property.The secure properties that are encrypted using the AES key are mapped against the property name with prefix 'AES' in SecureProperties dictionary. 
  + `encryption_key`:(string) The public key that was used to encrypt the values present in SecureProperties dictionary.If the given public key is not same as device connector's public key, an error reponse with appropriate error message is thrown back. 
  + `is_passphrase_set`:(bool)(ReadOnly) Indicates whether the value of the 'passphrase' property has been set. 
  + `is_password_set`:(bool)(ReadOnly) Indicates whether the value of the 'password' property has been set. 
  + `is_private_key_set`:(bool)(ReadOnly) Indicates whether the value of the 'privateKey' property has been set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `passphrase`:(string) Optional passphrase if provided while creating the private key. 
  + `password`:(string) Password to use in the SSH connection credentials (If empty, private key will be used). 
  + `private_key`:(string) PEM encoded private key to be used in the SSH connection credentials (Optional if password is given). 
  + `secure_properties`:(JSON as string) A dictionary of encrypted secure values mapped against the secure property name. The values that are encrypted using AES key must be mapped against the secure property name with a 'AES' prefixDevice connector expects the message body to be a golang template and the template can use the secure property names as placeholders. 
  + `target`:(string) The remote server to connect to. IPv4 address represented in dot decimal notation or hostname can bespecified. 
  + `user`:(string) Username for the remote SSH connection. 
* `ssh_op_timeout`:(string) SSH operation timeout value in seconds. The provided string value should be able to convert torespective integer value. 

### [workflow.WebApi](#argument-reference)
This models a single Web API request within a batch of requests that get
executed within a single workflow task.
* `cookies`:(JSON as string) Collection of key value pairs to set in the request header as Cookie list. 
* `endpoint_request_type`:(string) If the target type is Endpoint, this property determines whether the request isto be handled as internal request or external request by the device connector.* `Internal` - The endpoint API executed is an internal request handled by the device connector plugin.* `External` - The endpoint API request is passed through by the device connector. 
* `headers`:(JSON as string) Collection of key value pairs to set in the request header. 
* `method`:(string) The HTTP method to be executed in the given URL (GET, POST, PUT, etc).If the value is not specified, GET will be used as default.The supported values are GET, POST, PUT, DELETE, PATCH, HEAD. 
* `mo_type`:(string) The type of the intersight object for which API request is to be made.The property is valid in case of Intersight API calls and the base url is automatically prepended based on the value. 
* `protocol`:(string) The accepted web protocol values are http and https. 
* `target_type`:(string) If the web API is to be executed in a remote device connected to theIntersight through device connector, 'Endpoint' is expected as the valuewhereas if the API is an Intersight API, 'Local' is expected as the value. 
* `url`:(string) The URL of the resource in the target to which the API request is made. 

### [workflow.XmlApi](#argument-reference)
This models a single XML API request that can be sent to any Cisco UCS devices
that support Cisco UCS XML API interface.
  
