---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_backup_cluster"
description: |-
  BackupCluster object associated with Hyperflex cluster describing the backup related information.
---

# Data Source: intersight_hyperflex_backup_cluster
BackupCluster object associated with Hyperflex cluster describing the backup related information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_backup_cluster.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `backup_data_store`:(string) Defines the backup source cluster and its references. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 