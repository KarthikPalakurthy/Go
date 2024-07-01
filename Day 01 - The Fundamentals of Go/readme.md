# 1. Variables, Zero values and Blank Identifiers.

- Go is a statically typed Programming language where you’ll need to declare the type of a variable manually.
- Short declarations doesn’t need any additional declaration as they can be identified by the system.
- Blank identifiers are used when if you want to call a function and doesn’t want to return the value of one or more variables.
- If you dont assign any value to a variable, it assigns a `zero` value to itself. We cant use short declaration operator when declaring a zero value variable.
- Variables MUST be used when declared. It is called `polluting the code` when a variable is declared and not used. However, the compiler throws an error immediately.

## 2. Numerical Systems: decimal, binary and hexadecimal:

Decimal: value is calculated based on the exponential value of base 10.

Binary: value is calculated based on the exponential value of base  2.

Hexadecimal: value is calculated based on the exponential value of base  16.

[week+03+binary+numbering.pdf](https://prod-files-secure.s3.us-west-2.amazonaws.com/d1aecbb5-7dc8-4ce4-9575-56e9ba742748/9f6a05e8-7124-4f93-8463-89c8d51473bc/week03binarynumbering.pdf)

### 3. Values, types, conversion and scope:

- A short declaration can only be used inside a function.
- If you are looking for a `package scope variable`, it has to be declared outside the function and cannot be a short variable.

In Go, the scope of a variable is the region of the code where the variable is declared and accessible. Variables can be scoped at different levels:

1. **Package Scope:** Declared outside any function and accessible throughout the package.
2. **Function Scope:** Declared within a function and accessible only within that function.
3. **Block Scope:** Declared within a block (e.g., within `{}` in `if`, `for`, or `switch` statements) and accessible only within that block.

# Questions to Practice:

1. Write a code to Print the binary and hexadecimal values of the already declared variables in the code.