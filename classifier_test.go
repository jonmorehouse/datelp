package datelp

import (
        "testing"
        "time"
)


func TestClassifierOffsetContext(t *testing.T) {
        testCases := []struct{
                input string
                offset OffsetContext 
        } {
                {"3 weeks ago", OffsetContext{
                        interval: INTERVAL_WEEK,
                        direction: DIRECTION_LEFT,
                        count: 3,
                        value: 0,
                        size: 3,
                }},
                {"last wednesday", OffsetContext{
                        interval: INTERVAL_WEEKDAY,
                        direction: DIRECTION_LEFT,
                        count: 1,
                        value: WEEKDAY_WEDNESDAY,
                        size: 2,
                }},
                {"2 tuesdays ago", OffsetContext{
                        interval: INTERVAL_WEEKDAY,
                        direction: DIRECTION_LEFT,
                        count: 2,
                        value: WEEKDAY_TUESDAY,
                        size: 3,
                }},
                {"last month", OffsetContext{
                        interval: INTERVAL_MONTH,
                        direction: DIRECTION_LEFT,
                        count: 1,
                        value: 0,
                        size: 2,
                }},
                {"next june", OffsetContext{
                        interval: INTERVAL_MONTH,
                        direction: DIRECTION_RIGHT,
                        count: 1,
                        value: MONTH_JUNE,
                        size: 2,
                }},
                {"last july", OffsetContext{
                        interval: INTERVAL_MONTH,
                        direction: DIRECTION_LEFT,
                        count: 1,
                        value: MONTH_JULY,
                        size: 2,
                }},
                {"2 weeks from today", OffsetContext{
                        interval: INTERVAL_WEEK,
                        direction: DIRECTION_RIGHT,
                        count: 2,
                        size: 3,
                }},
                {"tuesday", OffsetContext{
                        interval: INTERVAL_WEEKDAY,
                        direction: DIRECTION_CURRENT,
                        value: WEEKDAY_TUESDAY,
                        count: 1,
                        size: 1,
                }},
        }

        for _, tc := range testCases {
                classifier := Classifier{}
                iterator := newWordIterator(tc.input)
                classifier.buildContexts(iterator)

                // since buildContexts is internal we build the context and
                // then fetch a pointer to the classifier.offset attribute
                offset := classifier.offset
                tests := []struct{
                        expected int
                        actual int
                        val string
                }{
                        {tc.offset.interval, offset.interval, "interval"},
                        {tc.offset.direction, offset.direction, "direction"},
                        {tc.offset.count, offset.count, "count"},
                        {tc.offset.value, offset.value, "value"},
                        {tc.offset.truncate, offset.truncate, "truncate"},
                        {tc.offset.size, offset.size, "size"},
                }

                for _, test := range tests {
                        if test.actual != test.expected {
                                t.Errorf("%s failed. expected: %d actual: %d", test.val, test.expected, test.actual)
                        }
                }
        }
}

func TestClassifierDateContext(t *testing.T) {
        testCases := []struct{
                input string
                date DateContext
        } {
                {"june 1st 2015", DateContext{
                        size: 3,
                        monthday: 1,
                        month: MONTH_JUNE,
                        weekday: -1,
                        year: 2015,
                        synonym: -1,
                }},
                {"tomorrow", DateContext{
                        size: 1,
                        weekday: -1,
                        synonym: SYNONYM_TOMORROW,
                }},
                {"day after tomorrow", DateContext{
                        size: 1,
                        weekday: -1,
                        synonym: SYNONYM_TOMORROW,
                }},
        }

        for _, tc := range testCases {
                c := Classifier{}
                i := newWordIterator(tc.input)
                c.buildContexts(i)
                date := c.date

                assertions := []struct {
                        expected int
                        actual int
                        message string
                } {
                        {tc.date.size, date.size, "size"},
                        {tc.date.synonym, date.synonym, "synonym"},
                        {tc.date.weekday, date.weekday, "weekday"},//TODO remove this
                        {tc.date.month, date.month, "month"},
                        {tc.date.monthday, date.monthday, "monthday"},
                        {tc.date.year, date.year, "year"},
                }

                for _, assertion := range assertions {
                        if assertion.actual != assertion.expected {
                                t.Errorf("Classifier built context with wrong %s. expected: %d. actual: %d.", 
                                          assertion.message, assertion.expected, assertion.actual)
                        }
                }
        }
}

func TestClassifierEndToEnd(t *testing.T) {
        testCases := []struct{
                input string
                expected time.Time
        }{
                {"june 2nd", time.Date(time.Now().Year(), time.June, 2, 0, 0, 0, 0, time.UTC)},
                {"day after tomorrow", time.Now().AddDate(0, 0, 2)},
                {"day before yesterday", time.Now().AddDate(0, 0, -2)},
                {"tuesday", time.Now().AddDate(0, 0, 2)},
                {"next tuesday", time.Now().AddDate(0, 0, 9)},
                {"this sunday", time.Now()},
                {"next sunday", time.Now().AddDate(0, 0, 7)},
                {"today", time.Now()},
        }

        for _, tc := range testCases {
                c := NewClassifier()
                i := newWordIterator(tc.input)
                res, _ := c.Parse(i)

                if res == nil {
                        t.Fatalf("Result object was nil.")
                }

                if tc.expected.Truncate(time.Hour*24) != res.Date.Truncate(time.Hour*24) {
                        t.Fatalf("Did not convert \"%s\". Expected: %s Actual: %s", tc.input, tc.expected, res.Date)
                }
        }
}

