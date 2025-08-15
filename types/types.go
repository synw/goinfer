package types

type GoInferConf struct {
	ModelsDir  string
	TasksDir   string
	Origins    []string
	ApiKey     string
	OpenAiConf OpenAiConf
}

type InferenceParams struct {
	Stream            bool     `json:"stream,omitempty" yaml:"stream,omitempty"`
	Threads           int      `json:"threads,omitempty" yaml:"threads,omitempty"`
	NPredict          int      `json:"n_predict,omitempty" yaml:"n_predict,omitempty"`
	TopK              int      `json:"top_k,omitempty" yaml:"top_k,omitempty"`
	TopP              float32  `json:"top_p,omitempty" yaml:"top_p,omitempty"`
	Temperature       float32  `json:"temperature,omitempty" yaml:"temperature,omitempty"`
	FrequencyPenalty  float32  `json:"frequency_penalty,omitempty" yaml:"frequency_penalty,omitempty"`
	PresencePenalty   float32  `json:"presence_penalty,omitempty" yaml:"presence_penalty,omitempty"`
	RepeatPenalty     float32  `json:"repeat_penalty,omitempty" yaml:"repeat_penalty,omitempty"`
	TailFreeSamplingZ float32  `json:"tfs_z,omitempty" yaml:"tfs_z,omitempty"`
	StopPrompts       []string `json:"stop,omitempty" yaml:"stop,omitempty"`
}

type InferenceStats struct {
	ThinkingTime       float64 `json:"thinkingTime"`
	ThinkingTimeFormat string  `json:"thinkingTimeFormat"`
	EmitTime           float64 `json:"emitTime"`
	EmitTimeFormat     string  `json:"emitTimeFormat"`
	TotalTime          float64 `json:"totalTime"`
	TotalTimeFormat    string  `json:"totalTimeFormat"`
	TokensPerSecond    float64 `json:"tokensPerSecond"`
	TotalTokens        int     `json:"totalTokens"`
}

type InferenceResult struct {
	Text  string         `json:"text"`
	Stats InferenceStats `json:"stats"`
}

type Task struct {
	Name        string          `json:"name" yaml:"name"`
	Template    string          `json:"template" yaml:"template"`
	ModelConf   ModelConf       `json:"modelConf" yaml:"modelConf"`
	InferParams InferenceParams `json:"inferParams" yaml:"inferParams"`
}

type ModelConf struct {
	Name      string `json:"name" yaml:"name"`
	Ctx       int    `json:"ctx,omitempty" yaml:"ctx,omitempty"`
	GPULayers int    `json:"gpu_layers,omitempty" yaml:"gpu_layers,omitempty"`
}

type TemplateInfo struct {
	Name string `json:"name" yaml:"name"`
	Ctx  int    `json:"ctx" yaml:"ctx"`
}

type MsgType string

const (
	TokenMsgType  MsgType = "token"
	SystemMsgType MsgType = "system"
	ErrorMsgType  MsgType = "error"
)

type StreamedMessage struct {
	Content string                 `json:"content"`
	Num     int                    `json:"num"` // number of tokens
	MsgType MsgType                `json:"msg_type"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type ApiType string

const (
	Llama  ApiType = "llama"
	OpenAi ApiType = "openai"
)
