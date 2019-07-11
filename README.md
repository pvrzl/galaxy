Implementation of galaxy merchant trading guide in go language.

### Requirement

- go

### How to compile and run test

```
$ ./bin/setup
```

### How to run the command and enter shell

```
$ ./bin/galaxy
```

### How to run the command and read input from file

```
$ ./bin/galaxy bin/sample
```

Please note,

- all of the command above are executed in the root directory of this project.
- this project use go module as dependency manager(this project actually didn't use any direct external library).
- golang version : 1.12.5

#### Step by step

unfortunately i didn't record every single step of this project with git commit. But if you still want to know, here's the step :

1. I start this project by creating convertRoman package first. Because, basically this is just another roman converter+alias.
2. I created method and test for convertRoman one per time.
3. The internal package is created along the romanConverter. This package is intended to use internally to support all package.
4. When convertRoman package finish, i start to create input package. At first i only create shell version, and starting to test the input and added the alias to convertRoman package. Later i added the file version.
5. As for input, i decided to split every command by "is", and split it again by space. Then categories every split into three condition: defining alias, calculate unknown word price, calculate the value. (the detail is in code, i commented a few of line code)
