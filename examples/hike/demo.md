# Examples

## Create a HelloWorld Project for hike

```bash
> create
ProjectName> HelloWorld
ProjectDirectory> .
# enter "cmd" to enter the command mode
# Make sure that the project directory contains the required csv files
# Run the hike commands
eg:
statement 1>> fake(20) | agg counter=count(age>30) | s1 = into();
statement 2>> s1 | stdout();;

#enter "end" to terminate the program
```

## Open the created project

```bash
> open
ProjectName> HelloWorld

# Run the hike commands
similar to above
```
