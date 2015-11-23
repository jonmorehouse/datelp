package datelp

import (
        "testing"
)

func TestMaxInt(t *testing.T) {
        testCases := []struct{
                args [2]int
                expected int
        } {
                {[2]int{1, 2}, 2},
                {[2]int{3, 2}, 3},
                {[2]int{2, 2}, 2},
        }

        for _, tc := range testCases {
                res := MaxInt(tc.args[0], tc.args[1])
                if res != tc.expected {
                        t.Fail()
                }
        }
}

