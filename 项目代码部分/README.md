# 大模型驱动的多层多Agent异构无人平台协同控制与任务调度系统

## 项目简介

本项目实现了一个三层多Agent架构的异构无人平台协同作战系统，解决以下四大核心问题：

1. **混合决策架构**：LLM+RAG+传统算法的结合，兼顾智能性与实时性
2. **三维冲突检测**：海空潜立体空间的工程化冲突检测与秒级消解
3. **多Agent管控**：严格层级的权限与边界管控，防止角色混淆
4. **自适应重规划**：支持故障转移与广域扩展的任务重规划

## 技术栈

- **后端核心**：Go (高性能、并发好)
- **LLM层**：Python (生态丰富、调用方便)
- **通信**：gRPC / HTTP
- **可视化**：ECharts / Three.js (可选)

## 项目结构

```
项目代码部分/
├── docs/                    # 文档
│   ├── architecture.md      # 架构设计
│   ├── protocol.md          # 通信协议
│   └── api.md               # API文档
│
├── go-backend/              # Go后端核心
│   ├── cmd/
│   │   └── simulator/       # 仿真器入口
│   │       └── main.go
│   │
│   ├── internal/
│   │   ├── agent/           # Agent实现
│   │   │   ├── l1/          # L1全局统筹Agent
│   │   │   ├── l2/          # L2设备Agent
│   │   │   └── statemgr/    # 状态管理模块
│   │   │
│   │   ├── simulation/      # 仿真引擎
│   │   │   ├── world/       # 世界模型
│   │   │   ├── device/      # 设备模拟
│   │   │   └── sensor/      # 传感器模拟
│   │   │
│   │   ├── planning/        # 规划算法
│   │   │   ├── path/        # 路径规划
│   │   │   ├── conflict/    # 冲突检测
│   │   │   └── assignment/  # 任务分配
│   │   │
│   │   ├── communication/   # 通信模块
│   │   │   ├── message/     # 消息定义
│   │   │   └── router/      # 消息路由
│   │   │
│   │   └── types/           # 公共类型定义
│   │
│   ├── pkg/                   # 可导出的包
│   │   └── protocol/         # 协议定义
│   │
│   └── go.mod
│   └── go.sum
│
├── python-llm/              # Python LLM层
│   ├── llm_service/
│   │   ├── __init__.py
│   │   ├── server.py        # gRPC/HTTP服务
│   │   ├── agent/
│   │   │   ├── l1_agent.py  # L1的LLM部分
│   │   │   └── prompts/     # Prompt模板
│   │   ├── rag/
│   │   │   ├── knowledge_base.py
│   │   │   └── documents/   # 知识库文档
│   │   └── client.py        # Go调用的客户端
│   │
│   ├── tests/
│   ├── requirements.txt
│   └── config.py
│
├── web-frontend/            # 可视化前端（可选）
│   ├── index.html
│   ├── app.js
│   └── style.css
│
├── configs/                 # 配置文件
│   ├── simulator.yaml
│   ├── llm.yaml
│   └── devices.yaml
│
├── scenarios/               # 测试场景
│   ├── scenario1.yaml       # 典型入侵场景
│   ├── scenario2.yaml       # 故障场景
│   └── scenario3.yaml       # 复杂场景
│
└── scripts/                 # 辅助脚本
    ├── start.sh
    ├── stop.sh
    └── run_experiment.py
```

## 快速开始

### 前置要求

- Go 1.21+
- Python 3.10+
- 大模型API密钥（OpenAI / Anthropic / 文心一言）

### 运行仿真

```bash
# 1. 启动Python LLM服务
cd python-llm
pip install -r requirements.txt
python llm_service/server.py

# 2. 启动Go仿真器
cd go-backend
go run cmd/simulator/main.go --config ../configs/simulator.yaml
```

## 核心模块

### 三层Agent架构

```
L1: 全局统筹Agent（混合决策）
    ↓
L2: 设备类型Agent（UAV/USV/UUV）
    ↓
L3: 实际物理设备
```

### 通信规则

- ✅ L1 ↔ L2 双向通信
- ✅ L2 ↔ 对应设备 双向通信
- ❌ L1 不能直接控制设备
- ❌ L2 之间不能直接通信
- ✅ 设备状态同时上报 L1+L2

## 论文创新点对应

| 创新点 | 对应代码模块 |
|--------|-------------|
| 混合决策架构 | `go-backend/internal/agent/l1/` + `python-llm/llm_service/agent/` |
| 三维冲突检测 | `go-backend/internal/planning/conflict/` |
| 多Agent边界管控 | `go-backend/internal/communication/router/` |
| 自适应任务重规划 | `go-backend/internal/planning/assignment/` |

## 开发计划

- [ ] 阶段1：定义类型与协议
- [ ] 阶段2：实现仿真引擎
- [ ] 阶段3：实现L2设备Agent
- [ ] 阶段4：实现L1混合决策架构
- [ ] 阶段5：实现冲突检测与消解
- [ ] 阶段6：实现边界管控机制
- [ ] 阶段7：实现自适应重规划
- [ ] 阶段8：实验验证与优化

## 许可证

本项目仅供学术研究使用。

## 联系方式

如有问题，请联系项目维护者。
