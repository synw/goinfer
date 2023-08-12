package state

import (
	llama "github.com/go-skynet/go-llama.cpp"

	"github.com/synw/goinfer/types"
)

// the language model instance
var Lm *llama.LLama

// models state
var ModelsDir = ""
var IsModelLoaded = false
var LoadedModel = ""
var ModelConf = types.ModelConf{
	Ctx: 1024,
}

// inference state
var ContinueInferingController = true
var IsInfering = false

// app state
var IsVerbose = false
var UseWs = true

// tasks
var TasksDir = "./tasks"

// OpenAi api
var OpenAiConf types.OpenAiConf
