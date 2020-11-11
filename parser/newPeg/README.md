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
    
    
    5. NetBuilder
    newNetBuilder() initializes the NetBuilder. The Build reciever 
    Build() builds the pipeline as per the Abstract tree
    
## FUNCTIONS
---
####Current Supported Functions

#####Source functions

    File Reader 
        - supported file format: csv only
        - syntax : file(FileName.csv)
        - usecase: 
            file(userinfo.csv) | filter(age<30) | stdout();;
        
    Fake
        - generates fake data
        - syntax : fake(integer)
        - usecase:
            fake(10) | filter(age<30) | stdout();;    

#####Transform functions
    
    Aggregator Functions
        1. count
            - counts the number of particular data in the stream
            - syntax alias = count() 
            - usecase:
                fake(10) | filter(age<30) | s1=into();
                s1 | count_of_age_less_than_30 = count() |stdout();; 
        2. max
            - 
        3. min
        4. sum
        5. variance
        6. avg  
    
    Do Functions
        1. filter
            -filters the data based on the condition provided.
            - syntax 
                filter(age<30) 
            - usecase:
                file(userinfo.csv) | filter(age<30) | stdout();;     
        
        2. select
            - select selects one or more fields form the stream
            - syntax 
                select(fieldName1,fieldName2)
            - usecase:
                source | select (field1,field2) | s1 = into();;
        
        3. sort
            - sorts the data
            - syntax
                sort() by alias
            - usecase:
                source | sort() by field1 | sink               
        
        4. batch
            -
            - syntax
                batch(integer)
            - usecase:
        
        5. pick
            - picks the given number of data
            - syntax
                pick(int)
            - usecase:
                source | pick(10) |sink
        
        6. enrich 
            - enriches the data
            - syntax
                newField = condtion       
            usecase:
                - used with map
                                  
                     - s1| map twice_age = 2*age |stdout();
                               -----------------
                     underlined one is the enrich syntax
                                            
    Map functions             
        map
            - maps the given condition to the field. Basically it 
            enriches the data
            - takes the EnrichDo function 
            - syntax
                map Enrich function
            - usecase:
                - s1| map twice_age = 2*age |stdout();                               
            
##### Sink Functions
    (currently supported)
    Stdout
        - prints the stream into the console
        - syntax : stdout()
        - usecase:
            file(userinfo.csv)| zero or more transform functions | stdout();;
        
    Plot
        - It is used to plot the data.
        - syntax: plot(widget)
        - usecase:
            plot(BarWidget)
    
    Blackhole
        
    
                                   