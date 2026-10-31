20. Valid Parentheses
Easy
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Для строки s, содержащей только символы '(', ')', '{', '}', '[' и ']', определите, является ли входная строка корректной.

Входная строка является корректной, если:

Открывающие скобки должны закрываться скобками того же типа.
Открывающие скобки должны закрываться в правильном порядке.
Каждой закрывающей скобке соответствует открывающая скобка того же типа.
 
Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Example 4:
Input: s = "([])"
Output: true

Example 5:
Input: s = "([)]"
Output: false

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.