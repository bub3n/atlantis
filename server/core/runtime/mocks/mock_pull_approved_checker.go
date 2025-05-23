// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: PullApprovedChecker)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockPullApprovedChecker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPullApprovedChecker(options ...pegomock.Option) *MockPullApprovedChecker {
	mock := &MockPullApprovedChecker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPullApprovedChecker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPullApprovedChecker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPullApprovedChecker) PullIsApproved(logger logging.SimpleLogging, baseRepo models.Repo, pull models.PullRequest) (models.ApprovalStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPullApprovedChecker().")
	}
	_params := []pegomock.Param{logger, baseRepo, pull}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsApproved", _params, []reflect.Type{reflect.TypeOf((*models.ApprovalStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 models.ApprovalStatus
	var _ret1 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(models.ApprovalStatus)
		}
		if _result[1] != nil {
			_ret1 = _result[1].(error)
		}
	}
	return _ret0, _ret1
}

func (mock *MockPullApprovedChecker) VerifyWasCalledOnce() *VerifierMockPullApprovedChecker {
	return &VerifierMockPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPullApprovedChecker) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPullApprovedChecker {
	return &VerifierMockPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPullApprovedChecker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPullApprovedChecker {
	return &VerifierMockPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPullApprovedChecker) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPullApprovedChecker {
	return &VerifierMockPullApprovedChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPullApprovedChecker struct {
	mock                   *MockPullApprovedChecker
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPullApprovedChecker) PullIsApproved(logger logging.SimpleLogging, baseRepo models.Repo, pull models.PullRequest) *MockPullApprovedChecker_PullIsApproved_OngoingVerification {
	_params := []pegomock.Param{logger, baseRepo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsApproved", _params, verifier.timeout)
	return &MockPullApprovedChecker_PullIsApproved_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPullApprovedChecker_PullIsApproved_OngoingVerification struct {
	mock              *MockPullApprovedChecker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPullApprovedChecker_PullIsApproved_OngoingVerification) GetCapturedArguments() (logging.SimpleLogging, models.Repo, models.PullRequest) {
	logger, baseRepo, pull := c.GetAllCapturedArguments()
	return logger[len(logger)-1], baseRepo[len(baseRepo)-1], pull[len(pull)-1]
}

func (c *MockPullApprovedChecker_PullIsApproved_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.SimpleLogging, _param1 []models.Repo, _param2 []models.PullRequest) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]logging.SimpleLogging, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(logging.SimpleLogging)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]models.Repo, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(models.Repo)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]models.PullRequest, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(models.PullRequest)
			}
		}
	}
	return
}
