package domain

type Warehouse struct {
	ID                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimumCapacity    int    `json:"minimum_capacity"`
	MinimumTemperature int    `json:"minimum_temperature"`
}
