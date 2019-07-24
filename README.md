## 1. 目标
学习go，并使用go编写
1. 客户端C
2. 中转服务器S
3. 计算服务器（集群），简单计算，loopback原样返回

注意事项：
1. 并发性能
2. 代码规范

## 2. 步骤和计划
1. 【DONE】熟悉go基础语法，绘制[思维导图](./go_xmind.pdf)
2. 【DONE】编写简单的CS socket代码 [simple_client.go](./simple_client.go) [simple_server.go](./simple_server.go) ，支持并发通信，测试效果[如图](./simple_cs_go_v1.0.png)
3. 【TODO】编写TCP反向代理服务器和计算服务器
4. 【TODO】调通整个链路
5. 【TODO】优化服务器端代码，提高并发性能（job缓存，任务队列）
6. 【TODO】测试现有go并发框架

## 3. 说明
1. 查阅很多资料后，发现这个设计架构类似nginx四层架构，TCP反向代理

为了简化计算，先假设
1. 集群的数量是固定的，IP也是固定的，手工分配
2. 简化负载均衡算法：找到空闲server，然后将TCP数据流转发到这个server
