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

type DistanceTo struct {
	id   uint32
	dist float64
}

type Box struct {
	id        uint32
	coord     Coordinates
	prev      *Box
	next      *Box
	distances []DistanceTo
}

func (b *Box) Compare(id uint32) bool {
	return b.id == id
}

func Day08Part01() {
	input := utils.ReadLines(8, true)

	boxes := getBoxes(input)
	boxes = calcDistancesForOther(boxes)

	circutParts := []DistanceTo{}

	for i := 0; i <= 10; {
		i++
		lowKey := uint32(0)

		for idx, b := range boxes {
			if lowKey == 0 {
				lowKey = idx
			}

			if b.distances[0].dist < boxes[lowKey].distances[0].dist {
				lowKey = idx
			}
		}

		toUse := boxes[lowKey].distances[0]
		fmt.Println(toUse)

		setNewDistances(boxes, lowKey, boxes[lowKey].distances[1:])
		setNewDistances(boxes, toUse.id, boxes[toUse.id].distances[1:])
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

func calcDistancesForOther(boxes map[uint32]Box) map[uint32]Box {
	for key, box := range boxes {
		distances := []DistanceTo{}

		for _, nextBox := range boxes {
			if box.id == nextBox.id {
				continue
			}

			distances = append(distances, DistanceTo{
				id:   nextBox.id,
				dist: calcDistance(box.coord.x, nextBox.coord.x, box.coord.y, nextBox.coord.y, box.coord.z, nextBox.coord.z),
			})

		}

		sort.Slice(distances, func(i, j int) bool {
			return distances[i].dist < distances[j].dist
		})

		setNewDistances(boxes, key, distances)
	}

	return boxes
}

func setNewDistances(boxes map[uint32]Box, key uint32, distances []DistanceTo) {
	if entry, ok := boxes[key]; ok {
		entry.distances = distances
		boxes[key] = entry
	}
}

func calcDistance(p1, q1, p2, q2, p3, q3 float64) float64 {
	xSub, ySub, zSub := p1-q1, p2-q2, p3-q3
	xPow, yPow, zPov := math.Pow(xSub, 2), math.Pow(ySub, 2), math.Pow(zSub, 2)
	distance := math.Sqrt(xPow + yPow + zPov)

	return distance
}

func getBoxes(input []string) map[uint32]Box {
	boxes := map[uint32]Box{}

	for _, l := range input {
		coordinates := getCoordinates(l)

		box := Box{
			id: uuid.New().ID(),
			coord: Coordinates{
				x: coordinates[0],
				y: coordinates[1],
				z: coordinates[2],
			},
			distances: []DistanceTo{},
			prev:      nil,
			next:      nil,
		}

		boxes[box.id] = box
	}

	return boxes
}
