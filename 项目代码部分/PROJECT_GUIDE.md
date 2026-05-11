# 项目开发指南

本文档描述项目结构、各模块职责，以及后续修改建议。

---

## 目录结构详解

### 1. `docs/` - 文档目录

存放所有设计文档、API文档、协议文档等。

**建议后续添加**：
- `architecture.md` - 详细架构设计
- `protocol.md` - 通信协议定义
- `api.md` - API接口文档
- `decision-flow.md` - 决策流程图
- `experiment-design.md` - 实验设计文档

---

### 2. `go-backend/` - Go后端核心

#### `cmd/simulator/` - 仿真器入口

**`main.go`** - 程序入口点，负责：
- 解析命令行参数
- 加载配置文件
- 初始化各模块
- 启动仿真循环

#### `internal/agent/` - Agent实现

**`l1/` - L1全局统筹Agent**
- `orchestrator.go` - 统筹器，协调L1各模块
- `decision_router.go` - **决策权动态路由算法**（创新点1）
  - 根据场景特征动态选择决策路径
  - 实现智能性与实时性的平衡
- `hybrid_decision.go` - 混合决策融合模块
  - 融合LLM决策和传统算法决策
  - 冲突仲裁机制
- `state_manager.go` - 状态管理模块
  - 设备状态聚合
  - 异常检测
  - 任务进度汇总

**`l2/` - L2设备类型Agent**
- `uav_agent.go` - 无人机集群Agent
  - 无人机航线规划
  - 编队控制
  - 传感器数据处理
- `usv_agent.go` - 无人艇Agent
  - 水面航行控制
  - 水面目标跟踪
- `uuv_agent.go` - 无人潜航器Agent
  - 水下航行控制
  - 深度控制

**`statemgr/` - 状态管理模块（可选独立）**
- 如果后续需要将状态管理独立出来，可以放在这里

#### `internal/simulation/` - 仿真引擎

**`world/` - 世界模型**
- `world.go` - 世界模型主类
  - 2.5D坐标系管理
  - 区域表示（海域、禁入区、敏感区）
  - 障碍物管理
  - 环境动态变化（天气、海况等）

**`device/` - 设备模拟**
- `device.go` - 设备基类
  - 设备状态机
  - 简化运动学模型
  - 状态上报机制
- `uav.go` - 无人机具体实现
- `usv.go` - 无人艇具体实现
- `uuv.go` - 无人潜航器具体实现

**`sensor/` - 传感器模拟**
- `sensor.go` - 传感器基类
- `radar.go` - 雷达模拟
- `camera.go` - 视觉传感器模拟
- `sonar.go` - 声呐模拟

#### `internal/planning/` - 规划算法

**`path/` - 路径规划**
- `astar.go` - A*算法实现
- `rrt.go` - RRT算法（可选）
- `trajectory.go` - 轨迹平滑

**`conflict/` - 冲突检测与消解**
- `detector.go` - **三维冲突检测算法**（创新点2）
  - 层级检测策略
  - 时间窗口优化
  - 秒级响应保证
- `resolver.go` - 冲突消解
  - 消解策略库
  - 优先级管理
  - 重规划触发

**`assignment/` - 任务分配**
- `allocator.go` - 任务分配算法
- `replaner.go` - 自适应任务重规划（创新点4）
  - 故障检测触发
  - 动态资源再分配
  - 平滑过渡策略

#### `internal/communication/` - 通信模块

**`message/` - 消息定义**
- `message.go` - 消息结构体定义
- `command.go` - 指令消息
- `status.go` - 状态消息
- `event.go` - 事件消息

**`router/` - 消息路由**
- `router.go` - 消息路由器
  - 严格层级通信规则实现
  - 指令签名与权限校验（创新点3）
  - 异常拦截机制
- `pubsub.go` - 发布订阅机制

#### `internal/types/` - 公共类型定义

- `types.go` - 通用类型定义
  - 坐标、姿态、速度等
  - 设备类型、任务类型枚举
  - 错误码定义

#### `pkg/protocol/` - 可导出的协议定义

- 供外部调用的协议定义
- Go和Python共用的类型定义

---

### 3. `python-llm/` - Python LLM层

#### `llm_service/`

**`server.py`** - gRPC/HTTP服务入口
- 启动LLM服务
- 注册API端点
- 处理Go后端的请求

**`agent/l1_agent.py`** - L1的LLM部分
- `situation_understanding()` - 态势理解
  - 解析战场态势
  - 理解目标意图
- `task_decomposition()` - 任务分解
  - 将大任务分解为子任务
  - 分配给各设备Agent
- `emergency_decision()` - 应急决策
  - 处理突发情况
  - 生成应急预案

**`agent/prompts/`** - Prompt模板目录
- `situation_understanding.prompt` - 态势理解Prompt
- `task_decomposition.prompt` - 任务分解Prompt
- `emergency.prompt` - 应急决策Prompt

**`rag/`** - RAG知识库模块
- `knowledge_base.py` - 知识库主类
  - 向量检索
  - 相关文档召回
- `documents/` - 知识库文档目录
  - `tactics.md` - 典型战术预案
  - `devices.md` - 设备能力手册
  - `procedures.md` - 标准作战流程
  - `cases.md` - 历史案例

**`client.py`** - Go调用的客户端封装
- 简化Go端调用LLM服务的代码

#### `tests/` - Python测试
- 单元测试
- 集成测试

#### `config.py` - Python配置
- LLM API配置
- RAG配置
- 服务端口配置

