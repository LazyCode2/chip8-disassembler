package main

import (
	"fmt"
	"io"
	"os"

	"github.com/LazyCode2/chip8-disassembler/disassembler"
)

func main() {
	file, error := os.Open("rom/PONG")
	if error != nil {
		panic(error)
	}

	defer file.Close()

	data, _ := io.ReadAll(file)

	out, err := os.Create("out.asm")
	if err != nil {
		panic(err)
	}

	defer out.Close()

	for i := 0; i+1 < len(data); i += 2 {
		opcode := uint16(data[i])<<8 | uint16(data[i+1])

		prefix := (opcode & 0xF000) >> 12

		if handler, exists := disassembler.OpcodeTable[uint16(prefix)]; exists {
			instruction := handler(uint16(opcode))
			fmt.Fprint(out, instruction)
		} else {
			return
		}
	}
}
