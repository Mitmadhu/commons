package helper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEmpty(t *testing.T) {
	tests := []struct{
		testName string
		data string
		name string
		errs []error
		isErrorExpected bool
		expectedErr []error
	}{
		{
			testName: "empty data",
			data: "",
			name: "username",
			errs: []error{},
			isErrorExpected: true,
			expectedErr: []error{errors.New("username can not be empty")},
		},
		{
			testName: "data",
			data : "ayush",
			name : "username",
			errs: []error{},
			isErrorExpected: false,
			expectedErr: []error{},

		},
	}

	for _, tc := range tests {
		CheckEmpty(tc.data, &tc.errs, tc.name)
		if tc.isErrorExpected{
			assert.Equal(t, 1, len(tc.errs))
			assert.Equal(t, tc.expectedErr[0].Error(), tc.errs[0].Error())
		}else{
			assert.Equal(t, 0, len(tc.errs))
		}

	}
}
