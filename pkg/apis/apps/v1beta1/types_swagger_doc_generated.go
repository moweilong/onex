// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using scripts/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_Condition = map[string]string{
	"":                   "Condition defines an observation of a cloud miner resource operational state.",
	"type":               "Type of condition in CamelCase or in foo.example.com/CamelCase. Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"severity":           "Severity provides an explicit classification of Reason code, so the users or machines can immediately understand the current situation and act accordingly. The Severity field MUST be set only when Status=False.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another. This should be when the underlying condition changed. If that is not known, then using the time when the API field changed is acceptable.",
	"reason":             "The reason for the condition's last transition in CamelCase. The specific API may choose whether or not this field is considered a guaranteed API. This field may not be empty.",
	"message":            "A human readable message indicating details about the transition. This field may be empty.",
}

func (Condition) SwaggerDoc() map[string]string {
	return map_Condition
}

var map_MinerAddress = map[string]string{
	"":        "MinerAddress contains information for the miner's address.",
	"type":    "Miner address type, one of Hostname, ExternalIP or InternalIP.",
	"address": "The machine address.",
}

func (MinerAddress) SwaggerDoc() map[string]string {
	return map_MinerAddress
}

var map_ObjectMeta = map[string]string{
	"":            "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create. This is a copy of customizable fields from metav1.ObjectMeta.\n\nObjectMeta is embedded in `Miner.Spec` and `MinerSet.Template`, which are not top-level Kubernetes objects. Given that metav1.ObjectMeta has lots of special cases and read-only fields which end up in the generated CRD validation, having it as a subset simplifies the API and some issues that can impact user experience.\n\nDuring the [upgrade to controller-tools@v2](https://github.com/kubernetes-sigs/cluster-api/pull/1054) for v1alpha2, we noticed a failure would occur running Cluster API test suite against the new CRDs, specifically `spec.metadata.creationTimestamp in body must be of type string: \"null\"`. The investigation showed that `controller-tools@v2` behaves differently than its previous version when handling types from [metav1](k8s.io/apimachinery/pkg/apis/meta/v1) package.\n\nIn more details, we found that embedded (non-top level) types that embedded `metav1.ObjectMeta` had validation properties, including for `creationTimestamp` (metav1.Time). The `metav1.Time` type specifies a custom json marshaller that, when IsZero() is true, returns `null` which breaks validation because the field isn't marked as nullable.\n\nIn future versions, controller-tools@v2 might allow overriding the type and validation for embedded types. When that happens, this hack should be revisited.",
	"labels":      "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
	"annotations": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
}

func (ObjectMeta) SwaggerDoc() map[string]string {
	return map_ObjectMeta
}

