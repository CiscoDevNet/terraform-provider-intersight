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

func dataSourceHyperflexDataProtectionPeer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexDataProtectionPeerRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
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
								},
							},
						},
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
								},
							},
						},
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
								},
							},
						},
					},
					"shared_scope": {
						Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
						Type:        schema.TypeString,
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
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
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
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexDataProtectionPeer object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexDataProtectionPeerList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexDataProtectionPeer: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexDataProtectionPeer: %s", responseErr.Error())
	}
	count := countResponse.HyperflexDataProtectionPeerList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HyperflexDataProtectionPeer data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hyperflexDataProtectionPeerResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexDataProtectionPeerList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexDataProtectionPeer: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexDataProtectionPeer: %s", responseErr.Error())
		}
		results := resMo.HyperflexDataProtectionPeerList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["er"] = flattenMapHyperflexEntityReference(s.GetEr(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["peer_info"] = flattenMapHyperflexReplicationPeerInfo(s.GetPeerInfo(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["src_cluster"] = flattenMapHyperflexClusterRelationship(s.GetSrcCluster(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
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
