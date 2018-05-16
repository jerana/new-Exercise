package main

//Problem Statement :
//System is revceiving events as below format from clusterNodes. You have
//build  2 API  as ProcessEvents(e events) and getAverageIndegree60Sec()

//10Sec : fork : {read , write , open} ; indegress= 3
//20Sec: open : { fork, read, write}   ; indegeres = 3
//30 Sec:read : {fork, open}	       ; indegress = 2
//Indegress means ; Number of edge from  one events to rest of other events ;

// Calculate Average indergess = 3 +3 +2/ 3(as syscalls context)
//like in above example ,
//GraphNode for each slot ; which describe given GraphNode for syscalls events

type events struct {
	timeSlot int
	syscalls map[string]struct{}
}
type tGraphNode struct {
	time uint64
	node map[string]map[string]bool
}

//get new GraphNode for given timeSlot
func getGraphNode(time uint64) tGraphNode {
	t := new(tGraphNode)
	t.time = time
	t.node = make(map[string]map[string]bool, 0)
	return t
}

var timeSlotMap map[int]tGraphNode

//Init timeslotHash Map
//this hashMap hold maximum 60 sec time slot data
func Init() {
	timeSlotMap = make(map[int]tGraphNode, 60)
}

//This function will process receive events and do lookup into timeSLot hashMap
//which is buckets based on 1 sec slot
func eventProcess(event events) {
	//Calculate where this events falls
	curSlot = (time.Now().Seconds - lastTimeStamp) % 60
	lastTimeStamp = time.Now().Seconds()
	if v, ok := timeSlotMap[curSlot]; ok { // Data Fall into slot which has older data
		if lastTimeStamp-v.time > 60 {
			//events falls in slot which is older data is timeStamp is more
			//than 60 sec ; therefore overwrite old data
			gNode := getGraphNode(lastTimeStamp)
			timeSlotMap[lastTimeStamp] = gNode //allocate new graph node and overwirte with older one
			gMap := gNode.node
			for sysEvent := range events.syscalls {
				var m map[string]int
				var ok bool
				if m, ok = gMap[sysEvent]; !ok {
					m = make(map[string]int, 0)
					gMap[sysEvent] = m
				}
				for sys := range sysEvent {
					m[sys]++
				}

			}
		} else { //Events falls in 1 sec slots; which have older data but older data timestamp is fall into 60sec
			gMap := v
			for sysEvent := range events.syscalls {
				var m map[string]int
				var ok bool
				if m, ok = gMap[sysEvent]; !ok {
					m = make(map[string]int, 0)
				}
				for sys := range sysEvent {
					m[sys]++
				}

			}

		}

	} else { //New data ; no Old data exist
		gNode := getGraphNode(lastTimeStamp)
		gMap = gNode.node
		for sysEvent := range events.syscalls {
			var m map[string]int
			var ok bool
			if m, ok = gMap[sysEvent]; !ok {
				m = make(map[string]int, 0)
			}
			m := gNode[sysEvent]
			for sys := range sysEvent {
				m[sys] = 1
			}

		}
	}
}

//Walk TimeSlot HashMap which is 60 slots longs
func findAverageIndegree60Sec() int {
	var calls map[string]bool
	var averageIndegree int
	for walk := range timeSlotMap {
		gMap := walk.Node
		for sysCalls, m := range gMap {
			calls[sysCalls] = true
			for _, v := range m {
				averageIndegree++
			}

		}
	}
	return averageIndegree
}
