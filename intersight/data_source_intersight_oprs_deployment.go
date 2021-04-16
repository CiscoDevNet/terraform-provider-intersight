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

func dataSourceOprsDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOprsDeploymentRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"available_replicas": {
				Description: "Available number of replicas.",
				Type:        schema.TypeInt,
				Optional:    true,
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
			"desired_replicas": {
				Description: "The expected number of replicas.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"event": {
				Description: "The type of event which was triggered.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"name": {
				Description: "Agent name for which the event is triggered.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"namespace": {
				Description: "Name space in which the agents are running.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"status": {
				Description: "Status which shows if the resource is healthy or not.\n* `` - An Unknown status indicates that the resource status is not known.\n* `Healthy` - A healthy status indicates that the resource is healthy and running as per spec.\n* `Unhealthy` - An unhealthy status indicates that the resource is down.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"time_stamp": {
				Description: "The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"unavailable_replicas": {
				Description: "Number of replicas Unavailable.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceOprsDeployment().Schema},
				Computed: true,
			}},
	}
}

func dataSourceOprsDeploymentRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.OprsDeployment{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("available_replicas"); ok {
		x := int64(v.(int))
		o.SetAvailableReplicas(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("desired_replicas"); ok {
		x := int64(v.(int))
		o.SetDesiredReplicas(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("event"); ok {
		x := (v.(string))
		o.SetEvent(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("namespace"); ok {
		x := (v.(string))
		o.SetNamespace(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("time_stamp"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetTimeStamp(x)
	}
	if v, ok := d.GetOk("unavailable_replicas"); ok {
		x := int64(v.(int))
		o.SetUnavailableReplicas(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of OprsDeployment object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.OprsApi.GetOprsDeploymentList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of OprsDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of OprsDeployment: %s", responseErr.Error())
	}
	count := countResponse.OprsDeploymentList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for OprsDeployment data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var oprsDeploymentResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.OprsApi.GetOprsDeploymentList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching OprsDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching OprsDeployment: %s", responseErr.Error())
		}
		results := resMo.OprsDeploymentList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["assist"] = flattenMapAssetDeviceRegistrationRelationship(s.GetAssist(), d)
				temp["available_replicas"] = (s.GetAvailableReplicas())
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["desired_replicas"] = (s.GetDesiredReplicas())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["event"] = (s.GetEvent())

				temp["labels"] = flattenListOprsKvpair(s.GetLabels(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["namespace"] = (s.GetNamespace())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["time_stamp"] = (s.GetTimeStamp()).String()
				temp["unavailable_replicas"] = (s.GetUnavailableReplicas())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				oprsDeploymentResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(oprsDeploymentResults))
	if err := d.Set("results", oprsDeploymentResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(oprsDeploymentResults[0]["moid"].(string))
	return de
}
