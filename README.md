# kubebdp

kubebdp是一个基于kubernetes的大数据管理平台。

ambari是我们经常使用到的大数据管理工具，支持hadoop、hive、hbase等众多大数据生态组价的部署和管理。

ambari是基于server + agent架构，基于主机系统设计的，而使用kubernetes进行管理的话，具有下列优势：

- 标准化的部署和输出
- 一个kubernetes集群内核部署多套大数据集群
- 一个大数据集群内多个应用服务实例
- 自动化扩缩容管理

其中多套大数据集群支持可有效支持在一批节点上部署不同的多套集群，如不同的测试环境，并结合自动扩缩容的动态管理，能够进一步缩减硬件成本。

要实现将大数据组件的自动化管理功能在kubernetes上实现，需要考虑下面的一些点：

- 服务的部署管理
- 服务的配置管理
- 服务的升级容管理
- 服务的本身的管理功能
- 服务的自动化运维功能
- kerberos和ldap服务的集成和配置
- 监控指标的收集
- 日志的收集

计划支持常用的大数据服务组件：

- kerberos
- zookeeper
- ldap
- hadoop
- hive
- hbase
- spack
- flink
- kafka
- elasticsearch
- cassandra

计划支持可选开源组件有：

- minio
- ceph
- juicefs

目前此项目处于设计和信息收集阶段，欢迎有兴趣的人加入，微信加`Blue_L`入群。
