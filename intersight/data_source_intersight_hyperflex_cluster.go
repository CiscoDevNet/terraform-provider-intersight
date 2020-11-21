package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHyperflexCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHyperflexClusterRead,
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
				Description: "The deployment type of the HyperFlex cluster.\nThe cluster can have one of the following configurations:\n1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.\n2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.\n3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.\nIf the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,\nthe deployment type is set as 'NA' (not available).\n* `NA` - The deployment type of the HyperFlex cluster is not available.\n* `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site.\n* `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites.\n* `Edge` - The deployment type of a HyperFlex cluster consisting of 2-4 standalone nodes.",
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
				Description: "The type of hypervisor running on this cluster.\n* `ESXi` - A Vmware ESXi hypervisor of any version.\n* `HXAP` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
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
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"boottime": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"cluster_access_policy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"compression_savings": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"data_replication_compliance": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"data_replication_factor": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"deduplication_savings": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"downtime": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"free_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"healing_info": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
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
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"in_progress": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"messages": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_iterator": {
										Type: schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										}, Optional: true,
										Computed: true,
									},
									"messages_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"percent_complete": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"resiliency_details": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
							Computed: true,
						},
						"resiliency_details_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"resiliency_info": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
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
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"messages": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_iterator": {
										Type: schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										}, Optional: true,
										Computed: true,
									},
									"messages_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"node_failures_tolerable": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"ssd_failures_tolerable": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"space_status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"total_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"total_savings": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"uptime": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"used_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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

func dataSourceHyperflexClusterRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
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
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.HyperflexApi.GetHyperflexClusterList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.HyperflexClusterList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to HyperflexCluster: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexCluster{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}

			if err := d.Set("alarm", flattenListHyperflexAlarmRelationship(s.GetAlarm(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Alarm: %+v", err)
			}

			if err := d.Set("alarm_summary", flattenMapHyperflexAlarmSummary(s.GetAlarmSummary(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property AlarmSummary: %+v", err)
			}
			if err := d.Set("capacity_runway", (s.GetCapacityRunway())); err != nil {
				return fmt.Errorf("error occurred while setting property CapacityRunway: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("cluster_name", (s.GetClusterName())); err != nil {
				return fmt.Errorf("error occurred while setting property ClusterName: %+v", err)
			}
			if err := d.Set("cluster_type", (s.GetClusterType())); err != nil {
				return fmt.Errorf("error occurred while setting property ClusterType: %+v", err)
			}
			if err := d.Set("cluster_uuid", (s.GetClusterUuid())); err != nil {
				return fmt.Errorf("error occurred while setting property ClusterUuid: %+v", err)
			}
			if err := d.Set("compute_node_count", (s.GetComputeNodeCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ComputeNodeCount: %+v", err)
			}
			if err := d.Set("converged_node_count", (s.GetConvergedNodeCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ConvergedNodeCount: %+v", err)
			}
			if err := d.Set("deployment_type", (s.GetDeploymentType())); err != nil {
				return fmt.Errorf("error occurred while setting property DeploymentType: %+v", err)
			}
			if err := d.Set("device_id", (s.GetDeviceId())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceId: %+v", err)
			}
			if err := d.Set("flt_aggr", (s.GetFltAggr())); err != nil {
				return fmt.Errorf("error occurred while setting property FltAggr: %+v", err)
			}

			if err := d.Set("health", flattenMapHyperflexHealthRelationship(s.GetHealth(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Health: %+v", err)
			}
			if err := d.Set("hx_version", (s.GetHxVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property HxVersion: %+v", err)
			}
			if err := d.Set("hxdp_build_version", (s.GetHxdpBuildVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property HxdpBuildVersion: %+v", err)
			}
			if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
				return fmt.Errorf("error occurred while setting property HypervisorType: %+v", err)
			}
			if err := d.Set("hypervisor_version", (s.GetHypervisorVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property HypervisorVersion: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}

			if err := d.Set("nodes", flattenListHyperflexNodeRelationship(s.GetNodes(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Nodes: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}

			if err := d.Set("storage_containers", flattenListStorageHyperFlexStorageContainerRelationship(s.GetStorageContainers(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property StorageContainers: %+v", err)
			}

			if err := d.Set("summary", flattenMapHyperflexSummary(s.GetSummary(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Summary: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("utilization_percentage", (s.GetUtilizationPercentage())); err != nil {
				return fmt.Errorf("error occurred while setting property UtilizationPercentage: %+v", err)
			}
			if err := d.Set("utilization_trend_percentage", (s.GetUtilizationTrendPercentage())); err != nil {
				return fmt.Errorf("error occurred while setting property UtilizationTrendPercentage: %+v", err)
			}
			if err := d.Set("vm_count", (s.GetVmCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VmCount: %+v", err)
			}

			if err := d.Set("volumes", flattenListStorageHyperFlexVolumeRelationship(s.GetVolumes(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Volumes: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
