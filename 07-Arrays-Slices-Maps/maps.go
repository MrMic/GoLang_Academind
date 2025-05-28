package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
	}
	fmt.Println("🪚 websites:", websites)
	fmt.Println("🪚 websites[Google]:", websites["Google"])

	// * INFO: Add K/V
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println("🪚 websites:", websites)

	delete(websites, "Linkedin")
	fmt.Println("🪚 websites:", websites)
}
