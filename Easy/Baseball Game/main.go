func calPoints(ops []string) int {
    works := make([]int, 1010)
    pointer := -1
    result := 0
    for _, op := range ops {
        if op != "C" && op != "D" && op != "+" {
            number, _ := strconv.Atoi(op)
            pointer++
            works[pointer] = number
            result += number
        } else if op == "+" {
            result += works[pointer - 1] + works[pointer]
            pointer++
            works[pointer] = works[pointer - 2] + works[pointer - 1]
        } else if op == "C" {
            result -= works[pointer]
            pointer--
        } else if op == "D" {
            result += 2 * works[pointer]
            pointer++
            works[pointer] = 2 * works[pointer - 1]
        }
    }
    return result
}