package datelp

import (
        "testing"
        "time"
)

func TestDateContextSynonym(t *testing.T) {
        testCases := []struct{
                context DateContext
                date time.Time
                truncate time.Duration
        }{
                {DateContext{synonym: SYNONYM_TODAY}, time.Now(), time.Second},
                {DateContext{synonym: SYNONYM_YESTERDAY}, time.Now().AddDate(0, 0, -1), time.Second},
                {DateContext{synonym: SYNONYM_TOMORROW}, time.Now().AddDate(0, 0, 1), time.Second},
        }

        for _, tc := range testCases {
                actual, err := tc.context.compileSynonym()
                if err != nil {
                        t.Errorf("Unexpected error returned.")
                }

                actual = actual.Truncate(tc.truncate)
                expected := tc.date.Truncate(tc.truncate)

                if actual != expected {
                        t.Errorf("DateContext (synonym) failed. expected: %s actual: %s", expected, actual)
                }
        }
}

func TestRelativeDateContext(t *testing.T) {
        // this is used to test something like 3rd wednesday of the month
        // first month of the year
        // first day of june
}

func TestDateContext(t *testing.T) {
        testCases := []struct{
                context DateContext
                date [3]int
        }{
                {DateContext{
                        month: MONTH_JUNE,
                        monthday: 25,
                        year: 2015,
                }, [3]int{2015, 6, 25}},
                {DateContext{
                        month: MONTH_JULY, 
                        year: 2015,
                }, [3]int{2015, 7, 1}},
                {DateContext{
                        month: MONTH_DECEMBER,
                        monthday: 20,
                }, [3]int{time.Now().Year(), 12, 20}},
        }

        for _, tc := range testCases {
                actual, err := tc.context.compile()
                expected := time.Date(tc.date[0], time.Month(tc.date[1]), tc.date[2], 0, 0, 0, 0, time.UTC)

                if err != nil {
                        t.Errorf("Unexpected error returned")
                }

                if actual != expected {
                        t.Errorf("DateContext failed. expected: %s actual: %s", expected, actual)
                }
        }
}
