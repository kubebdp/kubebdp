# kubebdp

## 项目状态

目前此项目处于设计和信息收集阶段，欢迎有兴趣的人加入，微信加`Blue_L`入群。

项目的整体设计文档[参考这里](doc/design.md)。

## 项目介绍

kubebdp是一个基于kubernetes的大数据管理平台。

通常我们使用ambari进行管理大数据服务组件，ambari支持hadoop、hive、hbase等众多大数据生态组价的部署和管理。ambari采用server + agent架构，基于主机系统而设计的，可以支持虚拟机，裸技术服务器上的服务管理，是一套成熟的系统。

伴随着kubernetes的日益成熟稳定和更加丰富的功能，我们看到了基于容器技术来管理大数据服务组件带来的各种优点，因而创建了这个项目，推进大数据平台服务运行在kubernetes上的探索与实践。

**优点**

将大数据服务组件运行在kubernetes上，有下列诸多优点：

- **完全标准化的部署配置模型**，完全标准化的部署与配置模型可支持集群的快速复制，标准化的输出配置。
- **多个大数据集群可共享相同的主机资源**，实现同一批物理资源支持多套大数据集群的运行。
- **一个大数据集群内多个服务实例的支持**，支持单个集群内多个服务实例的运行，可支持多租户场景。
- **完全自动化的扩缩容管理**，通过使用kubernetes的hpa机制，做到服务的完全自动化的扩缩容管理。
- **统一的计算调度框架**，使用pod作为最小计算单元，结合相关调度插件实现统一的资源管理。
- **自动化的服务与网络的配置**，自动化配置服务间的访问，集群内外的dns和网络访问关系。
- **一体化的运维管理方式**，通过exec接口直接进入各个pod执行管理命令，监控和日志等查看。

**大数据服务组件**

计划支持常用的大数据服务组件：

- kerberos
- zookeeper
- ldap
- hadoop
- hive
- impala
- hbase
- spark
- flink
- kafka
- elasticsearch

通用服务组件

- mysql/tidb
- redis
- juicefs
- minio

**平台服务组件**

平台与基础服务组件有：

- kubernetes
- prometheus
- elasticsearch
- kube-batch/vocalno
- flannel/calico/cilium
- metallb
- external-dns
- ceph
- minio
- juicefs

**主要工作**

要在kubernetes上运行和管理大数据服务组件，主要有下列工作：

- 服务的部署管理
- 服务的配置管理
- 服务的升级容管理
- 服务的本身的管理功能
- 服务的自动化运维功能
- kerberos和ldap服务的集成和配置
- 监控指标的收集
- 日志的收集
