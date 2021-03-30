package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexDataProtectionPeer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexDataProtectionPeerRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
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
					"er": {
						Description: "Entity reference to the cluster pair this DataProtectionPeer is associated with.",
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
					"peer_info": {
						Description: "Peer Cluster status details.",
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
								"datastores": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"ads": {
												Description: "First datastore in the HyperFlex datastore pair.",
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
														"cluster_er": {
															Description: "Entity reference to the cluster this datastore is attached to.",
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
														"datastore_er": {
															Description: "Entity reference to the datastore this object represents.",
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
													},
												},
											},
											"backup_only": {
												Description: "Boolean representing if this is a backup only pair.",
												Type:        schema.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"bds": {
												Description: "Second datastore in the HyperFlex datastore pair.",
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
														"cluster_er": {
															Description: "Entity reference to the cluster this datastore is attached to.",
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
														"datastore_er": {
															Description: "Entity reference to the datastore this object represents.",
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
													},
												},
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
											"quiesce": {
												Description: "Boolean representing if this datastore pairing has quiesce snapshots enabled.",
												Type:        schema.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"replication_interval_in_minutes": {
												Description: "The replication interval for this pair in minutes.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"sourceds": {
												Description: "HyperFlex source datastore for this pair.",
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
														"cluster_er": {
															Description: "Entity reference to the cluster this datastore is attached to.",
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
														"datastore_er": {
															Description: "Entity reference to the datastore this object represents.",
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
													},
												},
											},
											"storage_only": {
												Description: "HyperFlex datastore pair is used for storage only.",
												Type:        schema.TypeBool,
												Optional:    true,
												Computed:    true,
											},
										},
									},
									Computed: true,
								},
								"dcip": {
									Description: "Data Cluster IP for the replication peer.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"er": {
									Description: "Entity Reference to the replication peer info object.",
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
								"mcip": {
									Description: "Management Cluster IP for the replication peer.",
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
								"ports": {
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
											"i16": {
												Description: "Integer describing port type to port number map.",
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
											"string": {
												Description: "String describing port type to port number map.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
									Computed: true,
								},
								"repl_cip": {
									Description: "Replication Cluster IP for the replication peer.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"status": {
									Description: "Peer Cluster Status for the replication peer.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"status_details": {
									Description: "Peer Cluster Status Details for the replication peer.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexDataProtectionPeerRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexDataProtectionPeer{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexDataProtectionPeer object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexDataProtectionPeerList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexDataProtectionPeer: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexDataProtectionPeerList.GetCount()
	var i int32
	var hyperflexDataProtectionPeerResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexDataProtectionPeerList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexDataProtectionPeer: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexDataProtectionPeerList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexDataProtectionPeer data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["er"] = flattenMapHyperflexEntityReference(s.GetEr(), d)
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["peer_info"] = flattenMapHyperflexReplicationPeerInfo(s.GetPeerInfo(), d)

				temp["src_cluster"] = flattenMapHyperflexClusterRelationship(s.GetSrcCluster(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				hyperflexDataProtectionPeerResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexDataProtectionPeerResults))
	if err := d.Set("results", hyperflexDataProtectionPeerResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexDataProtectionPeerResults[0]["moid"].(string))
	return de
}
