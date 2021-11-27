package main

import (
    "math"
)

type Tooth struct {
    C1 *Circle
    C2 *Circle
    C3 *Circle
    C4 *Circle
}

func (t Tooth) rotate(a float64) {
    t.C1.Rotate(a)
    t.C2.Rotate(a)
    t.C3.Rotate(a)
    t.C4.Rotate(a)
}

func BasicTooth(rotate float64) Tooth {

    scale := 7.0
    tipCenter := Point{-0.06 * scale, 1.453 * scale}
    tipRadius := 0.356 * scale
    tipStop := 1.56 // 90 gr
    bottomCenter := Point{1.633 * scale, 1.176 * scale}
    bottomRadius := 0.85 * scale
    bottomStop := 1.56 // 90 grad

    tooth := Tooth{
        &Circle{&Point{-bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
        &Circle{&Point{tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{-tipCenter.X, tipCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{bottomCenter.X, bottomCenter.Y}, bottomRadius, 0.0, 0.0},
    }

    // symmetrical, just need one tan
    _, tan := tooth.C3.InnerTangentWith(tooth.C4)
    tooth.C1.Start = math.Pi*2.0 - bottomStop
    tooth.C1.Stop = math.Pi*2.0 - tan
    tooth.C2.Start = math.Pi - tipStop
    tooth.C2.Stop = math.Pi - tan
    tooth.C3.Start = tan
    tooth.C3.Stop = tipStop
    tooth.C4.Start = math.Pi + tan
    tooth.C4.Stop = math.Pi + bottomStop

    tooth.rotate(rotate)
    
    return tooth
}

func FlexTooth(p Point, rotate float64) []Point {

    tooth := BasicTooth(rotate)

    // append points in order
    points := []Point{}
    
    c := tooth.C4.Coordinates()
    for i := len(c) - 1; i >= 0; i-- {
        points = append(points, Point{c[i].X + p.X, c[i].Y + p.Y})
    }
    c = tooth.C3.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X + p.X, point.Y + p.Y})
    }
    c = tooth.C2.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X + p.X, point.Y + p.Y})
    }
    c = tooth.C1.Coordinates()
    for i := len(c) - 1; i >= 0; i-- {
        points = append(points, Point{c[i].X + p.X, c[i].Y + p.Y})
    }

    return points
}

func RigidTooth(p Point, rotate float64) []Point {

    tooth := BasicTooth(rotate)

    // append points in order
    points := []Point{}
        
    first := tooth.C4.PointInAngle(tooth.C4.Start)
    points = append(points, Point{first.X + p.X, first.Y + p.Y})

    c := tooth.C3.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X + p.X, point.Y + p.Y})
    }
    
    c = tooth.C2.Coordinates()
    for _, point := range c {
        points = append(points, Point{point.X + p.X, point.Y + p.Y})
    }
    
    first = tooth.C1.PointInAngle(tooth.C1.Stop)
    points = append(points, Point{first.X + p.X, first.Y + p.Y})
    
    return points
}