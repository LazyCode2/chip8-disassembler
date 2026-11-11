package main

import (
	"fmt"
	"github.com/LazyCode2/chip8-disassembler/disassembler"
)

func main() {
	programBytes := []uint16{0x00E0, 0x00EE, 0x1234, 0x3400}

	for _, opcode := range programBytes {
		prefix := (opcode & 0xF000) >> 12

		if handler, exists := disassembler.OpcodeTable[uint16(prefix)]; exists {
			handler(uint16(opcode))
		} else {
			fmt.Printf("-> Error: Unimplemented opcode prefix 0x%X\n", prefix)
			return
		}
	}
}
