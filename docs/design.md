集群管理

kubebdp以集群为单位管理大数据服务组件，一个集群内可部署多个服务组件。集群是一个逻辑层面的容器，各个服务组件之间以集群进行关联。集群管理功能再kubebdp中作为一个独立的控制器存在，负责创建和管理具体服务组件的各种资源，协调部署和组件变更的动作。集群管理控制器是在各个服务组件控制器功能之上定义和实现的，针对不同的服务组件定义特定的配置。

```yaml
apiVersion: manager.kubebdp.github.io/v1
kind: Cluster
metadata:
  name: ${cluster-name}
spec:
  stack: hadoop-3.1
  services:
    kerberos:
    zookeeper:
    hdfs:
status:
  ...
```

StackDefinition定义组件提供的配置和组件以及配置之间的依赖关系。

```yaml
apiVersion: manager.kubebdp.github.io/v1
kind: StackDefinition
metadata:
  name: default-3.1
spec:
  services:
    kerberos:
    zookeeper:
    hdfs:
status:
  ...
```



不同的集群之间的服务实例可以

服务管理

服务组件本身可以通过kubernetes的内置的statefulset或deployment控制器实现，或者可通过自定义控制器实现，尤其针对hdfs，zookeeper等，需要在启动pod后进行一些操作的服务。

```yaml
apiVersion: component.kubebdp.github.io/v1
kind: HdfsNameNode
metadata:
  name: ${cluster-name}-hdfs
  labels:
    kubebdp.github.io/cluster: ${cluster-name}
  # 如果是某个Cluster控制器创建的则要添加这些字段
  ownerReferences:
  # ...
spec:
  replicas: 2
  volumeClaimTemplates:
  - accessModes: ReadWriteOncePod
    storageClassName: large-capacity
  updateStrategy: RollingUpdate
  # 此处的serviceName会作为dns名称，会为每个pod创建dns
  serviceName: ${cluster-name}-hdfs-dn
  image: kubebdp.github.io/hdfs-namenode:v3.0
  # 此处可以配置额外的pod配置，controller会添加或修改必要配置
  template:
    metadata:
      labels: []
      spec:
        containers: []
  configurations:
    hdfsSite:
      configMapName: ${cluster-name}-hdfs-site
      configurationReference:
        apiVersion: configuration.kubebdp.github.io/v3.0
        kind: HdfsSite
        name: ${cluster-name}-hdfs-site
        namespace: default
    zookeeperSite: ${cluster-name}-zookeeper-site
status:
```

配置都是标准的yaml格式，组件可能依赖于xml或者ini等形式的配置。需要组件自己去做解析。

```yaml
apiVersion: configuration.kubebdp.github.io/v3.0
kind: HdfsSite
metadata:
  name: ${cluster-name}-hdfs-site
  labels:
    kubebdp.github.io/service: dev-hdfs
  ownerReferences:
  # ..
spec:
  nameNodeHeapSize: 2G
```

组件可实现自身的init container从apiserver查询配置并生成自己的格式，实现自己的sidecar container监视自己的资源，监视配置的变化，代理jmx指标等。

每个服务有自己的配置和依赖于其他服务的配置，依赖于其他服务的配置需要在资源文件中说明，依赖的资源配置可以来自于其他的cluster，或者是configmap。各个服务组件之间不紧密结合，各个服务组件可以单独部署。

- 部署
- 删除
- 升级/回滚
- 扩缩容
- 管理工具
- 配置管理
- 监控与告警

平台管理
