package util_test

import (
	"goPasswordGenerator/model"
	"goPasswordGenerator/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndCompareHash(t *testing.T) {
	p := "hola"
	t.Run("Testing func to generate hash and compare hash and password", func(t *testing.T) {
		h, err := util.GenerateHashPassword(p)
		assert.NoError(t, err)
		assert.NotEqual(t, h, "Esto debe ser diferente")

		assert.True(t, util.CompareHashPassword(h, p))
		assert.False(t, util.CompareHashPassword(h, "Adios"))
	})
}

func TestGenerateRandomPassword(t *testing.T) {
	cases := []struct {
		length int
		useNumbers, useSymbols, returnError bool
	} {
		{length: 8, useNumbers: true, useSymbols: true, returnError: false},
		{length: 0, useNumbers: true, useSymbols: true, returnError: true},
	}

	t.Run("Testing generate random password", func(t *testing.T) {
		for _, c := range cases {
			gp, err := util.GenerateRandomPassword(c.length, c.useNumbers, c.useSymbols)
			if !c.returnError {
				otherGp, _ := util.GenerateRandomPassword(c.length, c.useNumbers, c.useSymbols)
				assert.NoError(t, err)
				assert.Equal(t, c.length, len(gp))
				assert.NotEqual(t, gp, otherGp)
				
			} else {
				assert.Error(t, err)
				assert.Equal(t, 0, len(gp))
				assert.Equal(t, "", gp)
			}
		}
	})
}

func TestTemplatesFunction(t *testing.T) {
	cases := []struct {
		template string
		pArray []*model.Password
	} {
		{
			template: `{{.ID}}) {{.Label}} - {{.Password}}`,
			pArray: []*model.Password{{ID: 1, Label: "Google", Password: "Test", CreatedAt: time.Now()}},
		},
	}

	t.Run("Testing templates return value", func(t *testing.T) {
		for _, c := range cases {
			ts, err := util.PrintPasswordsTemplate(c.pArray, c.template)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, "1) Google - Test\n", ts)
			}
		}
	})
}