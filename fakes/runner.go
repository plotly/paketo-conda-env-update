package fakes

import "sync"

type Runner struct {
	ExecuteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			CondaEnvPath   string
			CondaCachePath string
			WorkingDir     string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string, string) error
	}
}

func (f *Runner) Execute(param1 string, param2 string, param3 string) error {
	f.ExecuteCall.Lock()
	defer f.ExecuteCall.Unlock()
	f.ExecuteCall.CallCount++
	f.ExecuteCall.Receives.CondaEnvPath = param1
	f.ExecuteCall.Receives.CondaCachePath = param2
	f.ExecuteCall.Receives.WorkingDir = param3
	if f.ExecuteCall.Stub != nil {
		return f.ExecuteCall.Stub(param1, param2, param3)
	}
	return f.ExecuteCall.Returns.Error
}
