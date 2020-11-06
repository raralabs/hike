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
    
New Query Structure
---
####command


    `Source | one or more Transform | Sink;
     #This is a comment
     Source  | one or more Transfrom | Sink;;`
     #Double semicolon denotes the end of command
        
        
source:
--- 

file(xyz.csv), fake(10), branch(some Identifier)


Transform:
---

Aggregator, DoFunction, Map


Sink:
---

Stdout,Into,BlackHole,Plot


steps:
    
    
    1. Preprocessor Initialization
    
    The command may consist of unwanted white spaces ,comments
    etc. This needs to be filtered before parsing it with the peg.
    Thus, to make command suitable for parsing, the preProcessor 
    from the LanguageProcessor package is initializd with the 
    command to be parsed and then the prepareQuery func of the 
    preprocessor is called to get the cleaned command.
    
    
    2. Parsing with Peg
    
    The cleaned command is passed to the peg parser to get the 
    stages from the command like source,transform and sink to pre
    pare the pipeline. For each statement the stages are grouped in
    the slice.
    eg: [[statement1 stages][statement2 stages][ statement3 stages]]
  
    3. Abstract Tree Initialization
    
    From the parsed stages Abstract tree is built. For that 
    NewATBuilder() is initialized. 
    
    4. Build Abstract Tree
    Then its BuildAT func is called to create the abstract tree.