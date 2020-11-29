<h2>Hello, I'm Berik <i>Alanapapa</i> Bazarov!</h2>
<img align='right' src="gopher-rainbow.gif" width="230">
<p><em>Freelancer and student at <a href="https://alem.school/" target="_blank">Alem Coding School</a><img src="https://media.giphy.com/media/WUlplcMpOCEmTGBtBW/giphy.gif" width="30"> 
</em></p>

[![Github: alanapapa](https://img.icons8.com/material-outlined/48/000000/github.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.github.com/alanapapa/)](https://www.github.com/alanapapa/)
[![Instagram: berikbazarov](https://img.icons8.com/fluent/48/000000/instagram-new.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.instagram.com/berikbazarov/)](https://www.instagram.com/berikbazarov/)
[![Facebook: berikbazarov](https://img.icons8.com/color/48/000000/facebook.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.facebook.com/bazarovberik/)](https://www.facebook.com/bazarovberik/)

![GitHub followers](https://img.shields.io/github/followers/alanapapa?label=Follow&style=social)


## ***ascii-art-output***
### This is my solution to the *ascii-art-output* educational task on [01 Edu System](https://www.01-edu.org/)

<br>

### Objectives

- You must follow the same [instructions](https://public.01-edu.org/subjects/ascii-art/) as in the first subject **while** writing the result into a file.

- The file must be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name.

This project will help you learn about :

- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Choices of outputs.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code presents a **test file**.

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard --output=banner.txt
student@ubuntu:~/ascii-art$ cat banner.txt
 _              _   _
| |            | | | |
| |__     ___  | | | |   ___
|  _ \   / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/



student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" shadow --output=banner.txt
student@ubuntu:~/ascii-art$ cat banner.txt

_|    _|          _| _|                _|_|_|_|_| _|                                  _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _|



student@ubuntu:~/ascii-art$
```
