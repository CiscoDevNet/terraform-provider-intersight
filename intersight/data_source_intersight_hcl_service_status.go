package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHclServiceStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHclServiceStatusRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"exemption_file_version": {
				Description: "Version of the last modified exemption file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identity": {
				Description: "A field to uniquely identify the document with the status.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_hcl_data_modified_time": {
				Description: "The timestamp of the last modified record in the HCL tool database. Used to query and get updated records.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_imported_data_checksum": {
				Description: "Checksum of the data dump used as the base for delta updates.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the service indicatating if the service is up or under maintenance due to data update.\n* `Unknown` - Default service status. Indicates that the full import of data has failed during service startup.\n* `Initializing` - The service starts or restarts.\n* `DataRefreshing` - Full import of data is in progress.\n* `Active` - The data import is successful and service is ready to serve API requests.",
				Type:        schema.TypeString,
				Optional:    true,
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
		},
	}
}

func dataSourceHclServiceStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewHclServiceStatusWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("exemption_file_version"); ok {
		x := (v.(string))
		o.SetExemptionFileVersion(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("last_hcl_data_modified_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastHclDataModifiedTime(x)
	}
	if v, ok := d.GetOk("last_imported_data_checksum"); ok {
		x := (v.(string))
		o.SetLastImportedDataChecksum(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.HclApi.GetHclServiceStatusList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.HclServiceStatusList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to HclServiceStatus: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewHclServiceStatusWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("exemption_file_version", (s.GetExemptionFileVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property ExemptionFileVersion: %+v", err)
			}
			if err := d.Set("identity", (s.GetIdentity())); err != nil {
				return fmt.Errorf("error occurred while setting property Identity: %+v", err)
			}

			if err := d.Set("last_hcl_data_modified_time", (s.GetLastHclDataModifiedTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property LastHclDataModifiedTime: %+v", err)
			}
			if err := d.Set("last_imported_data_checksum", (s.GetLastImportedDataChecksum())); err != nil {
				return fmt.Errorf("error occurred while setting property LastImportedDataChecksum: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property Status: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
