# 平台标准

平台是指包括kubernetes，以及生态内相关的基础服务组件。通过标准化这些基础服务组件，为上层服务提供一个一致的平台环境。

## kubernetes

整个平台的基础pass环境。

## flannel/calico/cilium/kube-router/metallb

这些组件为集群提供网络功能。在基础环境支持bgp的话，推荐使用calico或cilium+kube-router，如果基础网络环境不支持bgp的话，根据实际情况评估。

metallb主要实现为service供给LoadBlancer功能。

针对hdfs这种直接读写各个服务节点的情况，或者直接访问计算节点能这种情况，当有从集群外访问的需求时，则最好要求集群环境支持bgp协议，这样可以将这些pod的地址通过下面的externl-dns组件发布到集群外。如果不支持bgp，依赖于service的话，会需要很多service，增加对网络物理ip的依赖，不利于网络的管理维护。

## ceph

为整个平台提供对象存储，文件存储和块存储服务。

## openebs

为大数据组件和ceph提供本地磁盘。对于裸物理机部署场景，单台主机上往往具备多块磁盘，为了能够更加充分的使用这些磁盘，引入openebs。

例如，针对hdfs部署可以为其配置ceph的rbd设备，或者直接使用本地盘，使用rbd设备的话会明显增加网络带宽占用，hdfs写入的数据通过网络传输到各个rbd设备上，而直接使用本地磁盘，则数据可以直接落到本地的多个磁盘设备上，减少了网络带宽的浪费。

此外，针对于一些数据库场景，使用本地的ssd设备会具备比基于网络的磁盘具备更高的iops，可极大的提升数据库性能。

但是，当前的kubernetes针对本地pv的支持并不完美，kube-scheduler组件在调度pod的时候不会根据pvc进行过滤和打分。当前openebs对本地盘的供给也仅是采用延迟绑定的pvc功能，但这就很可能导致pod调度的节点上没有足够的磁盘资源，也不能实现针对不同的节点/磁盘类型作出不同的调度动作。如申请高性能ssd存储的pvc资源，kubernetes并不知道哪些节点有多少可用资源，哪些节点有ssd类型磁盘。这里需要对kube-scheduler和openebs做一定改造，在调度时如果pod的pvc的provisioner是openebs的话，在filter和score阶段查询openebs的接口来获取最合适的节点，

## external-dns

external-dns可以将集群内的pod，service，ingress等资源发布到集群外的dns服务器上。实现集群外的主机直接通过dns访问相关服务和pod，并且避免了pod或service的ip地址变化的问题。

## prometheus

云原生事实监控解决方案。

## elasticsearch

日志收集方案，elk套件。
