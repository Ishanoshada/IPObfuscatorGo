# IP Obfuscator in Go


This repository contains a Go program that provides functions for IP address validation and obfuscation. It includes different methods to obfuscate an IP address for security or privacy purposes.

## Usage

To use this program, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/ishanoshada/IPObfuscatorGo
   ```

2. Navigate to the project directory:

   ```bash
   cd IPObfuscatorGo
   ```

3. Run the program:

   ```bash
   go run main.go
   ```

4. Follow the prompts to enter an IP address for obfuscation.

## Code

The core functionality of this program is provided in the `main.go` file. It includes functions for validating an IP address and obfuscating it using various methods.

### Validating IP

The `isValidIP` function checks if an IP address is valid. It verifies that the address consists of four segments, each within the range of 0 to 255.

### Obfuscating IP

The `obfuscateIP` function takes a valid IP address as input and provides different obfuscated versions of it. It includes methods like converting to hexadecimal, octal, and binary representations.

### Running the Program

The program prompts the user to enter an IP address. It then displays the obfuscated versions along with the original IP address.

## Example

```plaintext
[+] Enter your IP: 192.168.0.1

[~] Obfuscated IPs :
[+] http://3232235521
[+] http://0xC0.0xA8.0x0.0x1
[+] http://0300.0250.00.01
[+] http://0xc0a80001

Original IP: 192.168.0.1
```



This program is written in Go.

