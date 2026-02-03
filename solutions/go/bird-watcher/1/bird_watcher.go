package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	result := 0

	for _, count := range birdsPerDay {
		result += count
	}

	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	result := 0
	weekStart := (week - 1) * 7
	weekEnd := weekStart + 7
	for i := weekStart; i < weekEnd; i++ {
		result += birdsPerDay[i]
	}
	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
