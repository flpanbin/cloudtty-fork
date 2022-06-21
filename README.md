# A Kubernates Cloud-Shell (Web Terminal) Operator.

English | [简体中文](https://github.com/cloudtty/cloudtty/blob/main/README_zh.md)

# Why needs cloudtty ?

Existing project [ttyd](https://github.com/tsl0922/ttyd) already provides great feature to share terminal over the web.
But Kubernetes cloud native enviroment requires a way to run web-tty via `kubernetes` way ( running as pod, and generated by CRD).
So you can try cloudtty.

# Applicable scene

1. Many enterprises use cloud platform to manage Kubernetes, but due to security reasons, they cannot use SSH to connect the node host to execute `kubectl` command.So it requires a cloud shell capability.

2. A running container on kubernetes can be "entered" (`Kubectl exec `) on a browser web page.

3. The container logs can be displayed realtime(scrolling) on a browser web page.

# Demo

![screenshot_gif](https://github.com/cloudtty/cloudtty/raw/main/docs/snapshot.gif)

# Quick Start

- Step 1: Install Operator and CRD

```
helm repo add daocloud  https://release.daocloud.io/chartrepo/cloudshell
helm install --version 0.1.0 daocloud/cloudtty --generate-name
```

- Step 2: Create a cloud-tty instance by applying CR, then monitor its status
```
kubectl apply -f ./config/samples/local-cluster_v1alpha1_cloudshell.yaml  && kubectl get cloudshell -w
```
By default, it will create a cloud-tty pod and expose `NodePort` type service.

Alternatively, `Cluster-IP`, `Ingress`, `Virtual Service`(for Istio) are all supported as `exposureMode`, please refer to `config/samples/` for more examples.

- Step 3: Observe CR status to obtain its web access url, such as:

```
 $kubectl get cloudshell -w
```

we can see the following information：

```shell
NAME                 USER   COMMAND  TYPE        URL                 PHASE   AGE
cloudshell-sample    root   bash     NodePort    192.168.4.1:30167   Ready   31s
cloudshell-sample2   root   bash     NodePort    192.168.4.1:30385   Ready   9s
```

When the status of cloudshell changes to `Ready` and the `URL` field appears, it can be accessed by the field in the browser, as shown below

![screenshot_png](https://github.com/cloudtty/cloudtty/raw/main/docs/snapshot.png)


# Advanced Usage Guide

## 1. Manage Multiple/Remote clusters

When if cloudtty will manage a remote cluster (another cluster than which cloudtty operator runs on), you will tell cloudtty the kube.conf of the remote cluster as below:

```
kubectl create configmap my-kubeconfig --from-file=kube.config # the kube.config can be copied from remote cluster `~/.kube/config`

# Be careful to ensure the /root/.kube/config:
#  1. contains the base64 encoded certs/secrets instead of local files.
#  2. the k8s api-server endpoint is reachable(host IP or cluster IP) instead of localhost
```

* If the cluster is remote, `cloudtty` needs to specify `kubeconfig` to access the cluster using the kubectl command tool. you need to provide the kubeconfig stored in configmap and specify the name to cloudshell `spec.configmapName` cr. kubeconfig will be automatically mounted to the cloudtty container. ensure that the server IP address is properly connected to the cluster network.

* If cloudtty runs on the same cluster which to be managed, you don't necessary to do this( a ServiceAccount with `cluster-admin` role permissions will be bind to the pod automaticlly. Inside the container, kubectl automatically detects 'CA'certificates and token. If there are a concerns with security, you can also provide your own kubeconfig to control the permissions of different users.)

## 2. More Exposure Mode

Cloudtty provides four modes to expose cloudtty services: `ClusterIP`, `NodePort`, `Ingress`, and `VitualService` to satisfy different usage scenarios:

* ClusterIP: [Service](https://kubernetes.io/docs/concepts/services-networking/service/) ClusterIP type in cluster. suitable for third-party integration of cloudtty server, users can choose a more flexible way to expose their services.

* NodePort (default): The simplest way to expose the service mode is to create a service resource of type NodePort in cluster. You can access the cloudtty service using the master node IP address and port number.

* Ingress: Create a Service resource of ClusterIP type in cluster and create an ingress resource to load the service based on routing rules. This works when the [Ingress Controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/) is used for traffic load in the cluster.

* VirtualService (Istio): Create a ClusterIP Service resource in cluster and create a `VirtaulService` resource. This mode is used when [Istio](https://github.com/istio/istio) is used to load traffic in a cluster.

#### Theory

1. Operator creates job and service with the same name in the corresponding NS. If Ingress or VitualService is used, it also creates routing information.

2. When pod phase to `Ready`, it set the access url to cloudshell status.

3. When [job](https://kubernetes.io/docs/concepts/workloads/controllers/job/) ends after the TTL times out or for other reasons, the cloudshell status changes to 'Completed' once the job changes to 'Completed'. we can set cloudshell to delete associated resources when the state is `Completed`.

4. When cloudshell had deleted, the corresponding job and service (through 'ownerReference') are automatically deleted. If Ingress or VitualService mode is used, the corresponding routing information is also deleted.

# SPECIAL THANKS
This project is based on https://github.com/tsl0922/ttyd. Many thanks to `tsl0922` `yudai` and the community.
The frontend UI code was originated from `ttyd` project, and the ttyd binary inside the container also comes from `ttyd` project.

# Developer Guide

## Run operator and install CRD
  
1. generate crd to [charts/_crds]()

```shell
make generate-yaml
```

2. install CRD

```shell
make install
```
  
3. running operator

  ```shell
make run
  ```
 
## create cloudshell

example: automatically prints logs for a container

```yaml
apiVersion: cloudshell.cloudtty.io/v1alpha1
kind: CloudShell
metadata:
  name: cloudshell-sample
spec:
  configmapName: "my-kubeconfig"
  runAsUser: "root"
  commandAction: "kubectl -n kube-system logs -f kube-apiserver-cn-stack"
  once: false
```

ToDo:

- Control permissions through `RBAC（generate file `/var/run/secret`）.

- For security, jobs should run in separate namespace, not in the same of namespace of clodushell.

- Check the pod is running and endpoint status changes to `Ready`, the cloudshell phase is set to Ready.

- TTL should be set to both job and shell.

- Job creation templates are currently hardcode and should provide a more flexible way to modify the job template.

More will be coming Soon. Welcome to your [issue](https://github.com/cloudtty/cloudtty/issues) and PR. 🎉🎉🎉
