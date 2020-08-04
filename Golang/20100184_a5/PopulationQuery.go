package main

import (
    "fmt"
    "os"
    "strconv"
    "math"
    "encoding/csv"
)

type CensusGroup struct {
    population int
    latitude, longitude float64
}

func ParseCensusData(fname string) ([]CensusGroup, error) {
    file, err := os.Open(fname)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    records, err := csv.NewReader(file).ReadAll()
    if err != nil {
        return nil, err
    }
    censusData := make([]CensusGroup, 0, len(records))

    for _, rec := range records {
        if len(rec) == 7 {
            population, err1 := strconv.Atoi(rec[4])
            latitude, err2 := strconv.ParseFloat(rec[5], 64)
            longitude, err3 := strconv.ParseFloat(rec[6], 64)
            if err1 == nil && err2 == nil && err3 == nil {
                latpi := latitude * math.Pi / 180
                latitude = math.Log(math.Tan(latpi) + 1 / math.Cos(latpi))
                censusData = append(censusData, CensusGroup{population, latitude, longitude})
            }
        }
    }

    return censusData, nil
}

func main () {
    if len(os.Args) < 4 {
        fmt.Printf("Usage:\nArg 1: file name for input data\nArg 2: number of x-dim buckets\nArg 3: number of y-dim buckets\nArg 4: -v1, -v2, -v3, -v4, -v5, or -v6\n")
        return
    }
    fname, ver := os.Args[1], os.Args[4]
    xdim, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println(err)
        return
    }
    ydim, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println(err)
        return
    }
    censusData, err := ParseCensusData(fname)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Some parts may need no setup code
    var sa, la, so, lo, block_width, block_height float64
    done := make(chan bool)
    arr := make([][]int, ydim)
    for i:= 0; i <ydim; i++{
        arr[i] = make([]int, xdim)
    }
    switch ver {
    case "-v1":
        sa = censusData[0].latitude
        la = censusData[0].latitude
        so = censusData[0].longitude
        lo = censusData[0].longitude
        for _, dat := range censusData {
            if dat.latitude > la{
                la = dat.latitude
            } else if dat.latitude < sa{
                sa = dat.latitude
            }
            if dat.longitude > lo{
                lo = dat.longitude
            } else if dat.longitude < so{
                so = dat.longitude
            }
        }
        block_width = (lo - so) / float64(xdim)
        block_height = (la - sa) / float64(ydim)
    case "-v2":
        
    case "-v3":
        sa = censusData[0].latitude
        la = censusData[0].latitude
        so = censusData[0].longitude
        lo = censusData[0].longitude
        for _, dat := range censusData {
            if dat.latitude > la{
                la = dat.latitude
            } else if dat.latitude < sa{
                sa = dat.latitude
            }
            if dat.longitude > lo{
                lo = dat.longitude
            } else if dat.longitude < so{
                so = dat.longitude
            }
        }
        sa += 0.000001
        la += 0.000001
        so += 0.000001
        lo += 0.000001
        block_width = (lo - so) / float64(xdim)
        block_height = (la - sa) / float64(ydim)
        for i:= 0; i < ydim; i++ {
            for j:=0; j< xdim; j++{
                arr[i][j] = 0;
            }
        }
        for _, dat := range censusData {
            long := int((dat.longitude - so) /block_width)
            lat := int((dat.latitude - sa)/block_height)
            arr[lat][long] += dat.population
        }
        for i:=0; i< ydim; i++{
            for j:=0; j<xdim; j++{
                if i>0 {
                    arr[i][j] += arr[i-1][j]
                    if j>0{
                        arr[i][j] += arr[i][j-1]
                        arr[i][j] -= arr[i-1][j-1]
                    }
                } else if j>0{
                    arr[i][j] += arr[i][j-1]
                }

            }
        }
        // for i:=0; i< ydim; i++{
        //     for j:=0; j<xdim; j++{
        //         fmt.Print(arr[i][j])
        //         fmt.Print(" ")
        //     }
        //     fmt.Println()
        // }
    case "-v4":
        // YOUR SETUP CODE FOR PART 4
    case "-v5":
        // YOUR SETUP CODE FOR PART 5
    case "-v6":
        // YOUR SETUP CODE FOR PART 6
    default:
        fmt.Println("Invalid version argument")
        return
    }

    for {
        var west, south, east, north int
        n, err := fmt.Scanln(&west, &south, &east, &north)
        if n != 4 || err != nil || west<1 || west>xdim || south<1 || south>ydim || east<west || east>xdim || north<south || north>ydim {
            break
        }

        var population int
        var percentage float64
        var total int
        switch ver {
        case "-v1":
            population = 0
            total = 0
            for _, dat := range censusData {
                total += dat.population
                long := int((dat.longitude - so) /block_width) + 1
                lat := int((dat.latitude - sa)/block_height) + 1
                if long <= east && long >= west && lat <= north && lat >= south{
                    population += dat.population
                }        
            }
            percentage = (float64(population)/float64(total)) * 100.0 
        case "-v2":
        //     <-done
        //     go func () {
        //         population = 0
        //         total = 0
        //         for i:=0; i<  {
        //             total += dat.population
        //             long := int((dat.longitude - so) /block_width) + 1
        //             lat := int((dat.latitude - sa)/block_height) + 1
        //             if long <= east && long >= west && lat <= north && lat >= south{
        //                 population += dat.population
        //             }        
        //         }
        //         percentage = (float64(population)/float64(total)) * 100.0 
        //     }()
        case "-v3":
            population = 0
            population += arr[north - 1][east - 1]
            if south > 1 && west > 1 {
                population += arr[south -2][west - 2]
            }
            if south > 1 {
                population -= arr[south][east - 1]
            }
            if west > 1{
                population -= arr[north - 1][west - 2]
            }
            percentage = float64(population)/float64(arr[ydim -1][xdim-1]) * 100.0
        case "-v4":
            // YOUR QUERY CODE FOR PART 4
        case "-v5":
            // YOUR QUERY CODE FOR PART 5
        case "-v6":
            // YOUR QUERY CODE FOR PART 6
        }

        fmt.Printf("%v %.2f%%\n", population, percentage)
    }
}
