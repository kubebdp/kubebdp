# 总体设计

kubebdp主要有两部分构成。第一个是集群管理控制器，负责对集群的部署、更配，释放等管理；第二个是各个服务组件的控制器，负责自身组件的部署管理，更配和释放等管理。

## 集群控制器

kubebdp以集群为单位管理大数据服务组件，一个集群内可部署多个服务组件。集群是一个逻辑层面的容器，各个服务组件之间以集群进行关联。集群管理功能在kubebdp中作为一个独立的控制器存在，负责创建和管理具体服务组件的各种资源，协调部署和组件变更的动作。集群管理控制器是在各个服务组件控制器功能之上定义和实现的，针对不同的服务组件定义特定的配置。可针对集群内是否部署了某些服务，自动给其他组件配置对应的参数，如集群资源中定义了kerberos组件，说明我们想部署一个安全集群，会自动给hdfs等服务组件配置kerberos认证参数。

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

StackDefinition定义组件提供的配置和组件以及配置之间的依赖关系，以及组价之间的版本要求。

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

通过Cluster资源，我们可以任意定义集群里想要安装的服务组件，同时集群内的服务组件可以与其他集群的服务组件互相访问。

## 服务控制器

服务组件本身可以通过kubernetes的内置的statefulset或deployment控制器实现，或者可通过自定义控制器实现，尤其针对hdfs，zookeeper等，需要在启动pod后进行一些操作的服务。下面以hdfs为例列举组件和配置的定义。

```yaml
apiVersion: component.hdfs.kubebdp.github.io/v1
kind: NameNode
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

下面是示例配置定义，可以通过编写工具软件将服务组件本身的默认配置生成kubernetes的crd资源。

```yaml
apiVersion: configuration.hdfs.kubebdp.github.io/v3.0
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

这个配置是kubernetes中的crd资源，并不是服务真实需要的格式。此处要求服务组件使用通过或自定义工具实现crd资源到真实配置文件的生成。此处可以考虑通过init container实现。如上面的namenode component中定义了一个configurationReference，该引用只向上述配置crd。控制器向pod内注入一个init container，将这个crd资源转换为真实配置并写入到对应的文件目录中。

服务组件本身还可实现sidecar container，监视配置上述配置文件的变更并重新生成并写入到文件中，此时sidecar container可修改pod加入annotation或label字段说明当前应用正在使用的配置版本和最新版本，这样我们可以知道在修改某些配置后是否需要重启某些pod。或者入股服务组件支持动态的重新加载配置，则可以直接reload程序重新加载配置。

上述sidecar container除了对配置的管理外，还可以实现代理jmx指标，将其转换为prometheus的监控指标。

每个服务有自己的配置和依赖于其他服务的配置，依赖于其他服务的配置需要在服务的资源文件中说明，依赖的资源配置可以来自于其他的Cluster集群，或者是configmap。各个服务组件之间不紧密结合，各个服务组件可以单独部署。

服务组件出了实现服务控制器逻辑和配置文件的crd定义，还可能会实现一些其他的自定义资源，如对于hdfs，可以实现下列资源对集群进行管理。

该资源用来对集群的数据进行rebalance。

```yaml
apiVersion: operate.hdfs.kubebdp.github.io/v3.0
kind: Rebalance
metadata:
  name: ${cluster-name}-hdfs-site
  labels:
    kubebdp.github.io/service: dev-hdfs
  ownerReferences:
  # ..
spec:
  selector:
    name: ${cluster-name}-hdfs
    namespace: default
```

此外我们还应该定义hpa字段以及实现扩缩容逻辑，如移除pod时，先把数据迁移走，扩容时可进行rebalance，并在部署的时候常见hpa资源，以便能够自动对集群进行扩缩容。

此外我们还需要实现服务的升级和回滚策略，并记录历史版本的定义。

此外我们还可以定义prometheus的监控和告警资源等。
