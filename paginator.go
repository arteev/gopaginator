package gopaginator

import (
	"fmt"
	"strconv"
)

type Pag struct {
	Name string
	Url string
}

const etc = "..."

var (
	firstPage = "<<"
	priorPage = "<"
	nextPage = ">"
	lastPage = ">>"
)

func formatUrl(pageName string, currentPage int, total int, templateUrl string) string {
	if total <= 0 {
		return ""
	}
	switch pageName {
	case firstPage:
		return fmt.Sprintf(templateUrl, "1")
	case priorPage:
		if currentPage <= 0 {
			return ""
		}
		if currentPage-1 <= 0 {
			return fmt.Sprintf(templateUrl, strconv.Itoa(currentPage))
		} else {
			return fmt.Sprintf(templateUrl, strconv.Itoa(currentPage-1))
		}

	case nextPage:
		if currentPage <= 0 {
			return ""
		}
		if currentPage+1 > total {
			return fmt.Sprintf(templateUrl, strconv.Itoa(currentPage))
		} else {
			return fmt.Sprintf(templateUrl, strconv.Itoa(currentPage+1))
		}
	case lastPage:
		return fmt.Sprintf(templateUrl, strconv.Itoa(total))
	case etc:
		return ""
	default:
		page, err := strconv.Atoi(pageName)
		if err != nil {
			return ""
		}
		if page <= 0 {
			if total > 0 {
				return fmt.Sprintf(templateUrl, "1")
			}
		}
		if page > total {
			return fmt.Sprintf(templateUrl, strconv.Itoa(total))
		}
		return fmt.Sprintf(templateUrl, pageName)
	}
}

func PagesArray(page int, total int, templateUrl string) []*Pag {
	if page <= 0 || total <= 0 || page > total {
		return []*Pag{
			&Pag{firstPage, formatUrl(firstPage, page, total, templateUrl)},
			&Pag{priorPage , formatUrl(priorPage, page, total, templateUrl)},
			&Pag{nextPage , formatUrl(nextPage, page, total, templateUrl)},
			&Pag{lastPage, formatUrl(lastPage, page, total, templateUrl)},
		}
	}
	arr := make([]string, 0)
	if total >= 7 {
		if (total - page) < 3 {
			arr = append(arr[:], "1", "2", "3", etc, strconv.Itoa(total-2), strconv.Itoa(total-1), strconv.Itoa(total))
		} else {
			cntadded := 3
			if page-1 > 0 {
				arr = append(arr[:], strconv.Itoa(page-1))
				cntadded -= 1
			}
			arr = append(arr[:], strconv.Itoa(page))
			cntadded -= 1
			for i := 1; i <= cntadded; i++ {
				arr = append(arr[:], strconv.Itoa(page+i))
			}
			arr = append(arr[:], etc, strconv.Itoa(total-2), strconv.Itoa(total-1), strconv.Itoa(total))
		}
	} else {
		for i := 1; i <= total; i++ {
			arr = append(arr[:], strconv.Itoa(i))
		}
	}
	larr := len(arr)
	resultmap := make([]*Pag, larr+4)
	resultmap[0] = &Pag{firstPage, formatUrl(firstPage, page, total, templateUrl)}
	resultmap[1] = &Pag{priorPage, formatUrl(priorPage, page, total, templateUrl)}
	resultmap[larr+2] = &Pag{nextPage, formatUrl(nextPage, page, total, templateUrl)}
	resultmap[larr+3] = &Pag{lastPage, formatUrl(lastPage, page, total, templateUrl)}
	for i:=0;i<larr; i++ {
		resultmap[i+2] = &Pag{arr[i], formatUrl(arr[i], page, total, templateUrl)}
	}
	return resultmap
}
