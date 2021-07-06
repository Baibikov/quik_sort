package main

func quickSort(n []int, l, r int) {
	var (
		lh, rh = l, r
		pivot = n[(l+r)/2]
	)

	for lh <= rh {
		for n[lh] < pivot {
			lh++
		}

		for n[rh] > pivot {
			rh--
		}

		if lh <= rh {
			if n[lh] > n[rh] {
				n[lh], n[rh] = n[rh], n[lh]
			}

			lh++
			rh--
		}
	}

	if l < rh {
		quickSort(n, l, rh)
	}

	if lh < r {
		quickSort(n, lh, r)
	}
}
