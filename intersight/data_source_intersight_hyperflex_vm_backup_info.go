package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexVmBackupInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexVmBackupInfoRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_status": {
				Description: "Description of the backup status of this VmBackupInfo.\n* `InitializingProtection` - Protection has started, but not completed.\n* `Protected` - Protection has completed successfully.\n* `ExceedsInterval` - Protection has not completed successfully in over two times the backup interval.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"protection_status": {
				Description: "Description of the protection status of this VmBackupInfo.\n* `PREPARE_FAILOVER_STARTED` - The protection status is prepare failover started.\n* `PREPARE_FAILOVER_FAILED` - The protection status is prepare failover failed.\n* `PREPARE_FAILOVER_COMPLETED` - The protection status is prepaire failover completed.\n* `FAILOVER_STARTED` - The protection status is failover started.\n* `FAILOVER_FAILED` - The protection status is failover failed.\n* `FAILOVER_COMPLETED` - The protection status is failover completed.\n* `PREPARE_REVERSEPROTECT_STARTED` - The protection status is prepare reverse protect started.\n* `PREPARE_REVERSEPROTECT_FAILED` - The protection status is prepare reverse protect failed.\n* `PREPARE_REVERSEPROTECT_COMPLETED` - The protection status is prepair reverse protect completed.\n* `REVERSEPROTECT_STARTED` - The protection status is reverse protect started.\n* `REVERSEPROTECT_FAILED` - The protection status is reverse protect failed.\n* `ACTIVE` - The protection status is active.\n* `CREATION_IN_PROGRESS` - The protection status is failover in progress.\n* `CREATION_FAILED` - The protection status is creation failed.\n* `LOCAL_RESTORE_STARTED` - The protection status is local restore started.\n* `LOCAL_RESTORE_FAILED` - The protection status is local restore failed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account_moid": {
					Description: "The Account ID for this managed object.",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"ancestors": {
						Description: "An array of relationships to moBaseMo resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"backup_status": {
						Description: "Description of the backup status of this VmBackupInfo.\n* `InitializingProtection` - Protection has started, but not completed.\n* `Protected` - Protection has completed successfully.\n* `ExceedsInterval` - Protection has not completed successfully in over two times the backup interval.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster_entity_reference": {
						Description: "Entity reference to the cluster this VmBackupInfo is associated with.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"confignum": {
									Description: "Configuration number for this reference.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"id": {
									Description: "Uuid of the entity for this reference.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"idtype": {
									Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the entity for this entity reference.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"type": {
									Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"cluster_id_protection_info_map": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"cluster_id": {
									Description: "The Cluster Id we are using to map to ProtectionInfo.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"protection_info": {
									Description: "ProtectionInfo that is being stored for this Virtual Machine.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"vm_current_protection_info": {
												Description: "Current snapshot protection information.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Computed:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"replication_status": {
															Description: "Replication status information for this particular snapshot.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Computed:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"bytes_replicated": {
																		Description: "Number of bytes currently replicated.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"end_time": {
																		Description: "Replication end time for this snapshot.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"error": {
																		Description: "Error information in case of failure.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"message": {
																					Description: "The error message string for this error stack.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"message_id": {
																					Description: "The error message ID for this error stack.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"pack_entity_reference": {
																		Description: "EntityReference for the Replication Pack.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"confignum": {
																					Description: "Configuration number for this reference.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"id": {
																					Description: "Uuid of the entity for this reference.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"idtype": {
																					Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"name": {
																					Description: "Name of the entity for this entity reference.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"type": {
																					Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"pct_complete": {
																		Description: "Completion percentage for the replication job.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"rpo_status": {
																		Description: "Status for timeliness of replication job in relation to the schedule interval.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"actual": {
																					Description: "Actual end time for the snapshot.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"expected": {
																					Description: "Expected end time for the snapshot.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"rpo_exceeded": {
																					Description: "A flag to determine if snapshot delivery is delayed.",
																					Type:        schema.TypeBool,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"start_time": {
																		Description: "Replication start time for this snapshot.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"status": {
																		Description: "Current replication state for a particular snapshot.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																},
															},
														},
														"site": {
															Description: "Cluster site for this snapshot.\n* `PRIMARY` - The cluster site for this backup is primary.\n* `SECONDARY` - The cluster site for this backup is secondary.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"snapshot_status": {
															Description: "Status for this Snapshot Point.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Computed:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"description": {
																		Description: "Description of this Snapshot Point.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"error": {
																		Description: "Error information in case of failure.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"message": {
																					Description: "The error message string for this error stack.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"message_id": {
																					Description: "The error message ID for this error stack.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"pct_complete": {
																		Description: "Completion percentage for this snapshot.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"status": {
																		Description: "Current snapshot state for this snapshot.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"timestamp": {
																		Description: "Timestamp at which the Snapshot is taken.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"used_space": {
																		Description: "Space Used by this Snapshot Point.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																},
															},
														},
														"vm_snapshot_entity_reference": {
															Description: "EntityReference of this VmSnapshot.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Computed:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"confignum": {
																		Description: "Configuration number for this reference.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"id": {
																		Description: "Uuid of the entity for this reference.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"idtype": {
																		Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"name": {
																		Description: "Name of the entity for this entity reference.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"type": {
																		Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																},
															},
														},
													},
												},
											},
											"vm_last_successful_protection_info": {
												Description: "Last successful snapshot protection information.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Computed:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"replication_status": {
															Description: "Replication status information for this particular snapshot.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Computed:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"bytes_replicated": {
																		Description: "Number of bytes currently replicated.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"end_time": {
																		Description: "Replication end time for this snapshot.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"error": {
																		Description: "Error information in case of failure.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"message": {
																					Description: "The error message string for this error stack.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"message_id": {
																					Description: "The error message ID for this error stack.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"pack_entity_reference": {
																		Description: "EntityReference for the Replication Pack.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"confignum": {
																					Description: "Configuration number for this reference.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"id": {
																					Description: "Uuid of the entity for this reference.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"idtype": {
																					Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"name": {
																					Description: "Name of the entity for this entity reference.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"type": {
																					Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"pct_complete": {
																		Description: "Completion percentage for the replication job.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"rpo_status": {
																		Description: "Status for timeliness of replication job in relation to the schedule interval.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"actual": {
																					Description: "Actual end time for the snapshot.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"expected": {
																					Description: "Expected end time for the snapshot.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"rpo_exceeded": {
																					Description: "A flag to determine if snapshot delivery is delayed.",
																					Type:        schema.TypeBool,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"start_time": {
																		Description: "Replication start time for this snapshot.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"status": {
																		Description: "Current replication state for a particular snapshot.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																},
															},
														},
														"site": {
															Description: "Cluster site for this snapshot.\n* `PRIMARY` - The cluster site for this backup is primary.\n* `SECONDARY` - The cluster site for this backup is secondary.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"snapshot_status": {
															Description: "Status for this Snapshot Point.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Computed:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"description": {
																		Description: "Description of this Snapshot Point.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"error": {
																		Description: "Error information in case of failure.",
																		Type:        schema.TypeList,
																		MaxItems:    1,
																		Optional:    true,
																		Computed:    true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"message": {
																					Description: "The error message string for this error stack.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"message_id": {
																					Description: "The error message ID for this error stack.",
																					Type:        schema.TypeInt,
																					Optional:    true,
																					Computed:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"pct_complete": {
																		Description: "Completion percentage for this snapshot.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"status": {
																		Description: "Current snapshot state for this snapshot.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"timestamp": {
																		Description: "Timestamp at which the Snapshot is taken.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"used_space": {
																		Description: "Space Used by this Snapshot Point.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																},
															},
														},
														"vm_snapshot_entity_reference": {
															Description: "EntityReference of this VmSnapshot.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Computed:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"confignum": {
																		Description: "Configuration number for this reference.",
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Computed:    true,
																	},
																	"id": {
																		Description: "Uuid of the entity for this reference.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"idtype": {
																		Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"name": {
																		Description: "Name of the entity for this entity reference.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"type": {
																		Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																},
															},
														},
													},
												},
											},
											"vm_space_usage": {
												Description: "Protection space usage for this snapshot.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Computed:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"space_usage": {
															Description: "Space usage of the VM from StDataServiceManager.",
															Type:        schema.TypeInt,
															Optional:    true,
															Computed:    true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
						Computed: true,
					},
					"create_time": {
						Description: "The time when this managed object was created.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"domain_group_moid": {
						Description: "The DomainGroup ID for this managed object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"error": {
						Description: "List of errors associated with this VmBackupInfo.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"message": {
									Description: "The error message string for this error stack.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"message_id": {
									Description: "The error message ID for this error stack.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"mod_time": {
						Description: "The time when this managed object was last modified.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"owners": {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"parent": {
						Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"permission_resources": {
						Description: "An array of relationships to moBaseMo resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"protection_status": {
						Description: "Description of the protection status of this VmBackupInfo.\n* `PREPARE_FAILOVER_STARTED` - The protection status is prepare failover started.\n* `PREPARE_FAILOVER_FAILED` - The protection status is prepare failover failed.\n* `PREPARE_FAILOVER_COMPLETED` - The protection status is prepaire failover completed.\n* `FAILOVER_STARTED` - The protection status is failover started.\n* `FAILOVER_FAILED` - The protection status is failover failed.\n* `FAILOVER_COMPLETED` - The protection status is failover completed.\n* `PREPARE_REVERSEPROTECT_STARTED` - The protection status is prepare reverse protect started.\n* `PREPARE_REVERSEPROTECT_FAILED` - The protection status is prepare reverse protect failed.\n* `PREPARE_REVERSEPROTECT_COMPLETED` - The protection status is prepair reverse protect completed.\n* `REVERSEPROTECT_STARTED` - The protection status is reverse protect started.\n* `REVERSEPROTECT_FAILED` - The protection status is reverse protect failed.\n* `ACTIVE` - The protection status is active.\n* `CREATION_IN_PROGRESS` - The protection status is failover in progress.\n* `CREATION_FAILED` - The protection status is creation failed.\n* `LOCAL_RESTORE_STARTED` - The protection status is local restore started.\n* `LOCAL_RESTORE_FAILED` - The protection status is local restore failed.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"schedule": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"schedule": {
									Description: "Replication schedule information.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"backup_interval": {
												Description: "Time interval between two copies in minutes.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"target_cluster_entity_reference": {
									Description: "EntityReference of target cluster.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"confignum": {
												Description: "Configuration number for this reference.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"id": {
												Description: "Uuid of the entity for this reference.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"idtype": {
												Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"name": {
												Description: "Name of the entity for this entity reference.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"type": {
												Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
							},
						},
						Computed: true,
					},
					"shared_scope": {
						Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"src_backup_cluster": {
						Description: "A reference to a hyperflexBackupCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"src_cluster": {
						Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"tgt_cluster": {
						Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"version_context": {
						Description: "The versioning info for this managed object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"interested_mos": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"moid": {
												Description: "The Moid of the referenced REST resource.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the remote type referred by this relationship.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"selector": {
												Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
									Computed: true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"ref_mo": {
									Description: "A reference to the original Managed Object.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"moid": {
												Description: "The Moid of the referenced REST resource.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the remote type referred by this relationship.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"selector": {
												Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"timestamp": {
									Description: "The time this versioned Managed Object was created.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"nr_version": {
									Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"version_type": {
									Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"vm_entity_reference": {
						Description: "Reference to the virtual machine this VmBackupInfo is associated with.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"confignum": {
									Description: "Configuration number for this reference.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"id": {
									Description: "Uuid of the entity for this reference.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"idtype": {
									Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the entity for this entity reference.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"type": {
									Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"vm_info": {
						Description: "More detailed information about this VmBackupInfo including runtime info.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"run_time_info": {
									Description: "Virtual machine runtime details.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"bios_uuid": {
												Description: "BiosUuid of the Protected Virtual Machine.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"connection_state": {
												Description: "Connection state of the VM.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"cpu_usage": {
												Description: "CPU Usage of Virtual Machine.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"folder": {
												Description: "Folder which VM belongs to.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"guest_family": {
												Description: "Guest operating system family, if known.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"guest_full_name": {
												Description: "Guest operating system full name, if known.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"guest_id": {
												Description: "Guest operating system identifier (short name), if known.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"guest_state": {
												Description: "VMware guest reset, powercycle, or boot action states.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"host_name": {
												Description: "Hostname of Virtual Machine.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"instance_uuid": {
												Description: "InstanceUuid of the Protected Virtual Machine.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"memory_mb": {
												Description: "CPU Memory in MB of Virtual Machine.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"memory_usage": {
												Description: "Memory usage of Virtual Machine.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"moid": {
												Description: "Virtual Machine unique MOID.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"name": {
												Description: "Name of the Virtual Machine.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"networks": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Schema{
													Type: schema.TypeString}},
											"num_cpu": {
												Description: "Number of CPUs for the VM.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"power_state": {
												Description: "Power state of the Virtual Machine.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"provisioned_size": {
												Description: "Provisioned Size of Virtual Machine.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"resource_pool": {
												Description: "Resource pool which VM belongs to.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"used_size": {
												Description: "Used Size of Virtual Machine.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"nr_version": {
												Description: "Version of the Virtual Machine.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"vmx_path": {
												Description: "Vmx Path in VC datastore format.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"status_code": {
									Description: "Virtual Machine Status Code.\n* `VM_ACCESSIBLE` - This virtual machine is accessible.\n* `VM_INACCESSIBLE` - This virtual machine is not accessible.\n* `VM_NOT_SUPPORTED` - This virtual machine is not supported.\n* `NONE` - This virtual machine does not have a status code.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"uuid": {
									Description: "Virtual machine unique UUID.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexVmBackupInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexVmBackupInfo{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("backup_status"); ok {
		x := (v.(string))
		o.SetBackupStatus(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("protection_status"); ok {
		x := (v.(string))
		o.SetProtectionStatus(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexVmBackupInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmBackupInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexVmBackupInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexVmBackupInfo: %s", responseErr.Error())
	}
	count := countResponse.HyperflexVmBackupInfoList.GetCount()
	var i int32
	var hyperflexVmBackupInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmBackupInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexVmBackupInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexVmBackupInfo: %s", responseErr.Error())
		}
		results := resMo.HyperflexVmBackupInfoList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexVmBackupInfo data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["backup_status"] = (s.GetBackupStatus())
				temp["class_id"] = (s.GetClassId())

				temp["cluster_entity_reference"] = flattenMapHyperflexEntityReference(s.GetClusterEntityReference(), d)

				temp["cluster_id_protection_info_map"] = flattenListHyperflexMapClusterIdToProtectionInfo(s.GetClusterIdProtectionInfoMap(), d)

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["error"] = flattenMapHyperflexErrorStack(s.GetError(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["protection_status"] = (s.GetProtectionStatus())

				temp["schedule"] = flattenListHyperflexReplicationClusterReferenceToSchedule(s.GetSchedule(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["src_backup_cluster"] = flattenMapHyperflexBackupClusterRelationship(s.GetSrcBackupCluster(), d)

				temp["src_cluster"] = flattenMapHyperflexClusterRelationship(s.GetSrcCluster(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["tgt_cluster"] = flattenMapHyperflexClusterRelationship(s.GetTgtCluster(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)

				temp["vm_entity_reference"] = flattenMapHyperflexEntityReference(s.GetVmEntityReference(), d)

				temp["vm_info"] = flattenMapHyperflexVirtualMachine(s.GetVmInfo(), d)
				hyperflexVmBackupInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexVmBackupInfoResults))
	if err := d.Set("results", hyperflexVmBackupInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexVmBackupInfoResults[0]["moid"].(string))
	return de
}
