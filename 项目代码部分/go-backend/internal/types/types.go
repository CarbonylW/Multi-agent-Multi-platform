package types

// ========== 基础类型定义 ==========

// Coordinate 3D坐标
type Coordinate struct {
	X float64 `json:"x"` // 米
	Y float64 `json:"y"` // 米
	Z float64 `json:"z"` // 米，正数表示高度，负数表示深度
}

// Pose 姿态
type Pose struct {
	Position Coordinate `json:"position"`
	Yaw      float64    `json:"yaw"`   // 偏航角（度）
	Pitch    float64    `json:"pitch"` // 俯仰角（度）
	Roll     float64    `json:"roll"`  // 翻滚角（度）
}

// Velocity 速度
type Velocity struct {
	X float64 `json:"x"` // m/s
	Y float64 `json:"y"` // m/s
	Z float64 `json:"z"` // m/s
}

// ========== 设备相关类型 ==========

// DeviceType 设备类型
type DeviceType string

const (
	DeviceTypeUAV DeviceType = "uav" // 无人机
	DeviceTypeUSV DeviceType = "usv" // 无人艇
	DeviceTypeUUV DeviceType = "uuv" // 无人潜航器
)

// DeviceStatus 设备状态
type DeviceStatus string

const (
	DeviceStatusOffline   DeviceStatus = "offline"   // 离线
	DeviceStatusStandby   DeviceStatus = "standby"   // 待命
	DeviceStatusMoving    DeviceStatus = "moving"    // 移动中
	DeviceStatusExecuting DeviceStatus = "executing" // 执行任务中
	DeviceStatusFault     DeviceStatus = "fault"     // 故障
)

// DeviceInfo 设备信息
type DeviceInfo struct {
	ID        string      `json:"id"`
	Type      DeviceType  `json:"type"`
	Status    DeviceStatus `json:"status"`
	Pose      Pose        `json:"pose"`
	Velocity  Velocity    `json:"velocity"`
	Battery   float64     `json:"battery"`   // 电量百分比 0-100
	Health    float64     `json:"health"`    // 健康度 0-100
	Timestamp int64       `json:"timestamp"` // Unix时间戳
}

// ========== 任务相关类型 ==========

// TaskType 任务类型
type TaskType string

const (
	TaskTypeSearch      TaskType = "search"      // 搜索
	TaskTypeTrack       TaskType = "track"       // 跟踪
	TaskTypeIntercept   TaskType = "intercept"   // 拦截
	TaskTypePatrol      TaskType = "patrol"      // 巡逻
	TaskTypeReturn      TaskType = "return"      // 返回
)

// TaskStatus 任务状态
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"   // 待执行
	TaskStatusRunning   TaskStatus = "running"   // 执行中
	TaskStatusCompleted TaskStatus = "completed" // 已完成
	TaskStatusFailed    TaskStatus = "failed"    // 失败
	TaskStatusCancelled TaskStatus = "cancelled" // 已取消
)

// Task 任务
type Task struct {
	ID          string     `json:"id"`
	Type        TaskType   `json:"type"`
	Status      TaskStatus `json:"status"`
	AssignedTo  []string   `json:"assigned_to"` // 分配给哪些设备
	Target      Coordinate `json:"target"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"` // 优先级，数字越小优先级越高
	CreatedAt   int64      `json:"created_at"`
	UpdatedAt   int64      `json:"updated_at"`
}

// ========== 目标相关类型 ==========

// TargetType 目标类型
type TargetType string

const (
	TargetTypeUnknown     TargetType = "unknown"
	TargetTypeUAV         TargetType = "uav"
	TargetTypeSurfaceShip TargetType = "surface_ship"
	TargetTypeUUV         TargetType = "uuv"
)

// Target 目标
type Target struct {
	ID          string      `json:"id"`
	Type        TargetType  `json:"type"`
	Pose        Pose        `json:"pose"`
	Velocity    Velocity    `json:"velocity"`
	Confidence  float64     `json:"confidence"` // 置信度 0-1
	IsFriendly  bool        `json:"is_friendly"`
	LastSeen    int64       `json:"last_seen"`
}

// ========== 事件相关类型 ==========

// EventType 事件类型
type EventType string

const (
	EventTypeTargetDetected  EventType = "target_detected"
	EventTypeConflictAlert   EventType = "conflict_alert"
	EventTypeDeviceFault     EventType = "device_fault"
	EventTypeCommunicationLoss EventType = "communication_loss"
	EventTypeTaskCompleted   EventType = "task_completed"
	EventTypeMissionSuccess  EventType = "mission_success"
	EventTypeMissionFailed   EventType = "mission_failed"
)

// Event 事件
type Event struct {
	ID        string      `json:"id"`
	Type      EventType   `json:"type"`
	Severity  string      `json:"severity"` // info, warning, error, critical
	Source    string      `json:"source"`   // 事件源
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

// ========== 区域相关类型 ==========

// ZoneType 区域类型
type ZoneType string

const (
	ZoneTypeNormal     ZoneType = "normal"     // 普通区域
	ZoneTypeRestricted ZoneType = "restricted" // 限制区域
	ZoneTypeForbidden  ZoneType = "forbidden"  // 禁入区域
)

// Zone 区域
type Zone struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Type      ZoneType    `json:"type"`
	Polygon   []Coordinate `json:"polygon"` // 多边形顶点
}
