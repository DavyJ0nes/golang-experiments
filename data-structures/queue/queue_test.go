package queue

import (
	"reflect"
	"testing"
)

// TestLen checks that we can get the length of the queue
func TestLen(t *testing.T) {
	q := &Queue{
		slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	// Testing Length of Queue is as expected
	if got, want := q.Len(), 9; got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

// TestEnqueue ensures that adding values to the queue works as expected
func TestEnqueue(t *testing.T) {
	q := &Queue{
		slice: []int{1, 2, 3},
	}

	want := []int{1, 2, 3, 4}

	q.Enqueue(4)

	// Testing Length of Queue
	if gotLen, wantLen := q.Len(), 4; gotLen != wantLen {
		t.Errorf("got: %v, want: %v", gotLen, wantLen)
	}

	// Testing Queue Values Are What We Want
	if !reflect.DeepEqual(want, q.slice) {
		t.Errorf("got: %v, want: %v", q.slice, want)
	}
}

// TestDequeue ensures that pulling values from the queue works as expected
func TestDequeue(t *testing.T) {
	testCases := []struct {
		name          string
		startingQueue []int
		want          []int
		err           bool
	}{
		{
			name:          "Simple Test",
			startingQueue: []int{1, 2, 3, 4},
			want:          []int{2, 3, 4},
			err:           false,
		},
		{
			name:          "Empty Queue",
			startingQueue: []int{},
			want:          []int{},
			err:           true,
		},
		{
			name:          "One Item In Queue",
			startingQueue: []int{1},
			want:          []int{},
			err:           false,
		},
		{
			name:          "Items Are Negative",
			startingQueue: []int{-1, -2},
			want:          []int{-2},
			err:           false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := &Queue{
				slice: tc.startingQueue,
			}

			val, err := q.Dequeue()

			if err != nil {
				// skip if we expect an error
				if tc.err {
					return
				}
				t.Fatalf("Unexpected Error: %v", err)
			}

			// Testing Length of Queue
			if gotLen, wantLen := q.Len(), len(tc.want); gotLen != wantLen {
				t.Errorf("got: %v, want: %v", gotLen, wantLen)
			}

			// Testing Pulled Value is the One We Want
			if val != tc.startingQueue[0] {
				t.Errorf("got: %v, want: %v", val, tc.startingQueue[0])
			}

			// Testing Queue Values Are What We Want
			if !reflect.DeepEqual(tc.want, q.slice) {
				t.Errorf("got: %v, want: %v", q.slice, tc.want)
			}
		})
	}
}
