package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Reverse function to reverse a string
func reverse(s string, ch chan string) {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	ch <- string(runes) // Send the result to the channel
}

// Function to sort numbers using Bubble Sort (runs in a goroutine)
func bubbleSort(arr []int, ch chan []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	ch <- arr // Send the sorted array to the channel
}

// Function to find the middle element of a linked list (runs in a goroutine)
type Node struct {
	Value int
	Next  *Node
}

func findMiddle(head *Node, ch chan *Node) {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	ch <- slow // Send the middle element to the channel
}

// Function to demonstrate input reading
func getUserInput(ch chan string) {
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	result := fmt.Sprintf("Hello %s! You are %d years old.\n", name, age)
	ch <- result // Send the result to the channel
}

// Main function to show menu and perform actions based on user choice
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Show the menu of options
		fmt.Println("\nChoose an operation:")
		fmt.Println("1. Reverse a string")
		fmt.Println("2. Get user details")
		fmt.Println("3. Sort numbers using Bubble Sort")
		fmt.Println("4. Find middle element of a linked list")
		fmt.Println("5. Exit")

		// Read user's choice
		fmt.Print("Enter choice (1-5): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			// Reverse a string in a goroutine
			var input string
			fmt.Print("Enter a string to reverse: ")
			fmt.Scanln(&input)

			ch := make(chan string)
			go reverse(input, ch) // Start the reverse function as a goroutine

			// Wait for the result and print it
			reversedString := <-ch
			fmt.Println("Reversed String:", reversedString)

		case "2":
			// Get user input concurrently
			ch := make(chan string)
			go getUserInput(ch) // Run input collection in a goroutine

			// Wait for the result and print it
			result := <-ch
			fmt.Println(result)

		case "3":
			// Sort numbers using Bubble Sort in a goroutine
			var n int
			fmt.Print("Enter the number of elements you want to sort: ")
			fmt.Scan(&n)

			arr := make([]int, n)
			fmt.Println("Enter the numbers:")
			for i := 0; i < n; i++ {
				fmt.Scan(&arr[i])
			}

			ch := make(chan []int)
			go bubbleSort(arr, ch) // Start the bubble sort as a goroutine

			// Wait for the sorted array and print it
			sortedArray := <-ch
			fmt.Println("Sorted numbers:", sortedArray)

		case "4":
			// Create a sample linked list: 1 -> 2 -> 3 -> 4 -> 5
			head := &Node{Value: 1}
			head.Next = &Node{Value: 2}
			head.Next.Next = &Node{Value: 3}
			head.Next.Next.Next = &Node{Value: 4}
			head.Next.Next.Next.Next = &Node{Value: 5}

			// Find and print the middle element concurrently
			ch := make(chan *Node)
			go findMiddle(head, ch) // Start finding the middle element as a goroutine

			// Wait for the result and print the middle element
			middle := <-ch
			if middle != nil {
				fmt.Println("Middle Element:", middle.Value)
			} else {
				fmt.Println("No middle element found!")
			}

		case "5":
			fmt.Println("Exiting the program.")
			return // Exit the program

		default:
			fmt.Println("Invalid choice! Please enter a number between 1 and 5.")
		}

		// Wait for a short time before showing the menu again (for better user experience)
		time.Sleep(1 * time.Second)
	}
}
