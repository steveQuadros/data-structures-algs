package doordash

// DeliveryOrder expected input to be a slice of strings
// order[i] = P/D then a number
// P == Pickup
// D == Delivery
// ex: []string{"P1", "P2", "D1", "D2"}
//     true
// delivery cannot happen before pickup
// cannot delivery a product that has not been picked up
// must deliver all Pickups

const (
	PickedUp  = false
	Delivered = true
)

func DeliveryOrder(order []string) bool {
	pickups := make(map[string]bool)
	for _, o := range order {
		if o[0] == 'P' {
			if _, ok := pickups[o]; ok {
				// cannot pick up same product if it's in list
				return false
			} else {
				pickups[o] = PickedUp
			}
		} else {
			pickup := "P" + o[1:]
			if status, ok := pickups[pickup]; !ok || status == Delivered {
				return false
			} else {
				pickups[pickup] = Delivered
			}
		}
	}

	for pu, d := range pickups {
		if d == Delivered {
			delete(pickups, pu)
		}
	}

	return len(pickups) == 0
}
