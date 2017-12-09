package main

func get_code(freq float64) int {
	const UNID = 1
	const AMATEUR = 1
	const BROADCAST = 2
	const PUBLIC = 3

	code := UNID

	// 2.2 km
	if ( 135700 <= freq && freq <= 137800 ) {
		code = AMATEUR
	}

	// 630 km
	if ( 472000 <= freq && freq <= 479000 ) {
		code = AMATEUR
	}

	// 160 m
	if ( 1800000 <= freq && freq <= 2000000 ) {
		code = AMATEUR
	}

	// 120 m bcb
	if ( 2300000 <= freq && freq <= 2495000 ) {
		code = BROADCAST
	}

	// 90 m bcb
	if ( 3200000 <= freq && freq <= 3400000 ) {
		code = BROADCAST
	}

	// 80 m
	if ( 3500000 <= freq && freq <= 4000000 ) {
		code = AMATEUR
	}

	// 75 m bcb
	if ( 3900000 <= freq && freq <= 4000000 ) {
		code = BROADCAST
	}

	// 60 m bcb
	if ( 4750000 <= freq && freq <= 5060000 ) {
		code = BROADCAST
	}

	// 60 m
	if ( 5332000 <= freq && freq <= 5405000 ) {
		code = AMATEUR
	}

	// 49 m bcb
	if ( 5800000 <= freq && freq <= 6200000 ) {
		code = BROADCAST
	}

	// 41 m bcb
	if ( 7200000 <= freq && freq <= 7450000 ) {
		code = BROADCAST
	}

	// 40 m
	if ( 7000000 <= freq && freq <= 7300000 ) {
		code = AMATEUR
	}

	// 31 m bcb
	if ( 9400000 <= freq && freq <= 9900000 ) {
		code = BROADCAST
	}

	// 30 m
	if ( 10100000 <= freq && freq <= 10150000 ) {
		code = AMATEUR
	}

	// 25 m bcb
	if ( 11600000 <= freq && freq <= 12100000 ) {
		code = BROADCAST
	}

	// 22 m bcb
	if ( 13570000 <= freq && freq <= 13870000 ) {
		code = BROADCAST
	}

	// 20 m
	if ( 14000000 <= freq && freq <= 14350000 ) {
		code = AMATEUR
	}

	// 19 m bcb
	if ( 15100000 <= freq && freq <= 15830000 ) {
		code = BROADCAST
	}

	// 17 m
	if ( 18068000 <= freq && freq <= 18168000 ) {
		code = AMATEUR
	}

	// 16 m bcb
	if ( 17480000 <= freq && freq <= 17900000 ) {
		code = BROADCAST
	}

	// 15 m bcb
	if ( 18900000 <= freq && freq <= 19020000 ) {
		code = BROADCAST
	}

	// 15 m
	if ( 21000000 <= freq && freq <= 21450000 ) {
		code = AMATEUR
	}

	// 13 m bcb
	if ( 21450000 <= freq && freq <= 21850000 ) {
		code = BROADCAST
	}

	// 12 m
	if ( 24890000 <= freq && freq <= 24990000 ) {
		code = AMATEUR
	}

	// 11 m bcb
	if ( 25600000 <= freq && freq <= 26100000 ) {
		code = BROADCAST
	}

	// 11 m cb
	if ( 26965000 <= freq && freq <= 27405000 ) {
		code = PUBLIC
	}

	// 10 m
	if ( 28000000 <= freq && freq <= 29700000 ) {
		code = AMATEUR
	}

	// 6 m
	if ( 50100000 <= freq && freq <= 54000000 ) {
		code = AMATEUR
	}

	// 6 m
	if ( 144100000 <= freq && freq <= 148000000 ) {
		code = AMATEUR
	}

	return code
}
