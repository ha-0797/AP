package main

import (
	"os"
	"fmt"
	"encoding/binary"
)

func lookup(codes []int, palette [][3]uint8) []uint8 {
    ret := make([]uint8, 0)
    for i:= 1; i < len(codes); i++{
        ret = append(ret, palette[codes[i]][2], palette[codes[i]][1], palette[codes[i]][0])
    }
    return ret
}
/* Part 3: Decompress the GIF codes to color codes
In this part, you are given a list of GIF codes and you will produce color codes that are then used by the lookup function you wrote in part 1/2. You will get a dictionary (actually a 2D array) that you need to index using the GIF codes. However, unlike parts 1/2 the number of color codes against a GIF code is variable. So the size of the resulting array is not known in advance. To avoid repeated allocations (with make or append), you are required to first calculate the size of the output array (by going over all GIF codes and finding the length of color codes against each GIF code) and then do a single allocation (using make) and finally traverse to fill in the color codes.
*/
func decompress(codes []int, dict [][]int) []int {
    ret := make ([]int,0)
    for i:= 0; i< len(codes); i++{
        for j:= 0; j < len(dict[codes[i]]); j++{
            ret = append(ret, dict[codes[i]][j])
        }
    }
    return ret
}

/* Part 4: Parallelize GIF decoding
In this part, you are to parallelize the above GIF decoding. It is easy to divide the task (like Part 2) but the complication is that we do not know where should the goroutine working on the right half of GIF codes start writing in the output array. Thus we use a PrefixSum kind of solution. We ask both halves how much space they need and then tell the total to the parent. The parent will then give a big enough slice. Since the left half told us how much space it will occupy, we can send the correct sub-slices to the left and right goroutines. Like part 2, our sequential cut-off will be 16 and the given Printf statement will help you see which tasks are generated.

Take it easy, make diagrams, go through the PrefixSum code, as this part will take some time. 

Once you finish the task, the given main function will produce "out4.bmp" which has the password for Part 5 
*/

func parallel_decompress(codes []int, dict [][]int, size chan<- int, slice <-chan []int) {
    fmt.Printf("Task to handle %d codes\n", len(codes))
    if(len(codes) > 16){ 
        size1:= make(chan int)
        slice1:= make(chan []int)
        size2:= make(chan int)
        slice2:= make(chan []int)
        go parallel_decompress(codes[0:len(codes)/2], dict, size1, slice1)
        go parallel_decompress(codes[len(codes)/2:len(codes)], dict, size2, slice2)
        len1:=(<-size1)
        len2:=(<-size2)
        size<- len1 + len2 // tell parent size
        colors:=<-slice // get slice from parent
        slice1<-colors[0:len1]
        slice2<-colors[len1:(len1+len2)]
        <-size1
        <-size2
    }    else{
        ret := make ([]int,0)
        for i:= 0; i< len(codes); i++{
            for j:= 0; j < len(dict[codes[i]]); j++{
                ret = append(ret, dict[codes[i]][j])
            }
        }
        size<-len(ret)
        colors:= <- slice
        for i:=0; i < len(ret); i++{
            colors[i] = ret[i]
        }
    }
    size<- 0 // reusing channel for done notification
}

