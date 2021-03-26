#  Valid Parentheses
### Question

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

### Example1 :

Input: s = "()"

Output: true

### Example 2 :
Input: s = "()[]{}"

Output: true

### Example 3 :
Input: s = "(]"

Output: false

### Exmaple 4 :
Input: s = "([)]"

Output: false

### Constraints:

- 1 <= s.length <= 104
- s consists of parentheses only '()[]{}'.

### 思路
先想一下用什麼資料結構做
主角有三對，可分為左邊三個，右邊三個 -> 用map，因為key value

考慮情況：
1. 空字串，回傳true
2. 字串個數是奇數，回傳false
3. 先用 map的key, value去存那六個主角
4. 再輪詢 s，設想情況去使用stack的 LIFO
5. 如果s的是左邊三個之一，先將s那左邊三個在map中所對應的value（右邊三個之一）存在stack中; 如果s都沒有左邊三個之一，可能存完了或一開始就沒有，繼續輪詢的情況如下:
- 先檢查stack是否有東西，沒有則回傳
false
- 有東西，就開始用map中左邊三個所對應的value，來跟比較stack最頂端的元素，看是否匹配，如果不匹配，則回傳false，有沒有發現就是用右邊的來找看看有沒有對應的左邊^^
- 有東西，而且跟stack的頂端元素匹配，那就換下一個stack 元素，其實就是stack的pop
- 最後要確定 stack中的東西都pop出來了，才能返回 true，就是len(stack)== 0
- 不然.... 如果s是兩個左括號，那就會都留在stack，都沒pop出來，但是還返回true

#### Topics : hash table, stack, 
#### Difficulty : easy




