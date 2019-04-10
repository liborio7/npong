package game

import (
    "fmt"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/golang/geo/r1"
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
    "github.com/liborio7/npong/game/player"
    "github.com/liborio7/npong/game/camera"
    "github.com/liborio7/npong/game/terrain"
    "github.com/liborio7/npong/system"
    "golang.org/x/image/colornames"
    "runtime"
    "time"
)

func Run() {
    // create window
    var width float64 = 600
    var height float64 = 400
    cfg := pixelgl.WindowConfig{
        Title:  "Pixel Rocks!",
        Bounds: pixel.R(0, 0, width, height),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    // create world
    world := core.NewWorld()

    // create managers
    boyManager := player.NewManager()
    cameraManager := camera.NewManager()
    groundManager := terrain.NewManager()

    // create player and its camera
    boyEntity := boyManager.NewEntity(win, &r2.Point{Y: 50})
    boyCamera := cameraManager.NewEntity(win, boyEntity)
    world.AddEntities(
        boyEntity,
        boyCamera,
    )

    // create terrain
    world.AddEntities(
        groundManager.NewEntity(
            &colornames.Mediumspringgreen,
            &r2.Point{},
            &r2.Rect{
                X: r1.Interval{Hi: 1000},
                Y: r1.Interval{Hi: 50},
            }),

        groundManager.NewEntity(
            &colornames.Palevioletred,
            &r2.Point{X: 200, Y: 200},
            &r2.Rect{
                X: r1.Interval{Hi: 200},
                Y: r1.Interval{Hi: 200},
            }),

        groundManager.NewEntity(&colornames.Blanchedalmond,
            &r2.Point{X: 200, Y: 510},
            &r2.Rect{
                X: r1.Interval{Hi: 200},
                Y: r1.Interval{Hi: 200},
            }),

        groundManager.NewEntity(&colornames.Cornflowerblue,
            &r2.Point{X: 510, Y: 510},
            &r2.Rect{
                X: r1.Interval{Hi: 200},
                Y: r1.Interval{Hi: 200},
            }),

        groundManager.NewEntity(&colornames.Lightgoldenrodyellow,
            &r2.Point{X: 510, Y: 200},
            &r2.Rect{
                X: r1.Interval{Hi: 200},
                Y: r1.Interval{Hi: 200},
            }),
    )

    // add systems
    world.AddSystems(
        system.NewPlayerControl(),
        system.NewGravity(),
        system.NewPlayerPhysics(),
        system.NewMovement(),
        system.NewFollow(),
        system.NewAutomata(),
        system.NewScripting(),
        system.NewAnimation(),
        system.NewGraphics(),
    )

    // start world
    world.Init()
    ticker := time.NewTicker(20 * time.Millisecond)
    time.Now()
    for t := range ticker.C {

        win.Clear(colornames.Forestgreen)

        world.Update(t)
        world.Render(t)

        win.Update()

        // printMemUsage()
    }
}

func printMemUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    // For info on each, see: https://golang.org/pkg/runtime/#MemStats
    fmt.Printf("Heap Alloc = %v MiB", bToMb(m.HeapAlloc))
    fmt.Printf("\tHeap Idle = %v MiB", bToMb(m.HeapIdle))
    fmt.Printf("\tHeap Objects = %v MiB", bToMb(m.HeapObjects))
    fmt.Printf("\tHeap Inuse = %v MiB", bToMb(m.HeapInuse))
    fmt.Printf("\tStack Inuse = %v MiB", bToMb(m.StackInuse))
    fmt.Printf("\tMallocs = %v MiB", bToMb(m.Mallocs))
    fmt.Printf("\tFrees = %v MiB", bToMb(m.Frees))
    fmt.Printf("\tAlloc = %v MiB", bToMb(m.Alloc))
    fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
    fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
    fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
