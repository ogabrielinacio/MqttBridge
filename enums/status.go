package enums 

type EStatus int

const (
    Online EStatus = iota
    Offline
    SensorError
)

func (s EStatus) String() string {
    return [...]string{"Online", "Offline", "SensorError"}[s]
}