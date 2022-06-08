package utils

func GetArtifactsPos(pos int64) string {
	var res string

	artpos := map[int64]string{
		1: "flower",
		2: "feather",
		3: "hourglass",
		4: "goblet",
		5: "crown",
	}
	for k, v := range artpos {
		if pos == k {
			res = v
		}
	}
	return res
}