#### `requirements.txt` - Python依赖
```
openai>=1.0.0
anthropic>=0.18.0
chromadb>=0.4.0
grpcio>=1.59.0
fastapi>=0.104.0
uvicorn>=0.24.0
pydantic>=2.5.0
```

---

### 4. `web-frontend/` - 可视化前端（可选）

**`index.html`** - 主页面
**`app.js`** - 前端逻辑
**`style.css`** - 样式

功能建议：
- 2.5D战场态势展示
- 实时位置追踪
- 任务状态显示
- 冲突告警提示
- 日志回放

---

### 5. `configs/` - 配置文件

**`simulator.yaml`** - 仿真器配置
```yaml
world:
  width: 10000    # 海域宽度（米）
  height: 10000   # 海域高度（米）
  depth: 500       # 最大深度（米）

simulation:
  tick_rate: 10    # 仿真频率（Hz）
  max_duration: 3600  # 最大仿真时长（秒）
```

**`llm.yaml`** - LLM配置
```yaml
provider: openai  # openai / anthropic / qianfan
model: gpt-4
api_key: ${API_KEY}
temperature: 0.7
max_tokens: 2000
```

**`devices.yaml`** - 设备配置
```yaml
uav:
  count: 3
  max_speed: 50     # m/s
  max_height: 1000  # m
  sensor_range: 2000 # m

usv:
  count: 2
  max_speed: 15     # m/s
  sensor_range: 3000 # m

uuv:
  count: 2
  max_speed: 10     # m/s
  max_depth: 300    # m
  sensor_range: 1500 # m
```

---

### 6. `scenarios/` - 测试场景

**`scenario1.yaml`** - 典型入侵场景
```yaml
name: 典型入侵场景
description: 单个可疑目标从东南方向入侵

intruder:
  type: ship
  position: [8000, 8000, 0]  # x, y, z
  velocity: [5, 0, 0]

mission:
  type: intercept
  rules:
    - search
    - track
    - warn
    - intercept

duration: 1800  # 30分钟
```

**`scenario2.yaml`** - 故障场景
```yaml
name: 设备故障场景
description: 跟踪过程中一架无人机故障

failures:
  - time: 300  # 第5分钟
    device: uav-1
    type: communication_loss

  - time: 600  # 第10分钟
    device: usv-1
    type: engine_failure
```

**`scenario3.yaml`** - 复杂场景
```yaml
name: 复杂协同场景
description: 多个目标同时入侵，需要多设备协同

intruders:
  - type: uav
    position: [2000, 8000, 500]
  - type: ship
    position: [8000, 2000, 0]
  - type: uuv
    position: [5000, 5000, -200]
```

---

### 7. `scripts/` - 辅助脚本

**`start.sh`** - 一键启动
```bash
#!/bin/bash
echo "启动LLM服务..."
cd python-llm && python llm_service/server.py &
LLM_PID=$!

sleep 3

echo "启动仿真器..."
cd go-backend && go run cmd/simulator/main.go

kill $LLM_PID
```

**`stop.sh`** - 停止服务

**`run_experiment.py`** - 批量运行实验
- 运行多个场景
- 收集指标
- 生成报告

---

## 后续修改建议

### 短期修改（架构确认后）

1. **补充类型定义**
   - 在 `go-backend/internal/types/` 定义核心数据结构
   - 在 `go-backend/pkg/protocol/` 定义跨语言协议

2. **添加配置管理**
   - 实现Go的配置加载
   - 实现Python的配置管理

3. **定义消息格式**
   - 在 `go-backend/internal/communication/message/` 定义消息结构

### 中期修改（开发中）

1. **根据需要调整模块边界**
   - 状态管理是否独立？
   - L1内部模块是否需要重新划分？

2. **添加更多规划算法**
   - 根据实验结果选择最优算法

3. **优化性能**
   - 冲突检测算法优化
   - 消息路由性能优化

### 长期修改（实验后）

1. **根据实验结果调整架构**
   - 哪些模块需要重构？
   - 是否需要新增模块？

2. **文档完善**
   - 补充API文档
   - 补充架构设计文档

---

## 开发顺序建议

### 第一阶段：基础框架（1-2周）
1. 定义类型与消息格式
2. 实现配置加载
3. 实现世界模型框架
4. 实现简单的设备模拟

### 第二阶段：仿真引擎（2-3周）
1. 完善设备模拟
2. 实现传感器模拟
3. 实现消息路由基础
4. 实现简单的可视化

### 第三阶段：L2 Agent（2-3周）
1. 实现设备Agent框架
2. 实现基础路径规划
3. 实现设备控制逻辑

### 第四阶段：L1 Agent（3-4周）
1. 实现状态管理
2. 实现传统规划算法
3. 集成Python LLM服务
4. 实现决策融合

### 第五阶段：冲突检测与管控（2-3周）
1. 实现三维冲突检测
2. 实现冲突消解
3. 实现权限校验机制
4. 实现边界管控

### 第六阶段：自适应重规划（2-3周）
1. 实现故障检测
2. 实现任务重规划
3. 实现故障转移

### 第七阶段：实验验证（3-4周）
1. 设计实验场景
2. 实现指标收集
3. 运行实验
4. 分析结果

---

## 注意事项

1. **不要过度设计** - 先实现MVP，再根据需要优化
2. **保持模块解耦** - 各模块通过接口通信，便于测试
3. **记录实验数据** - 每次运行都要记录日志，便于分析
4. **及时更新文档** - 架构有变化时及时更新本文档

---

## 联系方式

如有问题或建议，请更新本文档！
