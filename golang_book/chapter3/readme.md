# Types

## Numbers
 * Integer

    - Integers – like their mathematical counterpart – are numbers without a decimal component. (…, -3, -2, -1, 0, 1, …) 
    - Unlike the base-10 decimal system we use to represent numbers, computers use a **base-2** binary system.
    - Go's integer types are: **uint8**, **uint16**, **uint32**, **uint64**, **int8**, **int16**, **int32** and **int64**
    - Generally if you are working with integers you should just use the int type.

 * Floating Point Numbers
    - Floating point numbers are numbers that contain a decimal component (real numbers). (1.234, 123.4, 0.00001234, 12340000)
    - Floating point numbers are inexact. Occasionally it is not possible to represent a number. For example computing 1.01 - 0.99 results in 0.020000000000000018 – A number extremely close to what we would expect, but not exactly the same. ( **aproximation** )

    - Like integers floating point numbers have a certain size (32 bit or 64 bit). Using a larger sized floating point number increases it's precision. (how many digits it can represent)
    - In addition to numbers there are several other values which can be represented: “not a number” (NaN, for things like 0/0) and positive and negative infinity. (+∞ and −∞)
    - Go has two floating point types: **float32** and **float64**
* go has following operators

    operator| action
    ---     | ---
    +	    | addition      
    -	    | subtraction   
    *	    | multiplication
    /	    | division      
    %	    | remainder     


## Strings
 * Several common operations on strings include finding the length of a string: **len("Hello World")**, accessing an individual character in the string: **"Hello World"[1]**
 * see updated [main.go](./main.go)
 * Strings are “indexed” starting at 0 not 1. [1] gives you the 2nd element not the 1st. Also notice that you see 101 instead of e when you run this program. This is because the character is represented by a byte (remember a byte is an integer).
## Booleans
 * A boolean value (named after George Boole) is a special 1 bit integer type used to represent true and false (or on and off)

* Three logical operators are used with boolean values:

    operator    |   action
    :---:       |   :---:
    &&          |	and
    \|\|	    |   or
    !	        |   not

# Problems

* [x] How are integers stored on a computer?
    
    - The numbers are stored as a sequence of bits (binary digits)
    - int32 - sequence of bits with length 32  


* [x] We know that (in base 10) the largest 1 digit number is 9 and the largest 2 digit number is 99. Given that in binary the largest 2 digit number is 11 (3), the largest 3 digit number is 111 (7) and the largest 4 digit number is 1111 (15) what's the largest 8 digit number? 
    
    -  100000000 - 1 = 2<sup>8</sup> - 1 = 256 - 1 = 255
    -   1111 1111 = 2<sup>7</sup> + 2<sup>6</sup> + 2<sup>5</sup> + 2<sup>4</sup> + 2<sup>3</sup> + 2<sup>2</sup> + 2<sup>1</sup> + 2<sup>0</sup> = 128 + 64 + 32 + 16 + 8 + 4 + 2 + 1 = 255
               
             
* [x] Although overpowered for the task you can use Go as a calculator. Write a program that computes 32132 × 42452 and prints it to the terminal. (Use the * operator for multiplication)

* [x] What is a string? How do you find its length?
    - A string is a sequence of characters 8 bits represents one character
    - len("string")

* [x] What's the value of the expression (true && false) || (false && true) || !(false && false)?
    - true