package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"alarm": {
				Description: "An array of relationships to hyperflexAlarm resources.",
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
			"alarm_summary": {
				Description: "The summary of alarm counts based on the severity types (Critical or Warning).",
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
						"critical": {
							Description: "The count of alarms that have severity type Critical.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"warning": {
							Description: "The count of alarms that have severity type Warning.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
			"capacity_runway": {
				Description: "The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%.\nDefault value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_name": {
				Description: "The name of this HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_type": {
				Description: "The storage type of this cluster (All Flash or Hybrid).",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"cluster_uuid": {
				Description: "The unique identifier for this HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"compute_node_count": {
				Description: "The number of compute nodes that belong to this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"converged_node_count": {
				Description: "The number of converged nodes that belong to this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"deployment_type": {
				Description: "The deployment type of the HyperFlex cluster.\nThe cluster can have one of the following configurations:\n1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.\n2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.\n3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.\nIf the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,\nthe deployment type is set as 'NA' (not available).\n* `NA` - The deployment type of the HyperFlex cluster is not available.\n* `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site.\n* `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites.\n* `Edge` - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Description: "The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"flt_aggr": {
				Description: "The number of yellow (warning) and red (critical) alarms stored as an aggregate.\nThe first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"health": {
				Description: "A reference to a hyperflexHealth resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"hx_version": {
				Description: "The HyperFlex Data Platform version of this cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hxdp_build_version": {
				Description: "The version and build number of the HyperFlex Data Platform for this cluster.\nAfter a cluster upgrade, this version string will be updated on the next inventory cycle to reflect\nthe newly installed version.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hypervisor_type": {
				Description: "The type of hypervisor running on this cluster.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hypervisor_version": {
				Description: "The version of hypervisor running on this cluster.",
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
			"nodes": {
				Description: "An array of relationships to hyperflexNode resources.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"storage_containers": {
				Description: "An array of relationships to storageHyperFlexStorageContainer resources.",
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
			"summary": {
				Description: "The summary of HyperFlex cluster health, storage, and number of nodes.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_nodes": {
							Description: "The number of nodes currently participating in the storage cluster.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"address": {
							Description: "The data IP address of the HyperFlex cluster.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"boottime": {
							Description: "The time taken during last cluster startup in seconds.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"cluster_access_policy": {
							Description: "The cluster access policy for the HyperFlex cluster. An access policy of 'STRICT' means that the cluster becomes readonly once any fragment of data is reduced to one copy. 'LENIENT' means that the cluster stays in read-write mode even if any fragment of data is reduced to one copy.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"compression_savings": {
							Description: "The percentage of storage space saved using data compression.",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
						"data_replication_compliance": {
							Description: "The compliance with the data replication factor set for the HyperFlex cluster.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"data_replication_factor": {
							Description: "The number of data copies retained by the HyperFlex cluster.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"deduplication_savings": {
							Description: "The percentage of storage space saved using data deduplication.",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
						"downtime": {
							Description: "The amount of time the HyperFlex cluster has been offline.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"free_capacity": {
							Description: "The amount of storage capacity currently not in use, represented in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"healing_info": {
							Description: "The information about the HyperFlex cluster auto-healing process.",
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"estimated_completion_time_in_seconds": {
										Description: "The estimated time in seconds it will take to complete the auto-healing process.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"in_progress": {
										Description: "The status of the cluster's auto-healing process. If 'true', auto-healing is in progress for the cluster.",
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
									},
									"messages": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_iterator": {
										Description: "The current message describing the auto-healing process of the cluster.",
										Type:        schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										}, Optional: true,
										Computed: true,
									},
									"messages_size": {
										Description: "The number of elements in the messages collection.",
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
									"percent_complete": {
										Description: "The progress of the auto-healing process as a percentage.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Description: "The name of the HyperFlex cluster.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"resiliency_details": {
							Description: "The details about the resiliency health of the cluster. Includes information about the cluster healing status and the storage cluster health.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
							Computed: true,
						},
						"resiliency_details_size": {
							Description: "The number of elements in the resiliency details property.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"resiliency_info": {
							Description: "The information about the storage cluster resiliency. The resiliency info has details about the health status and number of device failures tolerable.",
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
									"hdd_failures_tolerable": {
										Description: "The number of persistent storage device failures tolerable before the storage cluster becomes offline.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"messages": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_iterator": {
										Description: "The current message describing the auto-healing process of the cluster.",
										Type:        schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										}, Optional: true,
										Computed: true,
									},
									"messages_size": {
										Description: "The number of elements in the messages collection.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"node_failures_tolerable": {
										Description: "The number of node failures tolerable before the storage cluster becomes offline.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"ssd_failures_tolerable": {
										Description: "The number of caching device failures tolerable before the storage cluster becomes offline.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "The resiliency state of the cluster. The resiliency status is 'HEALTHY' if there are no failures and the storage cluster is fully operational. The resiliency status is 'WARNING' when the cluster has experienced failures that may adversely affect the cluster. It is 'UNKNOWN' if the cluster is offline or if the state cannot be determined.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"space_status": {
							Description: "The space utilization status of the HyperFlex cluster.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"state": {
							Description: "The operational state of the HyperFlex cluster.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"total_capacity": {
							Description: "The total amount of storage capacity available for the HyperFlex cluster, represented in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"total_savings": {
							Description: "The percentage of storage space saved in total.",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
						"uptime": {
							Description: "The amount of time the HyperFlex cluster has been running since last startup.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"used_capacity": {
							Description: "The amount of storage capacity in use, represented in bytes.",
							Type:        schema.TypeInt,
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
			"utilization_percentage": {
				Description: "The storage utilization percentage is computed based on total capacity and current capacity utilization.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"utilization_trend_percentage": {
				Description: "The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"vm_count": {
				Description: "The number of virtual machines present on this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"volumes": {
				Description: "An array of relationships to storageHyperFlexVolume resources.",
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
		},
	}
}

func dataSourceHyperflexClusterRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexCluster{}
	if v, ok := d.GetOk("capacity_runway"); ok {
		x := int64(v.(int))
		o.SetCapacityRunway(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cluster_name"); ok {
		x := (v.(string))
		o.SetClusterName(x)
	}
	if v, ok := d.GetOk("cluster_type"); ok {
		x := int64(v.(int))
		o.SetClusterType(x)
	}
	if v, ok := d.GetOk("cluster_uuid"); ok {
		x := (v.(string))
		o.SetClusterUuid(x)
	}
	if v, ok := d.GetOk("compute_node_count"); ok {
		x := int64(v.(int))
		o.SetComputeNodeCount(x)
	}
	if v, ok := d.GetOk("converged_node_count"); ok {
		x := int64(v.(int))
		o.SetConvergedNodeCount(x)
	}
	if v, ok := d.GetOk("deployment_type"); ok {
		x := (v.(string))
		o.SetDeploymentType(x)
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.SetDeviceId(x)
	}
	if v, ok := d.GetOk("flt_aggr"); ok {
		x := int64(v.(int))
		o.SetFltAggr(x)
	}
	if v, ok := d.GetOk("hx_version"); ok {
		x := (v.(string))
		o.SetHxVersion(x)
	}
	if v, ok := d.GetOk("hxdp_build_version"); ok {
		x := (v.(string))
		o.SetHxdpBuildVersion(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("hypervisor_version"); ok {
		x := (v.(string))
		o.SetHypervisorVersion(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("utilization_percentage"); ok {
		x := v.(float32)
		o.SetUtilizationPercentage(x)
	}
	if v, ok := d.GetOk("utilization_trend_percentage"); ok {
		x := v.(float32)
		o.SetUtilizationTrendPercentage(x)
	}
	if v, ok := d.GetOk("vm_count"); ok {
		x := int64(v.(int))
		o.SetVmCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexCluster object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexCluster: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexCluster list: %s", err.Error())
	}
	var s = &models.HyperflexClusterList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexCluster list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for HyperflexCluster data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for HyperflexCluster data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexCluster{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("alarm", flattenListHyperflexAlarmRelationship(s.GetAlarm(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Alarm: %s", err.Error())
			}

			if err := d.Set("alarm_summary", flattenMapHyperflexAlarmSummary(s.GetAlarmSummary(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AlarmSummary: %s", err.Error())
			}
			if err := d.Set("capacity_runway", (s.GetCapacityRunway())); err != nil {
				return diag.Errorf("error occurred while setting property CapacityRunway: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("cluster_name", (s.GetClusterName())); err != nil {
				return diag.Errorf("error occurred while setting property ClusterName: %s", err.Error())
			}
			if err := d.Set("cluster_type", (s.GetClusterType())); err != nil {
				return diag.Errorf("error occurred while setting property ClusterType: %s", err.Error())
			}
			if err := d.Set("cluster_uuid", (s.GetClusterUuid())); err != nil {
				return diag.Errorf("error occurred while setting property ClusterUuid: %s", err.Error())
			}
			if err := d.Set("compute_node_count", (s.GetComputeNodeCount())); err != nil {
				return diag.Errorf("error occurred while setting property ComputeNodeCount: %s", err.Error())
			}
			if err := d.Set("converged_node_count", (s.GetConvergedNodeCount())); err != nil {
				return diag.Errorf("error occurred while setting property ConvergedNodeCount: %s", err.Error())
			}
			if err := d.Set("deployment_type", (s.GetDeploymentType())); err != nil {
				return diag.Errorf("error occurred while setting property DeploymentType: %s", err.Error())
			}
			if err := d.Set("device_id", (s.GetDeviceId())); err != nil {
				return diag.Errorf("error occurred while setting property DeviceId: %s", err.Error())
			}
			if err := d.Set("flt_aggr", (s.GetFltAggr())); err != nil {
				return diag.Errorf("error occurred while setting property FltAggr: %s", err.Error())
			}

			if err := d.Set("health", flattenMapHyperflexHealthRelationship(s.GetHealth(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Health: %s", err.Error())
			}
			if err := d.Set("hx_version", (s.GetHxVersion())); err != nil {
				return diag.Errorf("error occurred while setting property HxVersion: %s", err.Error())
			}
			if err := d.Set("hxdp_build_version", (s.GetHxdpBuildVersion())); err != nil {
				return diag.Errorf("error occurred while setting property HxdpBuildVersion: %s", err.Error())
			}
			if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorType: %s", err.Error())
			}
			if err := d.Set("hypervisor_version", (s.GetHypervisorVersion())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorVersion: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}

			if err := d.Set("nodes", flattenListHyperflexNodeRelationship(s.GetNodes(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Nodes: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}

			if err := d.Set("storage_containers", flattenListStorageHyperFlexStorageContainerRelationship(s.GetStorageContainers(), d)); err != nil {
				return diag.Errorf("error occurred while setting property StorageContainers: %s", err.Error())
			}

			if err := d.Set("summary", flattenMapHyperflexSummary(s.GetSummary(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Summary: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("utilization_percentage", (s.GetUtilizationPercentage())); err != nil {
				return diag.Errorf("error occurred while setting property UtilizationPercentage: %s", err.Error())
			}
			if err := d.Set("utilization_trend_percentage", (s.GetUtilizationTrendPercentage())); err != nil {
				return diag.Errorf("error occurred while setting property UtilizationTrendPercentage: %s", err.Error())
			}
			if err := d.Set("vm_count", (s.GetVmCount())); err != nil {
				return diag.Errorf("error occurred while setting property VmCount: %s", err.Error())
			}

			if err := d.Set("volumes", flattenListStorageHyperFlexVolumeRelationship(s.GetVolumes(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Volumes: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
