package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type SingleNode struct {
	cell string
	next *SingleNode
}

type Queue struct {
	head *SingleNode
	tail *SingleNode
}

func newQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) QPUSH(cell string) {
	node := &SingleNode{cell: cell, next: nil}

	if queue.head == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
}

func (queue *Queue) QPOP() {
	if queue.head == nil {
		return
	}
	queue.head = queue.head.next
	if queue.head == nil {
		queue.tail = nil
	}
}

func (queue *Queue) QREAD() {
	current := queue.head
	if current == nil {
		fmt.Println("Empty!")
	} else {
		for current != nil {
			fmt.Print(current.cell, " ")
			current = current.next
		}
		fmt.Println()
	}
}

func (queue *Queue) WritingFromFileToStructure(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		for _, element := range elements {
			queue.QPUSH(element)
		}
	}
}

func (queue *Queue) WritingFromStructureToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creatin file ", err)
		return
	}
	defer file.Close()

	current := queue.head
	for current != nil {
		_, err := fmt.Fprint(file, current.cell+" ")
		if err != nil {
			fmt.Println("Error writing file ", err)
			return
		}
		current = current.next
	}
}

func (queue *Queue) BinarySerialization(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	current := queue.head
	for current != nil {
		len := int32(len(current.cell))
		if err := binary.Write(file, binary.LittleEndian, len); err != nil {
			return err
		}

		_, err := file.Write([]byte(current.cell))
		if err != nil {
			return err
		}

		current = current.next
	}

	return nil
}

func (queue *Queue) BinaryDeserialization(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string

	for {
		var len int32
		err := binary.Read(file, binary.LittleEndian, &len)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		buffer := make([]byte, len)
		_, err = file.Read(buffer)
		if err != nil {
			return nil, err
		}

		text := string(buffer)
		result = append(result, text)
	}

	return result, nil
}
func (q *Queue) Clear() {
	q.head = nil
	q.tail = nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	queue := newQueue()
	queue.Clear()

	fmt.Println()
	fmt.Print("Enter command: ")
	scanner.Scan()
	command := scanner.Text()
	parts := strings.Fields(command)

	if len(parts) == 0 {
		return
	}

	if parts[0] == "exit" {
		return
	}

	if len(parts) < 2 {
		fmt.Println("Error: Missing filename!")
		return
	}

	filename := parts[1]
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist!")
		return
	}

	if strings.HasSuffix(filename, ".txt") {
		queue.WritingFromFileToStructure(filename)
	} else if strings.HasSuffix(filename, ".bin") {
		result, err := queue.BinaryDeserialization(filename)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Deserialized result:", result)
		}
	} else {
		fmt.Println("Error: Unsupported file format!")
		return
	}

	switch parts[0] {
	case "QPUSH":
		if len(parts) == 3 {
			queue.QPUSH(parts[2])
		} else {
			fmt.Println("Error: Missing value to push!")
		}
	case "QPOP":
		queue.QPOP()
	case "QREAD":
		queue.QREAD()
	default:
		fmt.Println("Error: Unknown command!")
	}

	if strings.HasSuffix(filename, ".txt") {
		queue.WritingFromStructureToFile(filename)
	} else if strings.HasSuffix(filename, ".bin") {
		if err := queue.BinarySerialization(filename); err != nil {
			fmt.Println("Error:", err)
		}
	}

}
