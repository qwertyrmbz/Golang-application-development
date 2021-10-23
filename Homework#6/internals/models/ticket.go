package models

type Ticket struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	SeatCount int     `json:"seat_count"`
	//Status      BookingStatus `json:"status"`
	//User        User          `json:"user"`
	//MovieShowID int           `json:"-"`
	//MovieShow   MovieShow     `json:"movie_show"`
}
