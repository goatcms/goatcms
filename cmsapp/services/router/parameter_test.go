package router

import "testing"

func TestParametersInject(t *testing.T) {
	var (
		expectedID = "111"
		err        error
		parameters *Parameters
		deps       struct {
			ID string `parameter:"ID"`
		}
	)
	t.Parallel()
	parameters = NewParameters(map[string]string{
		"ID": expectedID,
	})
	if err = parameters.InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if deps.ID != expectedID {
		t.Errorf("injected id is %v and expected %v", deps.ID, expectedID)
		return
	}
}
