# Welcome, welcome

## Install Go on Windows

### Install Git Bash (optional but recommended, don't skip if you are not sure) 

1. Download git binary: https://git-scm.com/download/win
2. Keep everything default is fine.

### Install Go binary

1. Go to https://golang.org/dl and download the windows installer.
2. Next. Next. Next...
3. Now fire up git bash / cmd, type the following:

   ```shell
   $ go verion
   go version go1.4.2 windows/amd64
   ```

4. If your output is something like the above, done!

### Set up GOPATH

GOPATH is where your go projects and libraries resides. Choose wisely :D

1. Make a new folder in an accessible location, take note of the paths.
2. Right click `My Computer` -> `Perperties`
3. Choose `Advanced system settings`
4. Open `Envionment Variables...`
5. Under User variables, click `New...`
6. `Variable name`: `GOPATH`
   `Variable value`: The folder chosen in step 1. Mine is `C:\Users\patat_000\go`.

### Add Go bin folder to PATH

It enable you to invoke go binary everywhere.

1. At step 4 of _Set up GOPATH_, look for a environment variable named `PATH`. If there are none, create one.
2. Edit and append the following to the end of the value (don't delete existing value!):
   `;<YOUR GOPATH HERE>\bin`

   Taking the previous example, mine would be like `;C:\Users\patat_000\go\bin`

## Install Go on Linux

### Install Go binary

See https://golang.org/doc/install#tarball

Of course you can also consult your distro packages...

### Set up GOPATH

```shell
$ mkdir $HOME/go
```

Now edit your `$HOME/.profile` to include the following lines:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

NOTE: if you are not using bash, the file to edit might be different. 

## Install Go on Mac OS X

### Install Go binary

See https://golang.org/doc/install#osx

If you have installed Homebrew:

```shell
$ brew install go
```

### Set up GOPATH

```shell
$ mkdir $HOME/go
```

Now edit your `$HOME/.profile` to include the following lines:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

NOTE: if you are not using terminal, the file to edit might be different. 
NOTE NOTE: same instructions as Linux's one :P
