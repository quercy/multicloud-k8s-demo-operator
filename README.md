# multicloud-k8s-demo-operator

This is a Kubernetes Operator that uses the Golang [Operator SDK](https://sdk.operatorframework.io/) to create a deployment and service that hosts a [Presto Coordinator](https://prestodb.io/docs/current/overview/concepts.html#:~:text=The%20Presto%20coordinator%20is%20the,to%20submit%20statements%20for%20execution.), based on the [Presto Helm chart](https://github.com/helm/charts/tree/master/stable/presto) (but not using Helm). The fictional company that is hosting this operator is `skittles`, so the API group is named `skittlesv1`.

This is not currently a polished tool, and is only for reference purposes.

Don't edit the files in `config` directly; most of them are managed by the Operator SDK.

## Deploying the operator to K8s from source
Run this after changing any kubebuilder annotations.
```
make manifests deploy
```
## Deploying a new Docker image
Run this after changing any Go code. May need to kill the current Controller pods to pull in the new image.
```
make docker-build docker-push
```

## Key files
* Reconciliation/controller logic: `controllers/coordinator.go`
* Sample deployment YAML: `config/samples/skittles_v1_presto.yaml`
* Deployment types: `api/v1/presto_types.go`