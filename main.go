package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isValidIP checks if an IP address is valid.
func isValidIP(ip string) bool {
	segments := strings.Split(ip, ".")

	if len(segments) != 4 {
		return false
	}

	for _, segment := range segments {
		num, err := strconv.Atoi(segment)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}

// obfuscateIP obfuscates the provided IP address in various ways.
func obfuscateIP(ip string) {
	if !isValidIP(ip) {
		fmt.Fprintln(os.Stderr, "[!] Enter a valid IP")
		os.Exit(-1)
	}

	dec := strings.Split(ip, ".")

	// Obfuscated IPs
	fmt.Printf("\n[~] Obfuscated IPs :\n")
	
	// Obfuscate IP using bitwise operations
	fmt.Printf("[+] http://%d\n", (toInt(dec[0]) << 24) | (toInt(dec[1]) << 16) | (toInt(dec[2]) << 8) | toInt(dec[3]))
	
	// Obfuscate IP by converting segments to hexadecimal
	fmt.Printf("[+] http://%s\n", strings.Join([]string{toHex(dec[0]), toHex(dec[1]), toHex(dec[2]), toHex(dec[3])}, "."))
	
	// Obfuscate IP by converting segments to octal
	fmt.Printf("[+] http://%s\n", strings.Join([]string{toOct(dec[0]), toOct(dec[1]), toOct(dec[2]), toOct(dec[3])}, "."))
	
	// Obfuscate IP by converting segments to hexadecimal with 8 characters
	fmt.Printf("[+] http://%s\n", strings.Join([]string{toHex2(dec[0]), toHex2(dec[1]), toHex2(dec[2]), toHex2(dec[3])}, "."))

	// Display the original IP address
	fmt.Printf("Original IP: %s\n", ip)
}

// toInt converts a string segment to an integer.
func toInt(segment string) int {
	num, _ := strconv.Atoi(segment)
	return num
}

// toHex converts a string segment to a hexadecimal representation.
func toHex(segment string) string {
	return fmt.Sprintf("0x%02X", toInt(segment))
}

// toOct converts a string segment to an octal representation.
func toOct(segment string) string {
	return fmt.Sprintf("%04o", toInt(segment))
}

// toHex2 converts a string segment to a hexadecimal representation with 8 characters.
func toHex2(segment string) string {
	return fmt.Sprintf("0x%08X", toInt(segment))
}

// main function prompts the user for an IP address and calls obfuscateIP.
func main() {
	var ip string
	fmt.Print("[+] Enter your IP: ")
	fmt.Scan(&ip)
	obfuscateIP(ip)
}
