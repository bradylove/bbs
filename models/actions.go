package models

import (
	"errors"
	"fmt"
	"time"
)

const (
	ActionTypeDownload     = "download"
	ActionTypeEmitProgress = "emit_progress"
	ActionTypeRun          = "run"
	ActionTypeUpload       = "upload"
	ActionTypeTimeout      = "timeout"
	ActionTypeTry          = "try"
	ActionTypeParallel     = "parallel"
	ActionTypeSerial       = "serial"
	ActionTypeCodependent  = "codependent"
)

var ErrInvalidActionType = errors.New("invalid action type")

type ActionInterface interface {
	ActionType() string
	Validate() error
}

func (a *Action) Validate() error {
	if a == nil {
		return nil
	}

	if inner := UnwrapAction(a); inner != nil {
		err := inner.Validate()
		if err != nil {
			return err
		}
	} else {
		return ErrInvalidField{"inner-action"}
	}
	return nil
}

func (a *DownloadAction) ActionType() string {
	return ActionTypeDownload
}

func (a DownloadAction) Validate() error {
	var validationError ValidationError

	if a.GetFrom() == "" {
		validationError = validationError.Append(ErrInvalidField{"from"})
	}

	if a.GetTo() == "" {
		validationError = validationError.Append(ErrInvalidField{"to"})
	}

	if a.GetUser() == "" {
		validationError = validationError.Append(ErrInvalidField{"user"})
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *UploadAction) ActionType() string {
	return ActionTypeUpload
}

func (a UploadAction) Validate() error {
	var validationError ValidationError

	if a.GetTo() == "" {
		validationError = validationError.Append(ErrInvalidField{"to"})
	}

	if a.GetFrom() == "" {
		validationError = validationError.Append(ErrInvalidField{"from"})
	}

	if a.GetUser() == "" {
		validationError = validationError.Append(ErrInvalidField{"user"})
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *RunAction) ActionType() string {
	return ActionTypeRun
}

func (a RunAction) Validate() error {
	var validationError ValidationError

	if a.Path == "" {
		validationError = validationError.Append(ErrInvalidField{"path"})
	}

	if a.User == "" {
		validationError = validationError.Append(ErrInvalidField{"user"})
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *TimeoutAction) ActionType() string {
	return ActionTypeTimeout
}

func (a TimeoutAction) Validate() error {
	var validationError ValidationError

	if a.Action == nil {
		validationError = validationError.Append(ErrInvalidField{"action"})
	} else {
		err := UnwrapAction(a.Action).Validate()
		if err != nil {
			validationError = validationError.Append(err)
		}
	}

	if a.GetTimeoutMs() <= 0 {
		validationError = validationError.Append(ErrInvalidField{"timeout_ms"})
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *TryAction) ActionType() string {
	return ActionTypeTry
}

func (a TryAction) Validate() error {
	var validationError ValidationError

	if a.Action == nil {
		validationError = validationError.Append(ErrInvalidField{"action"})
	} else {
		err := UnwrapAction(a.Action).Validate()
		if err != nil {
			validationError = validationError.Append(err)
		}
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *ParallelAction) ActionType() string {
	return ActionTypeParallel
}

func (a ParallelAction) Validate() error {
	var validationError ValidationError

	if a.Actions == nil || len(a.Actions) == 0 {
		validationError = validationError.Append(ErrInvalidField{"actions"})
	} else {
		for index, action := range a.Actions {
			if action == nil {
				errorString := fmt.Sprintf("action at index %d", index)
				validationError = validationError.Append(ErrInvalidField{errorString})
			} else {
				err := UnwrapAction(action).Validate()
				if err != nil {
					validationError = validationError.Append(err)
				}
			}
		}
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *CodependentAction) ActionType() string {
	return ActionTypeCodependent
}

func (a CodependentAction) Validate() error {
	var validationError ValidationError

	if a.Actions == nil || len(a.Actions) == 0 {
		validationError = validationError.Append(ErrInvalidField{"actions"})
	} else {
		for index, action := range a.Actions {
			if action == nil {
				errorString := fmt.Sprintf("action at index %d", index)
				validationError = validationError.Append(ErrInvalidField{errorString})
			} else {
				err := UnwrapAction(action).Validate()
				if err != nil {
					validationError = validationError.Append(err)
				}
			}
		}
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *SerialAction) ActionType() string {
	return ActionTypeSerial
}

func (a SerialAction) Validate() error {
	var validationError ValidationError

	if a.Actions == nil || len(a.Actions) == 0 {
		validationError = validationError.Append(ErrInvalidField{"actions"})
	} else {
		for index, action := range a.Actions {
			if action == nil {
				errorString := fmt.Sprintf("action at index %d", index)
				validationError = validationError.Append(ErrInvalidField{errorString})
			} else {
				err := UnwrapAction(action).Validate()
				if err != nil {
					validationError = validationError.Append(err)
				}
			}
		}
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (a *EmitProgressAction) ActionType() string {
	return ActionTypeEmitProgress
}

func (a EmitProgressAction) Validate() error {
	var validationError ValidationError

	if a.Action == nil {
		validationError = validationError.Append(ErrInvalidField{"action"})
	} else {
		err := UnwrapAction(a.Action).Validate()
		if err != nil {
			validationError = validationError.Append(err)
		}
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func EmitProgressFor(action ActionInterface, startMessage string, successMessage string, failureMessagePrefix string) *EmitProgressAction {
	return &EmitProgressAction{
		Action:               WrapAction(action),
		StartMessage:         startMessage,
		SuccessMessage:       successMessage,
		FailureMessagePrefix: failureMessagePrefix,
	}
}

func Timeout(action ActionInterface, timeout time.Duration) *TimeoutAction {
	return &TimeoutAction{
		Action:    WrapAction(action),
		TimeoutMs: (int64)(timeout / 1000000),
	}
}

func Try(action ActionInterface) *TryAction {
	return &TryAction{Action: WrapAction(action)}
}

func Parallel(actions ...ActionInterface) *ParallelAction {
	return &ParallelAction{Actions: WrapActions(actions)}
}

func Codependent(actions ...ActionInterface) *CodependentAction {
	return &CodependentAction{Actions: WrapActions(actions)}
}

func Serial(actions ...ActionInterface) *SerialAction {
	return &SerialAction{Actions: WrapActions(actions)}
}

func UnwrapAction(action *Action) ActionInterface {
	if action == nil {
		return nil
	}
	a := action.GetValue()
	if a == nil {
		return nil
	}
	return a.(ActionInterface)
}

func WrapActions(actions []ActionInterface) []*Action {
	wrappedActions := make([]*Action, 0, len(actions))
	for _, action := range actions {
		wrappedActions = append(wrappedActions, WrapAction(action))
	}
	return wrappedActions
}

func WrapAction(action ActionInterface) *Action {
	if action == nil {
		return nil
	}
	a := &Action{}
	a.SetValue(action)
	return a
}

func (action *Action) SetDeprecatedTimeoutNs() {
	if action == nil {
		return
	}

	a := action.GetValue()
	switch actionModel := a.(type) {
	case *RunAction, *DownloadAction, *UploadAction:
		return

	case *TimeoutAction:
		timeoutAction := actionModel
		timeoutAction.DeprecatedTimeoutNs = timeoutAction.TimeoutMs * int64(time.Millisecond)

	case *EmitProgressAction:
		actionModel.Action.SetDeprecatedTimeoutNs()

	case *TryAction:
		actionModel.Action.SetDeprecatedTimeoutNs()

	case *ParallelAction:
		for _, subaction := range actionModel.Actions {
			subaction.SetDeprecatedTimeoutNs()
		}

	case *SerialAction:
		for _, subaction := range actionModel.Actions {
			subaction.SetDeprecatedTimeoutNs()
		}

	case *CodependentAction:
		for _, subaction := range actionModel.Actions {
			subaction.SetDeprecatedTimeoutNs()
		}
	}
}

func (action *Action) SetTimeoutMsFromDeprecatedTimeoutNs() {
	if action == nil {
		return
	}

	a := action.GetValue()
	switch actionModel := a.(type) {
	case *RunAction, *DownloadAction, *UploadAction:
		return

	case *TimeoutAction:
		timeoutAction := actionModel
		timeoutAction.TimeoutMs = timeoutAction.DeprecatedTimeoutNs / int64(time.Millisecond)

	case *EmitProgressAction:
		actionModel.Action.SetDeprecatedTimeoutNs()

	case *TryAction:
		actionModel.Action.SetDeprecatedTimeoutNs()

	case *ParallelAction:
		for _, subaction := range actionModel.Actions {
			subaction.SetDeprecatedTimeoutNs()
		}

	case *SerialAction:
		for _, subaction := range actionModel.Actions {
			subaction.SetDeprecatedTimeoutNs()
		}

	case *CodependentAction:
		for _, subaction := range actionModel.Actions {
			subaction.SetDeprecatedTimeoutNs()
		}
	}
}
