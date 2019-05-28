// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func migrateStateNoop(v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	return is, nil
}

func GetComputeRegionBackendServiceCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeRegionBackendServiceApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/RegionBackendService",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "RegionBackendService",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeRegionBackendServiceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionBackendServiceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	healthChecksProp, err := expandComputeRegionBackendServiceHealthChecks(d.Get("health_checks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("health_checks"); !isEmptyValue(reflect.ValueOf(healthChecksProp)) && (ok || !reflect.DeepEqual(v, healthChecksProp)) {
		obj["healthChecks"] = healthChecksProp
	}
	backendsProp, err := expandComputeRegionBackendServiceBackend(d.Get("backend"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backend"); !isEmptyValue(reflect.ValueOf(backendsProp)) && (ok || !reflect.DeepEqual(v, backendsProp)) {
		obj["backends"] = backendsProp
	}
	descriptionProp, err := expandComputeRegionBackendServiceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeRegionBackendServiceFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	protocolProp, err := expandComputeRegionBackendServiceProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	sessionAffinityProp, err := expandComputeRegionBackendServiceSessionAffinity(d.Get("session_affinity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("session_affinity"); !isEmptyValue(reflect.ValueOf(sessionAffinityProp)) && (ok || !reflect.DeepEqual(v, sessionAffinityProp)) {
		obj["sessionAffinity"] = sessionAffinityProp
	}
	regionProp, err := expandComputeRegionBackendServiceRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	timeoutSecProp, err := expandComputeRegionBackendServiceTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	connectionDrainingProp, err := expandComputeRegionBackendServiceConnectionDraining(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connection_draining"); !isEmptyValue(reflect.ValueOf(connectionDrainingProp)) && (ok || !reflect.DeepEqual(v, connectionDrainingProp)) {
		obj["connectionDraining"] = connectionDrainingProp
	}
	loadBalancingSchemeProp, err := expandComputeRegionBackendServiceLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}

	return obj, nil
}

func expandComputeRegionBackendServiceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceHealthChecks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeRegionBackendServiceBackend(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeRegionBackendServiceBackendDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedGroup, err := expandComputeRegionBackendServiceBackendGroup(original["group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroup); val.IsValid() && !isEmptyValue(val) {
			transformed["group"] = transformedGroup
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionBackendServiceBackendDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceSessionAffinity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionBackendServiceTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceConnectionDraining(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedConnection_draining_timeout_sec, err := expandComputeRegionBackendServiceConnectionDrainingConnection_draining_timeout_sec(d.Get("connection_draining_timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["drainingTimeoutSec"] = transformedConnection_draining_timeout_sec
	}

	return transformed, nil
}

func expandComputeRegionBackendServiceConnectionDrainingConnection_draining_timeout_sec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
