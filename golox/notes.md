## Parsing Expressions

parsing techniques (future exploration):
    - combinations of L/R (LL(k), LR(1), LALR)
    - parser combinators
    - earley parsers
    - shunting yard alg
    - packrat parsing
    - recursive descent (golox)

looking ahead to determine how to parse -> *predictive parser*

given valid sequence, consume tokens to produce AST
given invalid, detect and errors
    - detect and report
    - no crashing or hanging
    - be fast
    - report each distinct error
    - minimize cascading errors

panic recovery
    - given an error, get current state and following sequence such that next token matches
      current rule, aka *synchronization*
        - pick rule in grammar as synchronization point
            - typically between statements
        - parse jumps out of nested productions until that point
        - discard tokens until successful match for sync rule
