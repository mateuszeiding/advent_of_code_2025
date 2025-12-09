package solutions

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Coordinates struct {
	x float64
	y float64
	z float64
}

type OtherBox struct {
	id   uint32
	ref  *Box
	dist float64
}

type Box struct {
	id        uint32
	coord     Coordinates
	distances []OtherBox
}

func (b *Box) Compare(id uint32) bool {
	return b.id == id
}

func Day08Part01() {
	input := utils.ReadLines(8, true)
	boxes := []Box{}

	for _, l := range input {
		coordinates := getCoordinates(l)

		box := Box{
			id: uuid.New().ID(),
			coord: Coordinates{
				x: coordinates[0],
				y: coordinates[1],
				z: coordinates[2],
			},
			distances: []OtherBox{},
		}

		boxes = append(boxes, box)
	}

	boxes = calcDistancesForOther(boxes)

	circutParts := []OtherBox{}

	for i := 0; i <= 10; {
		i++
		lowIdx := 0

		for idx, b := range boxes {
			if b.distances[0].dist < boxes[lowIdx].distances[0].dist {
				lowIdx = idx
			}
		}

		toUse := boxes[lowIdx].distances[0]

		circutParts = append(circutParts, toUse)

		toUse.ref.distances = toUse.ref.distances[1:]
		boxes[lowIdx].distances = boxes[lowIdx].distances[1:]

		fId, sId := toUse.id, boxes[lowIdx].id

		fCount := 0
		for _, cp := range circutParts {
			if cp.id == fId || cp.ref.id == fId {
				fCount++
			}
		}

		if fCount == 2 {
			for i, v := range boxes {
				if v.Compare(fId) {
					boxes = append(boxes[:i], boxes[i+1:]...)
				}
			}

		}

		sCount := 0
		for _, cp := range circutParts {
			if cp.id == sId || cp.ref.id == sId {
				sCount++
			}
		}

		if sCount == 2 {
			for i, v := range boxes {
				if v.Compare(sId) {
					boxes = append(boxes[:i], boxes[i+1:]...)
				}
			}

		}

	}

	for _, cp := range circutParts {

		fmt.Printf("%v %f \n", cp.id, cp.dist)
	}
}

func getCoordinates(str string) []float64 {
	split := strings.Split(str, ",")

	coordinates := []float64{}

	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		coordinates = append(coordinates, float64(n))
	}

	return coordinates
}

func calcDistancesForOther(boxes []Box) []Box {
	for i, b := range boxes {
		d := []OtherBox{}

		for j, ib := range boxes {
			if b.id == ib.id {
				continue
			}

			d = append(d, OtherBox{
				id:   ib.id,
				dist: calcDistance(b.coord.x, ib.coord.x, b.coord.y, ib.coord.y, b.coord.z, ib.coord.z),
				ref:  &boxes[j],
			})

		}

		sort.Slice(d, func(i, j int) bool {
			return d[i].dist < d[j].dist
		})

		boxes[i].distances = d
	}

	return boxes
}

func calcDistance(p1, q1, p2, q2, p3, q3 float64) float64 {
	xSub, ySub, zSub := p1-q1, p2-q2, p3-q3
	xPow, yPow, zPov := math.Pow(xSub, 2), math.Pow(ySub, 2), math.Pow(zSub, 2)
	distance := math.Sqrt(xPow + yPow + zPov)

	return distance
}
