from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        result = ListNode()
        currResult = result
        while list1 is not None and list2 is not None:
            if list1.val <= list2.val:
                currResult.next = list1
                list1 = list1.next
            else:
                currResult.next = list2
                list2 = list2.next
            currResult = currResult.next

        if list1 is not None:
            currResult.next = list1
        else:
            currResult.next = list2

        return result.next

if __name__ == "__main__":
    list1 = ListNode(1, ListNode(2, ListNode(4)))
    list2 = ListNode(1, ListNode(3, ListNode(4)))
    result = ListNode(1, ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(4))))))
    got = Solution().mergeTwoLists(list1, list2)
    if got != result:
        print("got error, want: {}, got: {}".format(result, got))
