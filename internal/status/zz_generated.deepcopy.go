// +build !ignore_autogenerated

/*
RabbitMQ Cluster Operator

Copyright 2021 VMware, Inc. All Rights Reserved.

This product is licensed to you under the Mozilla Public license, Version 2.0 (the "License").  You may not use this product except in compliance with the Mozilla Public License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by controller-gen. DO NOT EDIT.

package status

import (
	"k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAvailableConditionManager) DeepCopyInto(out *ClusterAvailableConditionManager) {
	*out = *in
	in.condition.DeepCopyInto(&out.condition)
	if in.endpoints != nil {
		in, out := &in.endpoints, &out.endpoints
		*out = new(v1.Endpoints)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAvailableConditionManager.
func (in *ClusterAvailableConditionManager) DeepCopy() *ClusterAvailableConditionManager {
	if in == nil {
		return nil
	}
	out := new(ClusterAvailableConditionManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterCondition) DeepCopyInto(out *RabbitmqClusterCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterCondition.
func (in *RabbitmqClusterCondition) DeepCopy() *RabbitmqClusterCondition {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterCondition)
	in.DeepCopyInto(out)
	return out
}
