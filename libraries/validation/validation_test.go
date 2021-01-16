package validation

import (
	"fmt"
	"testing"

	"github.com/edstell/lambda/libraries/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMissingString(t *testing.T) {
	t.Parallel()
	msg := &String{}
	err := Validate(msg)
	require.Error(t, err)
	assert.True(t, errors.PrefixMatches(err, fmt.Sprintf("bad request: missing param: field")))
}

func TestMissingMap(t *testing.T) {
	t.Parallel()
	msg := &Map{}
	err := Validate(msg)
	require.Error(t, err)
	assert.True(t, errors.PrefixMatches(err, fmt.Sprintf("bad request: missing param: field")))
}

func TestMissingList(t *testing.T) {
	t.Parallel()
	msg := &List{}
	err := Validate(msg)
	require.Error(t, err)
	assert.True(t, errors.PrefixMatches(err, fmt.Sprintf("bad request: missing param: field")))
}

func TestMissingBytes(t *testing.T) {
	t.Parallel()
	msg := &Bytes{}
	err := Validate(msg)
	require.Error(t, err)
	assert.True(t, errors.PrefixMatches(err, fmt.Sprintf("bad request: missing param: field")))
}

func TestMissingMessage(t *testing.T) {
	t.Parallel()
	msg := &Message{}
	err := Validate(msg)
	require.Error(t, err)
	assert.True(t, errors.PrefixMatches(err, fmt.Sprintf("bad request: missing param: field")))
}

func TestMissingOneof(t *testing.T) {
	t.Parallel()
	msg := &OneOf{}
	err := Validate(msg)
	require.Error(t, err)
	assert.True(t, errors.PrefixMatches(err, fmt.Sprintf("bad request: missing param: field")))
}