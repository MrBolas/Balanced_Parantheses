# Go_Balanced_Parathesis
Algorithm that validates that the parentheses in string are balanced.
- for each open parentheses there is a closing pair
- parentheses pairs should be correctly enclosed and sequenced.

## Classic Interview problem
Implementation in this repository uses a stack to validate the parathesis.
Also supports different types of parantheses.

## Sample Input/Output
| Input        | Output |
| ---          | ---    |
| "()()()"     | true   |
| "()()"       | true   |
| ""           | true   |
| "("          | false  |
| ")"          | false  |
| "()(())"     | true   |
| "()(())()()" | true   |
| "()(()))"    | false  |
| "()()())"    | false  |
| "()[]{}"     | true   |
| "()[()()]"   | true   |
| "()[(){}()]" | true   |
| "()[(){}(})]"| false  |
| "()[(){(})]" | false  |
| "()[()({})]" | true   |
| "(){}[()()]" | true   |
| "(){}[()(])" | false  |
| "{()[()()]}" | true   |

## Interview proposition
Interview proposition for Software engineers.
Structure the problem has a two part problem, being the first part the algorithm itself, and the second part an improvement to allow additional parentheses.
Part 1 should take around 20 min
Part 2 should take around 10 min
### Part 1 
Develop the balanced parenthese algorithm for a given set of example inputs.

| Input        | Output |
| ---          | ---    |
| "()()()"     | true   |
| "()()"       | true   |
| ""           | true   |
| "("          | false  |
| ")"          | false  |
| "()(())"     | true   |
| "()(())()()" | true   |
| "()(()))"    | false  |
| "()()())"    | false  |

### Part 2 
Modify the algorithm to allow parentheses strings with different types of parentheses.
| Input        | Output |
| ---          | ---    |
| "()[]{}"     | true   |
| "()[()()]"   | true   |
| "()[(){}()]" | true   |
| "()[(){}(})]"| false  |
| "()[(){(})]" | false  |
| "()[()({})]" | true   |
| "(){}[()()]" | true   |
| "(){}[()(])" | false  |
| "{()[()()]}" | true   |

## Observations
There is an optimal solution for this problem using a stack. Where the evaluated parantheses, if "closing" is compared with the top of the stack for a match, if "opening" added to the stack.