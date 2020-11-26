# Hike

A CLI for data exploration and visualization
### CHANGES ON IDIOM.PEG RULES

---
- Rules added as per following EBNF:
   1. command -->statement-->(statement)*-->TOK_SEMICOLON
   2. statement --> source -->(transform)*-->sink-->TOK_SEMICOLON
   3. Source --> file|fake ....
   4. transform --> Map|Do|Agg ...
   5. sink --> stdout/into/blackhole/....
   
   
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
    
    NOTE:A FilterMulti defines a filter that consists of either a single factor or
           multiple factors separated by secondary operators.
           FilterMulti consist of FilterFactor and restFilterFactors separated by the secondary filter operator
           secondary filter operator:
                - and operator
                - or operator
            primary filter operator:
                - equal to operator
                - not equal to operator
                - greather than
                - less than
                - greater or equal to
                - less or equal to
            A Factor is either a primary filter or a multi filter enclosed by parenthesis
            A FilterPrimary defines a filter that checks if the value defined by the
            variable is according to the primary operator.          
            
            ? - in any arg denotes its a optional parameter
            
    Aggregator Functions
        NOTE:A FilterMulti defines a filter that consists of either a single factor or
        multiple factors separated by secondary operators.
        1. count
            - counts the number of particular data in the stream
            - -args type:
                   - fieldName string
                   - FilterMulti expression (eg: age<30)
            
            - syntax alias = count()
                     count() #alias is optional
                     count(condition) # condition is expression or filter  
                     expression looks like
                        count(expression)
                        or
                        #case of multiple condition
                        count((expression) sec-operator (expression))
                        expression:
                            fieldName primary-operator value 
                            (string)    (operators)    (integer tested)
                        
                               
            - usecase:
                fake(10) | filter(age<30) | s1=into();
                s1 | agg count_of_age_less_than_30 = count() |stdout();;
                or
                #use of filter
                s1 | agg count_of_age = count(age<10) | stdout();;
                #use of multiple filter
                s1 | agg count_of_age = count(age<10 and age>5) | stdout();; 
                 
        2. max
            - calculates the max value
            -args type:
                - fieldName string
                - FilterMulti expression (eg: age<30) ?
            - syntax alias = agg max(fieldName)
                     alias = agg max(fieldName,FilterMulti)
            - usecase:  
                fake(10) | select(age) | s1=into();
                s1 | agg max(age) | stdout();;
                #one with filter
                s1 | agg max(age,age>50) |stdout();;
                
        3. min
            syntax are similar to max
            
        4. sum
            - perform the sum of the field values
            - args type:
                fieldName string
                filterMulti expression   ?
            - syntax alias = agg sum(fieldName)
                    alias = agg sum(fieldName, filterMulti?)
            -usecase:
                fake(10) | select(salary,empid) | s1=into();
                s1 | agg Total_salary = sum(salary) | stdout();;
                #one with filter
                s1 | agg Total_salary = sum(salary,salary>30000) |stdout();;   
                    
        5. variance
            - calculates the variance
            - args
                fieldName string
                FilterMulti expression ?
            - syntax
                var(fieldName,filterMulti)    
            - usecase
                same sum  
           
        6. average
            - calculates average
            - args
                fieldName string
            - syntax
                avg(fieldName,filterMulti)
            - usecase:
                same as above
        7. mode
            - calculates mode
            - args
                same as avg
            - syntax
                - mode(fieldName,filterMulti)
            - usecase:
                - same as average     
        
        8. quantile
            - calculates quantile
            - args
                 fieldName string
                 weight    string         ?
                 qth quartile Number        
                 FilterMulti expression   ?
            - syntax
                 quantile(quartile_of_data = quartile(age))
                 #one with filter
                 quantile(quartile(age,3,age>30)) #gets the 3rd quantile 
            - usecase:
                 similar to above sum,max etc
        
        9. quartile
            - calculates quartile
            - args
                 fieldName string
                 weight    string         ?
                 qth quartile Number        
                 FilterMulti expression   ?
            - syntax
                 quartile(quartile_of_data = quartile(age))
                 #one with filter
                 quartile(quartile(age,3,age>30)) #gets the 3rd quartile 
            - usecase:
                 similar to above sum,max etc
                   
        
        10. median
            - calculates median
            - args
                fieldName string
                weight    string         ?
                FilterMulti expression   ?
            - syntax
                median(median_of_data = median(age))
                #one with filter
                median(median(age,age>30))
            - usecase:
                similar to above sum,max etc    
                      
        11. covariance
            - calculates covariance between two fields
            - args
                fieldName1 string
                fieldName2 string
                filterMulti expression ?
            - syntax
                cov(fieldName1,fieldName2,condition)
            
            - usecase:
                file(something.csv) | select(field1,field2) | cov(field1,field2) | stdout();;
                
        12. correlation
            - calculates the correlation between two fields
            - syntax
                correlation(fieldName1,fieldName2)
            others same as covariance
        
        13. distinct count
            - count total unique items in the field
            - args
                fieldName string
                filterMulti expression ?
            - syntax
                dcount(age)
            - usecase:
                fake(10) | total_unique_ages = dcount(age) | stdout();;
                        
        14. pvalue count
            -
            - args
                fieldName1  string
                fieldName2  string
                filterMulti expression ?          
            - syntax
                pvalue(fieldName1,fieldName2)
            -                          
    Do Functions
        1. filter
            -filters the data based on the condition provided.
            - args
                filterMulti expression
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
   
    Join Funtions
        join
            - joins the two streams of the data
            - syntax 
                Inner Join :join alias = inner(query)
                Left Outer Join : join alias = leftouter(query)
                Right Outer JOin : join alias = rightouter(query)
                        
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
        
    
                                   