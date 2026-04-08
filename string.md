### Go Get Started
To start using Go, you need two things:

## A text editor, like VS Code, to write Go code
A compiler, like GCC, to translate the Go code into a language that the computer will understand
There are many text editors and compilers to choose from. In this tutorial, we will use an IDE (see below).

# Go Install
You can find the relevant installation files at https://golang.org/dl/.

Follow the instructions related to your operating system. To check if Go was installed successfully, you can run the following command in a terminal window:

# go version
Which should show the version of your Go installation.

# Go Install IDE
An IDE (Integrated Development Environment) is used to edit AND compile the code.

Popular IDE's include Visual Studio Code (VS Code), Vim, Eclipse, and Notepad. These are all free, and they can be used to both edit and debug Go code.

# Note: Web-based IDE's can work as well, but functionality is limited.

We will use VS Code in our tutorial, which we believe is a good place to start.

You can find the latest version of VS Code at https://code.visualstudio.com/.

*** Go Quickstart***
***Let's create our first Go program.***

```Launch the VS Code editor
Open the extension manager or alternatively, press Ctrl + Shift + x
In the search box, type "go" and hit enter
Find the Go extension by the GO team at Google and install the extension
After the installation is complete, open the command palette by pressing Ctrl + Shift + p
Run the Go: Install/Update Tools command
Select all the provided tools and click OK
VS Code is now configured to use Go.
```

Open up a terminal window and type:

go mod init example.com/hello
Do not worry if you do not understand why we type the above command. Just think of it as something that you always do, and that you will learn more about in a later chapter.

Create a new file (File > New File). Copy and paste the following code and save the file as helloworld.go (File > Save As):
```
helloworld.go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
Now, run the code: Open a terminal in VS Code and type:

go run .\helloworld.go
Hello World!
Congratulations! You have now written and executed your first Go program.
```
If you want to save the program as an executable, type and run:

# go build .\helloworld.go
When learning Go you can use our "Try it Yourself" tool. It shows both the code and the result. This will make it easier for you to understand every part as we move forward:

helloworld.go
Code:
```
package main
import ("fmt")

func main() {
 fmt.Println("Hello World!")
}
Result:

Hello World!
```


## Go Syntax
 **A Go file consists of the following parts:**

# Package declaration
**Import packages**
**Functions**
**Statements and expressions**
**Look at the following code, to understand it better:**

```Example
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
Example explained
Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

Line 2: import ("fmt") lets us import files included in the fmt package.

Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".

Note: In Go, any executable code belongs to the main package.
```
# Go Statements
fmt.Println("Hello World!") is a statement.

 # In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

# Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).

The left curly bracket { cannot come at the start of a line.

Run the following code and see what happens:

# Example
package main
import ("fmt")

func main()
{
  fmt.Println("Hello World!")
}


# Go Compact Code
You can write more compact code, like shown below (this is not recommended because it makes the code more difficult to read):

Example
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}
Go Exercises
Test Yourself With Exercises
Exercise:
Insert the missing part of the code below to output "Hello World".

package main   
import ("fmt")

func main() {  
  
("Hello World!")
}

Start the Exercise


Get Certified Offer
COLOR PICKER
colorpicker
## Go Comments
# Go Comments
# A comment is a text that is ignored upon execution.

# Comments can be used to explain the code, and to make it more readable.

# Comments can also be used to prevent code execution when testing an alternative code.

Go supports single-line or multi-line comments.

Go Single-line Comments
Single-line comments start with two forward slashes (//).

Any text between // and the end of the line is ignored by the compiler (will not be executed).

Example
// This is a comment
package main
import ("fmt")

func main() {
  // This is a comment
  fmt.Println("Hello World!")
}
The following example uses a single-line comment at the end of a code line:
```
Example
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!") // This is a comment
}
Go Multi-line Comments
Multi-line comments start with /* and ends with */.

Any text between /* and */ will be ignored by the compiler:

Example
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
}
Tip: It is up to you which you want to use. Normally, we use // for short comments, and /* */ for longer comments.
```


```

Comment to Prevent Code Execution
You can also use comments to prevent the code from being executed.

The commented code can be saved for later reference and troubleshooting.

Example
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
  // fmt.Println("This line does not execute")
}
Go Exercises
Test Yourself With Exercises
Exercise:
Comments in Go are written with a special character, which one?

package main   
import ("fmt") 

func main() {
  
 this is a comment  
  fmt.Println("Hello World!")
}
```



