package sort


func selectionSort(slice []int) []int {
  sortedList := make([]int, len(slice))
  sliceLength := len(slice)

  for i := 0; i < sliceLength; i++ {
    minIndex := findSmallest(slice)
    sortedList[i] = slice[minIndex]
    slice = append(slice[0:minIndex], slice[minIndex+1:len(slice)]...)
  }

  return sortedList
}

func findSmallest(slice []int) int {
  min := slice[0]
  minIndex := 0

  for i := 1; i < len(slice); i++ {
    if slice[i] < min {
      min = slice[i]
      minIndex = i
    }
  }

  return minIndex
}