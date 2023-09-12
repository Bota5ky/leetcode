package design_underground_system

// UndergroundSystem 1396. 设计地铁系统 https://leetcode.cn/problems/design-underground-system/
type UndergroundSystem struct {
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

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkInInfo: make(map[int]*info),
		record:      make(map[[2]string]*time),
	}
}

func (u *UndergroundSystem) checkIn(id int, stationName string, t int) {
	node := &info{
		checkInStationName: stationName,
		checkInTime:        t,
	}
	u.checkInInfo[id] = node
}

func (u *UndergroundSystem) checkOut(id int, stationName string, t int) {
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

func (u *UndergroundSystem) getAverageTime(startStation string, endStation string) float64 {
	time := float64(u.record[[2]string{startStation, endStation}].wholeTime)
	times := float64(u.record[[2]string{startStation, endStation}].times)
	return time / times
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
