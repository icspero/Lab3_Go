package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkQPUSH(bench *testing.B) {
	queue := newQueue()

	for i := 0; i < 1000; i++ {
		queue.QPUSH("element")
	}
}

func BenchmarkQPOP(bench *testing.B) {
	queue := newQueue()
	for i := 0; i < 1000; i++ {
		queue.QPUSH("element")
	}
	bench.ResetTimer()
	for i := 0; i < 1000; i++ {
		queue.QPOP()
	}
}
func TestQPUSH(test *testing.T) {
	queue := newQueue()
	queue.QREAD()
	queue.QPUSH("element1")
	queue.QPUSH("element2")
	queue.QPUSH("element3")

	assert.Equal(test, "element1", queue.head.cell)
	assert.Equal(test, "element2", queue.head.next.cell)
	assert.Equal(test, "element3", queue.tail.cell)

	queue.QREAD()
}

func TestQPOP(test *testing.T) {

	queue := newQueue()

	queue.QPUSH("element1")
	queue.QPUSH("element2")
	queue.QPUSH("element3")

	assert.Equal(test, "element1", queue.head.cell)

	queue.QPOP()

	assert.Equal(test, "element2", queue.head.cell)
	assert.Equal(test, "element3", queue.tail.cell)

	queue.QPOP()
	queue.QPOP()

	assert.Nil(test, queue.head)
	assert.Nil(test, queue.tail)
}

func TestQPOP_EMPTY(test *testing.T) {
	queue := newQueue()
	queue.QPOP()

	assert.Nil(test, queue.head)
	assert.Nil(test, queue.tail)
}

func TestBinarySerialization(test *testing.T) {
	queue := newQueue()
	queue.QPUSH("element1")
	queue.QPUSH("element2")
	queue.QPUSH("element3")
	queue.BinarySerialization("/root/notExist.txt")
	err := queue.BinarySerialization("file.bin")
	assert.Nil(test, err)
	queue.BinaryDeserialization("notExist.bin")
	result, err := queue.BinaryDeserialization("file.bin")
	assert.Nil(test, err)
	assert.Equal(test, []string{"element1", "element2", "element3"}, result)
}
func TestFileOperations(test *testing.T) {
	queue := newQueue()

	queue.QPUSH("element1")
	queue.QPUSH("element2")
	queue.WritingFromStructureToFile("/root/notExist.txt")
	queue.WritingFromStructureToFile("file.txt")

	newQueue := newQueue()
	newQueue.WritingFromFileToStructure("notExist.txt")
	newQueue.WritingFromFileToStructure("file.txt")

	assert.Equal(test, "element1", newQueue.head.cell)
	assert.Equal(test, "element2", newQueue.head.next.cell)
}
