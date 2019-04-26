package main

import (
	"bytes"
	"fmt"
	"strconv"
	"syscall"

	"github.com/xlab/tablewriter"
)

type context struct {
	pc uint
	register [10]int64
	memory   [10]int64
	lines  []op
}

type op interface {
	exec(ctx *context)
}

func parse(s string) op {
	switch s[0] {
	case '1':
		return &LoadMem{s}
	case '2':
		return &LoadRegister{s}
	case '3':
		return &StoreMem{s}
	case '4':
		return &Move{s}
	case '5':
		return &Add{s}
	case '6':
	case '7':
		return &Or{s}
	case '8':
		return &And{s}
	case '9':
		return &Xor{s}
	case 'A':
		return &Rot{s}
	case 'B':
		return &Jump{s}
	case 'C':
		return &Halt{}
	case 'P':
		return &Print{}
	}
	return nil
}

func main() {
	ctx := &context{}
	for {
		var line string
		fmt.Scanf("%s", &line)
		if line == "" {
			break
		}
		op := parse(line)
		if op == nil {
			fmt.Println("couldn't parse", line)
			continue
		}
		ctx.lines = append(ctx.lines, op)
	}

	for  {
		op := ctx.lines[ctx.pc]
		ctx.pc++
		op.exec(ctx)
	}
}


type LoadMem struct {
	line string
}

func (v *LoadMem) exec(ctx *context) {
	register, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	var t int64
	vv, err := strconv.Atoi(v.line[2:])
	if err != nil {
		t, err = strconv.ParseInt(v.line[2:], 0, 64)
	} else {
		t = int64(vv)
	}
	if err != nil {
		panic(err)
	}
	ctx.register[register] = ctx.memory[t]
}

type LoadRegister struct {
	line string
}

func (v *LoadRegister) exec(ctx *context) {
	register, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	var t int64
	vv, err := strconv.Atoi(v.line[2:])
	if err != nil {
		t, err = strconv.ParseInt(v.line[2:], 0, 64)
	} else {
		t = int64(vv)
	}
	if err != nil {
		panic(err)
	}
	ctx.register[register] = t
}

type StoreMem struct {
	line string
}

func (v *StoreMem) exec(ctx *context) {
	register, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	var t int64
	vv, err := strconv.Atoi(v.line[2:])
	if err != nil {
		t, err = strconv.ParseInt(v.line[2:], 0, 64)
	} else {
		t = int64(vv)
	}
	if err != nil {
		panic(err)
	}
	ctx.memory[t] = ctx.register[register]
}

type Print struct {
}

func (v *Print) exec(ctx *context) {
	t := tablewriter.CreateTable()
	t.UTF8Box()
	t.AddTitle("Machine")
	t.AddHeaders("", "R", "M")
	for i := 0; i < 10; i++ {
		t.AddRow(i, ctx.register[i], ctx.memory[i])
	}

	var buf bytes.Buffer
	buf.WriteString(t.Render())
	fmt.Println(buf.String())
}

type Move struct {
	line string
}

func (v *Move) exec(ctx *context) {
	from, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(v.line[2:3])
	if err != nil {
		panic(err)
	}
	ctx.register[to] = ctx.register[from]
}

type Add struct {
	line string
}

func (v *Add) exec(ctx *context) {
	a, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(v.line[2:3])
	if err != nil {
		panic(err)
	}
	c, err :=strconv.Atoi(v.line[3:4])
	ctx.register[a] = ctx.register[b] + ctx.register[c]
}

type Or struct {
	line string
}

func (v *Or) exec(ctx *context) {
	a, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(v.line[2:3])
	if err != nil {
		panic(err)
	}
	c, err :=strconv.Atoi(v.line[3:4])
	ctx.register[c] = ctx.register[a] | ctx.register[b]
}

type And struct {
	line string
}

func (v *And) exec(ctx *context) {
	a, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(v.line[2:3])
	if err != nil {
		panic(err)
	}
	c, err :=strconv.Atoi(v.line[3:4])
	ctx.register[c] = ctx.register[a] & ctx.register[b]
}

type Xor struct {
	line string
}

func (v *Xor) exec(ctx *context) {
	a, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(v.line[2:3])
	if err != nil {
		panic(err)
	}
	c, err :=strconv.Atoi(v.line[3:4])
	ctx.register[c] = ctx.register[a] ^ ctx.register[b]
}

type Rot struct {
	line string
}

func (v *Rot) exec(ctx *context) {
	a, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(v.line[2:3])
	if err != nil {
		panic(err)
	}
	c, err :=strconv.Atoi(v.line[3:4])
	ctx.register[c] = ctx.register[a] >> uint(ctx.register[b])
}

type Jump struct {
	line string
}

func (v *Jump) exec(ctx *context) {
	a, err := strconv.Atoi(v.line[1:2])
	if err != nil {
		panic(err)
	}
	pc, err := strconv.Atoi(v.line[2:])
	if ctx.register[0] == ctx.register[a] {
		ctx.pc = uint(pc)
	}
}

type Halt struct {
}

func (v *Halt) exec(ctx *context) {
	syscall.Exit(0)
}