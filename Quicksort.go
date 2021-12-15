package main

func SortNumbers(numbers []int) []int {
  if(len(numbers) <= 1 || numbers == nil) {
    return numbers
  }
  quickSort(numbers, 0, len(numbers) - 1)
  return numbers
}

func quickSort(numbers []int, start int, end int) {
  if start >= end {
    return
  }
  pivot := start
  left := start + 1
  right := end

  for left <= right {
    if(numbers[pivot] < numbers[left] && numbers[pivot] > numbers[right]) {
      swap(numbers, left, right)
    }
    
    if(numbers[pivot] >= numbers[left]) {
      left++
    }
    
    if(numbers[pivot] <= numbers[right]) {
      right--
    }
  }
    
  swap(numbers, pivot, right)
  quickSort(numbers, start, right - 1)
  quickSort(numbers, right + 1, end)
}

func swap(array []int, index int, target int) {
  temp := array[index]
  array[index] = array[target]
  array[target] = temp
}
