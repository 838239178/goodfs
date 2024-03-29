package hashslot

import (
	"common/util"
	"common/util/math"
	"common/util/slices"
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"strings"
)

const MaxSlot = 16383

var (
	ErrOverlap = errors.New("overlap slots")
)

func formatInt(slot string) (int, int, error) {
	rg := strings.Split(slot, "-")
	if len(rg) != 2 {
		return 0, 0, fmt.Errorf("slot '%s' must format like '0-100'", slot)
	}
	// convert start slot
	start, err := strconv.Atoi(rg[0])
	if err != nil {
		return 0, 0, err
	} else if start < 0 {
		return 0, 0, fmt.Errorf("slot must greater than 0")
	}
	// convert end slot
	end, err := strconv.Atoi(rg[1])
	if err != nil {
		return 0, 0, err
	} else if end < 0 {
		return 0, 0, fmt.Errorf("slot must greater than 0")
	}
	return start, end, nil
}

func sortAndValidate(edges EdgeList) error {
	sort.Sort(edges)
	// validate overlap
	for i, edge := range edges {
		if edge.Start >= edge.End {
			return fmt.Errorf("start must less than end: %s", edge)
		}
		if i == 0 {
			continue
		}
		pre := edges[i-1]
		if pre.Start == edge.Start || pre.End > edge.Start {
			return fmt.Errorf("%w: %s and %s", ErrOverlap, pre, edge)
		}
	}
	return nil
}

func wrapSlotsOriginal(slots []string, identify string) (EdgeList, error) {
	res := make(EdgeList, 0, len(slots))
	for _, slot := range slots {
		start, end, err := formatInt(slot)
		if err != nil {
			return nil, err
		}
		res = append(res, &Edge{
			Start: start,
			End:   end,
			Value: identify,
		})
	}
	return res, nil
}

// WrapSlotsToEdges Range: [0,100)
func WrapSlotsToEdges(slots []string, identify string) (EdgeList, error) {
	res, err := wrapSlotsOriginal(slots, identify)
	if err != nil {
		return nil, err
	}
	return res, sortAndValidate(res)
}

// WrapSlots slotsMap(key='identify', value=[]string{'0-100','110-221'}). Range: [0,100)
func WrapSlots(slotsMap map[string][]string) (IEdgeProvider, error) {
	if len(slotsMap) == 0 {
		return nil, fmt.Errorf("empty slotsMap")
	}
	res := make(EdgeList, 0, len(slotsMap))
	for value, slots := range slotsMap {
		edges, err := wrapSlotsOriginal(slots, value)
		if err != nil {
			return nil, err
		}
		res = append(res, edges...)
	}
	if err := sortAndValidate(res); err != nil {
		return nil, err
	}
	return &EdgeProvider{res}, nil
}

func GetStringIdentify(str string, provide IEdgeProvider) (string, error) {
	return GetDataIdentify([]byte(str), provide)
}

func GetDataIdentify(data []byte, provider IEdgeProvider) (string, error) {
	return GetSlotIdentify(CalcBytesSlot(data), provider)
}

func CalcBytesSlot(bt []byte) int {
	return int(crc32.ChecksumIEEE(bt) & MaxSlot)
}

func CopyOfEdges(identify string, provider IEdgeProvider) EdgeList {
	var res EdgeList
	list := provider.get()
	for _, v := range list {
		if identify == "" || v.Value == identify {
			res = append(res, &Edge{
				Start: v.Start,
				End:   v.End,
				Value: v.Value,
			})
		}
	}
	return res
}

func GetSlotIdentify(slot int, provider IEdgeProvider) (string, error) {
	edges := provider.get()
	if idx := sort.Search(len(edges), func(i int) bool {
		return slot < edges[i].End
	}); idx < len(edges) && slot >= edges[idx].Start {
		return edges[idx].Value, nil
	}
	return "", fmt.Errorf("slots assigment error, cannot find slot %d", slot)
}

func IsSlotInEdges(slot int, edges EdgeList) bool {
	if idx := sort.Search(len(edges), func(i int) bool {
		return slot < edges[i].End
	}); idx < len(edges) {
		return slot >= edges[idx].Start
	}
	return false
}

// FindRangeCurrentData find identifies in range [start,end)
func FindRangeCurrentData(start, end int, provider IEdgeProvider) (res EdgeList, fully bool) {
	edges := provider.get()
	rangeLen := end - start
	for _, edge := range edges {
		ed := math.MinInt(edge.End, end)
		if edge.Start < start && start < edge.End {
			res = append(res, &Edge{start, ed, edge.Value})
			rangeLen -= ed - start
		} else if edge.Start >= start && edge.Start < end {
			res = append(res, &Edge{edge.Start, ed, edge.Value})
			rangeLen -= ed - edge.Start
		}
	}
	fully = rangeLen == 0
	return
}

// IsValidEdge whether edge is contains in provider and fully belongs to single identify
func IsValidEdge(edge *Edge, provider IEdgeProvider) bool {
	res, fully := FindRangeCurrentData(edge.Start, edge.End, provider)
	if !fully {
		return false
	}
	if len(res) > 1 || len(res) == 0 {
		return false
	}
	return res[0].Value == edge.Value
}

// CombineEdges combine dest to src.
func CombineEdges(src EdgeList, dest EdgeList) EdgeList {
	newEdges := make(EdgeList, 0, len(src)+len(dest))
	src = append(src, dest...)
	sort.Sort(src)
	for _, edge := range src {
		if len(newEdges) <= 0 {
			newEdges = append(newEdges, edge)
			continue
		}
		last := slices.Last(newEdges)
		if edge.Start <= last.End {
			last.Start = math.MinInt(last.Start, edge.Start)
			last.End = math.MaxInt(last.End, edge.End)
			continue
		}
		newEdges = append(newEdges, edge)
	}
	return newEdges
}

// RemoveEdges remove edges contains in dest from src
func RemoveEdges(src EdgeList, dest EdgeList) EdgeList {
	buf1 := make(EdgeList, 0, len(src)+len(dest))
	buf2 := make(EdgeList, 0, len(src)+len(dest))
	for i, del := range dest {
		// use two buffer crossly to reduce memory allocates
		tmp := util.IfElse(i%2 == 0, buf1, buf2)
		for _, edge := range src {
			// edge no contains in del, leaves all
			if edge.End <= del.Start || del.End <= edge.Start {
				tmp = append(tmp, edge)
				continue
			}
			// edge contains del and leaves head part
			if edge.Start < del.Start {
				tmp = append(tmp, &Edge{Start: edge.Start, End: del.Start, Value: edge.Value})
			}
			// edge contains del and leaves tail part
			if edge.End > del.End {
				tmp = append(tmp, &Edge{Start: del.End, End: edge.End, Value: edge.Value})
			}
		}
		src = tmp
		// clear buffer
		tmp = slices.RemoveAll(tmp)
	}
	return src
}
