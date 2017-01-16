# golang-quickstart
This is a quick project to help creating a cleaner sub directory structure in a golang project

### Basic commands to run with this project.

To keep the working dir clean, generally you can do the following in the beginning in your shell:
```
export GOPATH=`pwd`
export PATH=$PATH:$GOPATH/bin
```

1. Test individual package:

    ```
    go test -v github.com/user/package1
    ```

2. Run all tests recursively in your packages:

    ```
    go test -v github.com/user/...  ( the 3 dots are to run tests recursively )
    ```

3. Definte the parallelism with which the tests would run:

    ```
    go test -p 1 -v github.com/user/...  ( In case you want your tests to run with parallelism 1, for e.g. issues when you have database connections and you want to maintain the integrity. )
    ```

4. Run your `main` package:

    ```
    go install github.com/user/
    user
    ```
    In cae you did not find `user` executable in `pwd` look in the `bin` directory.

5. If you want to build the current project with a different binary:

    ```
    go build -i -o mybinary github.com/user  ( -i will build all the dependencies as well)
    ./mybinary
    ```