func main() {
    // Part 3 testing code
    bmp("out3.bmp", lookup(decompress(part3_gif_codes, part3_gif_dict), part3_palette), 16, 8)

    // Part 4 testing code
    size := make(chan int)
    slice := make(chan []int)
    go parallel_decompress(part3_gif_codes, part3_gif_dict, size, slice)
    colors := make([]int, <-size)
    fmt.Println(len(colors))
    slice<- colors
    <-size // reusing channel for done notification
    bmp("out4.bmp", lookup(colors, part3_palette), 16, 8)
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

var part3_gif_codes = []int{
256, 251, 258, 259, 260, 261, 251, 200, 67, 26, 67, 34, 258, 199, 28, 258,
249, 68, 167, 261, 242, 125, 251, 243, 124, 111, 209, 249, 118, 152, 275, 259,
250, 111, 251, 250, 159, 208, 290, 258, 200, 287, 258, 277, 263, 119, 250, 117,
209, 258, 194, 298, 289, 305, 151, 68, 67, 68, 268, 251, 158, 298, 300, 288,
290, 250, 313, 67, 77, 262, 326, 260}

var part3_palette = [][3]uint8{
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

var part3_gif_dict = [][]int{
    {0},
    {1},
    {2},
    {3},
    {4},
    {5},
    {6},
    {7},
    {8},
    {9},
    {10},
    {11},
    {12},
    {13},
    {14},
    {15},
    {16},
    {17},
    {18},
    {19},
    {20},
    {21},
    {22},
    {23},
    {24},
    {25},
    {26},
    {27},
    {28},
    {29},
    {30},
    {31},
    {32},
    {33},
    {34},
    {35},
    {36},
    {37},
    {38},
    {39},
    {40},
    {41},
    {42},
    {43},
    {44},
    {45},
    {46},
    {47},
    {48},
    {49},
    {50},
    {51},
    {52},
    {53},
    {54},
    {55},
    {56},
    {57},
    {58},
    {59},
    {60},
    {61},
    {62},
    {63},
    {64},
    {65},
    {66},
    {67},
    {68},
    {69},
    {70},
    {71},
    {72},
    {73},
    {74},
    {75},
    {76},
    {77},
    {78},
    {79},
    {80},
    {81},
    {82},
    {83},
    {84},
    {85},
    {86},
    {87},
    {88},
    {89},
    {90},
    {91},
    {92},
    {93},
    {94},
    {95},
    {96},
    {97},
    {98},
    {99},
    {100},
    {101},
    {102},
    {103},
    {104},
    {105},
    {106},
    {107},
    {108},
    {109},
    {110},
    {111},
    {112},
    {113},
    {114},
    {115},
    {116},
    {117},
    {118},
    {119},
    {120},
    {121},
    {122},
    {123},
    {124},
    {125},
    {126},
    {127},
    {128},
    {129},
    {130},
    {131},
    {132},
    {133},
    {134},
    {135},
    {136},
    {137},
    {138},
    {139},
    {140},
    {141},
    {142},
    {143},
    {144},
    {145},
    {146},
    {147},
    {148},
    {149},
    {150},
    {151},
    {152},
    {153},
    {154},
    {155},
    {156},
    {157},
    {158},
    {159},
    {160},
    {161},
    {162},
    {163},
    {164},
    {165},
    {166},
    {167},
    {168},
    {169},
    {170},
    {171},
    {172},
    {173},
    {174},
    {175},
    {176},
    {177},
    {178},
    {179},
    {180},
    {181},
    {182},
    {183},
    {184},
    {185},
    {186},
    {187},
    {188},
    {189},
    {190},
    {191},
    {192},
    {193},
    {194},
    {195},
    {196},
    {197},
    {198},
    {199},
    {200},
    {201},
    {202},
    {203},
    {204},
    {205},
    {206},
    {207},
    {208},
    {209},
    {210},
    {211},
    {212},
    {213},
    {214},
    {215},
    {216},
    {217},
    {218},
    {219},
    {220},
    {221},
    {222},
    {223},
    {224},
    {225},
    {226},
    {227},
    {228},
    {229},
    {230},
    {231},
    {232},
    {233},
    {234},
    {235},
    {236},
    {237},
    {238},
    {239},
    {240},
    {241},
    {242},
    {243},
    {244},
    {245},
    {246},
    {247},
    {248},
    {249},
    {250},
    {251},
    {252},
    {253},
    {254},
    {255},
    {},
    {},
    {251,251},
    {251,251,251},
    {251,251,251,251},
    {251,251,251,251,251},
    {251,251,251,251,251,251},
    {251,200},
    {200,67},
    {67,26},
    {26,67},
    {67,34},
    {34,251},
    {251,251,199},
    {199,28},
    {28,251},
    {251,251,249},
    {249,68},
    {68,167},
    {167,251},
    {251,251,251,251,251,242},
    {242,125},
    {125,251},
    {251,243},
    {243,124},
    {124,111},
    {111,209},
    {209,249},
    {249,118},
    {118,152},
    {152,167},
    {167,251,251},
    {251,251,251,250},
    {250,111},
    {111,251},
    {251,250},
    {250,159},
    {159,208},
    {208,111},
    {111,251,251},
    {251,251,200},
    {200,167},
    {167,251,251,251},
    {251,251,242},
    {242,125,251},
    {251,200,119},
    {119,250},
    {250,117},
    {117,209},
    {209,251},
    {251,251,194},
    {194,167},
    {167,251,251,251,250},
    {250,111,209},
    {209,251,151},
    {151,68},
    {68,67},
    {67,68},
    {68,34},
    {34,251,251},
    {251,158},
    {158,167},
    {167,251,251,251,242},
    {242,125,251,251},
    {251,251,251,250,111},
    {111,251,250},
    {250,67},
    {67,68,67},
    {67,77},
    {77,251},
    {251,251,251,251,251,251,251},
    {251,251,251,251,251,251,251,251}}
