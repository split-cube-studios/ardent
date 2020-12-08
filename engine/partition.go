package engine

// PartitionMap handles spatial partitioning of PartitionEntry.
type PartitionMap struct {
	partitions    map[[2]int][]PartitionEntry
	partitionSize int

	buffer       map[string][]PartitionEntry
	linearBuffer []PartitionEntry
}

// Partionable is a type that can be
// used in a PartitionMap.
type PartitionEntry interface {
	IsDisposed() bool
	Position() Vec2
	Class() string
}

// NewPartitionMap returns a PartitionMap
// with a given range interval, and initial map
// bucket count.
func NewPartitionMap(partitionSize, bucketCount int) *PartitionMap {
	return &PartitionMap{
		partitions:    make(map[[2]int][]PartitionEntry, bucketCount),
		partitionSize: partitionSize,
		buffer:        make(map[string][]PartitionEntry),
	}
}

// Add inserts a PartitionEntry into the PartitionMap,
// and returns the map key.
func (pm *PartitionMap) Add(e PartitionEntry) [2]int {
	key := pm.positionToKey(e.Position())

	pm.partitions[key] = append(
		pm.partitions[key],
		e,
	)

	return key
}

// Tick updates partitions around a position.
func (pm *PartitionMap) Tick(
	pos Vec2,
	size int,
	tickFunc func([]PartitionEntry),
) {
	px, py := int(pos.X)/pm.partitionSize, int(pos.Y)/pm.partitionSize

	// update buffer and clear partitions
	for x := px - size; x <= px+size; x++ {
		for y := py - size; y <= py+size; y++ {

			key := [2]int{x, y}
			entries := pm.partitions[key]

			for _, e := range entries {
				if e.IsDisposed() {
					continue
				}

				pm.buffer[e.Class()] = append(
					pm.buffer[e.Class()],
					e,
				)
			}

			// clear partition
			for i := 0; i < len(entries); i++ {
				entries[i] = nil
			}
			pm.partitions[key] = entries[:0]
		}
	}

	// flatten buffer for updates
	for _, entries := range pm.buffer {
		pm.linearBuffer = append(pm.linearBuffer, entries...)

	}

	if tickFunc != nil {
		tickFunc(pm.linearBuffer)
	}

	// Clear buffers
	for class, entries := range pm.buffer {
		pm.buffer[class] = pm.buffer[class][:0]

		// return entry to partition map
		for _, entry := range entries {
			pm.Add(entry)
		}
	}
	pm.linearBuffer = pm.linearBuffer[:0]
}

// Class returns all entries of a given class in the current buffer.
func (pm *PartitionMap) Class(class string) []PartitionEntry {
	return pm.buffer[class]
}

func (pm *PartitionMap) positionToKey(pos Vec2) [2]int {
	return [2]int{int(pos.X) / pm.partitionSize, int(pos.Y) / pm.partitionSize}
}
