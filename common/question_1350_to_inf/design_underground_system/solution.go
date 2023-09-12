package design_underground_system

// https://leetcode.cn/problems/design-underground-system/
type undergroundSystem struct {
	checkInInfo map[int]*info
	record      map[[2]string]*time // 总时间  人次
}
type info struct {
	checkInStationName string
	checkInTime        int
}
type time struct {
	wholeTime int
	times     int
}

func constructor6() undergroundSystem {
	return undergroundSystem{
		checkInInfo: make(map[int]*info),
		record:      make(map[[2]string]*time),
	}
}
func (u *undergroundSystem) checkIn(id int, stationName string, t int) {
	node := &info{
		checkInStationName: stationName,
		checkInTime:        t,
	}
	u.checkInInfo[id] = node
}
func (u *undergroundSystem) checkOut(id int, stationName string, t int) {
	preStation := u.checkInInfo[id].checkInStationName
	preTime := u.checkInInfo[id].checkInTime
	if u.record[[2]string{preStation, stationName}] == nil {
		newtime := &time{
			wholeTime: t - preTime,
			times:     1,
		}
		u.record[[2]string{preStation, stationName}] = newtime
	} else {
		u.record[[2]string{preStation, stationName}].wholeTime += t - preTime
		u.record[[2]string{preStation, stationName}].times++
	}
}

func (u *undergroundSystem) getAverageTime(startStation string, endStation string) float64 {
	time := float64(u.record[[2]string{startStation, endStation}].wholeTime)
	times := float64(u.record[[2]string{startStation, endStation}].times)
	return time / times
}

/**
 * Your undergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
