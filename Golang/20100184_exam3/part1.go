/* CS300 Exam 3: 6-May-2018 --- Playing with the GIF and BMP Image formats

CAREFULLY READ THE FOLLOWING INSTRUCTIONS:
   ===> Parts must be done in order. If a part is not working, you will not be graded for later parts.
        Move to the next part only when the current part is working according to the description given.
   ===> You can make helper functions but CANNOT change prototypes of existing functions.
        You can use additional libraries you want, but none is needed to implement the exam.
   ===> No laptops, mobile phones, USBs, calculators, or electronic devices near you. 
        Mobiles completely turned off.
   ===> SUBMIT all .go files that you modify during the exam.
*/
package main

import (
	"os"
	"fmt"
	"encoding/binary"
)

/* Part 1: Lookup the color palette 
In this part, you are given a list of codes where each code refers to a color on the color pallete (a 2D array). Each color in the color palette is stored as an array of Red, Green, and Blue components. Your task is to lookup the palette for each color code and return an array containing RGB colors in the opposite order (i.e. Blue, then Green and then Red) because BMP format stores them in the opposite order. If the input array has 100 codes, your output array will have 300 color values. You do not have to handle out-of-range values.
*/
func lookup(codes []int, palette [][3]uint8) []uint8 {
    ret := make([]uint8, 3)

        ret[0] = palette[codes[0]][2]
        ret[1] = palette[codes[0]][1]
        ret[2] = palette[codes[0]][0]

    for i:= 1; i < len(codes); i++{
        ret = append(ret, palette[codes[i]][2], palette[codes[i]][1], palette[codes[i]][0])
    }
	return ret
}

/* Part 2: Parallel lookup
Parallelize the above function such that both halves of the codes array are recursively given to separate goroutines. There should be no subdivision for arrays of length 16 or less (sequential cut-off). The Printf statement written in the function can help you see which tasks are being generated. Make sure that both halves are running in goroutines (and not just sequential recursive calls) and wait for both halves to terminate using a channel.

Note that instead of returning the final colors, we have made it an argument, so that you can use its slices in recursive calls without new allocations in every recursive function.

Once you finish the task, the given main function will produce "out2.bmp" which has the password for Part 3 
*/
func parallel_helper(codes []int, palette [][3]uint8, colors []uint8, start int, done chan bool) {
    fmt.Printf("Task to handle %d codes\n", len(codes))
    done1 := make(chan bool)
    if len(codes) > 16 {
        go parallel_helper(codes[0:len(codes)/2],palette, colors, start,done1)
        go parallel_helper(codes[len(codes)/2: len(codes)],palette, colors,start+(len(codes)/2)*3, done1)
        <-done1
        <-done1
    } else{
        for i:= 0; i < len(codes); i++{
            colors[start+(i*3)] = palette[codes[i]][2]
            colors[start+(i*3)+1] = palette[codes[i]][1]
            colors[start+(i*3)+2] = palette[codes[i]][0]
        }
    }
    done<- true
}

func parallel_lookup(codes []int, palette [][3]uint8, colors []uint8, done chan<- bool) {
    fmt.Printf("Task to handle %d codes\n", len(codes))
    done1 := make(chan bool)
    if(len(codes) > 16){
        go parallel_helper(codes[0:len(codes)/2],palette, colors, 0, done1)
        go parallel_helper(codes[len(codes)/2: len(codes)],palette, colors, (len(codes)/2)*3, done1)
        <-done1
        <-done1
    } 
    done<- true
}

func main() {
    // Part 1 testing code
    bmp("out1.bmp", lookup(part1_color_codes, part1_palette), 16, 8)

    // Part 2 testing code
    colors := make([]uint8, len(part1_color_codes)*3)
    done := make(chan bool)
    go parallel_lookup(part1_color_codes, part1_palette, colors, done)
    <-done
    bmp("out2.bmp", colors, 16, 8)
    fmt.Println("Done")
}

// DO NOT WORRY ABOUT ANYTHING BELOW THIS LINE

func bmp(filename string, data []uint8, dimX, dimY uint16) {
    bmp, _ := os.Create(filename)
    defer bmp.Close()
    binary.Write(bmp, binary.LittleEndian, []byte{'B','M'})
    binary.Write(bmp, binary.LittleEndian, []int32{int32(54+len(data)), 0, 54, 40, int32(dimX), -1 * int32(dimY), 0x180001, 0, 0, 0, 0, 0, 0})
    binary.Write(bmp, binary.LittleEndian, data)
}

