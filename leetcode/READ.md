# 114. Flatten Binary Tree to Linked List

### Question
Given a root of binary tree, flatten the tree into a "linked list":
- The "linked list" should use the same "TreeNode" class where the right child pointer points the next node in the list and the left child pointer is alway null.

- The "linked list" should be in the same order as preorder traversal of the binary tree.

### Example1 :
Input : root = [1,2,5,3,4,null,6]
Output : [1,null,2,null,3,null,4,null,5,null,6]

### Example 2 :
Input: root = []
Output: []
### Example 3 :
Input: root = [0]
Output: [0]

### Constraints:
- The number of nodes in the tree is in the range [0, 2000].
- -100 <= Node.val <= 100
### 思路
利用遞迴函數, 把左右子樹拉平
左子樹作為root的右子樹, 原本的左子樹設為nil
再將右子樹接到當前右子樹的下面

### Hint


#### Topics : Array
#### Difficulty : Easy



