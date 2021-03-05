package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexVmSnapshotInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexVmSnapshotInfoRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"label": {
				Description: "The name of the Virtual Machine and the time stamp of the snapshot.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mode": {
				Description: "Quiesce Mode for snapshot taken on Hyperflex cluster.\n* `NONE` - The snapshot quiesce mode is none.\n* `CRASH` - The snapshot quiesce mode is crash.\n* `VMTOOLS` - The snapshot quiesce mode is VMTOOLS.\n* `APP_CONSISTENT` - The snapshot quiesce mode is app consistent.",
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
			"replication_status": {
				Description: "Replication status of the least successful target cluster.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"snapshot_status": {
				Description: "Snapshot status of the source cluster.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"source_timestamp": {
				Description: "Timestamp the snapshot was created on the source cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster_id_snap_map": {
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
									Description: "ClusterId of the snapshot point.",
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
								"snapshot_point": {
									Description: "Snapshot point details for this cluster.",
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
											"cluster_entity_reference": {
												Description: "Cluster to which this snapshot belongs.",
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
											"snapshot_files": {
												Description: "Files this snapshot is comprised of.",
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
														"name_tracked_files": {
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
																	"ds_info": {
																		Description: "Datastore for which this file resides.",
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
																				"ds_backend_id": {
																					Description: "This datastore's backend unique id.",
																					Type:        schema.TypeString,
																					Optional:    true,
																					Computed:    true,
																				},
																				"ds_frontend_id": {
																					Description: "This datastore's frontend id.",
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
																			},
																		},
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																	"relative_file_path": {
																		Description: "Relative file path within the datastore.",
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
														"uuid_tracked_disks_map": {
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
																	"tracked_disk": {
																		Description: "Tracked Disk for a snapshot.",
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
																				"disk_files": {
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
																							"file_path": {
																								Description: "File path and datastore information.",
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
																										"ds_info": {
																											Description: "Datastore for which this file resides.",
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
																													"ds_backend_id": {
																														Description: "This datastore's backend unique id.",
																														Type:        schema.TypeString,
																														Optional:    true,
																														Computed:    true,
																													},
																													"ds_frontend_id": {
																														Description: "This datastore's frontend id.",
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
																												},
																											},
																										},
																										"object_type": {
																											Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																											Type:        schema.TypeString,
																											Optional:    true,
																											Computed:    true,
																										},
																										"relative_file_path": {
																											Description: "Relative file path within the datastore.",
																											Type:        schema.TypeString,
																											Optional:    true,
																											Computed:    true,
																										},
																									},
																								},
																							},
																							"file_type": {
																								Description: "File type for the tracked file.",
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
																						},
																					},
																					Computed: true,
																				},
																				"disk_type": {
																					Description: "Disk type for a vm virtual disk.\n* `NONE` - The disk type for this VM is None.\n* `NATIVE` - The disk type for this VM is Native.\n* `NONNATIVE` - The disk type for this VM is Non-Native.",
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
																			},
																		},
																	},
																	"uuid": {
																		Description: "Disk unique id for a snapshot.",
																		Type:        schema.TypeString,
																		Optional:    true,
																		Computed:    true,
																	},
																},
															},
															Computed: true,
														},
													},
												},
											},
											"snapshot_point_entity_reference": {
												Description: "EntityReference for this snapshot point.",
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
											"vm_runtime_info": {
												Description: "Virtual machine runtime information.",
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
										},
									},
								},
							},
						},
						Computed: true,
					},
					"error_stack": {
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
					"label": {
						Description: "The name of the Virtual Machine and the time stamp of the snapshot.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"mode": {
						Description: "Quiesce Mode for snapshot taken on Hyperflex cluster.\n* `NONE` - The snapshot quiesce mode is none.\n* `CRASH` - The snapshot quiesce mode is crash.\n* `VMTOOLS` - The snapshot quiesce mode is VMTOOLS.\n* `APP_CONSISTENT` - The snapshot quiesce mode is app consistent.",
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
					"parent_snapshot": {
						Description: "Entity reference to the parent snapshot of this VmSnapshotInfo.",
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
					"replication_status": {
						Description: "Replication status of the least successful target cluster.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"snapshot_status": {
						Description: "Snapshot status of the source cluster.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"source_timestamp": {
						Description: "Timestamp the snapshot was created on the source cluster.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
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
					"vm_backup_info": {
						Description: "A reference to a hyperflexVmBackupInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"vm_entity_reference": {
						Description: "Entity reference to the virtual machine this snapshot info is attached to.",
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
					"vm_snapshot_entity_reference": {
						Description: "Entity reference to the snapshot this snapshot info is attached to.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexVmSnapshotInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexVmSnapshotInfo{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}
	if v, ok := d.GetOk("mode"); ok {
		x := (v.(string))
		o.SetMode(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("replication_status"); ok {
		x := (v.(string))
		o.SetReplicationStatus(x)
	}
	if v, ok := d.GetOk("snapshot_status"); ok {
		x := (v.(string))
		o.SetSnapshotStatus(x)
	}
	if v, ok := d.GetOk("source_timestamp"); ok {
		x := int64(v.(int))
		o.SetSourceTimestamp(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexVmSnapshotInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmSnapshotInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexVmSnapshotInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexVmSnapshotInfoList.GetCount()
	var i int32
	var hyperflexVmSnapshotInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmSnapshotInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexVmSnapshotInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexVmSnapshotInfoList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexVmSnapshotInfo data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_id_snap_map"] = flattenListHyperflexMapClusterIdToStSnapshotPoint(s.GetClusterIdSnapMap(), d)

				temp["error_stack"] = flattenMapHyperflexErrorStack(s.GetErrorStack(), d)
				temp["label"] = (s.GetLabel())
				temp["mode"] = (s.GetMode())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["parent_snapshot"] = flattenMapHyperflexEntityReference(s.GetParentSnapshot(), d)
				temp["replication_status"] = (s.GetReplicationStatus())
				temp["snapshot_status"] = (s.GetSnapshotStatus())
				temp["source_timestamp"] = (s.GetSourceTimestamp())

				temp["src_cluster"] = flattenMapHyperflexClusterRelationship(s.GetSrcCluster(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["tgt_cluster"] = flattenMapHyperflexClusterRelationship(s.GetTgtCluster(), d)

				temp["vm_backup_info"] = flattenMapHyperflexVmBackupInfoRelationship(s.GetVmBackupInfo(), d)

				temp["vm_entity_reference"] = flattenMapHyperflexEntityReference(s.GetVmEntityReference(), d)

				temp["vm_snapshot_entity_reference"] = flattenMapHyperflexEntityReference(s.GetVmSnapshotEntityReference(), d)
				hyperflexVmSnapshotInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexVmSnapshotInfoResults))
	if err := d.Set("results", hyperflexVmSnapshotInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexVmSnapshotInfoResults[0]["moid"].(string))
	return de
}
