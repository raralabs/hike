
#![](error.png)
    
###Error 
    
    cause = while subscribing msg from multiple stage
    
    cmd:=`fake(100) | filter(age>30) | sort() by age | s1 = into();
          fake(200) | filter(age>30) | s2 = into();
          # this is a comment
          s1,s2     | stdout();;
      
    Test Conducted:
        created the pipeline manually in canal - no errors 
        created the pipeline using hike parser - errors observed
              
    Reason of error:
        Bug exists in pretty.go when stdout recieves data
        from multiple stage.
     
        

                  