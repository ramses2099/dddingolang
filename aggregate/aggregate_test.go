package aggregate

import "testing"

func TestCustomer_NewCustomer(t *testing.T) {

	//Build our needed testcase data struct
	type testCase struct {
		test       string
		name       string
		expectdErr error
	}

	//create new test cases
	testCases := []testCase{
		{
			test:       "Empty Name validation",
			name:       "",
			expectdErr: ErrInvalidPerson,
		}, {
			test:       "valid name",
			name:       "Juan perez",
			expectdErr: nil,
		},
	}

	for _, tc := range testCases {
		//run the test
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)
			if err != tc.expectdErr {
				t.Errorf("Expected error %v, got %v", tc.expectdErr, err)
			}
		})
	}
}
