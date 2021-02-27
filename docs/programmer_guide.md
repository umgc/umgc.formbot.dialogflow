# Dialogflow Team Programmer Guide


## How to compile

Go compiles directly to machine code so it is super portable. This allows you to run on any system and hardware.

### Running a Go program

To run your go program (not compile), simply head to your directory and type: "go run main.go"

### Compiling multiple go files from the same package

"go run ."

# Dialogflow backend Algorithm

As the beckend services need to be engineered to operate in a dynamic way there is some special logic in place to make it available for anyone to customize.

The end user should be able to hook into the API and specify their own credentials.

1. Session ID is used to identify the session, which is written to a database.

