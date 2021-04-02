package main

func selection() {

	i := N
	for i > 0 {

		j := 1
		max := 0

		for j < N {

			if a[j] > a[max] {

				max = j
			}

			j++
		}

		t := a[j]
		a[max] = a[i]
		a[i] = t
		i--
	}
}