var map_Chain = map[string]string{
	"":         "Chain is the Schema for the chains API.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of the chain. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "Status is the most recently observed status of the Chain. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (Chain) SwaggerDoc() map[string]string {
	return map_Chain
}

var map_ChainList = map[string]string{
	"":         "ChainList is a list of Chain objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is a list of schema objects.",
}

func (ChainList) SwaggerDoc() map[string]string {
	return map_ChainList
}

var map_ChainSpec = map[string]string{
	"":                       "ChainSpec defines the desired state of Chain.",
	"displayName":            "The display name of the chain.",
	"minerType":              "Genesis node machine configuration.",
	"image":                  "Image specify the blockchain node image.",
	"minMineIntervalSeconds": "Minimum number of seconds for the miners to mine a block.",
	"bootstrapAccount":       "Default bootstrap OneX's Genesis account with 1M TBB tokens. This field is automatic generated by OneX, you should not set this field.",
}

func (ChainSpec) SwaggerDoc() map[string]string {
	return map_ChainSpec
}

var map_ChainStatus = map[string]string{
	"":                   "ChainStatus defines the observed state of Chain.",
	"observedGeneration": "ObservedGeneration is the latest generation observed by the controller.",
	"conditions":         "Conditions defines the current state of the Chain",
}

func (ChainStatus) SwaggerDoc() map[string]string {
	return map_ChainStatus
}

var map_ChargeRequest = map[string]string{
	"":         "ChargeRequest is the Schema for the chargerequests API.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of the chargerequest. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "Status is the most recently observed status of the ChargeRequest. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (ChargeRequest) SwaggerDoc() map[string]string {
	return map_ChargeRequest
}

var map_ChargeRequestList = map[string]string{
	"":         "ChargeRequestList is a list of ChargeRequest objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is a list of schema objects.",
}

func (ChargeRequestList) SwaggerDoc() map[string]string {
	return map_ChargeRequestList
}

var map_ChargeRequestSpec = map[string]string{
	"": "ChargeRequestSpec defines the desired state of ChargeRequest.",
}

func (ChargeRequestSpec) SwaggerDoc() map[string]string {
	return map_ChargeRequestSpec
}

var map_ChargeRequestStatus = map[string]string{
	"": "ChargeRequestStatus defines the observed state of ChargeRequest.",
}

func (ChargeRequestStatus) SwaggerDoc() map[string]string {
	return map_ChargeRequestStatus
}

var map_MinerSet = map[string]string{
	"":         "MinerSet ensures that a specified number of miners replicas are running at any given time.",
	"metadata": "If the Labels of a MinerSet are empty, they are defaulted to be the same as the Miner(s) that the MinerSet manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Spec defines the specification of the desired behavior of the MinerSet. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "Status is the most recently observed status of the MinerSet. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (MinerSet) SwaggerDoc() map[string]string {
	return map_MinerSet
}

var map_MinerSetList = map[string]string{
	"":         "MinerSetList contains a list of MinerSet.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
	"items":    "List of MinerSets.",
}

func (MinerSetList) SwaggerDoc() map[string]string {
	return map_MinerSetList
}

var map_MinerSetSpec = map[string]string{
	"":                        "MinerSetSpec defines the desired state of MinerSet.",
	"replicas":                "Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller",
	"selector":                "Selector is a label query over miners that should match the replica count. Label keys and values that must match in order to be controlled by this MinerSet. It must match the miner template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
	"template":                "Template is the object that describes the miner that will be created if insufficient replicas are detected.",
	"displayName":             "The display name of the minerset.",
	"deletePolicy":            "DeletePolicy defines the policy used to identify miners to delete when downscaling. Defaults to \"Random\". Valid values are \"Random, \"Newest\", \"Oldest\"",
	"minReadySeconds":         "Minimum number of seconds for which a newly created miner should be ready without any of its component crashing, for it to be considered available. Defaults to 0 (miner will be considered available as soon as it is ready)",
	"progressDeadlineSeconds": "The maximum time in seconds for a minerset to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.",
}

func (MinerSetSpec) SwaggerDoc() map[string]string {
	return map_MinerSetSpec
}

var map_MinerSetStatus = map[string]string{
	"":                     "MinerSetStatus represents the current status of a MinerSet.",
	"replicas":             "Replicas is the most recently observed number of replicas.",
	"fullyLabeledReplicas": "The number of miners that have labels matching the labels of the miner template of the minerset.",
	"readyReplicas":        "readyReplicas is the number of miners targeted by this MinerSet with a Ready Condition.",
	"availableReplicas":    "The number of available replicas (ready for at least minReadySeconds) for this minerset.",
	"observedGeneration":   "ObservedGeneration reflects the generation of the most recently observed MinerSet.",
	"failureReason":        "In the event that there is a terminal problem reconciling the replicas, both FailureReason and FailureMessage will be set. FailureReason will be populated with a succinct value suitable for miner interpretation, while FailureMessage will contain a more verbose string suitable for logging and human consumption.\n\nThese fields should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the MinerTemplate's spec or the configuration of the miner controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the miner controller, or the responsible miner controller itself being critically misconfigured.\n\nAny transient errors that occur during the reconciliation of Miners can be added as events to the MinerSet object and/or logged in the controller's output.",
	"failureMessage":       "FailureMessage will be set in the event that there is a terminal problem reconciling the MinerSet and will contain a more verbose string suitable for logging and human consumption.\n\nThis field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the MinerSet's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured.\n\nAny transient errors that occur during the reconciliation of MinerSets can be added as events to the MinerSet object and/or logged in the controller's output.",
	"conditions":           "Represents the latest available observations of a miner set's current state.",
}

func (MinerSetStatus) SwaggerDoc() map[string]string {
	return map_MinerSetStatus
}

var map_MinerTemplateSpec = map[string]string{
	"":         "MinerTemplateSpec describes the data needed to create a Miner from a template.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of the miner. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (MinerTemplateSpec) SwaggerDoc() map[string]string {
	return map_MinerTemplateSpec
}

var map_LocalObjectReference = map[string]string{
	"":     "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
	"name": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
}

func (LocalObjectReference) SwaggerDoc() map[string]string {
	return map_LocalObjectReference
}

var map_Miner = map[string]string{
	"":         "Miner is the Schema for the miners API.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of the miner. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "Most recently observed status of the miner. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (Miner) SwaggerDoc() map[string]string {
	return map_Miner
}

var map_MinerList = map[string]string{
	"":         "MinerList is a list of Miner objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is a list of schema objects.",
}

func (MinerList) SwaggerDoc() map[string]string {
	return map_MinerList
}

var map_MinerSpec = map[string]string{
	"":                   "MinerSpec defines the desired state of Miner.",
	"metadata":           "ObjectMeta will autopopulate the Pod created. Use this to indicate what labels, annotations, name prefix, etc., should be used when creating the Pod.",
	"displayName":        "The display name of the miner.",
	"minerType":          "Miner machine configuration.",
	"restartPolicy":      "Restart policy for the miner. One of Always, OnFailure, Never. Default to Always.",
	"podDeletionTimeout": "PodDeletionTimeout defines how long the controller will attempt to delete the Pod that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
}

func (MinerSpec) SwaggerDoc() map[string]string {
	return map_MinerSpec
}

var map_MinerStatus = map[string]string{
	"":                   "MinerStatus defines the observed state of Miner.",
	"podRef":             "PodRef will point to the corresponding Pod if it exists.",
	"lastUpdated":        "LastUpdated identifies when this status was last observed.",
	"failureReason":      "FailureReason will be set in the event that there is a terminal problem reconciling the Miner and will contain a succinct value suitable for miner interpretation.\n\nThis field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Miner's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured.\n\nAny transient errors that occur during the reconciliation of Miners can be added as events to the Miner object and/or logged in the controller's output.",
	"failureMessage":     "FailureMessage will be set in the event that there is a terminal problem reconciling the Miner and will contain a more verbose string suitable for logging and human consumption.\n\nThis field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Miner's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured.\n\nAny transient errors that occur during the reconciliation of Miners can be added as events to the Miner object and/or logged in the controller's output.",
	"addresses":          "Addresses is a list of addresses assigned to the miner. Queried from kind cluster, if available.",
	"phase":              "Phase represents the current phase of miner actuation. One of: Failed, Provisioning, Provisioned, Running, Deleting This field is maintained by miner controller.",
	"observedGeneration": "ObservedGeneration is the latest generation observed by the controller.",
	"conditions":         "Conditions defines the current state of the Miner",
}

func (MinerStatus) SwaggerDoc() map[string]string {
	return map_MinerStatus
}

var map_PodInfo = map[string]string{
	"":                "PodInfo is a set of ids/uuids to uniquely identify the pod.",
	"operatingSystem": "The Operating System reported by the pod",
	"architecture":    "The Architecture reported by the  pod",
}

func (PodInfo) SwaggerDoc() map[string]string {
	return map_PodInfo
}

// AUTO-GENERATED FUNCTIONS END HERE
