---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_metadata"
description: |-
  Workflow metadata is a collection of properties that are common across all the versions of a workflow.
---

# Data Source: intersight_workflow_workflow_metadata
Workflow metadata is a collection of properties that are common across all the versions of a workflow.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_workflow_metadata.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) The description for this workflow. 
* `label`:(string) A user friendly short name to identify the workflow metadata. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this workflow metadata. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
 