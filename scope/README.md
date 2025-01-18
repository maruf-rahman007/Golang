## Scope 

###### What is scope ?
*** It is the part of program which defines where a variable can be accessed. In go it is defined at compile time. A variable can only be accessed within the block where it is defined. ***

- Local Variables -- variables that are defined inside a function/loop/condition/block and they are not accessible outside that block
- Global Variables -- variables that are defined outside a function/loop/condition/block and they are accessible from anywhere of the program
- Local Variable Preference -- if a local variable shares the same name as a global variable the local variable will get the preferance/precedence within its scope    