var part1_color_codes = []int{
251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251,
251, 250, 223, 238, 251, 251, 230, 238, 251, 249, 226, 229, 239, 251, 236, 238,
251, 243, 224, 213, 251, 251, 250, 224, 244, 218, 251, 249, 225, 250, 225, 251,
251, 236, 203, 228, 197, 251, 251, 243, 210, 197, 251, 251, 230, 230, 190, 251,
250, 212, 251, 243, 232, 251, 251, 236, 182, 238, 251, 251, 250, 212, 251, 251,
236, 174, 211, 217, 169, 245, 250, 212, 251, 223, 245, 251, 250, 183, 251, 251,
229, 239, 251, 251, 236, 238, 223, 203, 251, 243, 225, 251, 250, 225, 251, 251,
251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251}

var part1_palette = [][3]uint8{
[3]uint8{0,0,0}, [3]uint8{0,0,51}, [3]uint8{0,0,102}, [3]uint8{0,0,153},
[3]uint8{0,0,204}, [3]uint8{0,0,255}, [3]uint8{0,43,0}, [3]uint8{0,43,51},
[3]uint8{0,43,102}, [3]uint8{0,43,153}, [3]uint8{0,43,204}, [3]uint8{0,43,255},
[3]uint8{0,85,0}, [3]uint8{0,85,51}, [3]uint8{0,85,102}, [3]uint8{0,85,153},
[3]uint8{0,85,204}, [3]uint8{0,85,255}, [3]uint8{0,128,0}, [3]uint8{0,128,51},
[3]uint8{0,128,102}, [3]uint8{0,128,153}, [3]uint8{0,128,204}, [3]uint8{0,128,255},
[3]uint8{0,170,0}, [3]uint8{0,170,51}, [3]uint8{0,170,102}, [3]uint8{0,170,153},
[3]uint8{0,170,204}, [3]uint8{0,170,255}, [3]uint8{0,213,0}, [3]uint8{0,213,51},
[3]uint8{0,213,102}, [3]uint8{0,213,153}, [3]uint8{0,213,204}, [3]uint8{0,213,255},
[3]uint8{0,255,0}, [3]uint8{0,255,51}, [3]uint8{0,255,102}, [3]uint8{0,255,153},
[3]uint8{0,255,204}, [3]uint8{0,255,255}, [3]uint8{51,0,0}, [3]uint8{51,0,51},
[3]uint8{51,0,102}, [3]uint8{51,0,153}, [3]uint8{51,0,204}, [3]uint8{51,0,255},
[3]uint8{51,43,0}, [3]uint8{51,43,51}, [3]uint8{51,43,102}, [3]uint8{51,43,153},
[3]uint8{51,43,204}, [3]uint8{51,43,255}, [3]uint8{51,85,0}, [3]uint8{51,85,51},
[3]uint8{51,85,102}, [3]uint8{51,85,153}, [3]uint8{51,85,204}, [3]uint8{51,85,255},
[3]uint8{51,128,0}, [3]uint8{51,128,51}, [3]uint8{51,128,102}, [3]uint8{51,128,153},
[3]uint8{51,128,204}, [3]uint8{51,128,255}, [3]uint8{51,170,0}, [3]uint8{51,170,51},
[3]uint8{51,170,102}, [3]uint8{51,170,153}, [3]uint8{51,170,204}, [3]uint8{51,170,255},
[3]uint8{51,213,0}, [3]uint8{51,213,51}, [3]uint8{51,213,102}, [3]uint8{51,213,153},
[3]uint8{51,213,204}, [3]uint8{51,213,255}, [3]uint8{51,255,0}, [3]uint8{51,255,51},
[3]uint8{51,255,102}, [3]uint8{51,255,153}, [3]uint8{51,255,204}, [3]uint8{51,255,255},
[3]uint8{102,0,0}, [3]uint8{102,0,51}, [3]uint8{102,0,102}, [3]uint8{102,0,153},
[3]uint8{102,0,204}, [3]uint8{102,0,255}, [3]uint8{102,43,0}, [3]uint8{102,43,51},
[3]uint8{102,43,102}, [3]uint8{102,43,153}, [3]uint8{102,43,204}, [3]uint8{102,43,255},
[3]uint8{102,85,0}, [3]uint8{102,85,51}, [3]uint8{102,85,102}, [3]uint8{102,85,153},
[3]uint8{102,85,204}, [3]uint8{102,85,255}, [3]uint8{102,128,0}, [3]uint8{102,128,51},
[3]uint8{102,128,102}, [3]uint8{102,128,153}, [3]uint8{102,128,204}, [3]uint8{102,128,255},
[3]uint8{102,170,0}, [3]uint8{102,170,51}, [3]uint8{102,170,102}, [3]uint8{102,170,153},
[3]uint8{102,170,204}, [3]uint8{102,170,255}, [3]uint8{102,213,0}, [3]uint8{102,213,51},
[3]uint8{102,213,102}, [3]uint8{102,213,153}, [3]uint8{102,213,204}, [3]uint8{102,213,255},
[3]uint8{102,255,0}, [3]uint8{102,255,51}, [3]uint8{102,255,102}, [3]uint8{102,255,153},
[3]uint8{102,255,204}, [3]uint8{102,255,255}, [3]uint8{153,0,0}, [3]uint8{153,0,51},
[3]uint8{153,0,102}, [3]uint8{153,0,153}, [3]uint8{153,0,204}, [3]uint8{153,0,255},
[3]uint8{153,43,0}, [3]uint8{153,43,51}, [3]uint8{153,43,102}, [3]uint8{153,43,153},
[3]uint8{153,43,204}, [3]uint8{153,43,255}, [3]uint8{153,85,0}, [3]uint8{153,85,51},
[3]uint8{153,85,102}, [3]uint8{153,85,153}, [3]uint8{153,85,204}, [3]uint8{153,85,255},
[3]uint8{153,128,0}, [3]uint8{153,128,51}, [3]uint8{153,128,102}, [3]uint8{153,128,153},
[3]uint8{153,128,204}, [3]uint8{153,128,255}, [3]uint8{153,170,0}, [3]uint8{153,170,51},
[3]uint8{153,170,102}, [3]uint8{153,170,153}, [3]uint8{153,170,204}, [3]uint8{153,170,255},
[3]uint8{153,213,0}, [3]uint8{153,213,51}, [3]uint8{153,213,102}, [3]uint8{153,213,153},
[3]uint8{153,213,204}, [3]uint8{153,213,255}, [3]uint8{153,255,0}, [3]uint8{153,255,51},
[3]uint8{153,255,102}, [3]uint8{153,255,153}, [3]uint8{153,255,204}, [3]uint8{153,255,255},
[3]uint8{204,0,0}, [3]uint8{204,0,51}, [3]uint8{204,0,102}, [3]uint8{204,0,153},
[3]uint8{204,0,204}, [3]uint8{204,0,255}, [3]uint8{204,43,0}, [3]uint8{204,43,51},
[3]uint8{204,43,102}, [3]uint8{204,43,153}, [3]uint8{204,43,204}, [3]uint8{204,43,255},
[3]uint8{204,85,0}, [3]uint8{204,85,51}, [3]uint8{204,85,102}, [3]uint8{204,85,153},
[3]uint8{204,85,204}, [3]uint8{204,85,255}, [3]uint8{204,128,0}, [3]uint8{204,128,51},
[3]uint8{204,128,102}, [3]uint8{204,128,153}, [3]uint8{204,128,204}, [3]uint8{204,128,255},
[3]uint8{204,170,0}, [3]uint8{204,170,51}, [3]uint8{204,170,102}, [3]uint8{204,170,153},
[3]uint8{204,170,204}, [3]uint8{204,170,255}, [3]uint8{204,213,0}, [3]uint8{204,213,51},
[3]uint8{204,213,102}, [3]uint8{204,213,153}, [3]uint8{204,213,204}, [3]uint8{204,213,255},
[3]uint8{204,255,0}, [3]uint8{204,255,51}, [3]uint8{204,255,102}, [3]uint8{204,255,153},
[3]uint8{204,255,204}, [3]uint8{204,255,255}, [3]uint8{255,0,0}, [3]uint8{255,0,51},
[3]uint8{255,0,102}, [3]uint8{255,0,153}, [3]uint8{255,0,204}, [3]uint8{255,0,255},
[3]uint8{255,43,0}, [3]uint8{255,43,51}, [3]uint8{255,43,102}, [3]uint8{255,43,153},
[3]uint8{255,43,204}, [3]uint8{255,43,255}, [3]uint8{255,85,0}, [3]uint8{255,85,51},
[3]uint8{255,85,102}, [3]uint8{255,85,153}, [3]uint8{255,85,204}, [3]uint8{255,85,255},
[3]uint8{255,128,0}, [3]uint8{255,128,51}, [3]uint8{255,128,102}, [3]uint8{255,128,153},
[3]uint8{255,128,204}, [3]uint8{255,128,255}, [3]uint8{255,170,0}, [3]uint8{255,170,51},
[3]uint8{255,170,102}, [3]uint8{255,170,153}, [3]uint8{255,170,204}, [3]uint8{255,170,255},
[3]uint8{255,213,0}, [3]uint8{255,213,51}, [3]uint8{255,213,102}, [3]uint8{255,213,153},
[3]uint8{255,213,204}, [3]uint8{255,213,255}, [3]uint8{255,255,0}, [3]uint8{255,255,51},
[3]uint8{255,255,102}, [3]uint8{255,255,153}, [3]uint8{255,255,204}, [3]uint8{255,255,255},
[3]uint8{0,0,0}, [3]uint8{0,0,0}, [3]uint8{0,0,0}, [3]uint8{0,0,0}}
