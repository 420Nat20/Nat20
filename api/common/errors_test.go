package common_test

import (
	"github.com/420Nat20/Nat20/nat-20/common"
	"github.com/stretchr/testify/assert"
	"gopkg.in/errgo.v2/fmt/errors"
	"testing"
)

func TestContext(t *testing.T) {

	err := common.BadRequest.New("an_error")
	errWithContext := common.AddErrorContext(err, "a_field", "the field is empty")

	expectedContext := map[string]string{"field": "a_field", "message": "the field is empty"}

	assert.Equal(t, common.BadRequest, common.GetType(errWithContext))
	assert.Equal(t, expectedContext, common.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestContextInNoTypeError(t *testing.T) {
	err := errors.New("a custom error")

	errWithContext := common.AddErrorContext(err, "a_field", "the field is empty")

	expectedContext := map[string]string{"field": "a_field", "message": "the field is empty"}

	assert.Equal(t, common.NoType, common.GetType(errWithContext))
	assert.Equal(t, expectedContext, common.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestWrapf(t *testing.T) {
	err := errors.New("an_error")
	wrappedError := common.BadRequest.Wrapf(err, "error %s", "1")

	assert.Equal(t, common.BadRequest, common.GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error")
}

func TestWrapfInNoTypeError(t *testing.T) {
	err := errors.Newf("an_error %s", "2")
	wrappedError := common.Wrapf(err, "error %s", "1")

	assert.Equal(t, common.NoType, common.GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error 2")
}
