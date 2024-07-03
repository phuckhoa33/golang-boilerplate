# Create makefile for the project
# Define the compiler and flags
CC := go
CFLAGS := build

# Define the target executable
TARGET := myapp

# Define the source files
SRCS := $(wildcard *.go)

# Define the object files
OBJS := $(SRCS:.go=.o)

# Define the default target
all: $(TARGET)

# Compile the source files into object files
%.o: %.go
	$(CC) $(CFLAGS) -o $@ -c $<

# Link the object files into the target executable
$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^

# Clean the generated files
clean:
	rm -f $(OBJS) $(TARGET)