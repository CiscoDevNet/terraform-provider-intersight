package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualizationVirtualDisk() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualizationVirtualDiskRead,
		Schema: map[string]*schema.Schema{
			"capacity": {
				Description: "Disk capacity to be provided with units example - 10Gi.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"discovered": {
				Description: "Flag to indicate whether the configuration is created from inventory object.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Description: "File mode of the disk  example - Filesystem, Block.\n* `Block` - It is a Block virtual disk.\n* `Filesystem` - It is a File system virtual disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the storage disk. Name must be unique within a Datastore.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"source_certs": {
				Description: "Base64 encoded CA certificates of the https source to check against.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"source_disk_to_clone": {
				Description: "Source disk from which the content is copied.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"source_file_path": {
				Description: "Image path used to import on the created disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVirtualizationVirtualDisk().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVirtualizationVirtualDiskRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VirtualizationVirtualDisk{}
	if v, ok := d.GetOk("capacity"); ok {
		x := (v.(string))
		o.SetCapacity(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("discovered"); ok {
		x := (v.(bool))
		o.SetDiscovered(x)
	}
	if v, ok := d.GetOk("mode"); ok {
		x := (v.(string))
		o.SetMode(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("source_certs"); ok {
		x := (v.(string))
		o.SetSourceCerts(x)
	}
	if v, ok := d.GetOk("source_disk_to_clone"); ok {
		x := (v.(string))
		o.SetSourceDiskToClone(x)
	}
	if v, ok := d.GetOk("source_file_path"); ok {
		x := (v.(string))
		o.SetSourceFilePath(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VirtualizationVirtualDisk object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualDiskList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VirtualizationVirtualDisk: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VirtualizationVirtualDiskList.GetCount()
	var i int32
	var virtualizationVirtualDiskResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualDiskList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VirtualizationVirtualDisk: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VirtualizationVirtualDiskList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VirtualizationVirtualDisk data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["capacity"] = (s.GetCapacity())
				temp["class_id"] = (s.GetClassId())

				temp["cluster"] = flattenMapVirtualizationBaseClusterRelationship(s.GetCluster(), d)
				temp["discovered"] = (s.GetDiscovered())

				temp["inventory"] = flattenMapVirtualizationBaseVirtualDiskRelationship(s.GetInventory(), d)
				temp["mode"] = (s.GetMode())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["source_certs"] = (s.GetSourceCerts())
				temp["source_disk_to_clone"] = (s.GetSourceDiskToClone())
				temp["source_file_path"] = (s.GetSourceFilePath())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["workflow_info"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)
				virtualizationVirtualDiskResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(virtualizationVirtualDiskResults))
	if err := d.Set("results", virtualizationVirtualDiskResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(virtualizationVirtualDiskResults[0]["moid"].(string))
	return de
}
