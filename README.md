**Exc is a tool for saving you command,args and execute it by alias**      

Usage:     
&emsp;exc [alias]  &emsp;- to execute you command    
&emsp;exc [command] [arguments...] &emsp;- to do other operations    

Commands:     
    &emsp;--add value, -a value        &emsp;to add a command named with it's alias to cache (The value should be: alias command [args])    
    &emsp;--remove value, -r value     &emsp;to remove a command from cache (The value should be: alias)     
    &emsp;--list, -l                   &emsp;to show command list, it will depend on you current directory     
    &emsp;--help, -h                   &emsp;to show help    
    &emsp;--version, -v                &emsp;to print the version      

Examples:    
    &emsp;exc -a gs git status        &emsp;set command "git status" as "gs", after it, you can execute it use "exc gs"    
    &emsp;exc -r gs                   &emsp;remove gs alias from cache     