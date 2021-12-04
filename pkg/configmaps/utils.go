// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package configmaps

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"

	deployerv1 "github.com/Azure/aks-deployer/pkg/api/v1"
)

// ParseConfig parses configuration data and return a list of runtime.Object
func ParseConfig(logger *logrus.Entry,
	scheme *runtime.Scheme, configData string) ([]runtime.Object, error) {
	decode := serializer.NewCodecFactory(scheme).UniversalDeserializer().Decode
	configs := strings.Split(configData, configDataDelimiter)
	var objs []runtime.Object
	for _, config := range configs {
		if config == "" {
			continue
		}

		obj, _, err := decode([]byte(config), nil, nil)
		if err != nil {
			logger.Errorf("unable to parse configuration data to runtime.Objects, %s", err.Error())
			logger.Errorf("raw data:\n%s", configData)
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

// ParseConfigToUnstructured parses configuration data and return a list of unstructured.Unstructured
func ParseConfigToUnstructured(logger *logrus.Entry, configData string) ([]*unstructured.Unstructured, error) {
	configs := strings.Split(configData, configDataDelimiter)
	var objs []*unstructured.Unstructured
	for _, config := range configs {
		if config == "" {
			continue
		}

		obj := &unstructured.Unstructured{}
		dec := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
		_, _, err := dec.Decode([]byte(config), nil, obj)
		if err != nil {
			logger.Errorf("unable to parse configuration data to unstructured, %s", err.Error())
			logger.Errorf("raw data:\n%s", config)
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

// GetNamespacedClusterConfigMapName returns namespaced cluster ConfigMap name
func GetNamespacedClusterConfigMapName(namespace string) string {
	return fmt.Sprintf(clusterConfigNameFormat, namespace, ClusterConfigMapName)
}

// GetAksAppConfigMapName returns AksApp ConfigMap name
func GetAksAppConfigMapName(aksapp deployerv1.AksApp) string {
	return fmt.Sprintf(aksAppConfigNameFormat, aksapp.Spec.Type, aksapp.Spec.Version)
}
