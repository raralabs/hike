### CHANGES ON IDIOM.PEG RULES

---
- Rules added as per following EBNF:
   1. command -->statement-->(statement)*-->TOK_SEMICOLON
   2. statement --> source -->(transform)*-->sink-->TOK_SEMICOLON
   3. Source --> file|fake ....
   4. transform --> Map|Do|Agg ...
   5. sink --> stdout/into/blackhole/....
   
### ADDED FILES
---
on utils 

    - NLP.go
    contains Tokenize and prepareQuery func for cleaning
    the commands before parsing with peg
    
on newPeg
    
    - helper.go
    added RemoveStatement func for removing the comments
    from the command
            