package main

func SplitAddress(address string) (string, string) {
	for i,_ := range address {
		if address[i] == ':'	{
			return address[:i], address[i:]
		}
	}

	return address	, ""
}