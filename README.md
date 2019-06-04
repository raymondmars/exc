[![MIT licensed][5]][6] [![LICENSE](https://img.shields.io/badge/license-NPL%20(The%20996%20Prohibited%20License)-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE) [![Build Status][1]][2] [![Go Report Card][7]][8] [![GoCover.io][11]][12] [![GoDoc][9]][10]

[1]: https://travis-ci.org/0RaymondJiang0/exc.svg?branch=master
[2]: https://travis-ci.org/0RaymondJiang0/exc
[5]: https://img.shields.io/dub/l/vibe-d.svg
[6]: LICENSE
[7]: https://goreportcard.com/badge/github.com/timothyye/skm
[8]: https://goreportcard.com/report/github.com/timothyye/skm
[9]: https://godoc.org/github.com/TimothyYe/skm?status.svg
[10]: https://godoc.org/github.com/TimothyYe/skm
[11]: https://img.shields.io/badge/gocover.io-81.8%25-green.svg
[12]: https://gocover.io/github.com/timothyye/skm

**Exc is a tool for saving you command,args and execute it by alias**      

Usage:     
&emsp;exc [alias]  &emsp;- to execute you command    
&emsp;exc [command] [arguments...] &emsp;- to do other operations    

Commands:     
    &emsp;--add value, -a value        &emsp;to add a command named with it's alias to cache (The value should be: alias command [args])    
    &emsp;--remove value, -r value     &emsp;to remove a command from cache (The value should be: alias)     
    &emsp;--list, -l                   &emsp;to show command list, it will depend on you current directory     
    &emsp;--list-all, -la              &emsp;to show all command list in all directory     
    &emsp;--help, -h                   &emsp;to show help    
    &emsp;--version, -v                &emsp;to print the version      

Examples:    
    &emsp;exc -a gs git status        &emsp;set command "git status" as "gs", after it, you can execute it use "exc gs"    
    &emsp;exc -r gs                   &emsp;remove gs alias from cache     