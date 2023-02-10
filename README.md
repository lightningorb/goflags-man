# Just an repo to experiment creating man pages in GoLang


## Running

```
$ make build generate_man install_man
$ man program
```

Output:

```
hellogo(1)                                                                                                                                       hellogo(1)

NAME
       hellogo -

SYNOPSIS
       hellogo [OPTIONS]

DESCRIPTION
OPTIONS
   Application Options
       -m, --message <default: "Hello, World!">
              The greeting message

       -g, --generate-man-page
              Generate man page

   Help Options
       -h, --help
              Show this help message

                                                                      10 February 2023                                                           hellogo(1)
```