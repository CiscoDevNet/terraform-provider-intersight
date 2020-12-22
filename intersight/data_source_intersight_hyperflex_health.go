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

func dataSourceHyperflexHealth() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHealthRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"arbitration_service_state": {
				Description: "The status of the HyperFlex cluster's connection to the Intersight arbitration service. The arbitration service state is only applicable to 2-node edge clusters.\n* `NOT_AVAILABLE` - The cluster does not require a connection to the arbitration service.\n* `UNKNOWN` - The cluster's connection state to the arbitration service cannot be determined.\n* `ONLINE` - The cluster is connected to the arbitration service.\n* `OFFLINE` - The cluster is disconnected from the arbitration service.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster": {
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
			"data_replication_compliance": {
				Description: "The HyperFlex cluster's compliance to the configured replication factor. It indicates that the compliance has degraded if the number of copies of data is reduced.\n* `UNKNOWN` - The replication compliance of the HyperFlex cluster is not known.\n* `COMPLIANT` - The HyperFlex cluster is compliant with the replication policy. All data on the cluster is replicated according to the configured replication factor.\n* `NON_COMPLIANT` - The HyperFlex cluster is not compliant with the replication policy. Some data on the cluster is not replicated in accordance with the configured replication factor.",
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
			"resiliency_details": {
				Description: "The details of the HyperFlex cluster's resiliency status.",
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
						"data_replication_factor": {
							Description: "The number of copies of data replicated by the cluster.\n* `ONE_COPY` - The HyperFlex cluster does not replicate data.\n* `TWO_COPIES` - The HyperFlex cluster keeps 2 copies of data.\n* `THREE_COPIES` - The HyperFlex cluster keeps 3 copies of data.\n* `FOUR_COPIES` - The HyperFlex cluster keeps 4 copies of data.\n* `SIX_COPIES` - The HyperFlex cluster keeps 6 copies of data.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"hdd_failures_tolerable": {
							Description: "The number of persistent device disruptions the HyperFlex storage cluster can handle at this point in time.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"messages": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"node_failures_tolerable": {
							Description: "The number of node disruptions the HyperFlex storage cluster can handle at this point in time.",
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
						"policy_compliance": {
							Description: "The state of the storage cluster's compliance with the cluster access policy.\n* `UNKNOWN` - The HyperFlex cluster's compliance with the data replication policy could not be determined.\n* `COMPLIANT` - The HyperFlex cluster is compliant with the data replication policy and data is replicated to the configured replication factor.\n* `NON_COMPLIANT` - The HyperFlex cluster is not compliant with the data replication policy because some data is not replicated.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"resiliency_state": {
							Description: "The resiliency state of the storage platform. The resiliency state is the storage cluster's current ability to maintain data.\n* `UNKNOWN` - The resiliency status of the HyperFlex cluster cannot be determined, or the cluster is transitioning into ONLINE state.\n* `HEALTHY` - The HyperFlex cluster is healthy. The cluster is able to perform IO operations and data is available.\n* `WARNING` - The HyperFlex cluster or data availability is adversely affected. This can happen if there are node or storage device failures beyond the tolerable failure threshold.\n* `OFFLINE` - The HyperFlex cluster is offline and not performing IO operations.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ssd_failures_tolerable": {
							Description: "The number of cache device disruptions the HyperFlex storage cluster can handle at this point in time.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"state": {
				Description: "The operational status of the HyperFlex cluster.\n* `UNKNOWN` - The operational status of the cluster cannot be determined.\n* `ONLINE` - The HyperFlex cluster is online and is performing IO operations.\n* `OFFLINE` - The HyperFlex cluster is offline and is not ready to perform IO operations.\n* `ENOSPACE` - The HyperFlex cluster is out of available storage capacity and cannot perform write transactions.\n* `READONLY` - The HyperFlex cluster is not accepting write transactions, but can still display static cluster information.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"uuid": {
				Description: "The unique identifier for the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"zk_health": {
				Description: "The health status of the HyperFlex cluster's zookeeper ensemble.\n* `NOT_AVAILABLE` - The operational status of the ZK ensemble is not provided by the HyperFlex cluster.\n* `UNKNOWN` - The operational status of the ZK ensemble cannot be determined.\n* `ONLINE` - The ZK ensemble is online and operational.\n* `OFFLINE` - The ZK ensemble is offline and not operational.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"zone_resiliency_list": {
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
						"name": {
							Description: "The name of the availability zone.",
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
						"resiliency_info": {
							Description: "The detailed resiliency info of the zone. The information includes the current health and number of device failures tolerable for this particular zone.",
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
									"data_replication_factor": {
										Description: "The number of copies of data replicated by the cluster.\n* `ONE_COPY` - The HyperFlex cluster does not replicate data.\n* `TWO_COPIES` - The HyperFlex cluster keeps 2 copies of data.\n* `THREE_COPIES` - The HyperFlex cluster keeps 3 copies of data.\n* `FOUR_COPIES` - The HyperFlex cluster keeps 4 copies of data.\n* `SIX_COPIES` - The HyperFlex cluster keeps 6 copies of data.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"hdd_failures_tolerable": {
										Description: "The number of persistent device disruptions the HyperFlex storage cluster can handle at this point in time.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"messages": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"node_failures_tolerable": {
										Description: "The number of node disruptions the HyperFlex storage cluster can handle at this point in time.",
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
									"policy_compliance": {
										Description: "The state of the storage cluster's compliance with the cluster access policy.\n* `UNKNOWN` - The HyperFlex cluster's compliance with the data replication policy could not be determined.\n* `COMPLIANT` - The HyperFlex cluster is compliant with the data replication policy and data is replicated to the configured replication factor.\n* `NON_COMPLIANT` - The HyperFlex cluster is not compliant with the data replication policy because some data is not replicated.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"resiliency_state": {
										Description: "The resiliency state of the storage platform. The resiliency state is the storage cluster's current ability to maintain data.\n* `UNKNOWN` - The resiliency status of the HyperFlex cluster cannot be determined, or the cluster is transitioning into ONLINE state.\n* `HEALTHY` - The HyperFlex cluster is healthy. The cluster is able to perform IO operations and data is available.\n* `WARNING` - The HyperFlex cluster or data availability is adversely affected. This can happen if there are node or storage device failures beyond the tolerable failure threshold.\n* `OFFLINE` - The HyperFlex cluster is offline and not performing IO operations.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"ssd_failures_tolerable": {
										Description: "The number of cache device disruptions the HyperFlex storage cluster can handle at this point in time.",
										Type:        schema.TypeInt,
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
		},
	}
}

func dataSourceHyperflexHealthRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHealth{}
	if v, ok := d.GetOk("arbitration_service_state"); ok {
		x := (v.(string))
		o.SetArbitrationServiceState(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("data_replication_compliance"); ok {
		x := (v.(string))
		o.SetDataReplicationCompliance(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}
	if v, ok := d.GetOk("zk_health"); ok {
		x := (v.(string))
		o.SetZkHealth(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHealth object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexHealth: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexHealth list: %s", err.Error())
	}
	var s = &models.HyperflexHealthList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexHealth list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for HyperflexHealth did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexHealth{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("arbitration_service_state", (s.GetArbitrationServiceState())); err != nil {
				return diag.Errorf("error occurred while setting property ArbitrationServiceState: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cluster", flattenMapHyperflexClusterRelationship(s.GetCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Cluster: %s", err.Error())
			}
			if err := d.Set("data_replication_compliance", (s.GetDataReplicationCompliance())); err != nil {
				return diag.Errorf("error occurred while setting property DataReplicationCompliance: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("resiliency_details", flattenMapHyperflexHxResiliencyInfoDt(s.GetResiliencyDetails(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ResiliencyDetails: %s", err.Error())
			}
			if err := d.Set("state", (s.GetState())); err != nil {
				return diag.Errorf("error occurred while setting property State: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property Uuid: %s", err.Error())
			}
			if err := d.Set("zk_health", (s.GetZkHealth())); err != nil {
				return diag.Errorf("error occurred while setting property ZkHealth: %s", err.Error())
			}

			if err := d.Set("zone_resiliency_list", flattenListHyperflexHxZoneResiliencyInfoDt(s.GetZoneResiliencyList(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ZoneResiliencyList: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
