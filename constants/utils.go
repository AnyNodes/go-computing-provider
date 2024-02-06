package constants

const StatusActive = "Active"
const StatusOffline = "Offline"

// bidding status
const BiddingCreated string = "created"
const BiddingAccepting string = "accepting_bids"
const BiddingProcessing string = "processing"
const BiddingSubmitted string = "submitted"
const BiddingCompleted string = "completed"
const BiddingCancelled string = "cancelled"

const TASK_DEPLOY string = "worker.deploy"

const K8S_NAMESPACE_NAME_PREFIX = "ns-"
const K8S_CONTAINER_NAME_PREFIX = "pod-"
const K8S_INGRESS_NAME_PREFIX = "ing-"
const K8S_SERVICE_NAME_PREFIX = "svc-"
const K8S_DEPLOY_NAME_PREFIX = "deploy-"

const REDIS_SPACE_PREFIX = "FULL:"
const REDIS_UBI_C2_PERFIX = "UBI-C2:"

const UBI_TASK_RECEIVED_STATUS = "received"
const UBI_TASK_RUNNING_STATUS = "running"
const UBI_TASK_SUCCESS_STATUS = "success"
const UBI_TASK_FAILED_STATUS = "failed"

const CPU_AMD = "AMD"
const CPU_INTEL = "INTEL"
