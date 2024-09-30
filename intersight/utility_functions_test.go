package intersight

import (
	"reflect"
	"testing"
)

func TestReorderPolicyBucket(t *testing.T) {
	tests := []struct {
		name     string
		old      []interface{}
		new      []interface{}
		expected []interface{}
	}{
		{
			name: "No changes",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Add new policy",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Remove policy",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Reorder policies",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Empty old array",
			old:  []interface{}{},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Empty new array",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new:      []interface{}{},
			expected: nil,
		},
		{
			name: "Modify existing policy",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "modified",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "modified",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReorderPolicyBucket(tt.old, tt.new)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReorderPolicyBucket() = %v, want %v", result, tt.expected)
			}
		})
	}
}
