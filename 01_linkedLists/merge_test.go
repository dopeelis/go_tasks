package main

import "testing"

func Test(t *testing.T) {
	testTable := []struct {
		l1  *ListNode
		l2  *ListNode
		res *ListNode
	}{
		{
			l1: &ListNode{
				Val: 4, Next: &ListNode{
					Val: 5, Next: nil}},
			l2: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 3, Next: nil}},
			res: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 4, Next: &ListNode{
							Val: 5, Next: nil}}}},
		},
		{
			l1:  nil,
			l2:  nil,
			res: nil,
		},
		{
			l1: nil,
			l2: &ListNode{
				Val: 4, Next: &ListNode{
					Val: 5, Next: nil}},
			res: &ListNode{
				Val: 4, Next: &ListNode{
					Val: 5, Next: nil}},
		},
		{
			l1: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 5, Next: &ListNode{
						Val: 9, Next: nil}}},
			l2: nil,
			res: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 5, Next: &ListNode{
						Val: 9, Next: nil}}},
		},
		{
			l1: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 5, Next: &ListNode{
						Val: 6, Next: nil}}},
			l2: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 5, Next: nil}}},
			res: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 5, Next: &ListNode{
								Val: 5, Next: &ListNode{
									Val: 6, Next: nil}}}}}},
		},
	}

	for _, testCase := range testTable {
		result := mergeTwoLists(testCase.l1, testCase.l2)

		if result != testCase.res {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.res, result)
		}
	}
}
