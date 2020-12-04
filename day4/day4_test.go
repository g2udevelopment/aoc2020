package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirth(t *testing.T) {
	assert.True(t, isValidByr("1921"))
	assert.False(t, isValidByr("1919"))
	assert.False(t, isValidByr("2003"))
}

func TestIssue(t *testing.T) {
	assert.True(t, isValidIyr("2016"))
	assert.False(t, isValidIyr("2009"))
	assert.False(t, isValidIyr("2021"))
}

func TestHeight(t *testing.T) {
	assert.True(t, IsValidHeight("151cm"))
	assert.False(t, IsValidHeight("149cm"))
	assert.False(t, IsValidHeight("15in"))
	assert.True(t, IsValidHeight("60in"))
}

func TestHC(t *testing.T) {
	assert.True(t, IsValidHC("#12ef67"))
	assert.False(t, IsValidHC("12ef67"))
	assert.False(t, IsValidHC("#12ef6y"))
}
